package cmd

import (
	"errors"
	"log"
	"strings"

	"github.com/jtrotsky/wise-cli/pkg/balance"
	"github.com/jtrotsky/wise-cli/pkg/client"
	"github.com/jtrotsky/wise-cli/pkg/util"
	"github.com/spf13/cobra"
)

type balanceCmd struct {
	cmd *cobra.Command

	client   client.Client
	currency string
}

func newBalanceCmd() *balanceCmd {
	bc := &balanceCmd{}

	bc.cmd = &cobra.Command{
		Use:   "balance",
		Short: "Manage multi-currency balances",
		Long: `The balance command can be used to check an account balance.

For example:

$ wise-cli balance --currency GBP`,
		RunE: bc.runBalanceCmd,
	}

	bc.cmd.PersistentFlags().StringVar(&bc.client.APIToken, "token", "", "API token")
	bc.cmd.Flags().StringVar(&bc.currency, "currency", "", "The account currency (e.g. GBP)")

	return bc
}

func (bc *balanceCmd) runBalanceCmd(cmd *cobra.Command, args []string) error {
	if bc.currency != "" && !util.ValidCurrencyCode(bc.currency) {
		return errors.New("invalid currency code, format like: GBP")
	}

	_, err := balance.Get(&bc.client, strings.ToUpper(bc.currency))
	if err != nil {
		log.Fatal(err)
	}

	return err
}

// Long: `The balance command can be used to create and manage balances in multiple currencies.
// You may also convert money between balances. For example:
// $ wise balance --currency GBP`,
