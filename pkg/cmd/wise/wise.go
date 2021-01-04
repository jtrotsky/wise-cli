package wise

import (
	"github.com/jtrotsky/wise-cli/pkg/client"
	"github.com/jtrotsky/wise-cli/pkg/quote"
	"github.com/jtrotsky/wise-cli/pkg/transfer"
	"github.com/spf13/cobra"
)

const (
	actionQuoteCreate = "quote.create"
	// actionTransfer = "transfer.create"
	// actionTransfer = "transfer.update"
	actionBalanceConvert = "balance.convert"
)

// Config is the cli configuration for the user
var Config client.Config

// NewCommand reads a user input and creates a commend to be executed.
func NewCommand(name string) *cobra.Command {
	c := &cobra.Command{
		Use:   name,
		Short: "A tool to play with Wise APIs",
		Long: `Use wise-cli to play with the Wise API and make understanding the Wise APIs and 
			integrating to them easier.`,
		// PersistentPreRun will run before all subcommands EXCEPT in the following conditions:
		//  - a subcommand defines its own PersistentPreRun function
		//  - the command is run without arguments or with --help and only prints the usage info
		// PersistentPreRun: func(cmd *cobra.Command, args []string) {
		// features.Enable(config.Features()...)
		// features.Enable(cmdFeatures...)
		// },
	}

	c.PersistentFlags().StringVar(&Config.APIKey, "api-key", "", "Your secret API key ")
	client := client.New(&Config)

	c.AddCommand(
		quote.NewCommand(*client),
		transfer.NewCommand(*client),
		// setup.NewCommand(client),
	)

	return c
}
