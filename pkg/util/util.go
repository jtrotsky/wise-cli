package util

import (
	"time"
)

func calcDeliveryTime(deliveryEstimate time.Time) time.Duration {
	// t, err := time.Parse(time.RFC3339, deliveryEstimate)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// duration is the time from now until the estimated delivery time in nanoseconds
	duration := time.Until(deliveryEstimate)

	return duration
}
