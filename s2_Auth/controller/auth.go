package controller

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/jaswanth-gorripati/PGK/s2_Auth/dto"
	"github.com/jaswanth-gorripati/PGK/s2_Auth/models"
	"github.com/jaswanth-gorripati/PGK/s2_Auth/services"
)

// LoginToken ...
func LoginToken(c *gin.Context) {
	fmt.Printf("\nGot request\n")
	var tokenDb models.TokenMasterDB
	binding.Validator = &defaultValidator{}
	err := c.ShouldBindWith(&tokenDb, binding.Form)
	fmt.Printf("\nErr: %+v\n", err)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Invalid Information ")
		return
	}
	tokenDb.AtExpires = time.Now().Add(time.Minute * 3600).Unix()
	token, err := services.CreateToken(&tokenDb)
	fmt.Printf("\ntoken: %v \n", token)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, dto.LoginTokenResp{Token: token})
	return
}

// CreateInternalToken ...
func CreateInternalToken(c *gin.Context) {
	fmt.Printf("\nGot request\n")
	var tokenDb models.TokenMasterDB
	binding.Validator = &defaultValidator{}
	err := c.ShouldBindWith(&tokenDb, binding.Form)
	fmt.Printf("\nErr: %+v\n", err)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Invalid Information ")
		return
	}
	tokenDb.TokenType = "INTERNAL"
	tokenDb.AtExpires = time.Now().Add(time.Minute * 2).Unix()
	token, err := services.CreateToken(&tokenDb)
	fmt.Printf("\ntoken: %v \n", token)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, dto.LoginTokenResp{Token: token})
	return
}

// AuthorizeToken ...
func AuthorizeToken(c *gin.Context) {
	var vrfTokenData dto.VrfToken
	binding.Validator = &defaultValidator{}
	err := c.ShouldBindWith(&vrfTokenData, binding.Form)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Cannot process required information for creating token")
		return
	}
	tokenClaims, err := services.VerifyToken(vrfTokenData.Token)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, tokenClaims)
	c.Abort()
	fmt.Println("======================= RETURNED FROM AUTH ================================")
	return
}

// DeleteAuthToken ...
func DeleteAuthToken(c *gin.Context) {
	fmt.Println("------>>>>>> Dleteing token")
	var vrfTokenData dto.VrfToken
	binding.Validator = &defaultValidator{}
	err := c.ShouldBindWith(&vrfTokenData, binding.Form)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Cannot process required information for Logout")
		return
	}
	err = services.DelToken(vrfTokenData.Token)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, "Token deleted")
	return
}
