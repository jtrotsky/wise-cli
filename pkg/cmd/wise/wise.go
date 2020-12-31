package wise

import (
	"fmt"
	"log"

	"github.com/jtrotsky/wise-cli/pkg/client"
	"github.com/jtrotsky/wise-cli/pkg/quote"
)

const Action (
	actionQuote = "quote"
	// actionTransfer = "transfer"
	// actionWebhook = "transfer"
)

// Command contains arguments and helpers
type Command struct {
	Action Action
	Client client.Client
}

// NewCommand uses args to preare a commend to be executed.
func NewCommand(action string) *Command {
	command := Command{}

	switch action {
	case action == actionQuote:
		command.Action = ActionQuote
		fmt.Println("command: quote")
	default:
		log.Fatal("no arg")
	}

	sourceCurrency := "NZD"
	targetCurrency := "AUD"
	targetAmount := 100.0000

	quote, err := quote.New(sourceCurrency, targetCurrency, targetAmount)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(quote)
}

// Execute performs the Command specified
func (c *Command) Execute() error {
	return client.Run(c)
}
