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

	sourceCurrency string
	targetCurrency string
	targetAmount   int
}

// LoadConfig loads the Wise client config and returns it as a WiseConfig.
func LoadConfig() *WiseConfig {
	config := WiseConfig{}

	flag.StringVar(&config.APIKey, "api-key", "", "your secret api key")
	flag.StringVar(&config.PrivateKey, "private-key", "", "path to your private signing key file")
	flag.StringVar(&config.ProfileID, "profile", "", "your wise profile id")

	flag.StringVar(&config.sourceCurrency, "from", "", "send from currency")
	flag.StringVar(&config.targetCurrency, "to", "", "send to currency")
	flag.IntVar(&config.targetAmount, "amount", 0, "amount to arrive")
	flag.Parse()

	if config.APIKey == "" {
		log.Fatal("please specify a secret api key with flag: --api-key")
	}
	if config.ProfileID == "" {
		log.Fatal("please specify a wise profile id flag: --profile")
	}
	if config.sourceCurrency == "" {
		log.Fatal("please specify a source currency")
	}
	if config.targetCurrency == "" {
		log.Fatal("please specify a target currency")
	}
	if config.targetAmount >= 0 {
		log.Fatal("please specify an amount using flat: amount")
	}

	return &config
}
