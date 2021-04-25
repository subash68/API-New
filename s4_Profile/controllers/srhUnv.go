// Package controllers ...
package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jaswanth-gorripati/PGK/s4_Profile/models"
)

// SearchUnv ...
func SearchUnv(c *gin.Context) {
	ctx, _, _, successResp := getFuncReq(c, "Search Universities")
	queryParams := c.Request.URL.Query()
	fmt.Println(queryParams.Get("universityName"))
	results, err := models.SearchUniversities(queryParams.Get("universityName"), queryParams.Get("hcID"), queryParams["skills"], queryParams["locations"], string(queryParams.Get("cutOff")))
	if err != nil {
		resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S4Profile", ErrTyp: "Invalid Search criteria", Err: err, SuccessResp: successResp})
		c.JSON(http.StatusUnprocessableEntity, resp)
		return
	}
	c.JSON(http.StatusOK, results.Universities)
	return
}

// GetUnvByID ...
func GetUnvByID(c *gin.Context) {
	ctx, ID, _, _ := getFuncReq(c, "Get University by ID")

	universityID := c.Param("unvID")
	if universityID == "" {
		resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Invalid information", Err: fmt.Errorf("Cannot find universityID"), SuccessResp: successResp})
		c.JSON(http.StatusUnprocessableEntity, resp)
		c.Abort()
		return
	}
	unvDb, err := models.GetUnvByID(universityID, ID)
	if err != nil {
		resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Failed to get Corporate information", Err: err, SuccessResp: successResp})
		c.JSON(http.StatusUnprocessableEntity, resp)
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, unvDb)
	return
}
