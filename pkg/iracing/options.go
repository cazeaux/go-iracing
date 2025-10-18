package iracing

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// Option représente une option de configuration pour NewClient.
type Option func(*Client) error

func WithBaseURL(raw string) Option {
	return func(c *Client) error {
		u, err := url.Parse(strings.TrimRight(raw, "/"))
		if err != nil {
			return err
		}
		c.baseURL = u
		return nil
	}
}

func WithHTTPClient(h *http.Client) Option {
	return func(c *Client) error { c.httpClient = h; return nil }
}

func WithUserAgent(ua string) Option {
	return func(c *Client) error { c.ua = ua; return nil }
}

// WithRateLimit fixe un simple seau à jetons (token bucket) par fenêtre.
func WithRateLimit(maxPerWindow int, window time.Duration) Option {
	return func(c *Client) error {
		if maxPerWindow <= 0 || window <= 0 {
			return fmt.Errorf("invalid rate limit")
		}
		c.ratePer = maxPerWindow
		c.rateWindow = window
		c.tokens = maxPerWindow
		c.windowStart = time.Now()
		return nil
	}
}

func WithAuthenticator(a Authenticator) Option {
	return func(c *Client) error { c.auth = a; return nil }
}

// WithRetry configure les relances automatiques pour 429/5xx.
func WithRetry(max int, backoff time.Duration) Option {
	return func(c *Client) error { c.retryMax, c.retryBackoff = max, backoff; return nil }
}
