package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// balanceCmd represents the balance command
var balanceCmd = &cobra.Command{
	Use:   "balance",
	Short: "Manage multi-currency balances",
	Long: `The balance command can be used to create and manage balances in multiple currencies.

You may also convert money between balances. For example:

$ wise balance convert --amount 100 --from GBP --to NZD`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("balance called")
	},
}

func init() {
	rootCmd.AddCommand(balanceCmd)

	balanceCmd.Flags().Int32("amount", 0, "The amount to send or convert")
	balanceCmd.Flags().String("from", "", "The currency to send from (e.g. GBP)")
	balanceCmd.Flags().String("to", "", "The currency to send to (e.g. NZD)")
}
