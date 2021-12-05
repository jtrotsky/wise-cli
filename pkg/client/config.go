package client

import (
	"flag"
	"log"
)

// Config contains account and authentication details.
type Config struct {
	APIToken    string
	PrivateKey  string
	ProfileID   int64
	ProfileType string
}

// LoadConfig loads the Wise client config and returns it as Config.
func LoadConfig() (*Config, error) {
	config := Config{}

	flag.StringVar(&config.APIToken, "api-token", "", "your secret api token")
	// flag.StringVar(&config.PrivateKey, "private-key", "", "path to your private signing key file")
	flag.Parse()

	if config.APIToken == "" {
		log.Fatal("please specify a secret api token with flag: --token")
	}

	return &config, nil
}
