package transfer

import (
	"fmt"
	"log"

	"github.com/jtrotsky/wise-cli/pkg/client"
)

// Transfer creates a fixed record of a to and from currency and amount which can later be used to
// create a transfer https://api-docs.transferwise.com/#quotes-create
type Transfer struct {
	QuoteID string `json:"quote_id,omitempty"`
	// SourceCurrency         string     `json:"source,omitempty"`
	// TargetCurrency         string     `json:"target,omitempty"`
	// SourceAmount           float64    `json:"sourceAmount,omitempty"`
	// TargetAmount           float64    `json:"targetAmount,omitempty"`
	// Rate                   float64    `json:"rate,omitempty"`
	// Type                   string     `json:"type,omitempty"`
	// RateType               string     `json:"rateType,omitempty"`
	// CreatedTime            time.Time  `json:"createdTime,omitempty"`
	// CreatedByUserID        string     `json:"createdByUserId,omitempty"`
	// Profile                int64      `json:"profile,omitempty"`
	// DeliveryEstimate       time.Time  `json:"deliveryEstimate,omitempty"`
	// Fee                    float64    `json:"fee,omitempty"`
	// FeeDetails             FeeDetails `json:"feeDetails,omitempty"`
	// AllowedProfileTypes    []string   `json:"allowedProfileTypes,omitempty"`
	// GuaranteedTargetAmount bool       `json:"GuaranteedTargetAmount,omitempty"`
	// OfSourceAmount         bool       `json:"OfSourceAmount,omitempty"`
}

// {

// }

// Prepare ...
func Prepare(profileID int64, quoteID string) *Transfer {
	return &Transfer{
		QuoteID: quoteID,
	}
}

// Create ...
func (t *Transfer) Create(client *client.Client) error {
	if t.QuoteID == "" {
		log.Fatal("Provide a quote to create a tranfer from")
	}

	fmt.Printf("new transfer created from quote: %s", t.QuoteID)

	return nil
}
