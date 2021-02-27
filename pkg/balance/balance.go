package balance

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/jtrotsky/wise-cli/pkg/client"
	"github.com/jtrotsky/wise-cli/pkg/profile"
	ccy "golang.org/x/text/currency"
)

const (
	// AvailableBalance ...
	AvailableBalance = "AVAILABLE"
	// UnavailableBalance ...
	UnavailableBalance = "UNAVAILABLE"
)

// Account contains info about all multi-currency accounts.
// https://api-docs.transferwise.com/#payouts-guide-check-account-balance
type Account struct {
	ID        int64     `json:"id"`
	ProfileID int64     `json:"profileId"`
	Balances  []Balance `json:"balances"`
}

// Balance has information about a specific account for a currency.
type Balance struct {
	Active         bool        `json:"active,omiitempty"`
	Amount         Amount      `json:"amount,omiitempty"`
	BankDetails    BankDetails `json:"bankDetails,omitempty"`
	Currency       string      `json:"currency,omiitempty"`
	Eligible       bool        `json:"eligible,omiitempty"`
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

// Get an account for a given profile
func Get(client *client.Client, currency string) (Account, error) {
	accounts := []Account{}

	// Get and then set the user's profile (business or personal)
	allProfiles, err := profile.Get(client)
	if err != nil {
		return Account{}, err
	}

	personalProfile, err := profile.GetProfileByType(allProfiles, profile.EntityPersonal)
	if err != nil {
		return Account{}, err
	}

	client.SetProfile(personalProfile.ID, profile.EntityPersonal)
	if client.ProfileID == 0 {
		return Account{}, errors.New("profile ID is required to get balances")
	}

	query := url.Values{}
	query.Add("profileId", fmt.Sprintf("%d", client.ProfileID))

	response, err := client.DoRequest(http.MethodGet, "/v1/borderless-accounts/", query.Encode())
	if err != nil {
		return Account{}, err
	}

	// Make sure response body is closed at end.
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return Account{}, err
	}

	err = json.Unmarshal(body, &accounts)
	if err != nil {
		return Account{}, err
	}

	// We know there is only 1 account returned because we specify a profile ID
	// but the Wise API returns an array anyway, so we accommodate.
	personalAccount := Account{}
	for _, account := range accounts {
		if account.ProfileID == client.ProfileID {
			personalAccount = account
			break
		}
	}

	personalAccount.printBalances(currency)

	return personalAccount, nil
}

// printBalances formats and prints all balances for a Wise account.
// Or if specified it prints only the balance asked for.
func (account Account) printBalances(currency string) {
	var emptyBalances []string
	for _, balance := range account.Balances {
		unit, _ := ccy.ParseISO(strings.ToUpper(balance.Amount.Currency))
		// unit.Amount(balance.Amount.Value)
		formatter := ccy.NarrowSymbol.Default(unit)

		if unit.String() == strings.ToUpper(currency) {
			fmt.Printf("\nYou have %s in your %s account", formatter(balance.Amount.Value), balance.Amount.Currency)
			return
		}

		if currency != "" {
			continue
		}

		if balance.Amount.Value == 0 {
			emptyBalances = append(emptyBalances, strings.ToUpper(balance.Amount.Currency))
			continue
		}

		fmt.Printf("\nYou have %s in your %s account", formatter(balance.Amount.Value), balance.Amount.Currency)
	}

	fmt.Printf("\nAnd these accounts are empty: %s", strings.Join(emptyBalances, ", "))
}
