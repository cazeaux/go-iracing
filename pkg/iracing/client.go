package iracing

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"
	"sync"
	"time"

	"github.com/cazeaux/go-iracing/pkg/config"
	"github.com/cazeaux/go-iracing/pkg/types"
)

// ----- Client core -----

type Client struct {
	baseURL    *url.URL
	httpClient *http.Client
	ua         string

	// retry/rate-limit
	mu           sync.Mutex
	ratePer      int           // nb de requêtes par fenêtre
	rateWindow   time.Duration // taille de fenêtre
	tokens       int
	windowStart  time.Time
	retryMax     int
	retryAuthMax int
	retryBackoff time.Duration

	auth Authenticator

	// Services
	Car                *CarService
	Track              *TrackService
	Member             *MemberService
	Stats              *StatsService
	Results            *ResultsService
	CarClass           *CarClassService
	ConstantsDivisions *ConstantsService
	LeagueService      *LeagueService
	LookupService      *LookupService
	SeriesService	   *SeriesService
}

// NewClient crée une nouvelle instance Client.
func NewClient(opts ...Option) (*Client, error) {
	base, _ := url.Parse(config.IR_API_URL)
	jar, _ := cookiejar.New(nil)
	c := &Client{
		baseURL:      base,
		httpClient:   &http.Client{Timeout: 30 * time.Second, Jar: jar},
		ua:           fmt.Sprintf("go-iracing/%v", config.Version),
		ratePer:      10,
		rateWindow:   time.Second,
		tokens:       10,
		windowStart:  time.Now(),
		retryMax:     3,
		retryAuthMax: 2,
		retryBackoff: 500 * time.Millisecond,
	}
	for _, opt := range opts {
		if err := opt(c); err != nil {
			return nil, err
		}
	}

	// Services
	c.Car = &CarService{client: c}
	c.Track = &TrackService{client: c}
	c.Member = &MemberService{client: c}
	c.Stats = &StatsService{client: c}
	c.Results = &ResultsService{client: c}
	c.CarClass = &CarClassService{client: c}
	c.ConstantsDivisions = &ConstantsService{client: c}
	c.LeagueService = &LeagueService{client: c}
	c.LookupService = &LookupService{client: c}
	c.SeriesService = &SeriesService{client: c}
	return c, nil
}

// buildURL combine baseURL + path + query.
func (c *Client) buildURL(path string, q url.Values) string {
	u := *c.baseURL
	u.Path = strings.TrimRight(c.baseURL.Path, "/") + "/" + strings.TrimLeft(path, "/")
	u.RawQuery = q.Encode()
	return u.String()
}

// throttle applique un contrôle simple de débit (token bucket par fenêtre).
func (c *Client) throttle(ctx context.Context) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	now := time.Now()
	if now.Sub(c.windowStart) >= c.rateWindow {
		c.windowStart = now
		c.tokens = c.ratePer
	}
	for c.tokens <= 0 {
		// dormir jusqu'à la prochaine fenêtre
		wait := c.rateWindow - now.Sub(c.windowStart)
		c.mu.Unlock()
		select {
		case <-ctx.Done():
			c.mu.Lock()
			return ctx.Err()
		case <-time.After(wait):
			c.mu.Lock()
			now = time.Now()
			c.windowStart = now
			c.tokens = c.ratePer
		}
	}
	c.tokens--
	return nil
}

// ErrorResponse représente une erreur JSON standardisée renvoyée par l'API.
type ErrorResponse struct {
	Status     int    `json:"status"`
	Code       string `json:"code"`
	Message    string `json:"message"`
	RequestID  string `json:"request_id,omitempty"`
	RetryAfter int    `json:"retry_after,omitempty"`
}

func (e *ErrorResponse) Error() string {
	return fmt.Sprintf("iracing: status=%d code=%s message=%s", e.Status, e.Code, e.Message)
}

func (c *Client) getAwsRessource(ctx context.Context, url string, out any) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	if out != nil {
		defer resp.Body.Close()
		dec := json.NewDecoder(resp.Body)
		if err := dec.Decode(out); err != nil && !errors.Is(err, io.EOF) {
			return resp, err
		}
	}
	return resp, nil

}

// do exécute une requête HTTP avec gestion de l'auth, du throttling et des retries.
func (c *Client) do(ctx context.Context, method, path string, query url.Values, body any, out any) (*http.Response, error) {
	if err := c.throttle(ctx); err != nil {
		return nil, err
	}

	var reader io.Reader
	if body != nil {
		b, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}
		reader = bytes.NewReader(b)
	}

	req, err := http.NewRequestWithContext(ctx, method, c.buildURL(path, query), reader)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	if c.ua != "" {
		req.Header.Set("User-Agent", c.ua)
	}
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	attempt := 0
	for {
		if c.auth != nil {
			if !c.auth.IsAuthenticated(ctx, c) {
				if err := c.auth.DoLogin(ctx, c); err != nil {
					return nil, err
				}
			}

			if err := c.auth.Apply(req); err != nil {
				return nil, err
			}
		}

		resp, err := c.httpClient.Do(req)
		if err != nil {
			if errors.Is(err, context.Canceled) || errors.Is(err, context.DeadlineExceeded) {
				return nil, err
			}
			if attempt >= c.retryMax {
				return nil, err
			}
			attempt++
			time.Sleep(c.retryBackoff * time.Duration(attempt))
			continue
		}

		if resp.StatusCode >= 200 && resp.StatusCode < 300 {
			if out != nil {
				defer resp.Body.Close()
				dec := json.NewDecoder(resp.Body)
				if err := dec.Decode(out); err != nil && !errors.Is(err, io.EOF) {
					return resp, err
				}
			}
			return resp, nil
		}

		if resp.StatusCode == http.StatusUnauthorized && attempt < c.retryAuthMax {
			attempt++
			if c.auth != nil {
				c.auth.ResetAuth()
				time.Sleep(c.retryBackoff)
				continue
			} else {
				return resp, fmt.Errorf("unauthorized and no authentication provider")
			}
		}

		if resp.StatusCode >= 500 {
			return resp, fmt.Errorf("service unavailable: %v", resp.StatusCode)
		}

		return resp, nil
	}
}

// get/do helpers
func (c *Client) get(ctx context.Context, path string, query url.Values, out any) (*http.Response, error) {
	return c.do(ctx, http.MethodGet, path, query, nil, out)
}
func (c *Client) post(ctx context.Context, path string, query url.Values, in, out any) (*http.Response, error) {
	return c.do(ctx, http.MethodPost, path, query, in, out)
}

func (c *Client) getRessource(ctx context.Context, path string, query url.Values, out any) (*http.Response, error) {
	var ressourceData types.RessourceLinkResp

	respIR, err := c.get(ctx, path, query, &ressourceData)
	if err != nil {
		return respIR, err
	}

	if respIR.StatusCode != http.StatusOK {
		return respIR, fmt.Errorf("error on iracing query %v: %v", path, respIR.StatusCode)
	}
	return c.getAwsRessource(ctx, ressourceData.Link, &out)
}

func GetRessourceChunks[T any](ctx context.Context, c *Client, path string, query url.Values, rows int, out *[]T) (*http.Response, error) {
	var ressourceData types.RessourceLinkChunksResp

	respIR, err := c.get(ctx, path, query, &ressourceData)
	if err != nil {
		return respIR, err
	}

	if respIR.StatusCode != http.StatusOK {
		return respIR, fmt.Errorf("error on iracing query %v: %v", path, respIR.StatusCode)
	}

	return GetAwsRessourceChunks(ctx, c, &ressourceData.Data.ChunkInfo, rows, out)

}

// ----- Pagination générique -----

type Page[T any] struct {
	Items    []T    `json:"items"`
	Next     string `json:"next,omitempty"` // URL ou cursor
	Previous string `json:"previous,omitempty"`
	Total    int    `json:"total,omitempty"`
}

// Pager itère sur toutes les pages d'un endpoint list.
func Pager[T any](ctx context.Context, fetch func(ctx context.Context, cursor string) (Page[T], *http.Response, error)) ([]T, *http.Response, error) {
	var (
		all  []T
		resp *http.Response
		cur  string
	)
	for {
		p, r, err := fetch(ctx, cur)
		if err != nil {
			return all, r, err
		}
		resp = r
		all = append(all, p.Items...)
		if p.Next == "" {
			break
		}
		cur = p.Next
	}
	return all, resp, nil
}

func GetAwsRessourceChunks[T any](ctx context.Context, c *Client, chunkInfo *types.ChunkInfoType, rows int, out *[]T) (*http.Response, error) {
	chunksToFetch := chunkInfo.NumChunks
	if rows <= 0 {
		rows = chunkInfo.Rows
		chunksToFetch = rows / chunkInfo.ChunkSize
	}

	for _, chunkURL := range chunkInfo.ChunkFileNames[:chunksToFetch] {
		fullURL := fmt.Sprintf("%s%s", chunkInfo.BaseDownloadURL, chunkURL)
		req, err := http.NewRequestWithContext(ctx, http.MethodGet, fullURL, nil)
		if err != nil {
			return nil, err
		}

		resp, err := c.httpClient.Do(req)
		if err != nil {
			return nil, err
		}

		var outSlice []T
		defer resp.Body.Close()
		dec := json.NewDecoder(resp.Body)
		if err := dec.Decode(&outSlice); err != nil && !errors.Is(err, io.EOF) {
			return resp, err
		}
		*out = append(*out, outSlice...)

	}
	return nil, nil
}
