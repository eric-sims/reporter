// Package openai connects to remote llm
package openai

import (
	"context"
	"fmt"

	"github.com/openai/openai-go/v2"
	"github.com/openai/openai-go/v2/option"
	"github.com/openai/openai-go/v2/responses"
)

// OpenAI - LLM client to interact with
type OpenAI struct {
	client openai.Client
}

// NewClient - returns a new client
func NewClient(apiKey string) *OpenAI {
	return &OpenAI{
		client: openai.NewClient(option.WithAPIKey(apiKey)),
	}
}

// Generate - generates a response from openai
func (o *OpenAI) Generate(ctx context.Context, prompt string) (string, error) {
	resp, err := o.client.Responses.New(ctx, responses.ResponseNewParams{
		Input: responses.ResponseNewParamsInputUnion{OfString: openai.String(prompt)},
		Model: openai.ChatModelGPT5,
	})

	if resp == nil || err != nil {
		return "", fmt.Errorf("openai: failed to generate request: %w", err)
	}

	return resp.OutputText(), nil
}
