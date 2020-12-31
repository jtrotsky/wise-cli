package client

import (
	"flag"
	"log"
)

// WiseConfig is a map of strings to interface{}
type WiseConfig struct {
	APIKey     string
	PrivateKey string
	ProfileID  string
}

// LoadConfig loads the Wise client config and returns it as a WiseConfig.
func LoadConfig() *WiseConfig {
	config := WiseConfig{}

	flag.StringVar(&config.APIKey, "api-key", "", "your secret api key")
	flag.StringVar(&config.PrivateKey, "private-key", "", "path to your private signing key file")
	flag.StringVar(&config.ProfileID, "profile", "", "your wise profile id")
	flag.Parse()

	if config.APIKey == "" {
		log.Fatal("please specify a secret api key with flag: --api-key")
	}
	if config.ProfileID == "" {
		log.Fatal("please specify a wise profile id flag: --profile")
	}

	return &config
}
