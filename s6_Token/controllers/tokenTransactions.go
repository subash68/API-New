package controllers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/jaswanth-gorripati/PGK/s6_Token/models"
)

// AddTokenTransactionsForID ...
func AddTokenTransactionsForID(c *gin.Context) {
	jobdb := make(chan models.DbModelError, 1)
	successResp = map[string]string{}

	ctx, cancel := context.WithCancel(context.Background())
	ctxkey = ctxFunc("Target")
	ctx = context.WithValue(ctx, ctxkey, "Add Token Transaction")

	defer cancel()

	var tokenTx models.TokenTransactionsModel
	binding.Validator = &defaultValidator{}
	err := c.ShouldBindWith(&tokenTx, binding.Form)
	if err == nil {
		go func() {
			select {
			case insertJobChan := <-tokenTx.TokenTransactionsToID():
				jobdb <- insertJobChan
			case <-ctx.Done():
				return
			}
		}()
		insertJob := <-jobdb
		fmt.Printf("\n insertjob: %+v\n", insertJob)
		if insertJob.ErrTyp != "000" {
			resp := ErrCheck(ctx, insertJob)
			c.Error(insertJob.Err)
			c.JSON(http.StatusInternalServerError, resp)
			c.Abort()
			return
		}
		c.JSON(http.StatusOK, models.TokenDbResp{"Token Transaction Added"})
		return
	}
	resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S6PJ", ErrTyp: "Required information not found", Err: err, SuccessResp: successResp})
	c.JSON(http.StatusUnprocessableEntity, resp)
	return
}

// GetTokenTxsOfID ...
func GetTokenTxsOfID(c *gin.Context) {
	ctx, ID, _, _ := getFuncReq(c, "Get Token Transactions")

	var tokenTxs models.TokenTransactions
	err := tokenTxs.GetTokenTransactionsForID(ID)

	if err != nil {
		resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S6TKN003", ErrTyp: "INTERNAL SERVER ERROR", Err: fmt.Errorf("Failed to get Token Transaction details %s", err.Error()), SuccessResp: successResp})
		c.JSON(http.StatusBadRequest, resp)
		return
	}
	c.JSON(http.StatusOK, tokenTxs.Transactions)
	return
}
