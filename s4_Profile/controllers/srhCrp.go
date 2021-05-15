// Package controllers ...
package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jaswanth-gorripati/PGK/s4_Profile/models"
)

// SearchCrp ...
func SearchCrp(c *gin.Context) {
	ctx, _, _, successResp := getFuncReq(c, "Search Corporates")
	queryParams := c.Request.URL.Query()
	fmt.Println(queryParams.Get("cutOff"))
	results, err := models.SeacrhCorporate(queryParams.Get("corporateName"), queryParams["industry"], queryParams["skills"], queryParams["locations"], string(queryParams.Get("cutOff")))
	if err != nil {
		resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S4Profile", ErrTyp: "Invalid Search criteria", Err: err, SuccessResp: successResp})
		c.JSON(http.StatusUnprocessableEntity, resp)
		return
	}
	c.JSON(http.StatusOK, results.Corporates)
	return
}

// GetCrpByID ...
func GetCrpByID(c *gin.Context) {
	ctx, ID, userType, _ := getFuncReq(c, "Get Corporate by ID")

	corporateID := c.Param("corporateID")
	if corporateID == "" {
		resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Invalid information", Err: fmt.Errorf("Cannot find CorporateID"), SuccessResp: successResp})
		c.JSON(http.StatusUnprocessableEntity, resp)
		c.Abort()
		return
	}
	corpDb, err := models.GetCorpByID(corporateID, 5, ID, userType)
	if err != nil {
		resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Failed to get Corporate information", Err: err, SuccessResp: successResp})
		c.JSON(http.StatusUnprocessableEntity, resp)
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, corpDb)
	return
}
