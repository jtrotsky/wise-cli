package client

import (
	"flag"
	// "fmt"
	"log"
)

// Config ...
type Config struct {
	APIKey      string
	Debug       bool
	PrivateKey  string
	ProfileID   int64
	ProfileType string
}

// LoadConfig loads the Wise client config and returns it as Config.
func LoadConfig() (*Config, error) {
	config := Config{}

	flag.StringVar(&config.APIKey, "api-key", "", "your secret api key")
	// flag.StringVar(&config.PrivateKey, "private-key", "", "path to your private signing key file")
	flag.Parse()

	if config.APIKey == "" {
		log.Fatal("please specify a secret api key with flag: --api-key")
	}

	// fmt.Println(config.APIKey)

	return &config, nil
}
