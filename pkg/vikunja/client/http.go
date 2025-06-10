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

	// Join baseURL and endpoint correctly using url.URL
	_, err := url.Parse(endpoint)
	if err != nil {
		return fmt.Errorf("invalid endpoint: %w", err)
	}
	fullURL := c.baseURL + endpoint
	log.Info("Full URL", "URL", fullURL)

	req, err := c.newRequest(ctx, http.MethodGet, fullURL, nil)
	if err != nil {
		return err
	}
	return c.do(req, result)
}

// Post performs a POST request with a JSON body and unmarshals the response into result.
func (c *Client) Post(ctx context.Context, endpoint string, body interface{}, result interface{}) error {
	log.Info("POST request", "endpoint", endpoint)

	var r io.Reader
	if body != nil {
		b, err := json.Marshal(body)
		if err != nil {
			return fmt.Errorf("marshal body: %w", err)
		}
		r = bytes.NewReader(b)
	}

	req, err := c.newRequest(ctx, http.MethodPost, endpoint, r)
	if err != nil {
		return err
	}
	return c.do(req, result)
}

// Put performs a PUT request with a JSON body and unmarshals the response into result.
func (c *Client) Put(ctx context.Context, endpoint string, body interface{}, result interface{}) error {
	log.Info("PUT request", "endpoint", endpoint)

	var r io.Reader
	if body != nil {
		b, err := json.Marshal(body)
		if err != nil {
			return fmt.Errorf("marshal body: %w", err)
		}
		r = bytes.NewReader(b)
	}

	req, err := c.newRequest(ctx, http.MethodPut, endpoint, r)
	if err != nil {
		return err
	}
	return c.do(req, result)
}

// Delete performs a DELETE request.
func (c *Client) Delete(ctx context.Context, endpoint string) error {
	log.Info("DELETE request", "endpoint", endpoint)

	req, err := c.newRequest(ctx, http.MethodDelete, endpoint, nil)
	if err != nil {
		return err
	}
	return c.do(req, nil)
}
