package controllers

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/jaswanth-gorripati/PGK/s5_Subcription/models"
)

// UnvStuDataController ...
var (
	UnvStuDataController unvStuDataController = unvStuDataController{}
)

type unvStuDataController struct{}

// SubscribeToStuData ...
func (usdc *unvStuDataController) SubscribeToStuData(c *gin.Context) {
	ctx, ID, _, successResp := getFuncReq(c, "Subscribe University Student Database")
	var usd models.UnvStuDataModel
	var usr models.UnvInsightSubsReqModel
	binding.Validator = &defaultValidator{}
	err := c.ShouldBindWith(&usr, binding.Form)
	if err == nil {
		usd.SubscribedStakeholderID = ID
		if usr.TransactionID == "" {
			usr.TransactionID = "TX" + GetRandomID(15)
		}
		tknReq, bonusPercent := getSubPayment(usr.SubscriberStakeholderID)
		if tknReq != (usr.BonusTokensUsed + usr.PaidTokensUsed) {
			resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S5SUB", ErrTyp: "Token Amount error", Err: fmt.Errorf("Required Tokens are not equal to TokensUsed in parameters"), SuccessResp: successResp})
			c.JSON(http.StatusBadRequest, resp)
			return
		}
		if usr.BonusTokensUsed > (tknReq / bonusPercent) {
			resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S5SUB", ErrTyp: "Token Amount error", Err: fmt.Errorf("Cannot use more than %v tokens for this transaction", (tknReq / bonusPercent)), SuccessResp: successResp})
			c.JSON(http.StatusBadRequest, resp)
			return
		}
		reqBody := map[string]string{"stakeholderID": ID, "transactionID": usr.TransactionID, "bonusTokensTransacted": fmt.Sprintf("%.2f", usr.BonusTokensUsed), "paidTokensTransacted": fmt.Sprintf("%.2f", usr.PaidTokensUsed)}
		resp, err := makeTokenServiceCall("/t/addTx", reqBody)

		if err != nil {
			resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S6PJ", ErrTyp: "Error while interacting with Tokens service", Err: fmt.Errorf("%v , %v", err, resp), SuccessResp: successResp})
			c.JSON(http.StatusUnprocessableEntity, resp)
			c.Abort()
			return
		}
		fmt.Println("==================== token resp ======", resp)
		usd.SubscriberStakeholderID = usr.SubscriberStakeholderID
		respUIM, err := usd.Subscribe()
		if err != nil {
			resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Internal Server Error", Err: err, SuccessResp: successResp})
			c.JSON(http.StatusInternalServerError, resp)
			return
		}
		c.JSON(http.StatusOK, respUIM)
		c.Abort()
		return
	}
	resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Required information not found", Err: err, SuccessResp: successResp})
	c.JSON(http.StatusUnprocessableEntity, resp)
	return
}

// QuerySubscribedStuData ...
func (usdc *unvStuDataController) QuerySubscribedStuData(c *gin.Context) {
	ctx, ID, _, successResp := getFuncReq(c, "Query University Student Database")
	var err error
	var unvStuQuery models.UnvStuDataQueryDataModel
	reqContentType := strings.Split(c.GetHeader("Content-Type"), ";")[0]
	if reqContentType != "application/json" || reqContentType == "" {
		resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S5Sub", ErrTyp: "Required information not found", Err: err, SuccessResp: successResp})
		c.JSON(http.StatusUnprocessableEntity, resp)
		c.Abort()
		return
	}
	binding.Validator = &defaultValidator{}

	err = c.ShouldBindWith(&unvStuQuery, binding.Default("POST", strings.Split(c.GetHeader("Content-Type"), ";")[0]))
	if err == nil {
		fmt.Printf("\n=======>>>==== %+v =========", unvStuQuery)
		var usd models.UnvStuDataModel
		usd.SubscribedStakeholderID = ID
		usd.SubscriberStakeholderID = unvStuQuery.SubscriberStakeholderID
		usd.SubscriptionID = unvStuQuery.SubscriptionID
		respData, err := usd.StoreStudentData("")
		if err != nil {
			resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S5Sub", ErrTyp: "Internal Server Error", Err: err, SuccessResp: successResp})
			c.JSON(http.StatusUnprocessableEntity, resp)
			c.Abort()
			return
		}
		c.JSON(http.StatusOK, respData)
		c.Abort()
		return
	}

	resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Required information not found", Err: err, SuccessResp: successResp})
	c.JSON(http.StatusUnprocessableEntity, resp)
	return

}

// FetchSubscribedStuData ...
func (usdc *unvStuDataController) FetchSubscribedStuData(c *gin.Context) {
	ctx, ID, _, successResp := getFuncReq(c, "Query University Student Database")
	var err error
	var usd models.UnvStuDataModel
	usd.SubscriptionID = c.Param("subID")
	if usd.SubscriptionID == "" {
		resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Invalid information", Err: fmt.Errorf("Cannot find subscription ID in params"), SuccessResp: successResp})
		c.JSON(http.StatusUnprocessableEntity, resp)
		return
	}
	usd.SubscribedStakeholderID = ID
	respData, err := usd.RetrieveStudentData()
	if err != nil {
		resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S5Sub", ErrTyp: "Internal Server Error", Err: err, SuccessResp: successResp})
		c.JSON(http.StatusUnprocessableEntity, resp)
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, respData)
	c.Abort()
	return
}
