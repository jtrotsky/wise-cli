package quote

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/jtrotsky/wise-cli/pkg/util"

	"github.com/jtrotsky/wise-cli/pkg/client"
)

// Quote creates a fixed record of a to and from currency and amount which can later be used to
// create a transfer https://api-docs.transferwise.com/#quotes-create
type Quote struct {
	ID                     string     `json:"id,omitempty"`
	SourceCurrency         string     `json:"source,omitempty"`
	TargetCurrency         string     `json:"target,omitempty"`
	SourceAmount           float64    `json:"sourceAmount,omitempty"`
	TargetAmount           float64    `json:"targetAmount,omitempty"`
	Rate                   float64    `json:"rate,omitempty"`
	Type                   string     `json:"type,omitempty"`
	RateType               string     `json:"rateType,omitempty"`
	CreatedTime            time.Time  `json:"createdTime,omitempty"`
	CreatedByUserID        string     `json:"createdByUserId,omitempty"`
	Profile                int64      `json:"profile,omitempty"`
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

// Prepare a quote to be sent to Wise
func Prepare(profileID int64, fromCurrency, toCurrency string, sourceAmount float64) *Quote {
	return &Quote{
		Profile:        profileID,
		SourceCurrency: fromCurrency,
		TargetCurrency: toCurrency,
		SourceAmount:   sourceAmount,
		RateType:       "FIXED", // TODO
	}
}

// Create a new quote from Wise based on currency pair and amount provided
func (q *Quote) Create(client *client.Client) error {
	query := url.Values{}
	query.Add("source", q.SourceCurrency)
	query.Add("target", q.TargetCurrency)
	query.Add("sourceAmount", fmt.Sprintf("%f", q.SourceAmount))
	query.Add("rateType", q.RateType)

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

	// Print exchange rate ASCII plot from rate history for last 30 days
	fmt.Print("\n")
	util.ExchangeRateGraph(client, q.SourceCurrency, q.TargetCurrency)

	// calculate time until the delivery estimate
	deliveryTime := util.CalculateDeliveryTime(q.DeliveryEstimate)

	fmt.Printf("\nQuote for %.0f %s to %s at 1=%f",
		q.SourceAmount, q.SourceCurrency, q.TargetCurrency, q.Rate)
	fmt.Printf("\n -> %.2f %s would arrive in %.0fh\n",
		q.TargetAmount, q.TargetCurrency, deliveryTime.Hours())

	return nil
}
