package quote

import (
	"errors"
	"fmt"
	"log"

	"github.com/jtrotsky/wise-cli/pkg/client"
	"github.com/spf13/cobra"
)

// Quote is used to create a transfer https://api-docs.transferwise.com/#quotes-create
type Quote struct {
	ID             string
	SourceCurrency string
	TargetCurrency string
	Amount         int
	Rate           float64
}

// NewCreateCommand creates a new quote
func NewCreateCommand(client *client.Client, action string) *cobra.Command {
	o := NewCreateOptions()

	c := &cobra.Command{
		Use:     action,
		Args:    cobra.NoArgs,
		Short:   "List quote commands",
		Example: "wise quote create --from GBP --amount 100 --to NZD",
		Run: func(c *cobra.Command, args []string) {
			o.Run(c, client)
		},
	}

	c.Flags().StringVar(&o.fromCurrency, "from", "", "The source / origin currency for the quote")
	c.Flags().StringVar(&o.toCurrency, "to", "", "The target / destination currency for the quote")
	c.Flags().IntVar(&o.amount, "amount", 0, "The amount of to / source currency to be sent")

	return c
}

// CreateOptions ...
type CreateOptions struct {
	fromCurrency string
	toCurrency   string
	amount       int

	client client.Client
}

// NewCreateOptions ...
func NewCreateOptions() *CreateOptions {
	return &CreateOptions{}
}

// Run ...
func (o *CreateOptions) Run(cmd *cobra.Command, client *client.Client) error {
	if o.fromCurrency == "" {
		log.Fatal("source / from currency is needed to create a quote")
	}

	if o.toCurrency == "" {
		log.Fatal("destination / to currency is needed to create a quote")
	}

	if o.amount >= 0 {
		log.Fatal("amount is needed to create a quote")
	}

	quote := prepare(o.fromCurrency, o.toCurrency, o.amount)
	err := quote.create(client)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(quote)
	return nil
}

// Prepare a quote to be sent to Wise
func prepare(fromCurrency, toCurrency string, amount int) *Quote {
	return &Quote{
		SourceCurrency: fromCurrency,
		TargetCurrency: toCurrency,
		Amount:         amount,
	}
}

// Create a new quote from Wise based on currency pair and amount provided
func (q *Quote) create(client *client.Client) error {

	q.Rate = 0.91435
	result := fmt.Sprintf("%s to %s with rate of %f and amount of %d", q.SourceCurrency, q.TargetCurrency, q.Rate, q.Amount)

	fmt.Println(result)

	return errors.New("fail")
}
