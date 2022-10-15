package util

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"time"

	"golang.org/x/text/currency"

	"github.com/guptarohit/asciigraph"
	"github.com/jtrotsky/wise-cli/pkg/client"
)

// CalculateDeliveryTime returns the time until an estimated future date
func CalculateDeliveryTime(deliveryEstimate time.Time) time.Duration {
	// duration is the time from now until the estimated delivery time in nanoseconds
	duration := time.Until(deliveryEstimate)

	return duration
}

// RateHistory is a collection of rate records for a currency pair
type RateHistory struct {
	Entries []ExchangeRateRecord
}

// ExchangeRateRecord is the recorded exchange rate at a point in time
type ExchangeRateRecord struct {
	Rate   float64 `json:"rate,omitempty"`
	Source string  `json:"source,omitempty"`
	Target string  `json:"target,omitempty"`
	Time   string  `json:"time,omitempty"`
}

// [
//     {
//         "rate": 1.166,
//         "source": "EUR",
//         "target": "USD",
//         "time": "2018-08-15T00:00:00+0000"
//     },
//     {
//         "rate": 1.168,
//         "source": "EUR",
//         "target": "USD",
//         "time": "2018-06-30T00:00:00+0000"
//     }
//     ...
// ]

// ExchangeRateGraph prints an ASCII plot for the last 30 days of rates for a currency pair
func ExchangeRateGraph(client *client.Client, sourceCurrency, targetCurrency string) {
	query := url.Values{}
	query.Add("source", sourceCurrency)
	query.Add("target", targetCurrency)
	query.Add("from", time.Now().UTC().AddDate(0, 0, -30).Format(time.RFC3339)) // 30 days ago
	query.Add("to", time.Now().UTC().Format(time.RFC3339))                      // until now
	query.Add("group", "day")                                                   // group data by day

	response, err := client.DoRequest(http.MethodGet, "/v1/rates/", query)
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	rateHistory := RateHistory{
		Entries: []ExchangeRateRecord{},
	}
	err = json.Unmarshal(body, &rateHistory.Entries)
	if err != nil {
		log.Fatal(err)
	}

	rates := rateHistory.Rates()
	graph := asciigraph.Plot(rates, asciigraph.Caption("30 days"))
	fmt.Println(graph)
}

// Rates returns all the number values for rate from a collection of entries
func (r *RateHistory) Rates() []float64 {
	rates := []float64{}
	for _, entry := range r.Entries {
		rates = append(rates, entry.Rate)
	}
	return rates
}

// ParseTime is to normalise different time formats
func ParseTime(timeStr string) time.Time {
	time, err := time.Parse("2006-01-02T15:04:05+0700", timeStr)
	if err != nil {
		log.Fatal(err)
	}
	return time
}

// Check that a currency code is a recognised ISO 4217 value and well formatted
func ValidCurrencyCode(code string) bool {
	_, err := currency.ParseISO(code)
	if err != nil {
		return false
	}
	return true
}
