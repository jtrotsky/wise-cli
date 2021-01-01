package quote

import (
	"github.com/jtrotsky/wise-cli/pkg/client"
	"github.com/spf13/cobra"
)

func NewCommand(client client.Client) *cobra.Command {
	c := &cobra.Command{
		Use:   "quote",
		Short: "Work with quotes",
		Long:  "Work with quotes",
	}

	c.AddCommand(
		NewCreateCommand(client, "create"),
		// NewGetCommand(f, "get"),
	)

	return c
}
