package client

import (
	"bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// mockRoundTripper implements http.RoundTripper for testing.
type mockRoundTripper struct {
	resp *http.Response
	err  error
}

func (m *mockRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	return m.resp, m.err
}

func newTestClient(resp *http.Response, err error) *Client {
	httpClient := &http.Client{
		Transport: &mockRoundTripper{resp: resp, err: err},
	}
	return &Client{
		httpClient: httpClient,
		baseURL:    "http://test/",
	}
}

func TestClient_Get_Success(t *testing.T) {
	expected := map[string]interface{}{"foo": "bar"}
	body := `{"foo":"bar"}`
	resp := &http.Response{
		StatusCode: 200,
		Body:       io.NopCloser(bytes.NewBufferString(body)),
		Header:     make(http.Header),
	}
	c := newTestClient(resp, nil)

	var result map[string]interface{}
	err := c.Get(context.Background(), "endpoint", &result)
	require.NoError(t, err)
	assert.Equal(t, expected, result)
}

func TestClient_Get_InvalidURL(t *testing.T) {
	c := &Client{baseURL: "http://[::1]:namedport/"}
	var result interface{}
	err := c.Get(context.Background(), "bad endpoint", &result)
	assert.Error(t, err)
}

func TestClient_Post_Success(t *testing.T) {
	expected := map[string]interface{}{"baz": "qux"}
	body := `{"baz":"qux"}`
	resp := &http.Response{
		StatusCode: 201,
		Body:       io.NopCloser(bytes.NewBufferString(body)),
		Header:     make(http.Header),
	}
	c := newTestClient(resp, nil)

	var result map[string]interface{}
	err := c.Post(context.Background(), "endpoint", map[string]string{"foo": "bar"}, &result)
	require.NoError(t, err)
	assert.Equal(t, expected, result)
}

func TestClient_Post_MarshalError(t *testing.T) {
	c := &Client{baseURL: "http://test/"}
	ch := make(chan int)
	var result interface{}
	err := c.Post(context.Background(), "endpoint", ch, &result)
	assert.Error(t, err)
}

func TestClient_Put_Success(t *testing.T) {
	expected := map[string]interface{}{"updated": true}
	body := `{"updated":true}`
	resp := &http.Response{
		StatusCode: 200,
		Body:       io.NopCloser(bytes.NewBufferString(body)),
		Header:     make(http.Header),
	}
	c := newTestClient(resp, nil)

	var result map[string]interface{}
	err := c.Put(context.Background(), "endpoint", map[string]string{"foo": "bar"}, &result)
	require.NoError(t, err)
	assert.Equal(t, expected, result)
}

func TestClient_Put_MarshalError(t *testing.T) {
	c := &Client{baseURL: "http://test/"}
	ch := make(chan int)
	var result interface{}
	err := c.Put(context.Background(), "endpoint", ch, &result)
	assert.Error(t, err)
}

func TestClient_Delete_Success(t *testing.T) {
	resp := &http.Response{
		StatusCode: 204,
		Body:       io.NopCloser(bytes.NewBuffer(nil)),
		Header:     make(http.Header),
	}
	c := newTestClient(resp, nil)

	err := c.Delete(context.Background(), "endpoint")
	require.NoError(t, err)
}

func TestClient_Delete_InvalidURL(t *testing.T) {
	c := &Client{baseURL: "http://[::1]:namedport/"}
	err := c.Delete(context.Background(), "bad endpoint")
	assert.Error(t, err)
}

func TestClient_Get_RequestError(t *testing.T) {
	c := &Client{
		baseURL: "http://test/",
		httpClient: &http.Client{
			Transport: &mockRoundTripper{resp: nil, err: errors.New("fail")},
		},
	}
	var result interface{}
	err := c.Get(context.Background(), "endpoint", &result)
	assert.Error(t, err)
}
