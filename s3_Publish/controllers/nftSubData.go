package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jaswanth-gorripati/PGK/s3_Publish/models"
)

// NftDataController ...
var (
	NftDataController nftDataController = nftDataController{}
)

type nftDataController struct{}

func (nd *nftDataController) GetNftData(c *gin.Context) {
	ctx, ID, userType, successResp := getFuncReq(c, "Add VolunteerExperience")
	publisherRole := c.Query("role")
	if publisherRole == "" {
		resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Invalid information", Err: fmt.Errorf("Cannot fing Hiring criteria ID"), SuccessResp: successResp})
		c.JSON(http.StatusUnprocessableEntity, resp)
		return
	}
	publishID := c.Query("publishID")
	if publishID == "" {
		resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Invalid information", Err: fmt.Errorf("Cannot fing Hiring criteria ID"), SuccessResp: successResp})
		c.JSON(http.StatusUnprocessableEntity, resp)
		return
	}
	resp, err := models.NftSubscriptionData.GetData(ID, userType, publisherRole, publishID)
	if err != nil {
		resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Internal Server Error", Err: err, SuccessResp: successResp})
		c.JSON(http.StatusUnprocessableEntity, resp)
		return
	}
	c.JSON(http.StatusOK, resp)
	c.Abort()
	return
}
