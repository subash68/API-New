package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jaswanth-gorripati/PGK/s6_Token/models"
)

// GetTokenBalanceForID ...
func GetTokenBalanceForID(c *gin.Context) {
	ctx, ID, _, _ := getFuncReq(c, "Get Token Balance")

	var tokenBalance models.TokenBalanceModel
	err := tokenBalance.GetGenTokenBalanceByID(ID)

	if err != nil {
		resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S6TKN003", ErrTyp: "Invalid File", Err: fmt.Errorf("Failed to get token balance %s", err.Error()), SuccessResp: successResp})
		c.JSON(http.StatusBadRequest, resp)
		return
	}
	c.JSON(http.StatusOK, tokenBalance)
	return
}
