/*
Copyright © 2021 Joe Armstrong <joe.armstrong@transferwise.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// quoteCmd represents the quote command
var quoteCmd = &cobra.Command{
	Use:   "quote",
	Short: "Get a quote for sending an amount between currencies",
	Long: `A quote provides a summary of the cost and exchange rate to send money. 
Quotes can also hold a fixed exchange rate for a period of time (e.g. 24hrs).

The quote reference should also be used as an input to create a transfer. For example:
// create the quote
./wise-cli quote create --amount 100 --from GBP --to NZD

// use the quote to create a transfer 
./wise-cli transfer create --quote 12345678`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("quote called")
	},
}

func init() {
	rootCmd.AddCommand(quoteCmd)

	balanceCmd.PersistentFlags().Int("amount", 0, "The amount to send or convert")
	balanceCmd.PersistentFlags().String("from", "", "The currency to send from (e.g. GBP)")
	balanceCmd.PersistentFlags().String("to", "", "The currency to send to (e.g. NZD)")
}
