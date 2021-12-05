package cmd

import (
	"log"

	"github.com/jtrotsky/wise-cli/pkg/client"
	"github.com/jtrotsky/wise-cli/pkg/profile"
	"github.com/jtrotsky/wise-cli/pkg/recipient"
	"github.com/spf13/cobra"
)

type recipientCmd struct {
	cmd *cobra.Command

	client client.Client
}

func newRecipientCmd() *recipientCmd {
	rc := &recipientCmd{}

	rc.cmd = &cobra.Command{
		Use:   "recipient",
		Short: "Create and manage the people / bank accounts to send money to",
		Long: `A recipient must be specified as a destination for a transfer to be sent. 
	You may create and manage recipients with this command`,

		RunE: rc.runRecipientCmd,
	}

	rc.cmd.PersistentFlags().StringVar(&rc.client.APIToken, "token", "", "API token")

	// rc.cmd.Flags().Float64Var(&rc.amount, "amount", 0, "The amount to send or convert")
	// rc.cmd.Flags().StringVar(&rc.fromCurrency, "from", "", "The currency to send from (e.g. GBP)")
	// rc.cmd.Flags().StringVar(&rc.toCurrency, "to", "", "The currency to send to (e.g. NZD)")

	return rc
}

func (rc *recipientCmd) runRecipientCmd(cmd *cobra.Command, args []string) error {
	profiles, err := profile.Get(&rc.client)
	if err != nil {
		log.Fatal(err)
	}

	personalProfile, err := profile.GetProfileByType(profiles, profile.EntityPersonal)
	if err != nil {
		log.Fatal(err)
	}
	rc.client.SetProfile(personalProfile.ID, profile.EntityPersonal)

	recipient := recipient.Prepare(rc.client.ProfileID, "GBP", "231470", "28821822")

	err = recipient.Create(&rc.client)
	if err != nil {
		log.Fatal(err)
	}
	return err
}
