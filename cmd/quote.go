package cmd

import (
	"log"

	"github.com/jtrotsky/wise-cli/pkg/client"
	"github.com/jtrotsky/wise-cli/pkg/quote"
	"github.com/spf13/cobra"
)

type quoteCmd struct {
	cmd *cobra.Command

	amount       float64
	apiToken     string
	client       client.Client
	fromCurrency string
	profileID    int64
	toCurrency   string
}

func newQuoteCmd() *quoteCmd {
	qc := &quoteCmd{}

	qc.cmd = &cobra.Command{
		Use:   "quote",
		Short: "Get a quote for sending an amount between currencies",
		Long: `A quote provides a summary of the cost and exchange rate to send money. 
	Quotes can also hold a fixed exchange rate for a period of time (e.g. 24hrs).
	
	The quote reference should also be used as an input to create a transfer. For example:
	
	// create the quote
	$ wise quote create --amount 100 --from GBP --to NZD
	
	// use the quote to create a transfer 
	$ wise transfer create --quote 12345678`,
		RunE: qc.runQuoteCmd,
	}

	qc.cmd.PersistentFlags().StringVar(&qc.client.APIKey, "token", "", "API token")
	qc.cmd.PersistentFlags().BoolVar(&qc.client.Debug, "debug", false, "Toggle to debug")

	qc.cmd.Flags().Float64Var(&qc.amount, "amount", 0, "The amount to send or convert")
	qc.cmd.Flags().StringVar(&qc.fromCurrency, "from", "", "The currency to send from (e.g. GBP)")
	qc.cmd.Flags().StringVar(&qc.toCurrency, "to", "", "The currency to send to (e.g. NZD)")

	return qc
}

func (qc *quoteCmd) runQuoteCmd(cmd *cobra.Command, args []string) error {
	if qc.fromCurrency == "" {
		log.Fatal("from currency is needed to create a quote")
	}

	if qc.toCurrency == "" {
		log.Fatal("to currency is needed to create a quote")
	}

	if qc.amount <= 0 {
		log.Fatal("amount greater than 0 is needed to create a quote")
	}

	qc.profileID = qc.client.ProfileID

	quote := quote.Prepare(qc.profileID, qc.fromCurrency, qc.toCurrency, qc.amount)
	err := quote.Create(&qc.client)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}
