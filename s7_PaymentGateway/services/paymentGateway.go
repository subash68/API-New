package services

import (
	"fmt"
	"strconv"
	"time"

	"github.com/jaswanth-gorripati/PGK/s7_PaymentGateway/configuration"
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
	fmt.Printf("\n========== %v %+v\n", amount, notes)
	data := map[string]interface{}{
		"amount":          amount * 100,
		"currency":        "INR",
		"receipt":         "Receipt#" + strconv.Itoa(int(time.Now().UnixNano()/int64(time.Millisecond))),
		"payment_capture": 1,
		"notes":           notes,
	}
	//fmt.Println(RazorClient)
	body, err := RazorClient.Order.Create(data, nil)
	fmt.Println("=====body", body)
	if err != nil {
		fmt.Println("Got Error ", err.Error())
		return "", err
	}

	//fmt.Println("====Error", err)
	return body["id"].(string), err
}

// CheckPaymentStatus ...
func CheckPaymentStatus(id string) (map[string]interface{}, error) {
	var notes map[string]interface{}
	body, err := RazorClient.Order.Fetch(id, nil, nil)
	if err != nil {
		return notes, err
	}
	if body["amount_due"].(float64) == 0 && body["status"].(string) == "paid" {
		notes = body["notes"].(map[string]interface{})
		notes["amountPaid"] = body["amount_paid"]
		return notes, nil
	} else {
		return notes, fmt.Errorf("Payment status %v" + body["status"].(string))
	}

	// // For testing
	// if body["amount_due"].(float64) != 0 || body["status"].(string) == "paid" {
	// 	notes = body["notes"].(map[string]interface{})
	// 	notes["amountPaid"] = body["amount_due"]
	// 	return notes, nil
	// } else {
	// 	return notes, fmt.Errorf("Payment status %v" + body["status"].(string))
	// }

}
