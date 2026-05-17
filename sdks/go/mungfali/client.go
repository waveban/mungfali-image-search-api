// Package mungfali provides a lightweight client for the Mungfali Image Search API.
package mungfali

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

const baseURL = "https://mungfali.net/v1/search"

// Client calls the Mungfali Image Search API.
type Client struct {
	APIKey     string
	HTTPClient *http.Client
}

// NewClient returns a client with a default HTTP timeout.
func NewClient(apiKey string) *Client {
	return &Client{
		APIKey: apiKey,
		HTTPClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

// SearchOptions configures a search request.
type SearchOptions struct {
	Page               int
	PerPage            int
	SafeSearch         string
	FilterTransparent  *bool
}

// ImageResult is a single image in the response.
type ImageResult struct {
	Name            string `json:"name"`
	FamilyFriendly  bool   `json:"FamilyFriendly"`
	ImageURL        string `json:"imageUrl"`
	HostURL         string `json:"hostUrl"`
	ContentSize     string `json:"contentSize"`
	Width           int    `json:"width"`
	Height          int    `json:"height"`
	IsTransparent   bool   `json:"isTransparent"`
	AccentColor     string `json:"accentColor"`
}

// SearchResponse is the API JSON payload.
type SearchResponse struct {
	Name    string        `json:"name"`
	Value   []ImageResult `json:"value"`
	Page    int           `json:"page"`
	PerPage int           `json:"per_page"`
	Total   int           `json:"total"`
}

// Search executes GET /v1/search.
func (c *Client) Search(ctx context.Context, query string, opts SearchOptions) (*SearchResponse, error) {
	u, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}

	q := u.Query()
	q.Set("q", query)
	if opts.Page > 0 {
		q.Set("page", strconv.Itoa(opts.Page))
	}
	if opts.PerPage > 0 {
		q.Set("per_page", strconv.Itoa(opts.PerPage))
	}
	if opts.SafeSearch != "" {
		q.Set("safeSearch", opts.SafeSearch)
	}
	if opts.FilterTransparent != nil {
		q.Set("filterTransparent", strconv.FormatBool(*opts.FilterTransparent))
	}
	u.RawQuery = q.Encode()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+c.APIKey)
	req.Header.Set("Accept", "application/json")

	client := c.HTTPClient
	if client == nil {
		client = http.DefaultClient
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	if res.StatusCode < 200 || res.StatusCode >= 300 {
		return nil, fmt.Errorf("API error %d: %s", res.StatusCode, string(body))
	}

	var data SearchResponse
	if err := json.Unmarshal(body, &data); err != nil {
		return nil, err
	}
	return &data, nil
}
