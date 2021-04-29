package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/jaswanth-gorripati/PGK/s7_PaymentGateway/models"
	"github.com/jaswanth-gorripati/PGK/s7_PaymentGateway/services"
)

// CreateOrder ...
func CreateOrder(c *gin.Context) {
	ctx, ID, userType, successResp := getFuncReq(c, "Validating Payment")

	var newPaymentOrder models.CreatePaymentModel
	binding.Validator = &defaultValidator{}
	err := c.ShouldBindWith(&newPaymentOrder, binding.Form)
	if err == nil {
		newPaymentOrder.StakeholderID = ID
		newPaymentOrder.StakeholderType = userType
		newPaymentOrder.PayAmount, err = getPaymentForType(newPaymentOrder.StakeholderType, newPaymentOrder.PayType, newPaymentOrder.TokensUsed, newPaymentOrder.TokensToAdd)
		if err != nil {
			resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S7PG", ErrTyp: "Invalid information", Err: err, SuccessResp: successResp})
			c.JSON(http.StatusBadRequest, resp)
			c.Abort()
			return
		}
		newPaymentOrder.Notes = map[string]interface{}{}
		newPaymentOrder.Notes["stakeholderID"] = newPaymentOrder.StakeholderID
		newPaymentOrder.Notes["stakeholderType"] = newPaymentOrder.StakeholderType
		newPaymentOrder.Notes["tokensUsed"] = newPaymentOrder.TokensUsed
		newPaymentOrder.Notes["tokensToAdd"] = newPaymentOrder.TokensToAdd
		newPaymentOrder.Notes["payType"] = newPaymentOrder.PayType
		newPaymentOrder.Notes["referenceObject"] = newPaymentOrder.ReferenceObject
		orderID, err := services.CreateOrder(newPaymentOrder.PayAmount, newPaymentOrder.Notes)

		if err != nil {
			resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S7PG", ErrTyp: "Internal Server error", Err: err, SuccessResp: successResp})
			c.JSON(http.StatusInternalServerError, resp)
			c.Abort()
			return
		}
		coResp := models.CreatePayRespModel{
			orderID, newPaymentOrder.PayAmount,
		}
		c.JSON(http.StatusOK, coResp)
		return

	}
	resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S7PG", ErrTyp: "Required information not found", Err: err, SuccessResp: successResp})
	c.JSON(http.StatusUnprocessableEntity, resp)
	return

}

func getPaymentForType(st string, paytype string, tokensUsed float64, tokensToAdd float64) (float64, error) {
	var amount float64
	switch paytype {
	case "REG_FEE":
		switch st {
		case "Corporate":
			amount = 15000.00
			break
		case "University":
			amount = 10000.00
			break
		case "Student":
			amount = 5000.00
			break
		default:
			return 0, fmt.Errorf("Invalid Stakeholder for REG_FEE transaction")
		}
		if tokensUsed > 0 {
			if tokensUsed > amount {
				return 0, fmt.Errorf("Reduce Token Usage to " + fmt.Sprintf("%.2f", (amount-tokensUsed)) + " ")
			}
			amount = amount - tokensUsed
		}
		break
	case "ADD_TKN":

		if tokensToAdd <= 0 {
			return 0, fmt.Errorf("Required No of tokens to Purchase")
		}
		amount = tokensToAdd * 5
		break
	default:
		return 0, fmt.Errorf("Invalid Payment type , Expecting REG_FEE / ADD_TKN")
	}

	return amount, nil
}
