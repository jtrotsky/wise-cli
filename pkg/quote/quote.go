package quote

import (
	"errors"
	"fmt"
)

// Get a new quote from TransferWise based on currency pair and amount provided
func Get(sourceCurrency string, targetCurrency string, amount float64) (string, error) {
	// lookup tw api for quote
	// get latest rate

	rate := 0.91435
	quoteResult := fmt.Sprintf("%s to %s with rate of %f and amount of %f", sourceCurrency, targetCurrency, rate, amount)

	return quoteResult, errors.New("fail")
}
