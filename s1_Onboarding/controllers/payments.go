package controllers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jaswanth-gorripati/PGK/s1_Onboarding/models"
	"github.com/jaswanth-gorripati/PGK/s1_Onboarding/services"
)

// ShPaymentModel ...
type ShPaymentModel struct {
	Stakeholder string `form:"stakeholder" binding:"required"`
	UserID      string `form:"userID" binding:"required"`
}

// RegPaymentBreakup ...
type RegPaymentBreakup struct {
	RegFee float64 `json:"registrationFee"`
	Tax    float64 `json:"tax"`
	Total  float64 `json:"total"`
}

// PaymentResponse ...
type PaymentResponse struct {
	RegPaymentBreakup
	OrderID string `json:"orderID"`
}

// GetPayment ...
func GetPayment(c *gin.Context) {
	var shPamentInfo ShPaymentModel
	ctx, cancel := context.WithCancel(context.Background())
	ctxkey = ctxFunc("Target")
	ctx = context.WithValue(ctx, ctxkey, "Payment")
	shPamentInfo.Stakeholder = c.GetString("userType")
	shPamentInfo.UserID = c.GetString("userID")
	if shPamentInfo.Stakeholder == "" || shPamentInfo.UserID == "" {
		resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S1REGPAY", ErrTyp: "Required information not found", Err: fmt.Errorf("Cannot get User details from token"), SuccessResp: successResp})
		c.JSON(http.StatusUnprocessableEntity, resp)
		return
	}
	defer cancel()
	notes := map[string]interface{}{"userId": shPamentInfo.UserID}
	orderID := ""
	var payBreakUp RegPaymentBreakup
	var err error
	switch shPamentInfo.Stakeholder {
	case "Corporate":
		payBreakUp = RegPaymentBreakup{20000, 240, 20240}
		orderID, err = services.CreateOrder(payBreakUp.Total, notes)
		break
	case "University":
		payBreakUp = RegPaymentBreakup{10000, 120, 10120}
		orderID, err = services.CreateOrder(payBreakUp.Total, notes)
		break
	case "Student":
		payBreakUp = RegPaymentBreakup{5000, 60, 5060}
		orderID, err = services.CreateOrder(payBreakUp.Total, notes)
		break
	default:
		resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S1REGPAY", ErrTyp: "Invalid Stakehodler type", Err: fmt.Errorf("" + shPamentInfo.Stakeholder + " is invaild,  Expecting Corporate,University or Student"), SuccessResp: successResp})
		c.JSON(http.StatusUnprocessableEntity, resp)
		return
	}
	if err != nil {
		resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S1REGPAY", ErrTyp: "Cannot fetch Payment", Err: fmt.Errorf("Cannot fetch payment due to payment gateway error : %v ", err.Error()), SuccessResp: successResp})
		c.JSON(http.StatusInternalServerError, resp)
		return
	}
	paymenResp := PaymentResponse{payBreakUp, orderID}
	c.JSON(http.StatusOK, paymenResp)
	return

}
