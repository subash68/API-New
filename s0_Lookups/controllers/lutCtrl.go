package controllers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jaswanth-gorripati/PGK/s0_Lookups/models"
)

// LutController ...
var (
	LutController lutController = lutController{}
)

type lutController struct{}

// GetLookUpData ...
func (lc *lutController) GetLookUpData(c *gin.Context) {
	successResp = map[string]string{}
	ctx, cancel := context.WithCancel(context.Background())
	ctxkey = ctxFunc("Target")
	ctx = context.WithValue(ctx, ctxkey, "LUT data")

	defer cancel()
	log.Debug("c query %v", c.QueryArray("lutList"))
	lutList := c.QueryArray("lutList")
	log.Debugf("lut query %v ", lutList)
	if len(lutList) > 0 {
		var ald models.AllLutData
		err := ald.GetAllLutData(lutList)
		if err != nil {
			if fmt.Sprintf("%v", err) == "Internal" {
				c.JSON(http.StatusInternalServerError, "Try again")
				c.Abort()
				return
			}
			resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S4PRF", ErrTyp: "Error while getting data", Err: err, SuccessResp: successResp})
			c.JSON(http.StatusBadRequest, resp)
			return
		}
		c.JSON(http.StatusOK, ald)
		c.Abort()
		return
	}
	resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S0LUT", ErrTyp: "Required information not found", Err: fmt.Errorf("Required lutList in query params"), SuccessResp: successResp})
	c.JSON(http.StatusUnprocessableEntity, resp)
	c.Abort()
	return
	return
}
