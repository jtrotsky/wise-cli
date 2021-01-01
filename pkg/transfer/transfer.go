package transfer

import (
	"github.com/jtrotsky/wise-cli/pkg/client"
	"github.com/spf13/cobra"
)

// NewCommand ...
func NewCommand(client client.Client) *cobra.Command {
	c := &cobra.Command{
		Use:   "transfer",
		Short: "Work with transfers",
		Long:  "Work with transfers",
	}

	c.AddCommand(
		NewCreateCommand(&client, "create"),
	// NewGetCommand(f, "get"),
	)

	return c
}

// NewCreateCommand creates a new quote
func NewCreateCommand(client *client.Client, action string) *cobra.Command {
	o := NewCreateOptions()

	c := &cobra.Command{
		Use:     action,
		Args:    cobra.NoArgs,
		Short:   "Create a transfer",
		Example: "wise transfer create --quote-id 325346345",
		// Run: func(c *cobra.Command, args []string) {
		// 	o.Run(c, client)
		// },
	}

	c.Flags().StringVar(&o.quoteID, "quote-id", "", "The quote to create a transfer from")

	return c
}

// CreateOptions ...
type CreateOptions struct {
	quoteID string

	client client.Client
}

// NewCreateOptions ...
func NewCreateOptions() *CreateOptions {
	return &CreateOptions{}
}
