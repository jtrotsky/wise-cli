package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/jtrotsky/wise-cli/pkg/quote"
)

var apiKey string

func init() {
	flag.StringVar(&apiKey, "api-key", "", "your secret api key")
	flag.Parse()
}

func main() {
	if apiKey == "" {
		log.Fatal("please specify a secret api key with flag: --api-key")
	}

	sourceCurrency := "NZD"
	targetCurrency := "AUD"
	targetAmount := 100

	quote, err := quote.Get(sourceCurrency, targetCurrency, targetAmount)
	if err != nil {
		log.Fatalf(err)
	}
	fmt.Println(quote)
}
