// Package ollama is the interface for interacting with the ollama api
package ollama

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// Client - LLM client to interact with
type Client struct {
	base string
	c    *http.Client
}

// NewClient - returns a new client
func NewClient(base string) *Client {
	return &Client{base: base, c: &http.Client{Timeout: 120 * time.Second}}
}

type generateReq struct {
	Model  string `json:"model"`
	Prompt string `json:"prompt"`
	Stream bool   `json:"stream"`
}

type generateResp struct {
	Response string `json:"response"`
}

// Generate calls Ollama's /api/generate with stream=false and returns the final text response.
func (cl *Client) Generate(ctx context.Context, model, prompt string) (string, error) {
	body, _ := json.Marshal(generateReq{Model: model, Prompt: prompt, Stream: false})
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, cl.base+"/api/generate", bytes.NewReader(body))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := cl.c.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		b, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("ollama: %s: %s", resp.Status, string(b))
	}
	var gr generateResp
	if err := json.NewDecoder(resp.Body).Decode(&gr); err != nil {
		return "", err
	}
	return gr.Response, nil
}
