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
func (c *Client) Get(ctx *context.Context, endpoint string, result interface{}) error {
	log.Info("GET request", "endpoint", endpoint)
	log.Debug("Base URL", "baseURL", c.baseURL)
	fullURL := c.baseURL + endpoint
	log.Debug("Full URL", "URL", fullURL)
	_, err := url.Parse(fullURL)
	if err != nil {
		log.Error("Invalid endpoint URL", "url", fullURL, "error", err)
		return fmt.Errorf("invalid endpoint: %w", err)
	}
	log.Debug("Building GET request", "url", fullURL)
	req, err := c.newRequest(ctx, http.MethodGet, fullURL, nil)
	if err != nil {
		log.Error("Failed to build GET request", "error", err)
		return err
	}
	log.Debug("Executing GET request", "url", fullURL)
	return c.do(req, result)
}

// Post performs a POST request with a JSON body and unmarshals the response into result.
func (c *Client) Post(ctx *context.Context, endpoint string, body interface{}, result interface{}) error {
	log.Info("POST request", "endpoint", endpoint)
	var r io.Reader
	if body != nil {
		b, err := json.Marshal(body)
		if err != nil {
			log.Error("Failed to marshal POST body", "error", err)
			return fmt.Errorf("marshal body: %w", err)
		}
		log.Debug("POST body marshaled", "body", string(b))
		r = bytes.NewReader(b)
	}
	log.Debug("Building POST request", "endpoint", endpoint)
	req, err := c.newRequest(ctx, http.MethodPost, endpoint, r)
	if err != nil {
		log.Error("Failed to build POST request", "error", err)
		return err
	}
	log.Debug("Executing POST request", "endpoint", endpoint)
	return c.do(req, result)
}

// Put performs a PUT request with a JSON body and unmarshals the response into result.
func (c *Client) Put(ctx *context.Context, endpoint string, body interface{}, result interface{}) error {
	log.Info("PUT request", "endpoint", endpoint)
	var r io.Reader
	if body != nil {
		b, err := json.Marshal(body)
		if err != nil {
			log.Error("Failed to marshal PUT body", "error", err)
			return fmt.Errorf("marshal body: %w", err)
		}
		log.Debug("PUT body marshaled", "body", string(b))
		r = bytes.NewReader(b)
	}
	log.Debug("Building PUT request", "endpoint", endpoint)
	req, err := c.newRequest(ctx, http.MethodPut, endpoint, r)
	if err != nil {
		log.Error("Failed to build PUT request", "error", err)
		return err
	}
	log.Debug("Executing PUT request", "endpoint", endpoint)
	return c.do(req, result)
}

// Delete performs a DELETE request.
func (c *Client) Delete(ctx *context.Context, endpoint string) error {
	log.Info("DELETE request", "endpoint", endpoint)
	fullURL := c.baseURL + endpoint
	log.Debug("Full URL", "URL", fullURL)
	_, err := url.Parse(fullURL)
	if err != nil {
		log.Error("Invalid endpoint URL for DELETE", "url", fullURL, "error", err)
		return fmt.Errorf("invalid endpoint: %w", err)
	}
	log.Debug("Building DELETE request", "url", fullURL)
	req, err := c.newRequest(ctx, http.MethodDelete, fullURL, nil)
	if err != nil {
		log.Error("Failed to build DELETE request", "error", err)
		return err
	}
	log.Debug("Executing DELETE request", "url", fullURL)
	return c.do(req, nil)
}
