package client

import (
	"net/http"
	"os"
	"testing"
	"time"
)

func TestNewClient_envVars(t *testing.T) {
	os.Setenv("VIKUNJA_URL", "https://example.com")
	os.Setenv("VIKUNJA_TOKEN", "testtoken")
	c, err := NewClient()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if c.baseURL != "https://example.com" {
		t.Errorf("unexpected baseURL: %s", c.baseURL)
	}
	if c.token != "testtoken" {
		t.Errorf("unexpected token: %s", c.token)
	}
}

func TestNewClientWithCredentials_valid(t *testing.T) {
	c, err := NewClientWithCredentials("https://foo.com", "tok")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if c.baseURL != "https://foo.com" {
		t.Errorf("unexpected baseURL: %s", c.baseURL)
	}
	if c.token != "tok" {
		t.Errorf("unexpected token: %s", c.token)
	}
}

func TestClient_url(t *testing.T) {
	c := &Client{baseURL: "https://api.example.com"}
	full := c.url("/api/v1/user")
	if full != "https://api.example.com/api/v1/user" {
		t.Errorf("unexpected url: %s", full)
	}
}

func TestAuthRoundTripper(t *testing.T) {
	rt := authRoundTripper{token: "tok", rt: http.DefaultTransport}
	req, _ := http.NewRequest("GET", "https://foo", nil)
	resp := &http.Response{StatusCode: 200, Body: http.NoBody}
	// Wrap a dummy RoundTripper
	dummy := func(req *http.Request) (*http.Response, error) {
		if req.Header.Get("Authorization") != "Bearer tok" {
			t.Error("missing or wrong Authorization header")
		}
		if req.Header.Get("User-Agent") != "vikunja-mcp/0.1.0" {
			t.Error("missing or wrong User-Agent header")
		}
		return resp, nil
	}
	rt.rt = roundTripperFunc(dummy)
	_, _ = rt.RoundTrip(req)
}

type roundTripperFunc func(*http.Request) (*http.Response, error)

func (f roundTripperFunc) RoundTrip(r *http.Request) (*http.Response, error) {
	return f(r)
}

func TestClient_do_error_status(t *testing.T) {
	c := &Client{httpClient: &http.Client{Timeout: 1 * time.Second}}
	req, _ := http.NewRequest("GET", "https://example.com", nil)
	resp := &http.Response{StatusCode: 500, Body: http.NoBody}
	c.httpClient = &http.Client{Transport: roundTripperFunc(func(r *http.Request) (*http.Response, error) {
		return resp, nil
	})}
	err := c.do(req, nil)
	if err == nil {
		t.Error("expected error for 500 status")
	}
}
