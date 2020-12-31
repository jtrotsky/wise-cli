package cmd

import (
	"github.com/jtrotsky/wise-cli/pkg/client"
	"github.com/spf13/cobra"
)

const (
	actionQuoteCreate = "quote.create"
	// actionTransfer = "transfer.create"
	// actionTransfer = "transfer.update"
	actionBalanceConvert = "balance.convert"
	// actionWebhook = "webhook.send"
)

// Command contains arguments and helpers
type Command struct {
	Action string
	client *client.Client
}

// NewCommand reads a user input and creates a commend to be executed.
func NewCommand(name string) *cobra.Command {
	// config := client.LoadConfig()

	command := &cobra.Command{
		Use:   name,
		Short: "play with the Wise API",
		Long:  "play with the Wise API",
		// PersistentPreRun will run before all subcommands EXCEPT in the following conditions:
		//  - a subcommand defines its own PersistentPreRun function
		//  - the command is run without arguments or with --help and only prints the usage info
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			// features.Enable(config.Features()...)
			// features.Enable(cmdFeatures...)
		},
	}

	command.AddCommand(
		newQuoteCmd().cmd,
		// transfer.newTransferCmd(),
	)

	return command
}
