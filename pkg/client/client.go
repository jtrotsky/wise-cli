package client

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

// DefaultAPIBaseURL is the default base URL for API requests
const DefaultAPIBaseURL = "https://api.transferwise.com"

// Client contains the authorisation and config for the Wise account
type Client struct {
	APIToken   string
	ProfileID  int64
	httpClient *http.Client
}

// New creates a new instance of Client and loads config
func New(config *Config) *Client {
	return &Client{config.APIToken, 0, newHTTPClient()}
}

// SetProfile accepts a profile ID and type to set as a global configuration.
func (client *Client) SetProfile(profileID int64, profileType string) error {
	if profileID != 0 {
		client.ProfileID = profileID
		// client.Config.ProfileType = profileType
		return nil
	}
	return errors.New("missing profile ID or type")
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

	// TODO: testing
	fmt.Printf("\n%s\n", url.String())

	request, err := http.NewRequest(method, url.String(), nil)
	if err != nil {
		return nil, err
	}

	request.Header.Set("User-Agent", "wise-cli")
	if client.APIToken == "" {
		return nil, errors.New("no api token set")
	}
	request.Header.Set("Authorization", "Bearer "+client.APIToken)

	if client.httpClient == nil {
		client.httpClient = newHTTPClient()
	}

	response, err := client.httpClient.Do(request)
	if err != nil {
		return nil, err
	}

	switch response.StatusCode {
	case 400:
		return nil, fmt.Errorf("HTTP: %d , something failed", response.StatusCode)
	case 401:
		return nil, fmt.Errorf("HTTP: %d , unauthorised check token is valid", response.StatusCode)
	case 500:
		return nil, fmt.Errorf("HTTP: %d , something failed", response.StatusCode)
	}

	// TODO: checkErrors
	// {"timestamp":"2021-01-02T12:57:07.659+0000","status":400,"error":"Bad Request","message":"Bad Request","path":"/v1/rates/"}

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
