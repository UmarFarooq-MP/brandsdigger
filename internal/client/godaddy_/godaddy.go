package godaddy_

import (
	"brandsdigger/internal/config"
	"encoding/json"
	"fmt"
	"sync"

	"github.com/valyala/fasthttp"
)

// DomainResponse represents the API response
type DomainResponse struct {
	Domain    string `json:"domain"`
	Available bool   `json:"available"`
}

// Client holds the reusable HTTP client
type Client struct {
	APIKey    string
	APISecret string
	BaseURL   string
	HTTP      *fasthttp.Client
}

var (
	instance *Client
	once     sync.Once
)

// NewClient initializes the singleton client (thread-safe)
func NewClient(apiKey, apiSecret string) *Client {
	once.Do(func() {
		instance = &Client{
			APIKey:    config.GODADDY_API_KEY,
			APISecret: apiSecret,
			BaseURL:   "https://api.godaddy.com/v1",
			HTTP: &fasthttp.Client{
				MaxConnsPerHost: 100, // Optimize for high concurrency
			},
		}
	})
	return instance
}

// createRequest prepares the HTTP request
func (c *Client) createRequest(method, path string) (*fasthttp.Request, *fasthttp.Response) {
	req := fasthttp.AcquireRequest()
	req.SetRequestURI(fmt.Sprintf("%s%s", c.BaseURL, path))
	req.Header.SetMethod(method)
	req.Header.Set("Authorization", fmt.Sprintf("sso-key %s:%s", c.APIKey, c.APISecret))
	req.Header.Set("Accept", "application/json")

	resp := fasthttp.AcquireResponse()
	return req, resp
}

// CheckDomainAvailability checks domain availability concurrently
func (c *Client) CheckDomainAvailability(domain string) (*DomainResponse, error) {
	req, resp := c.createRequest("GET", fmt.Sprintf("/domains/available?domain=%s&checkType=FAST", domain))
	defer fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(resp)

	// Thread-safe HTTP call
	if err := c.HTTP.Do(req, resp); err != nil {
		return nil, fmt.Errorf("request error: %w", err)
	}

	if resp.StatusCode() != fasthttp.StatusOK {
		return nil, fmt.Errorf("API error: %s", resp.Body())
	}

	var domainResp DomainResponse
	if err := json.Unmarshal(resp.Body(), &domainResp); err != nil {
		return nil, fmt.Errorf("JSON parse error: %w", err)
	}

	return &domainResp, nil
}
