package quote

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/jtrotsky/wise-cli/pkg/client"
	"github.com/spf13/cobra"
)

// Quote is used to create a transfer https://api-docs.transferwise.com/#quotes-create
type Quote struct {
	// ID                     string
	// ProfileID              string
	SourceCurrency         string     `json:"source,omitempty"`
	TargetCurrency         string     `json:"target,omitempty"`
	SourceAmount           float64    `json:"sourceAmount,omitempty"`
	TargetAmount           float64    `json:"targetAmount,omitempty"`
	Rate                   float64    `json:"rate,omitempty"`
	Type                   string     `json:"type,omitempty"`
	RateType               string     `json:"rateType,omitempty"`
	CreatedTime            time.Time  `json:"createdTime,omitempty"`
	DeliveryEstimate       time.Time  `json:"deliveryEstimate,omitempty"`
	Fee                    float64    `json:"fee,omitempty"`
	FeeDetails             FeeDetails `json:"feeDetails,omitempty"`
	AllowedProfileTypes    []string   `json:"allowedProfileTypes,omitempty"`
	GuaranteedTargetAmount bool       `json:"GuaranteedTargetAmount,omitempty"`
	OfSourceAmount         bool       `json:"OfSourceAmount,omitempty"`
}

// FeeDetails contains the breakdown of fee components
type FeeDetails struct {
	TransferWise float64 `json:"transferwise,omitempty"`
	PayIn        float64 `json:"payIn,omitempty"`
	Discount     float64 `json:"discount,omitempty"`
	PriceSetID   int     `json:"priceSetId,omitempty"`
	Partner      float64 `json:"partner,omitempty"`
}

// {
//   "source": "GBP",
//   "target": "NZD",
//   "sourceAmount": 100,
//   "targetAmount": 188.26,
//   "type": "REGULAR",
//   "rate": 1.89929,
//   "createdTime": "2021-01-01T21:28:17.192Z",
//   "rateType": "FIXED",
//   "deliveryEstimate": "2021-01-04T23:15:00.000Z",
//   "fee": 0.88,
//   "feeDetails": {
//     "transferwise": 0.88,
//     "payIn": 0,
//     "discount": 0,
//     "priceSetId": 146,
//     "partner": 0
//   },
//   "allowedProfileTypes": [
//     "PERSONAL",
//     "BUSINESS"
//   ],
//   "guaranteedTargetAmount": false,
//   "ofSourceAmount": true
// }

// NewCommand ...
func NewCommand(client client.Client) *cobra.Command {
	c := &cobra.Command{
		Use:   "quote",
		Short: "Work with quotes",
		Long:  "Work with quotes",
	}

	c.AddCommand(
		NewCreateCommand(&client, "create"),
		// NewGetCommand(f, "get"),
		// NewDescribeCommand(f, "describe"),
		// NewDeleteCommand(f, "delete"),
	)

	return c
}

// NewCreateCommand creates a new quote
func NewCreateCommand(client *client.Client, action string) *cobra.Command {
	o := NewCreateOptions()

	c := &cobra.Command{
		Use:     action,
		Args:    cobra.NoArgs,
		Short:   "Create a quote",
		Example: "wise quote create --from GBP --amount 100 --to NZD",
		Run: func(c *cobra.Command, args []string) {
			o.Run(c, client)
		},
	}

	c.Flags().StringVar(&o.fromCurrency, "from", "", "The source / origin currency for the quote")
	c.Flags().StringVar(&o.toCurrency, "to", "", "The target / destination currency for the quote")
	c.Flags().Float64Var(&o.sourceAmount, "amount", 0, "The amount of to / source currency to be sent")

	return c
}

// CreateOptions ...
type CreateOptions struct {
	fromCurrency string
	toCurrency   string
	sourceAmount float64

	client client.Client
}

// NewCreateOptions ...
func NewCreateOptions() *CreateOptions {
	return &CreateOptions{}
}

// Run ...
func (o *CreateOptions) Run(cmd *cobra.Command, client *client.Client) error {
	if o.fromCurrency == "" {
		log.Fatal("source / from currency is needed to create a quote")
	}

	if o.toCurrency == "" {
		log.Fatal("destination / to currency is needed to create a quote")
	}

	if o.sourceAmount <= 0 {
		log.Fatal("amount greater than 0 is needed to create a quote")
	}

	quote := prepare(o.fromCurrency, o.toCurrency, o.sourceAmount)
	err := quote.create(client)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}

// Prepare a quote to be sent to Wise
func prepare(fromCurrency, toCurrency string, sourceAmount float64) *Quote {
	return &Quote{
		SourceCurrency: fromCurrency,
		TargetCurrency: toCurrency,
		SourceAmount:   sourceAmount,
		RateType:       "FIXED", // TODO
	}
}

// Create a new quote from Wise based on currency pair and amount provided
func (q *Quote) create(client *client.Client) error {
	query := url.Values{}
	query.Add("source", fmt.Sprintf("%s", q.SourceCurrency))
	query.Add("target", fmt.Sprintf("%s", q.TargetCurrency))
	query.Add("sourceAmount", fmt.Sprintf("%f", q.SourceAmount))
	query.Add("rateType", fmt.Sprintf("%s", q.RateType))

	response, err := client.DoRequest(http.MethodGet, "/v1/quotes/", query.Encode())
	if err != nil {
		log.Fatal(err)
	}

	// Make sure response body is closed at end.
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(body, &q)
	if err != nil {
		log.Fatal(err)
	}

	prettyJSON, nil := json.MarshalIndent(q, "", "\t")
	fmt.Printf("%s", prettyJSON)

	return nil
}
