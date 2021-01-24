package client

import (
	"bytes"
	"encoding/json"
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
	httpClient *http.Client

	APIKey    string
	Debug     bool
	ProfileID int64
}

// New creates a new instance of Client and loads config
// func New(config *Config) *Client {
// return &Client{config.APIKey, 0, newHTTPClient()}
// }

// SetProfile accepts a profile ID and type to set as a global configuration.
func (client *Client) SetProfile(profileID int64, profileType string) error {
	if profileID != 0 {
		client.ProfileID = profileID
		return nil
	}
	return errors.New("missing profile ID or type")
}

// DoRequest performs the HTTP request
func (client *Client) DoRequest(method, path string, params url.Values) (*http.Response, error) {
	// TODO: cleanup
	urlJoined := fmt.Sprintf("%s%s", DefaultAPIBaseURL, path)
	url, err := url.Parse(urlJoined)
	if err != nil {
		return nil, err
	}

	request := &http.Request{}

	// if get add params
	if method == http.MethodGet {
		url.RawQuery = params.Encode()
		request, err = http.NewRequest(method, url.String(), nil)
		if err != nil {
			return nil, err
		}
	}

	if method == http.MethodPost {
		// print URL requested when debugging
		body, err := json.Marshal(params)

		if client.Debug {
			fmt.Printf("\n%s\n", body)
		}

		request, err = http.NewRequest(method, url.String(), bytes.NewBuffer(body))
		if err != nil {
			return nil, err
		}
	}

	if client.Debug {
		fmt.Printf("\n%s %s\n", method, url.String())
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("User-Agent", "wise-cli")
	if client.APIKey == "" {
		return nil, errors.New("no api key set")
	}
	request.Header.Set("Authorization", "Bearer "+client.APIKey)

	if client.httpClient == nil {
		client.httpClient = newHTTPClient()
	}

	response, err := client.httpClient.Do(request)
	if err != nil {
		return nil, err
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
