package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/jaswanth-gorripati/PGK/s5_Subcription/models"
)

// HiringCriteriaController ...
var (
	HiringCriteriaController hiringCriteriaController = hiringCriteriaController{}
)

type hiringCriteriaController struct{}

// Subscribe ...
func (hcc *hiringCriteriaController) Subscribe(c *gin.Context) {
	ctx, ID, _, successResp := getFuncReq(c, "Subscribe Hiring Insight")
	var uim models.CrpHiringInsightsModel
	var usr models.CorpHiringInsightsSubsReqModel
	binding.Validator = &defaultValidator{}
	err := c.ShouldBindWith(&usr, binding.Form)
	if err == nil {
		uim.SubscribedStakeholderID = ID
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
		uim.SubscriberStakeholderID = usr.SubscriberStakeholderID
		respUIM, err := uim.Insert()
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

// GetHiringCriteriaByID ...
func (hcc *hiringCriteriaController) GetHiringCriteriaByID(c *gin.Context) {
	ctx, ID, _, successResp := getFuncReq(c, "Get Subscribed Hiring Insight")
	var uim models.CrpHiringInsightsModel

	uim.SubscribedStakeholderID = ID
	var ok bool
	uim.SubscriptionID, ok = c.Params.Get("subscriptionID")
	if !ok {
		resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Required information not found", Err: fmt.Errorf("subscriptionID not found in Params"), SuccessResp: successResp})
		c.JSON(http.StatusUnprocessableEntity, resp)
		return
	}
	respUIM, err := uim.GetHiringInsightBySubID()
	if err != nil {
		resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Internal Server Error", Err: err, SuccessResp: successResp})
		c.JSON(http.StatusInternalServerError, resp)
		return
	}
	c.JSON(http.StatusOK, respUIM)
	c.Abort()
	return
}

// GetAllHiringInsights ...
func (hcc *hiringCriteriaController) GetAllHiringInsights(c *gin.Context) {
	ctx, ID, _, successResp := getFuncReq(c, "Get Subscribed Hiring Insight")
	var uim models.CrpHiringInsightsModel

	uim.SubscribedStakeholderID = ID

	respUIM, err := uim.GetHiringInsightAll()
	if err != nil {
		resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Internal Server Error", Err: err, SuccessResp: successResp})
		c.JSON(http.StatusInternalServerError, resp)
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, respUIM)
	c.Abort()
	return
}
