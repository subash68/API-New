package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/jaswanth-gorripati/PGK/s5_Subcription/models"
)

// CampusDriveInvitationsController ....
var (
	CampusDriveInvitationsController campusDriveInvitationsController = campusDriveInvitationsController{}
	CDNotificationType               string                           = "CampusHiring Request"
	CDNotificationID                 string                           = "CHI"
	CDRNotificationType              string                           = "CampusHiring Response"
	CDRNotificationID                string                           = "CHR"
)

// Corporate ...
const (
	Corporate       string = "Corporate"
	University      string = "University"
	LastCrpID       string = "CORP_CD_Get_Last_ID"
	LastUnvID       string = "UNV_CD_Get_Last_ID"
	CrpCode         string = "CCDI"
	UnvCode         string = "UCPI"
	CrpInsCmd       string = "CORP_CD_INIT"
	UnvInsCmd       string = "UNV_CD_INIT"
	CrpInviteCmd    string = "CORP_CD_SUB_UPDATE"
	UnvInviteCmd    string = "UNV_CD_SUB_UPDATE"
	CrpRespondCmd   string = "CORP_CD_UNV_RESP"
	UnvRespondCmd   string = "UNV_CD_UNV_RESP"
	CrpGetByIDCmd   string = "CORP_CD_GET_BY_ID"
	UnvGetByIDCmd   string = "UNV_CD_GET_BY_ID"
	CrpGetIRByIDCmd string = "CORP_CD_GET_INITIATOR_FR_CD"
	UnvGetIRByIDCmd string = "UNV_CD_GET_INITIATOR_FR_CD"
)

type campusDriveInvitationsController struct{}

type cdRespModel struct {
	Message       string `json:"message,omitempty"`
	CampusDriveID string `json:"campusDriveID,omitempty"`
	EmailTo       string `json:"emailTo"`
	EmailFrom     string `json:"emailFrom"`
}

// NftMessageResp ...
type NftMessageResp struct {
	Message string `json:"message"`
	NftID   string `json::"nftID"`
}

// Subscribe ...
func (cdi *campusDriveInvitationsController) Subscribe(c *gin.Context) {
	ctx, ID, userType, successResp := getFuncReq(c, "Campus Drive Subscription")
	var usr models.CampusDriveSubInitModel
	binding.Validator = &defaultValidator{}
	err := c.ShouldBindWith(&usr, binding.Form)
	if err == nil {
		var lastQueryCmd, code, insCmd, pubUserType string
		//var initiatorEmail,receiverEmail string
		switch userType {
		case Corporate:
			lastQueryCmd, code, insCmd, pubUserType = LastCrpID, CrpCode, CrpInsCmd, "UNIV"
			break
		case University:
			lastQueryCmd, code, insCmd, pubUserType = LastUnvID, UnvCode, UnvInsCmd, "CORP"
			break
		default:
			resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S5SUB", ErrTyp: "Internal Server Error", Err: fmt.Errorf("Invalid user type %s", userType), SuccessResp: successResp})
			c.JSON(http.StatusInternalServerError, resp)
			return
		}
		var err error
		var cdm models.CampusDriveDataModel
		cdm.InitiatorID = ID
		cdm.CampusDriveID, err = models.CreateSudID(cdm.InitiatorID, lastQueryCmd, code)

		if err != nil {
			resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S5SUB", ErrTyp: "Internal Server Error", Err: err, SuccessResp: successResp})
			c.JSON(http.StatusInternalServerError, resp)
			return
		}

		if usr.TransactionID == "" {
			usr.TransactionID = "TX" + GetRandomID(15)
		}
		tknReq, bonusPercent := getSubPayment(ID)
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
		reqBody := map[string]string{"stakeholderID": ID, "transactionID": usr.TransactionID, "bonusTokensTransacted": fmt.Sprintf("%.2f", usr.BonusTokensUsed), "paidTokensTransacted": fmt.Sprintf("%.2f", usr.PaidTokensUsed), "publisherType": pubUserType, "publisherID": usr.ReceiverID, "subscriptionID": cdm.CampusDriveID, "subscriptionType": "CR"}
		resp, err := makeTokenServiceCall("/t/addTx", reqBody)

		if err != nil {
			resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S5SUB", ErrTyp: "Error while interacting with Tokens service", Err: fmt.Errorf("%v , %v", err, resp), SuccessResp: successResp})
			c.JSON(http.StatusUnprocessableEntity, resp)
			c.Abort()
			return
		}
		fmt.Println("==================== token res ======", resp)

		cdm.ReceiverID = usr.ReceiverID
		err = cdm.SubscribeToInviteForCD(lastQueryCmd, code, insCmd)

		if err != nil {
			resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S5SUB", ErrTyp: "Internal Server Error", Err: err, SuccessResp: successResp})
			c.JSON(http.StatusInternalServerError, resp)
			return
		}
		var cdResp cdRespModel
		cdResp.CampusDriveID = cdm.CampusDriveID
		if userType == "Corporate" {
			cdResp.EmailTo, cdResp.EmailFrom = models.GetEmailsForCH(ID, usr.ReceiverID)
		} else {
			cdResp.EmailTo, cdResp.EmailFrom = models.GetEmailsForCH(usr.ReceiverID, ID)
		}
		c.JSON(http.StatusOK, cdResp)
		c.Abort()
		return
	}
	resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S5SUB", ErrTyp: "Required information not found", Err: err, SuccessResp: successResp})
	c.JSON(http.StatusUnprocessableEntity, resp)
	return
}

// Invite ...
func (cdi *campusDriveInvitationsController) Invite(c *gin.Context) {
	ctx, ID, userType, successResp := getFuncReq(c, "Campus Drive Invite")
	var usr models.CampusDriveInviteEmailModel
	binding.Validator = &defaultValidator{}
	err := c.ShouldBindWith(&usr, binding.Form)
	if err == nil {
		var cdm models.CampusDriveDataModel
		cdm.CampusDriveID = usr.CampusDriveID
		i, r, err := cdm.GetIRByID(userType, ID, false)
		if i != ID {
			resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S6PJ", ErrTyp: "Unauthorized", Err: fmt.Errorf("Unauthorized request"), SuccessResp: successResp})
			c.JSON(http.StatusUnprocessableEntity, resp)
			c.Abort()
			return
		}
		var reqAsBytes []byte
		if userType == "University" {
			usr.UniversityDetails = models.GetUnvDetailsByID(i)
		} else if userType == "Corporate" {
			usr.CorporateDetails = models.GetCorpDetailsByID(i)
		}
		reqAsBytes, _ = json.Marshal(usr)
		reqBody := map[string]string{"senderID": ID, "senderUserRole": userType, "notificationType": CDNotificationType, "content": fmt.Sprint(string(reqAsBytes)), "publishFlag": "false", "publishID": "", "receiverID": r, "redirectedURL": "", "isGeneric": "false", "notificationTypeID": CDNotificationID}
		nftResp, err := makeNFTServiceCall("/nft/addNotification", reqBody)
		if err != nil {
			resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S6PJ", ErrTyp: "Error while interacting with Notification service", Err: fmt.Errorf("%v , %v", err, nftResp), SuccessResp: successResp})
			c.JSON(http.StatusUnprocessableEntity, resp)
			c.Abort()
			return
		}
		fmt.Printf("\n==========>>>>>>> NFT Response : %s \n", nftResp)

		cdm.InitiatorID = ID
		cdm.ReceiverID = r

		cdm.RequestedNftID = nftResp

		fmt.Printf("\n====== CDM ==== %+v ===\n", cdm)

		err = cdm.SendInvitationToReceiver(userType)
		if err != nil {
			resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S5SUB", ErrTyp: "Internal Server Error", Err: err, SuccessResp: successResp})
			c.JSON(http.StatusInternalServerError, resp)
			return
		}
		var cdResp cdRespModel
		cdResp.CampusDriveID = cdm.CampusDriveID
		cdResp.Message = "Invitation sent to Receiver"
		c.JSON(http.StatusOK, cdResp)
		c.Abort()
		return
	}
	resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S5SUB", ErrTyp: "Required information not found", Err: err, SuccessResp: successResp})
	c.JSON(http.StatusUnprocessableEntity, resp)
	return
}

// Respond ...
func (cdi *campusDriveInvitationsController) Respond(c *gin.Context) {
	ctx, ID, userType, successResp := getFuncReq(c, "Responding Campus Drive")
	var usr models.CampusDriveRespondDataModel
	binding.Validator = &defaultValidator{}
	err := c.ShouldBindWith(&usr, binding.Form)
	if err == nil {
		if usr.Accepted == false && usr.ReasonToReject == "" {
			resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S6PJ", ErrTyp: "Required field not found", Err: fmt.Errorf("Require reasonToReject field when rejecting"), SuccessResp: successResp})
			c.JSON(http.StatusUnprocessableEntity, resp)
			c.Abort()
			return
		}
		var cdm models.CampusDriveDataModel
		cdm.CampusDriveID = usr.CampusDriveID
		reqContent, err := models.GetContentByNftID(usr.NftID, ID)
		if err != nil {
			resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S6PJ", ErrTyp: "Unauthorized", Err: fmt.Errorf("Unauthorized or Invalid Notification ID " + usr.NftID), SuccessResp: successResp})
			c.JSON(http.StatusUnprocessableEntity, resp)
			c.Abort()
			return
		}

		reqAsBytes, _ := json.Marshal(usr)
		i, r, err := cdm.GetIRByID(userType, ID, true)
		if r != ID {
			resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S6PJ", ErrTyp: "Unauthorized", Err: fmt.Errorf("Unauthorized request"), SuccessResp: successResp})
			c.JSON(http.StatusUnprocessableEntity, resp)
			c.Abort()
			return
		}
		nftRespContent := struct {
			Request  string `json:"requestContent"`
			Response string `json:"responseContent"`
		}{
			Request:  strings.ReplaceAll(reqContent, "\\", ""),
			Response: strings.ReplaceAll(string(reqAsBytes), "\\", ""),
		}
		nftRespContentAsBytes, _ := json.Marshal(nftRespContent)
		reqBody := map[string]string{"senderID": ID, "senderUserRole": userType, "notificationType": CDRNotificationType, "content": strings.ReplaceAll(string(nftRespContentAsBytes), "\\", ""), "publishFlag": "false", "publishID": "", "receiverID": i, "redirectedURL": "", "isGeneric": "false", "notificationTypeID": CDRNotificationID}
		nftResp, err := makeNFTServiceCall("/nft/addNotification", reqBody)
		if err != nil {
			resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S6PJ", ErrTyp: "Error while interacting with Notification service", Err: fmt.Errorf("%v , %s", err, nftResp), SuccessResp: successResp})
			c.JSON(http.StatusUnprocessableEntity, resp)
			c.Abort()
			return
		}
		fmt.Printf("\n==========>>>>>>> NFT Response : %s \n", nftResp)

		if err != nil {
			resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S6PJ", ErrTyp: "Error while Decoding notification response", Err: fmt.Errorf("%v , %s", err, nftResp), SuccessResp: successResp})
			c.JSON(http.StatusInternalServerError, resp)
			c.Abort()
			return
		}

		cdm.InitiatorID = i
		cdm.ReceiverID = ID
		cdm.AccOrRejectNftID = nftResp
		cdm.Accepted = usr.Accepted
		cdm.ReasonToReject = usr.ReasonToReject
		fmt.Printf("\n====== CDM ==== %+v ===\n", cdm)
		err = cdm.Respond(userType)
		if err != nil {
			resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S5SUB", ErrTyp: "Internal Server Error", Err: err, SuccessResp: successResp})
			c.JSON(http.StatusInternalServerError, resp)
			return
		}
		var cdResp cdRespModel
		cdResp.CampusDriveID = cdm.CampusDriveID
		cdResp.Message = "Response Sent"
		c.JSON(http.StatusOK, cdResp)
		c.Abort()
		return
	}
	resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S5SUB", ErrTyp: "Required information not found", Err: err, SuccessResp: successResp})
	c.JSON(http.StatusUnprocessableEntity, resp)
	return
}

// GetAllCD ...
func (cdi *campusDriveInvitationsController) GetAllCD(c *gin.Context) {
	//ctx, ID, userType, successResp := getFuncReq(c, "Responding Campus Drive")
	return
}
