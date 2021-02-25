package cmd

import (
	"log"

	"github.com/jtrotsky/wise-cli/pkg/balance"
	"github.com/jtrotsky/wise-cli/pkg/client"
	"github.com/spf13/cobra"
)

type balanceCmd struct {
	cmd *cobra.Command

	client    client.Client
	profileID int64
	currency  string
}

func newBalanceCmd() *balanceCmd {
	bc := &balanceCmd{}

	bc.cmd = &cobra.Command{
		Use:   "balance",
		Short: "Manage multi-currency balances",
		Long: `The balance command can be used to create and manage balances in multiple currencies.

You may also convert money between balances. For example:

$ wise balance convert --amount 100 --from GBP --to NZD`,
		RunE: bc.runBalanceCmd,
	}

	bc.cmd.PersistentFlags().StringVar(&bc.client.APIToken, "token", "", "API token")
	bc.cmd.Flags().StringVar(&bc.currency, "currency", "GBP", "The account currency (e.g. GBP)")

	return bc
}

func (bc *balanceCmd) runBalanceCmd(cmd *cobra.Command, args []string) error {
	_, err := balance.Get(&bc.client, bc.currency)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}
