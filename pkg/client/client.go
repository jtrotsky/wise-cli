package client

import (
	"fmt"
	"log"

	"github.com/jtrotsky/wise-cli/pkg/cmd/wise"
	"github.com/jtrotsky/wise-cli/pkg/quote"
)

// Client contains the authorisation and config for the Wise account
type Client struct {
	Config *WiseConfig
}

// New creates a new instance of Client and loads config
func New() Client {
	return Client{LoadConfig()}
}

// Run uses the right factory to complete the action requested
func (c *Client) Run(command *wise.Command) error {
	sourceCurrency := "NZD"
	targetCurrency := "AUD"
	targetAmount := 100

	quote, err := quote.Prepare(sourceCurrency, targetCurrency, targetAmount)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(quote)
}
