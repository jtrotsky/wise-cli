package client

// Client contains the authorisation and config for the Wise account
type Client struct {
	Config *WiseConfig
}

// New creates a new instance of Client and loads config
func New() Client {
	return Client{LoadConfig()}
}

// Run uses the right factory to complete the action requested
// func (c *Client) Run(action string) {
// 	sourceCurrency := "NZD"
// 	targetCurrency := "AUD"
// 	targetAmount := 100

// 	quote := quote.Prepare(sourceCurrency, targetCurrency, targetAmount)

// 	err := quote.Create()
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Println(quote)
// }
