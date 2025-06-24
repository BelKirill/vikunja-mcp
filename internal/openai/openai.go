// Package openai provides OpenAI-powered decision engine for task selection
package openai

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/charmbracelet/log"
)

// OpenAIConfig holds configuration for OpenAI API
type OpenAIConfig struct {
	APIKey      string
	BaseURL     string        // Default: "https://api.openai.com/v1"
	Model       string        // Default: "gpt-4o-mini"
	MaxTokens   int           // Default: 1000
	Temperature float64       // Default: 0.3 for consistent decisions
	Timeout     time.Duration // Default: 30s
}

// OpenAIDecisionEngine implements the DecisionEngine interface using OpenAI
type OpenAIDecisionEngine struct {
	config     OpenAIConfig
	httpClient *http.Client
}

// NewOpenAIDecisionEngine creates a new OpenAI-powered decision engine
func NewOpenAIDecisionEngine(config OpenAIConfig) *OpenAIDecisionEngine {
	// Set defaults
	if config.BaseURL == "" {
		config.BaseURL = "https://api.openai.com/v1"
	}
	if config.Model == "" {
		config.Model = "gpt-4o-mini"
	}
	if config.MaxTokens == 0 {
		config.MaxTokens = 1000
	}
	if config.Temperature == 0 {
		config.Temperature = 0.3
	}
	if config.Timeout == 0 {
		config.Timeout = 30 * time.Second
	}

	return &OpenAIDecisionEngine{
		config: config,
		httpClient: &http.Client{
			Timeout: config.Timeout,
		},
	}
}

// OpenAI API request/response structures
type openAIRequest struct {
	Model       string    `json:"model"`
	Messages    []message `json:"messages"`
	MaxTokens   int       `json:"max_tokens"`
	Temperature float64   `json:"temperature"`
}

type message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type openAIResponse struct {
	Choices []choice  `json:"choices"`
	Error   *apiError `json:"error,omitempty"`
}

type choice struct {
	Message message `json:"message"`
}

type apiError struct {
	Message string `json:"message"`
	Type    string `json:"type"`
	Code    string `json:"code"`
}

// sanitizeResponse removes anything other than JSON. Handles both objects {} and arrays []
func (e *OpenAIDecisionEngine) sanitizeResponse(response string) (string, error) {
	// Find the start of JSON object or array
	startObj := strings.Index(response, "{")
	startArr := strings.Index(response, "[")

	var start int
	var endChar string

	if startObj == -1 && startArr == -1 {
		log.Error("No JSON found in response", "response", response)
		return "", fmt.Errorf("no JSON found in response")
	} else if startObj == -1 {
		// Only array found
		start = startArr
		endChar = "]"
	} else if startArr == -1 {
		// Only object found
		start = startObj
		endChar = "}"
	} else if startObj < startArr {
		// Object comes first
		start = startObj
		endChar = "}"
	} else {
		// Array comes first
		start = startArr
		endChar = "]"
	}

	// Find the corresponding end character
	end := strings.LastIndex(response, endChar)
	if end == -1 || start > end {
		log.Error("No matching JSON end found in response", "response", response, "expected_end", endChar)
		return "", fmt.Errorf("no matching JSON end found in response")
	}

	cleanResponse := response[start : end+1]
	return cleanResponse, nil
}

// callOpenAI makes the actual API call to OpenAI
func (e *OpenAIDecisionEngine) callOpenAI(ctx context.Context, prompt string) (string, error) {
	log.Info("Calling OpenAI API", "model", e.config.Model, "max_tokens", e.config.MaxTokens)
	log.Debug("Prompt sent to OpenAI", "prompt", prompt)

	reqBody := openAIRequest{
		Model:       e.config.Model,
		MaxTokens:   e.config.MaxTokens,
		Temperature: e.config.Temperature,
		Messages:    []message{{Role: "user", Content: prompt}},
	}

	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		log.Error("Failed to marshal OpenAI request body", "error", err)
		return "", fmt.Errorf("failed to marshal request: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "POST", e.config.BaseURL+"/chat/completions", bytes.NewBuffer(jsonBody))
	if err != nil {
		log.Error("Failed to create OpenAI HTTP request", "error", err)
		return "", fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+e.config.APIKey)

	log.Debug("Sending HTTP request to OpenAI", "url", req.URL.String())
	resp, err := e.httpClient.Do(req)
	if err != nil {
		log.Error("HTTP request to OpenAI failed", "error", err)
		return "", fmt.Errorf("HTTP request failed: %w", err)
	}
	defer func() {
		cerr := resp.Body.Close()
		if cerr != nil {
			log.Warn("warning: failed to close response body", "error", cerr)
		}
	}()

	log.Debug("Received HTTP response from OpenAI", "status", resp.StatusCode)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Error("Failed to read OpenAI response body", "error", err)
		return "", fmt.Errorf("failed to read response: %w", err)
	}

	log.Debug("OpenAI response body", "body", string(body))
	if resp.StatusCode != http.StatusOK {
		log.Error("OpenAI API returned error status", "status", resp.StatusCode, "body", string(body))
		return "", fmt.Errorf("OpenAI API error (status %d): %s", resp.StatusCode, string(body))
	}

	var openAIResp openAIResponse
	if err := json.Unmarshal(body, &openAIResp); err != nil {
		log.Error("Failed to unmarshal OpenAI response JSON", "error", err)
		return "", fmt.Errorf("failed to unmarshal response: %w", err)
	}

	if openAIResp.Error != nil {
		log.Error("OpenAI API returned error", "error", openAIResp.Error.Message)
		return "", fmt.Errorf("OpenAI API error: %s", openAIResp.Error.Message)
	}

	if len(openAIResp.Choices) == 0 {
		log.Error("No choices in OpenAI response")
		return "", fmt.Errorf("no choices in OpenAI response")
	}

	log.Info("OpenAI API call successful, returning content")
	return openAIResp.Choices[0].Message.Content, nil
}

// callOpenAIWithSystem makes an OpenAI API call with both system and user prompts
func (e *OpenAIDecisionEngine) callOpenAIWithSystem(ctx context.Context, systemPrompt, userPrompt string) (string, error) {
	log.Info("Calling OpenAI API with system/user prompts", "model", e.config.Model)
	log.Debug("System prompt", "prompt", systemPrompt)
	log.Debug("User prompt", "prompt", userPrompt)

	reqBody := openAIRequest{
		Model:       e.config.Model,
		MaxTokens:   e.config.MaxTokens,
		Temperature: e.config.Temperature,
		Messages: []message{
			{Role: "system", Content: systemPrompt},
			{Role: "user", Content: userPrompt},
		},
	}

	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		log.Error("Failed to marshal OpenAI request body", "error", err)
		return "", fmt.Errorf("failed to marshal request: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "POST", e.config.BaseURL+"/chat/completions", bytes.NewBuffer(jsonBody))
	if err != nil {
		log.Error("Failed to create OpenAI HTTP request", "error", err)
		return "", fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+e.config.APIKey)

	log.Debug("Sending HTTP request to OpenAI", "url", req.URL.String())
	resp, err := e.httpClient.Do(req)
	if err != nil {
		log.Error("HTTP request to OpenAI failed", "error", err)
		return "", fmt.Errorf("HTTP request failed: %w", err)
	}
	defer func() {
		cerr := resp.Body.Close()
		if cerr != nil {
			log.Warn("warning: failed to close response body", "error", cerr)
		}
	}()

	log.Debug("Received HTTP response from OpenAI", "status", resp.StatusCode)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Error("Failed to read OpenAI response body", "error", err)
		return "", fmt.Errorf("failed to read response: %w", err)
	}

	log.Debug("OpenAI response body", "body", string(body))
	if resp.StatusCode != http.StatusOK {
		log.Error("OpenAI API returned error status", "status", resp.StatusCode, "body", string(body))
		return "", fmt.Errorf("OpenAI API error (status %d): %s", resp.StatusCode, string(body))
	}

	var openAIResp openAIResponse
	if err := json.Unmarshal(body, &openAIResp); err != nil {
		log.Error("Failed to unmarshal OpenAI response JSON", "error", err)
		return "", fmt.Errorf("failed to unmarshal response: %w", err)
	}

	if openAIResp.Error != nil {
		log.Error("OpenAI API returned error", "error", openAIResp.Error.Message)
		return "", fmt.Errorf("OpenAI API error: %s", openAIResp.Error.Message)
	}

	if len(openAIResp.Choices) == 0 {
		log.Error("No choices in OpenAI response")
		return "", fmt.Errorf("no choices in OpenAI response")
	}

	log.Info("OpenAI API call with system/user prompts successful")
	return openAIResp.Choices[0].Message.Content, nil
}
