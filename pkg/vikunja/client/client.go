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
	log.Info("NewClient called for Vikunja API")
	rawURL := os.Getenv("VIKUNJA_URL")
	if rawURL == "" {
		log.Error("VIKUNJA_URL environment variable is not set")
		return nil, fmt.Errorf("VIKUNJA_URL environment variable is not set")
	}
	log.Debug("VIKUNJA_URL loaded", "url", rawURL)
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
	log.Debug("VIKUNJA_TOKEN loaded")
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
	log.Info("NewClientWithCredentials called", "base_url", baseURL)
	// Validate the URL but store as string
	_, err := url.Parse(baseURL)
	if err != nil {
		log.Error("invalid base URL", "url", baseURL, "error", err)
		return nil, fmt.Errorf("invalid base URL: %w", err)
	}
	log.Debug("Token provided for NewClientWithCredentials")
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
	log.Debug("authRoundTripper injecting headers")
	req.Header.Set("Authorization", "Bearer "+a.token)
	req.Header.Set("User-Agent", "vikunja-mcp/0.1.0")
	return a.rt.RoundTrip(req)
}

// url resolves a relative API path against the base URL.
func (c *Client) url(path string) string {
	log.Debug("url called", "baseURL", c.baseURL, "path", path)
	base, err := url.Parse(c.baseURL)
	if err != nil {
		log.Error("Failed to parse baseURL", "baseURL", c.baseURL, "error", err)
		return c.baseURL + path
	}
	rel, err := url.Parse(path)
	if err != nil {
		log.Error("Failed to parse path for url", "path", path, "error", err)
		return c.baseURL + path
	}
	return base.ResolveReference(rel).String()
}

// newRequest builds an HTTP request with context and common headers.
func (c *Client) newRequest(ctx *context.Context, method, path string, body io.Reader) (*http.Request, error) {
	log.Debug("newRequest called", "method", method, "path", path)
	req, err := http.NewRequestWithContext(ctx, method, c.url(path), body)
	if err != nil {
		log.Error("Failed to create new HTTP request", "error", err)
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	if body != nil {
		log.Debug("Setting Content-Type header for request body")
		req.Header.Set("Content-Type", "application/json")
	}
	return req, nil
}

// do executes the request and decodes JSON into v if provided.
// Adds retry logic and improved error handling.
func (c *Client) do(req *http.Request, v any) error {
	log.Debug("do called for HTTP request", "method", req.Method, "url", req.URL.String())
	const maxRetries = 3
	const retryDelay = 500 * time.Millisecond

	var lastErr error
	for attempt := 0; attempt < maxRetries; attempt++ {
		log.Debug("HTTP request attempt", "attempt", attempt+1)
		resp, err := c.httpClient.Do(req)
		if err != nil {
			lastErr = err
			log.Error("HTTP request failed, will retry if possible", "attempt", attempt+1, "error", err)
			time.Sleep(retryDelay * time.Duration(attempt+1))
			continue
		}
		defer func() {
			cerr := resp.Body.Close()
			if cerr != nil {
				log.Error("error closing response body", "error", cerr)
			}
		}()

		if resp.StatusCode >= 500 && resp.StatusCode < 600 {
			lastErr = fmt.Errorf("server error: %s", resp.Status)
			log.Error("Server error, will retry", "status", resp.Status, "attempt", attempt+1)
			time.Sleep(retryDelay * time.Duration(attempt+1))
			continue
		}
		if resp.StatusCode >= 400 {
			log.Error("vikunja API error", "status", resp.Status)
			return fmt.Errorf("vikunja API error: %s", resp.Status)
		}
		if v == nil {
			log.Debug("No output variable provided for response body")
			return nil
		}
		log.Debug("Decoding response body to output variable")
		return json.NewDecoder(resp.Body).Decode(v)
	}
	log.Error("request failed after max retries", "retries", maxRetries, "lastErr", lastErr)
	return fmt.Errorf("request failed after %d attempts: %w", maxRetries, lastErr)
}
