package cmd

import (
	"log"

	"github.com/jtrotsky/wise-cli/pkg/quote"
	"github.com/spf13/cobra"
)

type quoteCmd struct {
	cmd *cobra.Command
}

func newQuoteCmd() *quoteCmd {
	qc := &quoteCmd{}

	qc.cmd = &cobra.Command{
		Use: "quote",
		// Args:  validators.NoArgs,
		Short: "List quote commands",
	}
	// qc.cmd.SetHelpTemplate(getResourcesHelpTemplate())

	return qc
}

func (lc *quoteCmd) runQuoteCmd(cmd *cobra.Command, args []string) error {
	q := quote.Prepare("NZD", "AUD", 100)
	err := q.Create()
	if err != nil {
		log.Fatal(err)
	}
	return nil
}
