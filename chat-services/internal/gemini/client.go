package gemini

import (
	"context"
	"os"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

type Client struct {
	genaiClient *genai.Client
	model       *genai.GenerativeModel
}

// NewClient membuat instance baru Gemini Client
func NewClient() (*Client, error) {
	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(os.Getenv("GEMINI_API_KEY_EKSPOR")))
	if err != nil {
		return nil, err
	}

	return &Client{
		genaiClient: client,
		model:       client.GenerativeModel("gemini-1.5-flash"),
	}, nil
}

func (c *Client) GenerateContent(ctx context.Context, prompt string) (string, error) {
	resp, err := c.model.GenerateContent(ctx, genai.Text(prompt))
	if err != nil {
		return "", err
	}

	if len(resp.Candidates) > 0 && len(resp.Candidates[0].Content.Parts) > 0 {
		if text, ok := resp.Candidates[0].Content.Parts[0].(genai.Text); ok {
			return string(text), nil
		}
	}

	return "", nil
}

func (c *Client) Close() {
	if c.genaiClient != nil {
		c.genaiClient.Close()
	}
}
