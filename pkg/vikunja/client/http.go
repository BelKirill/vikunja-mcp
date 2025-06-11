package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/charmbracelet/log"
)

// Get performs a GET request and unmarshals the response into result.
func (c *Client) Get(ctx context.Context, endpoint string, result interface{}) error {
	log.Info("GET request", "endpoint", endpoint)
	log.Info(c.baseURL)

	fullURL := c.baseURL + endpoint
	log.Info("Full URL", "URL", fullURL)
	_, err := url.Parse(fullURL)
	if err != nil {
		return fmt.Errorf("invalid endpoint: %w", err)
	}

	req, err := c.newRequest(ctx, http.MethodGet, fullURL, nil)
	if err != nil {
		return err
	}
	return c.do(req, result)
}

// Post performs a POST request with a JSON body and unmarshals the response into result.
func (c *Client) Post(ctx context.Context, endpoint string, body interface{}, result interface{}) error {
	log.Info("POST request", "endpoint", endpoint)

	fullURL := c.baseURL + endpoint
	log.Info("Full URL", "URL", fullURL)
	_, err := url.Parse(fullURL)
	if err != nil {
		return fmt.Errorf("invalid endpoint: %w", err)
	}

	var r io.Reader
	if body != nil {
		b, err := json.Marshal(body)
		if err != nil {
			return fmt.Errorf("marshal body: %w", err)
		}
		r = bytes.NewReader(b)
	}

	req, err := c.newRequest(ctx, http.MethodPost, fullURL, r)
	if err != nil {
		return err
	}
	return c.do(req, result)
}

// Put performs a PUT request with a JSON body and unmarshals the response into result.
func (c *Client) Put(ctx context.Context, endpoint string, body interface{}, result interface{}) error {
	log.Info("PUT request", "endpoint", endpoint)

	fullURL := c.baseURL + endpoint
	log.Info("Full URL", "URL", fullURL)
	_, err := url.Parse(fullURL)
	if err != nil {
		return fmt.Errorf("invalid endpoint: %w", err)
	}

	var r io.Reader
	if body != nil {
		b, err := json.Marshal(body)
		if err != nil {
			return fmt.Errorf("marshal body: %w", err)
		}
		r = bytes.NewReader(b)
	}

	req, err := c.newRequest(ctx, http.MethodPut, fullURL, r)
	if err != nil {
		return err
	}
	return c.do(req, result)
}

// Delete performs a DELETE request.
func (c *Client) Delete(ctx context.Context, endpoint string) error {
	log.Info("DELETE request", "endpoint", endpoint)

	fullURL := c.baseURL + endpoint
	log.Info("Full URL", "URL", fullURL)
	_, err := url.Parse(fullURL)
	if err != nil {
		return fmt.Errorf("invalid endpoint: %w", err)
	}

	req, err := c.newRequest(ctx, http.MethodDelete, fullURL, nil)
	if err != nil {
		return err
	}
	return c.do(req, nil)
}
