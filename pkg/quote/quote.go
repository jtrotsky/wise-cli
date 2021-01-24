package quote

import (
	"encoding/json"
	"fmt"
	"github.com/jtrotsky/wise-cli/pkg/util"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/jtrotsky/wise-cli/pkg/client"
	"github.com/jtrotsky/wise-cli/pkg/profile"
)

const (
	// BankTransferPayIn represents the cheap pay-in option of bank transfer
	BankTransferPayIn = "BANK_TRANSFER"
	// BalancePayIn represents the cheap pay-in option of Wise balance funding
	BalancePayIn = "BALANCE"
	// FixedRateType is when a quote has an exchange rate that is locked for a period of time
	FixedRateType = "FIXED"
	// VariableRateType is when a quote exchange rate stays live
	VariableRateType = "VARIABLE"
)

// Quote is a V2 quote object which is similar to Quote V1 in fields but includes rate fix info.
type Quote struct {
	ID                 string    `json:"id,omitempty"`
	SourceCurrency     string    `json:"sourceCurrency,omitempty"`
	TargetCurrency     string    `json:"targetCurrency,omitempty"`
	SourceAmount       float64   `json:"sourceAmount,omitempty"`
	TargetAmount       float64   `json:"targetAmount,omitempty"`
	PayOut             string    `json:"payOut,omitempty"`
	Rate               float64   `json:"rate,omitempty"`
	CreatedTime        time.Time `json:"createdTime,omitempty"`
	User               int64     `json:"user,omitempty"`
	Profile            int64     `json:"profile,omitempty"`
	RateType           string    `json:"rateType,omitempty"`
	RateExpirationTime time.Time `json:"rateExpirationTime,omitempty"`

	// GuaranteedTargetAmountAllowed bool       `json:"GuaranteedTargetAmountAllowed,omitempty"`
	// TargetAmountAllowed bool       `json:"TargetAmountAllowed,omitempty"`
	// GuaranteedTargetAmount bool       `json:"GuaranteedTargetAmount,omitempty"`
	// ProvidedAmountType         string       `json:"providedAmountType,omitempty"`

	PaymentOptions []PaymentOption `json:"paymentOptions,omitempty"`
	Status         string          `json:"status,omitempty"`
	ExpirationTime time.Time       `json:"expirationTime,omitempty"`
	Notices        Notice          `json:"notices,omitempty"`
}

// PaymentOption describes the various ways to fund a transfer, e.g. from balance.
type PaymentOption struct {
	Disabled                   bool      `json:"disabled,omitempty"`
	EstimatedDelivery          time.Time `json:"estimatedDelivery,omitempty"`
	FormattedEstimatedDelivery string    `json:"formattedEstimatedDelivery,omitempty"`
	EstimatedDeliveryDelays    []string  `json:"estimatedDeliveryDelays,omitempty"`
	Fee                        Fee       `json:"fee,omitempty"`
	SourceAmount               float64   `json:"sourceAmount,omitempty"`
	TargetAmount               float64   `json:"targetAmount,omitempty"`
	SourceCurrency             string    `json:"sourceCurrency,omitempty"`
	TargetCurrency             string    `json:"targetCurrency,omitempty"`
	PayIn                      string    `json:"payIn,omitempty"`
	PayOut                     string    `json:"payOut,omitempty"`
	AllowedProfileTypes        []string  `json:"allowedProfileTypes,omitempty"`
	PayInProduct               string    `json:"payInProduct,omitempty"`
	FeePercentage              float64   `json:"feePercentage,omitempty"`
}

// Fee contains the breakdown of fee components
type Fee struct {
	TransferWise float64 `json:"transferwise,omitempty"`
	PayIn        float64 `json:"payIn,omitempty"`
	Discount     float64 `json:"discount,omitempty"`
	PriceSetID   int     `json:"priceSetId,omitempty"`
	Partner      float64 `json:"partner,omitempty"`
}

// Notice contains special information about a quote.
type Notice struct {
	Text string `json:"text,omitempty"`
	Link string `json:"link,omitempty"`
	Type string `json:"type,omitempty"`
}

// {
// "id": "11144c35-9fe8-4c32-b7fd-d05c2a7734bf",
// "sourceCurrency": "GBP",
// "targetCurrency": "USD",
// "sourceAmount": 100,
// "payOut": "BANK_TRANSFER",
// "rate": 1.30445,
// "createdTime": "2019-04-05T13:18:58Z",
// "user": 55,
// "profile": 101,
// "rateType": "FIXED",
// "rateExpirationTime": "2019-04-08T13:18:57Z",
// "guaranteedTargetAmountAllowed": true,
// "targetAmountAllowed": true,
// "guaranteedTargetAmount": false,
// "providedAmountType": "SOURCE",
// "paymentOptions": [
// 		{
// 				"disabled": false,
// 				"estimatedDelivery": "2019-04-08T12:30:00Z",
// 				"formattedEstimatedDelivery": "by Apr 8",
// 				"estimatedDeliveryDelays": [],
// 				"fee": {
// 						"transferwise": 0.92,
// 						"payIn": 0,
// 						"discount": 0,
// 						"partner": 0,
// 						"total": 0.92
// 				},
// 				"sourceAmount": 100,
// 				"targetAmount": 129.24,
// 				"sourceCurrency": "GBP",
// 				"targetCurrency": "USD",
// 				"payIn": "BANK_TRANSFER",
// 				"payOut": "BANK_TRANSFER",
// 				"allowedProfileTypes": [
// 						"PERSONAL",
// 						"BUSINESS"
// 				],
// 				"payInProduct": "CHEAP",
// 				"feePercentage": 0.0092
// 		},
// 		{
// 				"disabled": true,
// 				"estimatedDelivery": null,
// 				"formattedEstimatedDelivery": null,
// 				"estimatedDeliveryDelays": [],
// 				"fee": {
// 						"transferwise": 1.11,
// 						"payIn": 0,
// 						"discount": 0,
// 						"partner": 0,
// 						"total": 1.11
// 				},
// 				"sourceAmount": 100,
// 				"targetAmount": 129,
// 				"sourceCurrency": "GBP",
// 				"targetCurrency": "USD",
// 				"payIn": "BALANCE",
// 				"payOut": "BANK_TRANSFER",
// 				"allowedProfileTypes": [
// 						"PERSONAL",
// 						"BUSINESS"
// 				],
// 				"disabledReason": {
// 						"code": "error.payInmethod.disabled",
// 						"message": "Open a borderless account and add funds to instantly pay for your transfers."
// 				},
// 				"payInProduct": "BALANCE",
// 				"feePercentage": 0.0111
// 		}
// ],
// "status": "PENDING",
// "expirationTime": "2019-04-05T13:48:58Z",
// "notices": [{
// 		"text": "You can have a maximum of 3 open transfers with a guaranteed rate. After that, they'll be transferred using the live rate. Complete or cancel your other transfers to regain the use of guaranteed rate.",
// 		"link": null,
// 		"type": "WARNING"
// }]
// }

// below is a v1 quote example

// // QuoteV1 creates a fixed record of a to and from currency and amount which can later be used to
// // create a transfer https://api-docs.transferwise.com/#quotes-create
// type QuoteV1 struct {
// 	ID                     string     `json:"id,omitempty"`
// 	SourceCurrency         string     `json:"source,omitempty"`
// 	TargetCurrency         string     `json:"target,omitempty"`
// 	SourceAmount           float64    `json:"sourceAmount,omitempty"`
// 	TargetAmount           float64    `json:"targetAmount,omitempty"`
// 	Rate                   float64    `json:"rate,omitempty"`
// 	Type                   string     `json:"type,omitempty"`
// 	RateType               string     `json:"rateType,omitempty"`
// 	RateExpirationTime     time.Time  `json:"rateExpirationTime,omitempty"`
// 	CreatedTime            time.Time  `json:"createdTime,omitempty"`
// 	CreatedByUserID        string     `json:"createdByUserId,omitempty"`
// 	Profile                int64      `json:"profile,omitempty"`
// 	DeliveryEstimate       time.Time  `json:"deliveryEstimate,omitempty"`
// 	Fee                    float64    `json:"fee,omitempty"`
// 	FeeDetails             FeeDetails `json:"feeDetails,omitempty"`
// 	AllowedProfileTypes    []string   `json:"allowedProfileTypes,omitempty"`
// 	GuaranteedTargetAmount bool       `json:"GuaranteedTargetAmount,omitempty"`
// 	OfSourceAmount         bool       `json:"OfSourceAmount,omitempty"`
// }

// // FeeDetails contains the breakdown of fee components
// type FeeDetails struct {
// 	TransferWise float64 `json:"transferwise,omitempty"`
// 	PayIn        float64 `json:"payIn,omitempty"`
// 	Discount     float64 `json:"discount,omitempty"`
// 	PriceSetID   int     `json:"priceSetId,omitempty"`
// 	Partner      float64 `json:"partner,omitempty"`
// }

// {
//   "source": "GBP",
//   "target": "NZD",
//   "sourceAmount": 100,
//   "targetAmount": 188.26,
//   "type": "REGULAR",
//   "rate": 1.89929,
// "rateExpirationTime": "2019-04-08T13:18:57Z", // does this actually exist
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
		RateType:       FixedRateType, // TODO
	}
}

// Create a new quote from Wise based on currency pair and amount provided
func (q *Quote) Create(client *client.Client) error {
	// Get and then set the user's profile (business or personal)
	allProfiles, err := profile.Get(client)
	if err != nil {
		return err
	}

	personalProfile, err := profile.GetProfileByType(allProfiles, profile.EntityPersonal)
	if err != nil {
		return err
	}

	// If profile is empty we can still proceed to quote
	client.SetProfile(personalProfile.ID, profile.EntityPersonal)

	query := url.Values{}
	query.Add("sourceCurrency", fmt.Sprintf("%s", q.SourceCurrency))
	query.Add("targetCurrency", fmt.Sprintf("%s", q.TargetCurrency))
	query.Add("sourceAmount", fmt.Sprintf("%f", q.SourceAmount))
	query.Add("profile", fmt.Sprintf("%d", client.ProfileID))

	response, err := client.DoRequest(http.MethodPost, "/v2/quotes", query)
	if err != nil {
		log.Fatal(err)
	}

	// Make sure response body is closed at end.
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	if client.Debug {
		fmt.Printf("%s", body)
	}

	err = json.Unmarshal(body, &q)
	if err != nil {
		log.Fatal(err)
	}

	// Print exchange rate ASCII plot from rate history for last 30 days
	fmt.Print("\n")
	util.ExchangeRateGraph(client, q.SourceCurrency, q.TargetCurrency)

	// calculate time until the money should arrive assuming bank transfer payment option is used
	var deliveryTime time.Duration
	for _, paymentOption := range q.PaymentOptions {
		if paymentOption.PayIn == BankTransferPayIn {
			deliveryTime, err = util.TimeUntil(paymentOption.EstimatedDelivery)
			if err != nil {
				return err
			}
		}
	}

	// calculate how long the quoted exchange rate will be valid for
	expirationTime, err := util.TimeUntil(q.RateExpirationTime)
	if err != nil {
		return err
	}

	fmt.Printf("\nQuote for %.0f %s to %s at 1=%f (rate fixed for %.0fh)",
		q.SourceAmount, q.SourceCurrency, q.TargetCurrency, q.Rate, expirationTime.Hours())
	fmt.Printf("\n -> %.2f %s will arrive in %.0fh\n",
		q.TargetAmount, q.TargetCurrency, deliveryTime.Hours())

	return nil
}
