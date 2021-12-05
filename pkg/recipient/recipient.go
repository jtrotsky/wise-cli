package recipient

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	"github.com/jtrotsky/wise-cli/pkg/client"
)

const (
	RecipientTypeSortCode = "sort_code"
	LegalTypePrivate      = "PRIVATE"
)

// Recipient is an individual or business you can send money to.
type Recipient struct {
	ID                string
	ProfileID         int64   `json:"profile,omitempty"`
	Currency          string  `json:"currency,omitempty"`
	Type              string  `json:"type,omitempty"`
	AccountHolderName string  `json:"accountHolderName,omitempty"`
	Details           Details `json:"details,omitempty"`

	// Useless field per https://api-docs.transferwise.com/#recipient-accounts-create
	// OwnedByCustomer   string  `json:"ownedByCustomer,omitempty"`
}

// Details contains the type-specific information about a bank record.
type Details struct {
	LegalType     string  `json:"legalType,omitempty"`
	SortCode      string  `json:"sortCode,omitempty"`
	AccountNumber string  `json:"accountNumber,omitempty"`
	AccountType   string  `json:"accountType,omitempty"`
	ABARTN        string  `json:"abartn,omitempty"` // American Bankers Association Routing Transit Number (US)
	IBAN          string  `json:"iban,omitempty"`   // International Bank Account Number (EU)
	BIC           string  `json:"bic,omitempty"`    // Bank Identification Code
	Address       Address `json:"address,omitempty"`
}

type Address struct {
	Country   string `json:"country,omitempty"`
	State     string `json:"state,omitempty"`
	City      string `json:"city,omitempty"`
	PostCode  string `json:"postCode,omitempty"`
	FirstLine string `json:"firstLine,omitempty"`
}

//{
// 	"currency": "GBP",
// 	"type": "sort_code",
// 	"profile": <your profile id>,
// 	"ownedByCustomer": true,
// 	"accountHolderName": "Ann Johnson",
// 	"details": {
//    "legalType": "PRIVATE",
// 	  "sortCode": "231470",
// 	  "accountNumber": "28821822"
//   }
//}

// Prepare a recipient to be sent to Wise
func Prepare(profileID int64, currency, sortCode, accountNumber string) *Recipient {
	return &Recipient{
		ProfileID: profileID,
		Currency:  currency,
		Details: Details{
			SortCode:      sortCode,
			AccountNumber: accountNumber,
		},
	}
}

// Create a new recipient record
func (r *Recipient) Create(client *client.Client) error {
	query := url.Values{}
	query.Add("profile", fmt.Sprintf("%d", r.ProfileID))
	query.Add("currency", r.Currency)
	query.Add("details", r.Details.SortCode)
	query.Add("details", r.Details.AccountNumber)

	response, err := client.DoRequest(http.MethodPost, "/v1/accounts", query)
	if err != nil {
		log.Fatal(err)
	}

	// Make sure response body is closed at end
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(body, &r)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Recipient created with ID: %s\n", r.ID)

	return nil
}
