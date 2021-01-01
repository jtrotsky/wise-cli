package client

// Client contains the authorisation and config for the Wise account
type Client struct {
	Config *WiseConfig
}

// New creates a new instance of Client and loads config
func New() *Client {
	return &Client{LoadConfig()}
}
