package client

import (
	"fmt"
	"net/http"
	"net/url"
	"time"
)

// DefaultAPIBaseURL is the default base URL for API requests
const DefaultAPIBaseURL = "https://api.transferwise.com"

// Client contains the authorisation and config for the Wise account
type Client struct {
	Config     *Config
	httpClient *http.Client
}

// New creates a new instance of Client and loads config
func New(config *Config) *Client {
	return &Client{config, newHTTPClient()}
}

// NewFromConfig creates a new instance of Client and loads config
func NewFromConfig() *Client {
	return &Client{LoadConfig(), newHTTPClient()}
}

// DoRequest performs the HTTP request
func (client *Client) DoRequest(method, path, params string) (*http.Response, error) {
	// TODO: cleanup
	urlJoined := fmt.Sprintf("%s%s", DefaultAPIBaseURL, path)
	url, err := url.Parse(urlJoined)
	if err != nil {
		return nil, err
	}

	url.RawQuery = params

	request, err := http.NewRequest(method, url.String(), nil)
	if err != nil {
		return nil, err
	}

	request.Header.Set("User-Agent", "wise-cli")

	if client.Config.APIKey != "" {
		request.Header.Set("Authorization", "Bearer "+client.Config.APIKey)
	}

	if client.httpClient == nil {
		client.httpClient = newHTTPClient()
	}

	// DEBUG: remove
	fmt.Printf("\n %s %s\n\n", method, url.String())

	response, err := client.httpClient.Do(request)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func newHTTPClient() *http.Client {
	var httpTransport *http.Transport

	httpTransport = &http.Transport{
		ResponseHeaderTimeout: 30 * time.Second,
		ExpectContinueTimeout: 10 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
	}
	return &http.Client{
		Transport: httpTransport,
	}
}
