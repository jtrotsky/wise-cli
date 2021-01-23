package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// transferCmd represents the transfer command
var transferCmd = &cobra.Command{
	Use:   "transfer",
	Short: "Send money using a quote and recipient",
	Long: `Once you have a quote and recipient, these can be used to create a transfer.
	And when that transfer is funded (e.g. from one of your Wise balances)
	then the transfer will be completed and the money sent.
	
	For example, to fund a transfer using your balance run:

	$ wise balance fund --transfer 12345678`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("transfer called")
	},
}

func init() {
	rootCmd.AddCommand(transferCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// transferCmd.PersistentFlags().String("foo", "", "A help for foo")

	// local flags
	transferCmd.Flags().Int("quote", 0, "Quote ID to create a transafer from")
}
