package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Logout ...
func Logout(c *gin.Context) {
	fmt.Printf("\n -----> Logout started \n")
	token, ok := c.Get("token")
	if !ok {
		c.JSON(http.StatusInternalServerError, "Cannot decode token from request")
		return
	}
	fmt.Println("--> token : ", token.(string))
	resp, err := DeleteToken(token.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, resp)
	return
	return
}
