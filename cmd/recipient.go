package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// recipientCmd represents the recipient command
var recipientCmd = &cobra.Command{
	Use:   "recipient",
	Short: "Create and manage the people / bank accounts to send money to",
	Long: `A recipient must be specified as a destination for a transfer to be sent. 
You may create and manage recipients with this command`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("recipient called")
	},
}

func init() {
	rootCmd.AddCommand(recipientCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// recipientCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// recipientCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
