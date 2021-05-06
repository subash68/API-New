// Package controllers ...
package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/jaswanth-gorripati/PGK/s5_Subcription/configuration"
	"github.com/jaswanth-gorripati/PGK/s5_Subcription/middleware"
	"github.com/jaswanth-gorripati/PGK/s5_Subcription/models"
)

// TokenDbResp ...
type TokenDbResp struct {
	Message string `json:"message"`
	NftID   string `json:"nftID"`
}

// Subscribe ...
func Subscribe(c *gin.Context) {
	successResp = map[string]string{}
	var newSubscriptions models.SubscriptionReq
	jobdb := make(chan models.DbModelError, 1)

	ctx, cancel := context.WithCancel(context.Background())
	ctxkey = ctxFunc("Target")
	ctx = context.WithValue(ctx, ctxkey, "New Subscription")

	defer cancel()
	defer close(jobdb)

	fmt.Printf("\n +++++++---->  %+v\n", newSubscriptions)
	binding.Validator = &defaultValidator{}
	err := c.ShouldBindWith(&newSubscriptions, binding.Form)
	if err == nil {
		ID, ok := c.Get("userID")
		fmt.Println("-----> Got ID", ID.(string))
		if !ok {
			resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Invalid information", Err: fmt.Errorf("Cannot decode User ID from the request"), SuccessResp: successResp})
			c.JSON(http.StatusUnprocessableEntity, resp)
			return
		}

		stakeholder, ok := c.Get("userType")
		fmt.Println("-----> Got ID", ID.(string))
		if !ok {
			resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Invalid information", Err: fmt.Errorf("Cannot decode User Type from the request"), SuccessResp: successResp})
			c.JSON(http.StatusUnprocessableEntity, resp)
			return
		}

		//newSubscriptions.Publisher = ID.(string)
		if newSubscriptions.TransactionID == "" {
			newSubscriptions.TransactionID = "TX" + GetRandomID(15)
		}
		tknReq, bonusPercent := getSubPayment(newSubscriptions.PublishID)
		if tknReq != (newSubscriptions.BonusTokensUsed + newSubscriptions.PaidTokensUsed) {
			resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S5SUB", ErrTyp: "Token Amount error", Err: fmt.Errorf("Required Tokens are not equal to TokensUsed in parameters"), SuccessResp: successResp})
			c.JSON(http.StatusBadRequest, resp)
			return
		}
		if newSubscriptions.BonusTokensUsed > (tknReq / bonusPercent) {
			resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S5SUB", ErrTyp: "Token Amount error", Err: fmt.Errorf("Cannot use more than %v tokens for this transaction", (tknReq / bonusPercent)), SuccessResp: successResp})
			c.JSON(http.StatusBadRequest, resp)
			return
		}
		reqBody := map[string]string{"stakeholderID": ID.(string), "transactionID": newSubscriptions.TransactionID, "bonusTokensTransacted": fmt.Sprintf("%.2f", newSubscriptions.BonusTokensUsed), "paidTokensTransacted": fmt.Sprintf("%.2f", newSubscriptions.PaidTokensUsed)}
		resp, err := makeTokenServiceCall("/t/addTx", reqBody)

		if err != nil {
			resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S6PJ", ErrTyp: "Error while interacting with Tokens service", Err: fmt.Errorf("%v , %v", err, resp), SuccessResp: successResp})
			c.JSON(http.StatusUnprocessableEntity, resp)
			c.Abort()
			return
		}
		fmt.Println("==================== token resp ======", resp)
		go func() {
			select {
			case insertJobChan := <-newSubscriptions.Insert(stakeholder.(string), ID.(string)):
				jobdb <- insertJobChan
			case <-ctx.Done():
				return
			}
		}()
		insertJob := <-jobdb
		fmt.Printf("\n insertjob: %+v\n", insertJob)
		if insertJob.ErrTyp != "000" {
			resp := ErrCheck(ctx, insertJob)
			c.Error(insertJob.Err)
			c.JSON(http.StatusInternalServerError, resp)
			return
		}

		c.JSON(http.StatusOK, models.SubSuccessResp{"Successfully subscribed"})
		return
	}
	resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Required information not found", Err: err, SuccessResp: successResp})
	c.JSON(http.StatusUnprocessableEntity, resp)
	return
}

// GetSubscriptionPayment ...
func GetSubscriptionPayment(c *gin.Context) {
	successResp = map[string]string{}

	ctx, cancel := context.WithCancel(context.Background())
	ctxkey = ctxFunc("Target")
	ctx = context.WithValue(ctx, ctxkey, "New Subscription")
	defer cancel()

	pubID, ok := c.Params.Get("publishID")
	if !ok {
		resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Required information not found", Err: fmt.Errorf("Publish ID not found in Params"), SuccessResp: successResp})
		c.JSON(http.StatusUnprocessableEntity, resp)
		return
	}
	subsPayment := models.SubscriptionPaymentModel{"Payment details are not implemented, sending sample tokens required for " + pubID, 40.00, 10}
	c.JSON(http.StatusOK, subsPayment)
	return
}

// GetAllSubscriptions ...
func GetAllSubscriptions(c *gin.Context) {
	successResp = map[string]string{}
	jobdb := make(chan models.DbModelError, 1)
	ctx, cancel := context.WithCancel(context.Background())
	ctxkey = ctxFunc("Target")
	ctx = context.WithValue(ctx, ctxkey, "Get All Subscriptions")
	defer cancel()
	defer close(jobdb)

	ID, ok := c.Get("userID")
	fmt.Println("-----> Got ID", ID.(string))
	if !ok {
		resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Invalid information", Err: fmt.Errorf("Cannot decode User ID from the request"), SuccessResp: successResp})
		c.JSON(http.StatusUnprocessableEntity, resp)
		return
	}

	stakeholder, ok := c.Get("userType")
	fmt.Println("-----> Got ID", stakeholder.(string))
	if !ok {
		resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Invalid information", Err: fmt.Errorf("Cannot decode User Type from the request"), SuccessResp: successResp})
		c.JSON(http.StatusUnprocessableEntity, resp)
		return
	}
	var subs models.AllSubscriptionsModel
	go func() {
		select {
		case insertJobChan := <-subs.GetAllSubscriptions(ID.(string), stakeholder.(string)):
			jobdb <- insertJobChan
		case <-ctx.Done():
			return
		}
	}()
	insertJob := <-jobdb
	fmt.Printf("\n insertjob: %+v\n", insertJob)
	if insertJob.ErrTyp != "000" {
		resp := ErrCheck(ctx, insertJob)
		c.Error(insertJob.Err)
		c.JSON(http.StatusInternalServerError, resp)
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, subs.Subscriptions)
	c.Abort()
	return
}

// getSubPayment ...
func getSubPayment(pID string) (float64, float64) {

	return 40, 10
}

func makeTokenServiceCall(endpoint string, reqData map[string]string) (string, error) {
	tokenConfig := configuration.TokenConfig()
	resBody, err := middleware.MakeInternalServiceCall(tokenConfig.Host, tokenConfig.Port, "POST", endpoint, reqData)
	fmt.Println(err, "==========================")
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
