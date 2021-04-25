package services

import (
	"fmt"

	"github.com/jaswanth-gorripati/PGK/s4_Profile/configuration"
	razorpay "github.com/razorpay/razorpay-go"
)

// RazorClient ...
var RazorClient *razorpay.Client

// ConfigPaymentClient ...
func ConfigPaymentClient() {
	razorConfig := configuration.PaymentConfig()
	RazorClient = razorpay.NewClient(razorConfig.KeyID, razorConfig.KeySecret)
}

// CreateOrder ...
func CreateOrder(amount float64, notes map[string]interface{}) (string, error) {
	data := map[string]interface{}{
		"amount":          amount * 100,
		"currency":        "INR",
		"receipt":         "some_receipt_id",
		"payment_capture": 1,
		"notes":           notes,
	}
	body, err := RazorClient.Order.Create(data, nil)
	return body["id"].(string), err
}

// CheckPaymentStatus ...
func CheckPaymentStatus(id string) (bool, error) {
	body, err := RazorClient.Order.Fetch(id, nil, nil)
	if err != nil {
		return false, err
	}
	if body["amount_due"].(float64) == 0 && body["status"].(string) == "paid" {
		return true, nil
	} else {
		return false, fmt.Errorf("Payment incomplete")
	}

}
