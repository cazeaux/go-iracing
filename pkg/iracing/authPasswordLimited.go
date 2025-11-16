package iracing

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/cazeaux/go-iracing/pkg/config"
	"github.com/google/go-querystring/query"
)

type IrOAuthRequest struct {
	Username     string `url:"username"`
	Password     string `url:"password"`
	GrantType    string `url:"grant_type"`
	ClientID     string `url:"client_id"`
	ClientSecret string `url:"client_secret"`
	Scope        string `url:"scope"`
}

type IrOAuthResp struct {
	AccessToken        string `json:"access_token"`
	TokenType          string `json:"token_type"`
	ExpiresIn          int    `json:"expires_in"`
	RefreshToken       string `json:"refresh_token"`
	RefreshTokenExpiry int    `json:"refresh_token_expires_in"`
}

// IR Legacy Auth
func NewIrOAuth(username, password, clientID, clientSecret string) Authenticator {
	return &IrOauthService{
		username:        username,
		password:        password,
		clientID:        clientID,
		clientSecret:    clientSecret,
		isAuthenticated: false,
	}
}

type IrOauthService struct {
	username, password, clientID, clientSecret string
	isAuthenticated                            bool
	tokenExpiryDate                            time.Time
	refreshTokenExpiryDate                     time.Time
	token, refreshToken                        string
	mu                                         sync.Mutex
}

func (a *IrOauthService) Apply(req *http.Request) error {
	req.Header.Set("Authorization", "Bearer "+a.token)
	return nil
}

func (a *IrOauthService) ResetAuth() {
	a.mu.Lock()
	a.isAuthenticated = false
	a.mu.Unlock()
}

func (a *IrOauthService) IsTokenExpired() bool {
	return time.Now().After(a.tokenExpiryDate)
}

func (a *IrOauthService) DoLogin(ctx context.Context, c *Client) error {
	a.mu.Lock()
	defer a.mu.Unlock()

	var reader io.Reader

	params := IrOAuthRequest{
		Username:     a.username,
		Password:     EncodePassword(a.username, a.password),
		GrantType:    "password_limited",
		ClientID:     a.clientID,
		ClientSecret: EncodePassword(a.clientID, a.clientSecret),
		Scope:        "iracing.auth",
	}

	if err := c.throttle(ctx); err != nil {
		return err
	}

	dataReq, err := query.Values(params)
	if err != nil {
		return err
	}
	reader = strings.NewReader(dataReq.Encode())

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, config.IR_OAUTH_URL, reader)
	if err != nil {
		return err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "x-www-form-urlencoded")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode >= 400 {
		body, _ := io.ReadAll(resp.Body)
		_ = resp.Body.Close()
		return fmt.Errorf("oauth login failed: %v %s", resp.StatusCode, body)
	}

	var authResp IrOAuthResp
	body, _ := io.ReadAll(resp.Body)
	defer func() { _ = resp.Body.Close() }()
	if err := json.Unmarshal(body, &authResp); err != nil {
		return err
	}

	a.token = authResp.AccessToken
	a.refreshToken = authResp.RefreshToken
	a.tokenExpiryDate = time.Now().Add(time.Duration(authResp.ExpiresIn) * time.Second)
	a.refreshTokenExpiryDate = time.Now().Add(time.Duration(authResp.RefreshTokenExpiry) * time.Second)
	a.isAuthenticated = true

	return nil
}

func (a *IrOauthService) IsAuthenticated(ctx context.Context, c *Client) bool {
	a.mu.Lock()
	defer a.mu.Unlock()
	return a.isAuthenticated && !a.IsTokenExpired()
}
