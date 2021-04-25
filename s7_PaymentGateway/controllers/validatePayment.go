package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/jaswanth-gorripati/PGK/s7_PaymentGateway/configuration"
	"github.com/jaswanth-gorripati/PGK/s7_PaymentGateway/middleware"
	"github.com/jaswanth-gorripati/PGK/s7_PaymentGateway/models"
	"github.com/jaswanth-gorripati/PGK/s7_PaymentGateway/services"
)

// TokenDbResp ...
type TokenDbResp struct {
	Message string `json:"message"`
}

// ValidatePayment ...
func ValidatePayment(c *gin.Context) {

	// jobdb := make(chan models.DbModelError, 1)
	ctx, ID, _, _ := getFuncReq(c, "Validating Payment")
	var paymentOrder models.PaySuccessReqModel
	binding.Validator = &defaultValidator{}
	err := c.ShouldBindWith(&paymentOrder, binding.Form)
	if err == nil {
		paymentDetails, err := services.CheckPaymentStatus(paymentOrder.OrderID)
		fmt.Println(paymentDetails)
		if err != nil {
			resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S6PJ", ErrTyp: "Internal Server Error", Err: err, SuccessResp: successResp})
			c.JSON(http.StatusUnprocessableEntity, resp)
			c.Abort()
			return
		}
		var paymentModel models.PaymentDbModel
		paymentModel.StakeholderID = paymentDetails["stakeholderID"].(string)
		if paymentModel.StakeholderID != ID {
			resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S6PJ", ErrTyp: "Internal Server Error", Err: fmt.Errorf("Bad request"), SuccessResp: successResp})
			c.JSON(http.StatusBadRequest, resp)
			c.Abort()
			return
		}
		paymentModel.PaymentID = paymentOrder.OrderID
		paymentModel.PaymentMode = "Credit Card"
		paymentModel.PayedAmount = paymentDetails["amountPaid"].(float64)

		errLogChan := make(chan error, 3)
		successLogChan := make(chan string, 3)

		chanLength := 6

		switch paymentDetails["payType"].(string) {

		case "REG_FEE":

			go func() {
				resp, err := addPayment(ctx, paymentModel, paymentDetails["stakeholderType"].(string))
				fmt.Println(resp, " ==================== AddPayment")
				errLogChan <- err
				successLogChan <- resp
			}()
			go func() {
				reqBody := map[string]string{"stakeholder": paymentDetails["stakeholderType"].(string), "stakeholderID": paymentDetails["stakeholderID"].(string)}
				resp, err := markAccountActive(reqBody)
				fmt.Println(resp, " ==================== Mark Active")
				errLogChan <- err
				successLogChan <- resp
			}()
			go func() {
				reqBody := map[string]string{"modeOfTokenissue": "Bonus", "stakeholderID": paymentDetails["stakeholderID"].(string), "paymentID": paymentOrder.OrderID, "allocatedTokens": "1000"}
				resp, err := makeTokenServiceCall("/t/addAllocation", reqBody)
				fmt.Println(resp, " ==================== Give Bonus tokens")
				errLogChan <- err
				successLogChan <- resp
			}()

			break

		case "ADD_TKN":

			go func() {
				resp, err := addPayment(ctx, paymentModel, paymentDetails["stakeholderType"].(string))
				errLogChan <- err
				successLogChan <- resp
			}()
			go func() {
				reqBody := map[string]string{"modeOfTokenissue": "Paid", "stakeholderID": paymentDetails["stakeholderID"].(string), "paymentID": paymentOrder.OrderID, "allocatedTokens": fmt.Sprintf("%.2f", paymentDetails["tokensToAdd"].(float64))}
				resp, err := makeTokenServiceCall("/t/addAllocation", reqBody)
				errLogChan <- err
				successLogChan <- resp
			}()
			chanLength = 4

			break

		default:
			resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S6PJ", ErrTyp: "Required information not found from Payment", Err: fmt.Errorf("Invalid Payment category %v", paymentDetails["payType"].(string)), SuccessResp: successResp})
			c.JSON(http.StatusUnprocessableEntity, resp)
			c.Abort()
			return
		}
		var successMessage []string
		for i := 0; i < chanLength; i++ {
			fmt.Println("=========== SELECT Statement EXEC ============> \n", i, chanLength, "\n==============\n")
			select {
			case errorMsg := <-errLogChan:
				if errorMsg != nil {
					fmt.Println("=========== Error in validate Payment ============> \n", errorMsg, "\n==============\n")
					resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S6PJ", ErrTyp: "Error occurred", Err: fmt.Errorf("%v", errorMsg), SuccessResp: successResp})
					c.JSON(http.StatusUnprocessableEntity, resp)
					c.Abort()
					return
				}
			case successMsg := <-successLogChan:
				fmt.Println("=========== Success in validate Payment ============> \n", successMsg, "\n==============\n")
				successMessage = append(successMessage, successMsg)
			}
		}
		close(successLogChan)
		close(errLogChan)
		c.JSON(http.StatusOK, successMessage)
		c.Abort()
		return

	}
	resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S6PJ", ErrTyp: "Required information not found", Err: err, SuccessResp: successResp})
	c.JSON(http.StatusUnprocessableEntity, resp)
	return
}

func addPayment(ctx context.Context, paymentModel models.PaymentDbModel, userType string) (string, error) {
	insertJobChan := paymentModel.AddPayment(userType)
	if insertJobChan.ErrTyp != "000" {
		return "", insertJobChan.Err
	}
	fmt.Println("AccountPayment is working")
	return "Payment has been Successful", nil
}

func markAccountActive(reqData map[string]string) (string, error) {
	onboardConfig := configuration.OnboardConfig()
	resBody, err := middleware.MakeInternalServiceCall(onboardConfig.Host, onboardConfig.Port, "POST", "/o/changeAccStatus", reqData)
	if err != nil {
		return "", err
	}
	fmt.Println(string(resBody))
	return "Account status updated", nil

}

func makeTokenServiceCall(endpoint string, reqData map[string]string) (string, error) {
	tokenConfig := configuration.TokenConfig()
	resBody, err := middleware.MakeInternalServiceCall(tokenConfig.Host, tokenConfig.Port, "POST", endpoint, reqData)
	if err != nil {
		return "", err
	}
	var tokenResp TokenDbResp
	err = json.Unmarshal(resBody, &tokenResp)
	if err != nil {
		return "", err
	}
	return tokenResp.Message, nil
}
