package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// webhookCmd represents the webhook command
var webhookCmd = &cobra.Command{
	Use:   "webhook",
	Short: "Create and listen to webhooks",
	Long: `Use webhooks to listen for updates to the Wise account.
	This command can also be useful to see examples of webhook payloads to help with development.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("webhook called")
	},
}

func init() {
	rootCmd.AddCommand(webhookCmd)
}
