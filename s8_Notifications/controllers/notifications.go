package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/jaswanth-gorripati/PGK/s8_Notifications/models"
)

// NftController ...
var (
	NftController nftController = nftController{}
)

type nftController struct{}

// AddNewNotification ...
func (nft *nftController) AddNewNotification(c *gin.Context) {
	ctx, _, _, successResp := getFuncReq(c, "Add Notification")
	binding.Validator = &defaultValidator{}
	var newNft models.NotificationsModel
	err := c.ShouldBindWith(&newNft, binding.Form)
	if err == nil {
		nftID, err := models.NftPersistance.AddNotification(newNft)
		if err != nil {
			resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Failed to Process request", Err: err, SuccessResp: successResp})
			c.JSON(http.StatusInternalServerError, resp)
			c.Abort()
			return
		}
		fmt.Println("\n==============>>> NEW NFT ID ---- %s \n", nftID)
		c.JSON(http.StatusOK, models.NftMessageResp{"Notification Saved", nftID})
		c.Abort()
		return
	}
	resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Required information not found", Err: err, SuccessResp: successResp})
	c.JSON(http.StatusUnprocessableEntity, resp)
	return
}

// GetAllNotification ...
func (nft *nftController) GetAllNotification(c *gin.Context) {
	ctx, ID, userType, successResp := getFuncReq(c, "Get Notifications")

	page, err := strconv.Atoi(c.Param("page"))
	if page <= 0 || err != nil {
		resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Invalid information", Err: fmt.Errorf("Cannot find Page number in params"), SuccessResp: successResp})
		c.JSON(http.StatusUnprocessableEntity, resp)
		return
	}
	size, err := strconv.Atoi(c.Param("perPage"))
	if size < 0 || err != nil {
		resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Invalid information", Err: fmt.Errorf("Cannot find Size of the page in params"), SuccessResp: successResp})
		c.JSON(http.StatusUnprocessableEntity, resp)
		return
	}
	filter := ""
	if c.Query("senderID") != "" {
		filter += " AND Initiator_Stakeholder_ID='" + c.Query("senderID") + "'"
	}
	if c.Query("senderUserRole") != "" {
		filter += " AND SenderUserRole='" + c.Query("senderUserRole") + "'"
	}
	if c.Query("notificationType") != "" {
		filter += " AND NotificationType='" + c.Query("notificationType") + "'"
	}
	if c.Query("content") != "" {
		filter += " AND Notification_Content LIKE '%" + c.Query("content") + "%'"
	}
	if c.Query("publishFlag") != "" {
		filter += " AND PublishFlag=" + c.Query("publishFlag") + ""
	}

	notifications, err := models.NftPersistance.GetAllNotifications(ID, filter, page, size, userType)
	if err != nil {
		resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Failed to Get Notifications", Err: err, SuccessResp: successResp})
		c.JSON(http.StatusInternalServerError, resp)
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, notifications)
	c.Abort()
	return
}

// GetNotificationByID ...
func (nft *nftController) GetNotificationByID(c *gin.Context) {
	ctx, ID, _, successResp := getFuncReq(c, "Get Notifications")

	nftID := c.Param("nftID")
	if nftID == "" {
		resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Invalid information", Err: fmt.Errorf("Cannot find nftID"), SuccessResp: successResp})
		c.JSON(http.StatusUnprocessableEntity, resp)
		return
	}
	notification, err := models.NftPersistance.GetNotificationByID(ID, nftID)
	if err != nil {
		resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Failed to Get Notifications", Err: err, SuccessResp: successResp})
		c.JSON(http.StatusInternalServerError, resp)
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, notification)
	c.Abort()
	return
}
