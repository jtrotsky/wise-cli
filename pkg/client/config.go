package client

import (
	"flag"
	"log"
)

// Config contains all the details needed to make API calls for a Wise account.
type Config struct {
	APIToken    string
	PrivateKey  string
	ProfileID   int64
	ProfileType string
}

// LoadConfig loads the Wise client config and returns it as Config.
func LoadConfig() (*Config, error) {
	config := Config{}

	flag.StringVar(&config.APIToken, "token", "", "your secret api token")
	flag.StringVar(&config.PrivateKey, "key", "", "path to your private signing key file")
	flag.Parse()

	if config.APIToken == "" {
		log.Fatal("please specify a secret api token with flag: --token")
	}

	return &config, nil
}
