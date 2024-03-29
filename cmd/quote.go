package cmd

import (
	"log"

	"github.com/jtrotsky/wise-cli/pkg/client"
	"github.com/jtrotsky/wise-cli/pkg/profile"
	"github.com/jtrotsky/wise-cli/pkg/quote"
	"github.com/spf13/cobra"
)

type quoteCmd struct {
	cmd *cobra.Command

	amount float64
	//apiToken     string
	client       client.Client
	fromCurrency string
	//profileID    int64
	toCurrency string
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
	$ wise quote --amount 100 --from GBP --to NZD --token <TOKEN>
	
	// use the quote to create a transfer 
	$ wise transfer --quote 12345678`,
		RunE: qc.runQuoteCmd,
	}

	qc.cmd.PersistentFlags().StringVar(&qc.client.APIToken, "token", "", "API token")

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

	profiles, err := profile.Get(&qc.client)
	if err != nil {
		log.Fatal(err)
	}

	personalProfile, err := profile.GetProfileByType(profiles, profile.EntityPersonal)
	if err != nil {
		log.Fatal(err)
	}
	qc.client.SetProfile(personalProfile.ID, profile.EntityPersonal)

	quote := quote.Prepare(qc.client.ProfileID, qc.fromCurrency, qc.toCurrency, qc.amount)
	err = quote.Create(&qc.client)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}
