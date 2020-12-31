package quote

import (
	"errors"
	"fmt"
)

// Quote is used to create a transfer https://api-docs.transferwise.com/#quotes-create
type Quote struct {
	ID             string
	SourceCurrency string
	TargetCurrency string
	Amount         int
	Rate           float64
}

// Prepare a quote to be sent to Wise
func Prepare(sourceCurrency, targetCurrency string, targetAmount int) *Quote {
	return &Quote{
		SourceCurrency: sourceCurrency,
		TargetCurrency: targetCurrency,
		Amount:         targetAmount,
	}
}

// Create a new quote from Wise based on currency pair and amount provided
func (q *Quote) Create() error {

	q.Rate = 0.91435
	result := fmt.Sprintf("%s to %s with rate of %f and amount of %d", q.SourceCurrency, q.TargetCurrency, q.Rate, q.Amount)

	fmt.Println(result)

	return errors.New("fail")
}
