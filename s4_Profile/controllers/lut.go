// Package controllers

package controllers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jaswanth-gorripati/PGK/s4_Profile/models"
)

// LutsReq ...
type LutsReq struct {
	LutList []string `form:"lutList" binding:"required"`
}

// GetLutData ...
func GetLutData(c *gin.Context) {
	successResp = map[string]string{}
	jobdb := make(chan models.DbModelError, 1)

	ctx, cancel := context.WithCancel(context.Background())
	ctxkey = ctxFunc("Target")
	ctx = context.WithValue(ctx, ctxkey, "Lut data")

	defer cancel()
	defer close(jobdb)
	var lr LutsReq
	fmt.Printf("\n%+v\n", c.QueryArray("lutList"))
	lr.LutList = c.QueryArray("lutList")

	if len(lr.LutList) > 0 {
		var ld models.LutResponse
		go func() {
			select {
			case insertJobChan := <-ld.GetData(lr.LutList):
				jobdb <- insertJobChan
			case <-ctx.Done():
				return
			}
		}()
		insertJob := <-jobdb

		if insertJob.ErrTyp != "000" {
			resp := ErrCheck(ctx, insertJob)
			c.Error(insertJob.Err)
			c.JSON(http.StatusInternalServerError, resp)
			c.Abort()
			return
		}
		c.JSON(http.StatusOK, ld)
		c.Abort()
		return
	}
	resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S4PRF", ErrTyp: "Required information not found", Err: fmt.Errorf("Required lutList in query params"), SuccessResp: successResp})
	c.JSON(http.StatusUnprocessableEntity, resp)
	c.Abort()
	return
}
