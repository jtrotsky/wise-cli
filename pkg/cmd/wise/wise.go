package wise

// import (
// 	"fmt"
// 	"log"

// 	"github.com/jtrotsky/wise-cli/pkg/client"
// 	"github.com/jtrotsky/wise-cli/pkg/profile"
// 	"github.com/jtrotsky/wise-cli/pkg/quote"
// 	"github.com/jtrotsky/wise-cli/pkg/transfer"
// 	"github.com/spf13/cobra"
// 	"github.com/spf13/viper"
// )

// const (
// 	actionQuoteCreate = "quote.create"
// 	// actionTransfer = "transfer.create"
// 	// actionTransfer = "transfer.update"
// 	actionBalanceConvert = "balance.convert"
// )

// // Config ...
// var Config client.Config

// // NewCommand reads a user input and creates a commend to be executed.
// func NewCommand(name string) *cobra.Command {
// 	// config, err := client.LoadConfig()
// 	// if err != nil {
// 	// log.Fatal(err)
// 	// }

// 	c := &cobra.Command{
// 		Use:   name,
// 		Short: "A tool to play with Wise APIs",
// 		Long: `Use wise-cli to play with the Wise API and make understanding the Wise APIs and
// 			integrating to them easier.`,
// 		// PersistentPreRun will run before all subcommands EXCEPT in the following conditions:
// 		//  - a subcommand defines its own PersistentPreRun function
// 		//  - the command is run without arguments or with --help and only prints the usage info
// 		// PersistentPreRun: func(cmd *cobra.Command, args []string) {
// 		// features.Enable(config.Features()...)
// 		// features.Enable(cmdFeatures...)
// 		// },
// 	}

// 	c.PersistentFlags().StringVar(&Config.APIKey, "api-key", "", "your secret API key")
// 	viper.BindPFlag(Config.APIKey, c.PersistentFlags().Lookup(Config.APIKey))
// 	fmt.Println(Config.APIKey)
// 	client := client.New(&Config)

// 	// fmt.Println(client.APIKey)

// 	// Get and then set the user's profile (business or personal)
// 	allProfiles, err := profile.Get(client)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// TODO: test
// 	fmt.Println(allProfiles)

// 	// TODO: user input to choose profile
// 	personalProfile, err := profile.GetProfileByType(allProfiles, profile.EntityPersonal)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	client.SetProfile(personalProfile.ID, profile.EntityPersonal)

// 	c.AddCommand(
// 		quote.NewCommand(*client),
// 		transfer.NewCommand(*client),
// 		// setup.NewCommand(client),
// 	)

// 	return c
// }
