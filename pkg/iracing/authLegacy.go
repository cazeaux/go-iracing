package iracing

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"sync"

	"github.com/cazeaux/go-iracing/pkg/config"
)

// Authenticator applique un mécanisme d'authentification à une requête sortante.
type Authenticator interface {
	Apply(req *http.Request) error
	DoLogin(ctx context.Context, c *Client) error
	IsAuthenticated(ctx context.Context, c *Client) bool
	ResetAuth()
}

type IrAuthRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type IrAuthResp struct {
	Authcode        string `json:"authcode"`
	AutoLoginSeries string `json:"autoLoginSeries"`
	AutoLoginToken  string `json:"autoLoginToken"`
	CustID          int    `json:"custId"`
	Email           string `json:"email"`
	SsoCookieDomain string `json:"ssoCookieDomain"`
	SsoCookieName   string `json:"ssoCookieName"`
	SsoCookiePath   string `json:"ssoCookiePath"`
	SsoCookieValue  string `json:"ssoCookieValue"`
}

// IR Legacy Auth
func NewIrLegacyAuth(username, password string) Authenticator {
	return &IrLegacyAuth{
		username:        username,
		password:        password,
		isAuthenticated: false,
	}
}

type IrLegacyAuth struct {
	username, password string
	isAuthenticated    bool
	cookie             string
	mu                 sync.Mutex
}

func (a *IrLegacyAuth) Apply(req *http.Request) error {
	return nil
}

func (a *IrLegacyAuth) ResetAuth() {
	a.mu.Lock()
	a.isAuthenticated = false
	a.mu.Unlock()
}

func (a *IrLegacyAuth) DoLogin(ctx context.Context, c *Client) error {
	a.mu.Lock()
	defer a.mu.Unlock()
	cookie, err := doLogin(ctx, c, a.username, a.password)
	if err != nil {
		a.isAuthenticated = false
		return err
	}
	a.cookie = cookie
	a.isAuthenticated = true
	return nil
}

func (a *IrLegacyAuth) IsAuthenticated(ctx context.Context, c *Client) bool {
	a.mu.Lock()
	defer a.mu.Unlock()
	return a.isAuthenticated
}

func doLogin(ctx context.Context, c *Client, username, password string) (string, error) {
	var reader io.Reader

	params := IrAuthRequest{
		Email:    username,
		Password: EncodePassword(username, password),
	}

	if err := c.throttle(ctx); err != nil {
		return "", err
	}

	b, err := json.Marshal(params)
	if err != nil {
		return "", err
	}
	reader = bytes.NewReader(b)

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, config.IR_API_URL+config.IR_PATH_LEGACY_AUTH, reader)
	if err != nil {
		return "", err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return "", err
	}
	defer func() { _ = resp.Body.Close() }()

	var authResp IrAuthResp
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &authResp)
	if err != nil {
		return err
	}

	return authResp.SsoCookieValue, nil
}

func EncodePassword(username, password string) string {
	// Concatène le mot de passe et le nom d'utilisateur en minuscules
	input := password + strings.ToLower(username)

	// Calcul du SHA256
	hash := sha256.Sum256([]byte(input))

	// Encodage en Base64
	return base64.StdEncoding.EncodeToString(hash[:])
}
