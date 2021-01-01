package client

import (
	"flag"
	"log"
)

// Config is a map of strings to interface{}
type Config struct {
	APIKey     string
	PrivateKey string
	profileID  string
}

// LoadConfig loads the Wise client config and returns it as a WiseConfig.
func LoadConfig() *Config {
	config := Config{}

	flag.StringVar(&config.APIKey, "api-key", "", "your secret api key")
	flag.StringVar(&config.PrivateKey, "private-key", "", "path to your private signing key file")
	flag.Parse()

	if config.APIKey == "" {
		log.Fatal("please specify a secret api key with flag: --api-key")
	}

	return &config
}
