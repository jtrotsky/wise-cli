package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

// DefaultAPIBaseURL is the default base URL for API requests
const DefaultAPIBaseURL = "https://api.wise.com"

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
		return nil
	}
	return errors.New("missing profile ID or type")
}

// DoRequest performs the HTTP request
func (client *Client) DoRequest(method, path string, params url.Values) (*http.Response, error) {
	urlJoined := fmt.Sprintf("%s%s", DefaultAPIBaseURL, path)
	url, err := url.Parse(urlJoined)
	if err != nil {
		return nil, err
	}

	var body io.Reader

	switch method {
	case http.MethodGet:
		url.RawQuery = params.Encode()
	case http.MethodPost:
		queryData := map[string]string{}
		for k, v := range params {
			queryData[k] = v[0]
		}
		postBody, err := json.Marshal(queryData)
		if err != nil {
			return nil, err
		}
		fmt.Printf("%s", postBody)

		body = bytes.NewBuffer(postBody)
	}

	request, err := http.NewRequest(method, url.String(), body)
	if err != nil {
		return nil, err
	}

	if method == http.MethodPost {
		request.Header.Add("Content-Type", "application/json")
	}
	request.Header.Set("User-Agent", "wise-cli")
	if client.APIToken == "" {
		return nil, errors.New("no api token set")
	}
	request.Header.Set("Authorization", "Bearer "+client.APIToken)

	if client.httpClient == nil {
		client.httpClient = newHTTPClient()
	}

	fmt.Printf("%s %s\n", method, url.String())

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
	httpTransport := &http.Transport{
		ResponseHeaderTimeout: 30 * time.Second,
		ExpectContinueTimeout: 10 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
	}
	return &http.Client{
		Transport: httpTransport,
	}
}
