package transfer

import (
	"github.com/jtrotsky/wise-cli/pkg/client"
	"github.com/spf13/cobra"
)

func NewCommand(f client.Factory) *cobra.Command {
	c := &cobra.Command{
		Use:   "transfer",
		Short: "Work with transfers",
		Long:  "Work with transfers",
	}

	c.AddCommand(
	// NewCreateCommand(f, "create"),
	// NewGetCommand(f, "get"),
	)

	return c
}
