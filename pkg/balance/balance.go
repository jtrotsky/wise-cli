package balance

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/jtrotsky/wise-cli/pkg/client"
	"github.com/jtrotsky/wise-cli/pkg/profile"
)

const (
	// AvailableBalance ...
	AvailableBalance = "AVAILABLE"
	// UnavailableBalance ...
	UnavailableBalance = "UNAVAILABLE"
)

// Accounts contains info about all multi-currency accounts.
// https://api-docs.transferwise.com/#payouts-guide-check-account-balance
// // /v1/borderless-accounts?profileId={profileId}
type Accounts struct {
	ID        int64     `json:"id"`
	ProfileID int64     `json:"profileId"`
	Balances  []Balance `json:"balances"`
}

// Balance has information about a specific account for a currency.
type Balance struct {
	Active         bool        `json:"active,omitempty"`
	Amount         Amount      `json:"amount,omitempty"`
	BankDetails    BankDetails `json:"bankDetails,omitempty"`
	Currency       string      `json:"currency,omitempty"`
	Eligible       bool        `json:"eligible,omitempty"`
	Type           string      `json:"balanceType,omitempty"`
	ReservedAmount Amount      `json:"reservedAmount,omitempty"`
}

// Amount contains the balance held for a currency.
type Amount struct {
	Value    float64 `json:"value,omitempty"`
	Currency string  `json:"currency,omitempty"`
}

// BankDetails contain the info needed to deposit into an account.
type BankDetails struct {
	ID                int64       `json:"id,omitempty"`
	Currency          string      `json:"currency,omitempty"`
	BankCode          string      `json:"bankCode,omitempty"`
	AccountNumber     string      `json:"accountNumber,omitempty"`
	BankName          string      `json:"bankName,omitempty"`
	AccountHolderName string      `json:"accountHolderName,omitempty"`
	BankAddress       BankAddress `json:"bankAddress,omitempty"`
}

// BankAddress is a physical address.
type BankAddress struct {
	AddressFirstLine string `json:"addressFirstLine,omitempty"`
	PostCode         string `json:"postCode,omitempty"`
	City             string `json:"city,omitempty"`
	Country          string `json:"country,omitempty"`
	StateCode        string `json:"stateCode,omitempty"`
}

// [
//     {
//         "id": 64,
//         "profileId": <your profile id>,
//         "recipientId": 13828530,
//         "creationTime": "2018-03-14T12:31:15.678Z",
//         "modificationTime": "2018-03-19T15:19:42.111Z",
//         "active": true,
//         "eligible": true,
//         "balances": [
//             {
//                 "balanceType": "AVAILABLE",
//                 "currency": "GBP",
//                 "amount": {
//                     "value": 10999859,
//                     "currency": "GBP"
//                 },
//                 "reservedAmount": {
//                     "value": 0,
//                     "currency": "GBP"
//                 },
//                 "bankDetails": null
//             },
//             {
//                 "balanceType": "AVAILABLE",
//                 "currency": "EUR",
//                 "amount": {
//                     "value": 9945236.2,
//                     "currency": "EUR"
//                 },
//                 "reservedAmount": {
//                     "value": 0,
//                     "currency": "EUR"
//                 },
//                 "bankDetails": null
//             }
//         ]
//     }
// ]

// Get all accounts for a given profile
func Get(client *client.Client, currency string) ([]Accounts, error) {
	accounts := []Accounts{}

	// Get and then set the user's profile (business or personal)
	allProfiles, err := profile.Get(client)
	if err != nil {
		return accounts, err
	}

	personalProfile, err := profile.GetProfileByType(allProfiles, profile.EntityPersonal)
	if err != nil {
		return accounts, err
	}

	client.SetProfile(personalProfile.ID, profile.EntityPersonal)
	if client.ProfileID == 0 {
		return accounts, errors.New("profile ID is required to get balances")
	}

	query := url.Values{}
	query.Add("profileId", fmt.Sprintf("%d", client.ProfileID))

	response, err := client.DoRequest(http.MethodGet, "/v1/borderless-accounts/", query)
	if err != nil {
		return accounts, err
	}

	// Make sure response body is closed at end.
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return accounts, err
	}

	err = json.Unmarshal(body, &accounts)
	if err != nil {
		return accounts, err
	}

	// DEBUG
	// log.Println(accounts)

	// TODO: [0] is filthy
	accounts[0].printBalance(currency)

	return accounts, nil
}

func (accounts Accounts) printBalance(currency string) {
	for _, balance := range accounts.Balances {
		if balance.Type == AvailableBalance {
			if currency != "" {
				if balance.Currency == currency {
					fmt.Printf("You have %.2f %s in cash\n\n", balance.Amount.Value, balance.Amount.Currency)
					return
				} else {
					fmt.Printf("%s balance unavailable\n\n", currency)
					return
				}
			}

			if balance.Amount.Value > 0 {
				fmt.Printf("- %.2f %s in cash\n", balance.Amount.Value, balance.Amount.Currency)
			}
			continue
		}
	}
}
