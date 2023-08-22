package catfacts

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
)

// Client that interacts with Cat Facts
type Client struct {
	BaseURL   *url.URL
	UserAgent string

	httpClient *http.Client
}

// Generate a new Client type
func NewClient(baseURL string, userAgent string) *Client {
	u, err := url.Parse(baseURL)
	if err != nil {
		log.Fatal("Unable to parse URL")
	}
	c := Client{
		BaseURL:    u,
		UserAgent:  userAgent,
		httpClient: http.DefaultClient,
	}

	return &c
}

func (c *Client) newRequest(method, path string, queryValues url.Values, body interface{}) (*http.Request, error) {
	rel := &url.URL{Path: path, RawQuery: queryValues.Encode()}
	u := c.BaseURL.ResolveReference(rel)
	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}
	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", c.UserAgent)
	return req, nil
}
func (c *Client) do(req *http.Request, v interface{}) (*http.Response, error) {
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(v)
	return resp, err
}

// The response type that comes with paged responses
type PagedResponse struct {
	CurrentPage int `json:"current_page"`

	FirstPageURL string `json:"first_page_url"`
	From         int    `json:"from"`
	LastPage     int    `json:"last_page"`
	LastPageURL  string `json:"last_page_url"`
	Links        []struct {
		URL    any    `json:"url"`
		Label  string `json:"label"`
		Active bool   `json:"active"`
	} `json:"links"`
	NextPageURL string `json:"next_page_url"`
	Path        string `json:"path"`
	PerPage     int    `json:"per_page"`
	PrevPageURL any    `json:"prev_page_url"`
	To          int    `json:"to"`
	Total       int    `json:"total"`
}
