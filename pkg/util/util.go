package util

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/guptarohit/asciigraph"
	"github.com/jtrotsky/wise-cli/pkg/client"
)

// type CurrencyCode string

// const (
//   AED CurrencyCode = "د.إ",
//   AFN CurrencyCode = "؋",
//   ALL CurrencyCode = "L",
//   AMD CurrencyCode = "֏",
//   ANG CurrencyCode = "ƒ",
//   AOA CurrencyCode = "Kz",
//   ARS CurrencyCode = "$",
//   AUD CurrencyCode = "$",
//   AWG CurrencyCode = "ƒ",
//   AZN CurrencyCode = "₼",
//   BAM CurrencyCode = "KM",
//   BBD CurrencyCode = "$",
//   BDT CurrencyCode = "৳",
//   BGN CurrencyCode = "лв",
//   BHD CurrencyCode = ".د.ب",
//   BIF CurrencyCode = "FBu",
//   BMD CurrencyCode = "$",
//   BND CurrencyCode = "$",
//   BOB CurrencyCode = "$b",
//   BOV CurrencyCode = "BOV",
//   BRL CurrencyCode = "R$",
//   BSD CurrencyCode = "$",
//   BTC CurrencyCode = "₿",
//   BTN CurrencyCode = "Nu.",
//   BWP CurrencyCode = "P",
//   BYN CurrencyCode = "Br",
//   BYR CurrencyCode = "Br",
//   BZD CurrencyCode = "BZ$",
//   CAD CurrencyCode = "$",
//   CDF CurrencyCode = "FC",
//   CHE CurrencyCode = "CHE",
//   CHF CurrencyCode = "CHF",
//   CHW CurrencyCode = "CHW",
//   CLF CurrencyCode = "CLF",
//   CLP CurrencyCode = "$",
//   CNY CurrencyCode = "¥",
//   COP CurrencyCode = "$",
//   COU CurrencyCode = "COU",
//   CRC CurrencyCode = "₡",
//   CUC CurrencyCode = "$",
//   CUP CurrencyCode = "₱",
//   CVE CurrencyCode = "$",
//   CZK CurrencyCode = "Kč",
//   DJF CurrencyCode = "Fdj",
//   DKK CurrencyCode = "kr",
//   DOP CurrencyCode = "RD$",
//   DZD CurrencyCode = "دج",
//   EEK CurrencyCode = "kr",
//   EGP CurrencyCode = "£",
//   ERN CurrencyCode = "Nfk",
//   ETB CurrencyCode = "Br",
//   ETH CurrencyCode = "Ξ",
//   EUR CurrencyCode = "€",
//   FJD CurrencyCode = "$",
//   FKP CurrencyCode = "£",
//   GBP CurrencyCode = "£",
//   GEL CurrencyCode = "₾",
//   GGP CurrencyCode = "£",
//   GHC CurrencyCode = "₵",
//   GHS CurrencyCode = "GH₵",
//   GIP CurrencyCode = "£",
//   GMD CurrencyCode = "D",
//   GNF CurrencyCode = "FG",
//   GTQ CurrencyCode = "Q",
//   GYD CurrencyCode = "$",
//   HKD CurrencyCode = "$",
//   HNL CurrencyCode = "L",
//   HRK CurrencyCode = "kn",
//   HTG CurrencyCode = "G",
//   HUF CurrencyCode = "Ft",
//   IDR CurrencyCode = "Rp",
//   ILS CurrencyCode = "₪",
//   IMP CurrencyCode = "£",
//   INR CurrencyCode = "₹",
//   IQD CurrencyCode = "ع.د",
//   IRR CurrencyCode = "﷼",
//   ISK CurrencyCode = "kr",
//   JEP CurrencyCode = "£",
//   JMD CurrencyCode = "J$",
//   JOD CurrencyCode = "JD",
//   JPY CurrencyCode = "¥",
//   KES CurrencyCode = "KSh",
//   KGS CurrencyCode = "лв",
//   KHR CurrencyCode = "៛",
//   KMF CurrencyCode = "CF",
//   KPW CurrencyCode = "₩",
//   KRW CurrencyCode = "₩",
//   KWD CurrencyCode = "KD",
//   KYD CurrencyCode = "$",
//   KZT CurrencyCode = "₸",
//   LAK CurrencyCode = "₭",
//   LBP CurrencyCode = "£",
//   LKR CurrencyCode = "₨",
//   LRD CurrencyCode = "$",
//   LSL CurrencyCode = "M",
//   LTC CurrencyCode = "Ł",
//   LTL CurrencyCode = "Lt",
//   LVL CurrencyCode = "Ls",
//   LYD CurrencyCode = "LD",
//   MAD CurrencyCode = "MAD",
//   MDL CurrencyCode = "lei",
//   MGA CurrencyCode = "Ar",
//   MKD CurrencyCode = "ден",
//   MMK CurrencyCode = "K",
//   MNT CurrencyCode = "₮",
//   MOP CurrencyCode = "MOP$",
//   MRO CurrencyCode = "UM",
//   MRU CurrencyCode = "UM",
//   MUR CurrencyCode = "₨",
//   MVR CurrencyCode = "Rf",
//   MWK CurrencyCode = "MK",
//   MXN CurrencyCode = "$",
//   MXV CurrencyCode = "MXV",
//   MYR CurrencyCode = "RM",
//   MZN CurrencyCode = "MT",
//   NAD CurrencyCode = "$",
//   NGN CurrencyCode = "₦",
//   NIO CurrencyCode = "C$",
//   NOK CurrencyCode = "kr",
//   NPR CurrencyCode = "₨",
//   NZD CurrencyCode = "$",
//   OMR CurrencyCode = "﷼",
//   PAB CurrencyCode = "B/.",
//   PEN CurrencyCode = "S/.",
//   PGK CurrencyCode = "K",
//   PHP CurrencyCode = "₱",
//   PKR CurrencyCode = "₨",
//   PLN CurrencyCode = "zł",
//   PYG CurrencyCode = "Gs",
//   QAR CurrencyCode = "﷼",
//   RMB CurrencyCode = "￥",
//   RON CurrencyCode = "lei",
//   RSD CurrencyCode = "Дин.",
//   RUB CurrencyCode = "₽",
//   RWF CurrencyCode = "R₣",
//   SAR CurrencyCode = "﷼",
//   SBD CurrencyCode = "$",
//   SCR CurrencyCode = "₨",
//   SDG CurrencyCode = "ج.س.",
//   SEK CurrencyCode = "kr",
//   SGD CurrencyCode = "S$",
//   SHP CurrencyCode = "£",
//   SLL CurrencyCode = "Le",
//   SOS CurrencyCode = "S",
//   SRD CurrencyCode = "$",
//   SSP CurrencyCode = "£",
//   STD CurrencyCode = "Db",
//   STN CurrencyCode = "Db",
//   SVC CurrencyCode = "$",
//   SYP CurrencyCode = "£",
//   SZL CurrencyCode = "E",
//   THB CurrencyCode = "฿",
//   TJS CurrencyCode = "SM",
//   TMT CurrencyCode = "T",
//   TND CurrencyCode = "د.ت",
//   TOP CurrencyCode = "T$",
//   TRL CurrencyCode = "₤",
//   TRY CurrencyCode = "₺",
//   TTD CurrencyCode = "TT$",
//   TVD CurrencyCode = "$",
//   TWD CurrencyCode = "NT$",
//   TZS CurrencyCode = "TSh",
//   UAH CurrencyCode = "₴",
//   UGX CurrencyCode = "USh",
//   USD CurrencyCode = "$",
//   UYI CurrencyCode = "UYI",
//   UYU CurrencyCode = "$U",
//   UYW CurrencyCode = "UYW",
//   UZS CurrencyCode = "лв",
//   VEF CurrencyCode = "Bs",
//   VES CurrencyCode = "Bs.S",
//   VND CurrencyCode = "₫",
//   VUV CurrencyCode = "VT",
//   WST CurrencyCode = "WS$",
//   XAF CurrencyCode = "FCFA",
//   XBT CurrencyCode = "Ƀ",
//   XCD CurrencyCode = "$",
//   XOF CurrencyCode = "CFA",
//   XPF CurrencyCode = "₣",
//   XSU CurrencyCode = "Sucre",
//   XUA CurrencyCode = "XUA",
//   YER CurrencyCode = "﷼",
//   ZAR CurrencyCode = "R",
//   ZMW CurrencyCode = "ZK",
//   ZWD CurrencyCode = "Z$",
//   ZWL CurrencyCode = "$"
// )

// CalculateDeliveryTime returns the time until an estimated future date
func CalculateDeliveryTime(deliveryEstimate time.Time) time.Duration {
	// Duration is the time from now until the estimated delivery time in nanoseconds
	duration := time.Until(deliveryEstimate)

	return duration
}

// RateHistory is a collection of rate records for a currency pair
type RateHistory struct {
	Entries []ExchangeRateRecord
}

// ExchangeRateRecord is the rate Wise recorded at a point in time
//	{
//		"rate": 1.166,
//		"source": "EUR",
//		"target": "USD",
//		"time": "2018-08-15T00:00:00+0000"
//	}
type ExchangeRateRecord struct {
	Rate   float64 `json:"rate,omitempty"`
	Source string  `json:"source,omitempty"`
	Target string  `json:"target,omitempty"`
	Time   string  `json:"time,omitempty"`
}

// ExchangeRateGraph prints an ASCII plot for the last 30 days of rates for a currency pair
func ExchangeRateGraph(client *client.Client, sourceCurrency, targetCurrency string) {
	query := url.Values{}
	query.Add("source", fmt.Sprintf("%s", sourceCurrency))
	query.Add("target", fmt.Sprintf("%s", targetCurrency))
	query.Add("from", fmt.Sprintf("%s", time.Now().UTC().AddDate(0, 0, -30).Format(time.RFC3339))) // 30 days ago
	query.Add("to", fmt.Sprintf("%s", time.Now().UTC().Format(time.RFC3339)))                      // until now
	query.Add("group", "day")                                                                      // group data by day

	response, err := client.DoRequest(http.MethodGet, "/v1/rates/", query.Encode())
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
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

// // GetSymbolFromCurrencyCode looks up a 3 letter currency code provided
// // by the user and returns its symbol e.g. $ or £.
// func GetSymbolFromCurrencyCode(currencyCode string) CurrencySymbol {
// 	if CurrencySymbol[strings.ToUpper(currencyCode)] != "" {
// 		return CurrencySymbol
// 	} else {
// 		return ""
// 	}
// }
