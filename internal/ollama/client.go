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

const (
	defaultTimeout = 60 * time.Second
)

type Client struct {
	endpoint   string
	model      string
	httpClient *http.Client
}

type Options struct {
	Endpoint string

	Model string

	Timeout time.Duration
}

func New(
	options Options,
) *Client {

	return &Client{
		endpoint: options.Endpoint,
		model:    options.Model,
		httpClient: &http.Client{
			Timeout: options.Timeout,
		},
	}
}

func (c *Client) Embedding(
	ctx context.Context,
	request EmbeddingRequest,
) ([]float32, error) {

	body, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		c.endpoint+"/api/embeddings",
		bytes.NewReader(body),
	)
	if err != nil {
		return nil, err
	}

	req.Header.Set(
		"Content-Type",
		"application/json",
	)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {

		var apiErr ErrorResponse

		if json.Unmarshal(data, &apiErr) == nil &&
			apiErr.Message != "" {

			return nil, apiErr
		}

		return nil, fmt.Errorf(
			"ollama returned status %d",
			resp.StatusCode,
		)
	}

	var embedding EmbeddingResponse

	if err := json.Unmarshal(
		data,
		&embedding,
	); err != nil {

		return nil, err
	}

	return embedding.Embedding, nil
}
