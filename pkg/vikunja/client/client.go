package client

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/charmbracelet/log"
)

// Client represents a Vikunja API client
// baseURL is now a string, not *url.URL
type Client struct {
	baseURL    string
	httpClient *http.Client
	token      string
}

// NewClient creates a new Vikunja client from environment variables
func NewClient() (*Client, error) {
	rawURL := os.Getenv("VIKUNJA_URL")
	if rawURL == "" {
		log.Error("VIKUNJA_URL environment variable is not set")
		return nil, fmt.Errorf("VIKUNJA_URL environment variable is not set")
	}

	// Validate the URL but store as string
	_, err := url.Parse(rawURL)
	if err != nil {
		log.Error("invalid VIKUNJA_URL", "url", rawURL, "error", err)
		return nil, fmt.Errorf("invalid VIKUNJA_URL: %w", err)
	}

	token := os.Getenv("VIKUNJA_TOKEN")
	if token == "" {
		log.Error("VIKUNJA_TOKEN environment variable is not set")
		return nil, fmt.Errorf("VIKUNJA_TOKEN environment variable is not set")
	}

	log.Info("creating new Vikunja client", "base_url", rawURL)
	return &Client{
		baseURL: rawURL,
		httpClient: &http.Client{
			Timeout:   30 * time.Second,
			Transport: authRoundTripper{token: token, rt: http.DefaultTransport},
		},
		token: token,
	}, nil
}

// NewClientWithCredentials creates a client with explicit credentials
func NewClientWithCredentials(baseURL, token string) (*Client, error) {
	// Validate the URL but store as string
	_, err := url.Parse(baseURL)
	if err != nil {
		log.Error("invalid base URL", "url", baseURL, "error", err)
		return nil, fmt.Errorf("invalid base URL: %w", err)
	}

	log.Info("creating new Vikunja client with explicit credentials", "base_url", baseURL)
	return &Client{
		baseURL: baseURL,
		httpClient: &http.Client{
			Timeout:   30 * time.Second,
			Transport: authRoundTripper{token: token, rt: http.DefaultTransport},
		},
		token: token,
	}, nil
}

// authRoundTripper injects the Authorization header and User‑Agent.
type authRoundTripper struct {
	token string
	rt    http.RoundTripper
}

func (a authRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Set("Authorization", "Bearer "+a.token)
	req.Header.Set("User-Agent", "vikunja-mcp/0.1.0")
	return a.rt.RoundTrip(req)
}

// url resolves a relative API path against the base URL.
func (c *Client) url(path string) string {
	// Use url.Parse to join baseURL and path safely
	base, err := url.Parse(c.baseURL)
	if err != nil {
		// fallback: just concatenate
		return c.baseURL + path
	}
	rel, err := url.Parse(path)
	if err != nil {
		return c.baseURL + path
	}
	return base.ResolveReference(rel).String()
}

// newRequest builds an HTTP request with context and common headers.
func (c *Client) newRequest(ctx context.Context, method, path string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequestWithContext(ctx, method, c.url(path), body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	return req, nil
}

// do executes the request and decodes JSON into v if provided.
func (c *Client) do(req *http.Request, v any) error {
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer func() {
		cerr := resp.Body.Close()
		if cerr != nil {
			log.Error("error closing response body", "error", cerr)
		}
	}()

	if resp.StatusCode >= 400 {
		return fmt.Errorf("vikunja API error: %s", resp.Status)
	}
	if v == nil {
		return nil
	}
	return json.NewDecoder(resp.Body).Decode(v)
}

// Me fetches the current authenticated user.
func (c *Client) Me(ctx context.Context) (*User, error) {
	req, err := c.newRequest(ctx, http.MethodGet, "/api/v1/user", nil)
	if err != nil {
		return nil, err
	}
	var u User
	if err := c.do(req, &u); err != nil {
		return nil, err
	}
	return &u, nil
}
