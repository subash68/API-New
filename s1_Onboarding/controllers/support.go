package controllers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/jaswanth-gorripati/PGK/s1_Onboarding/models"
)

// AddSupport ...
func AddSupport(c *gin.Context) {
	successResp = map[string]string{}
	var supportInfo models.SupportDataModel

	ctx, cancel := context.WithCancel(context.Background())
	ctxkey = ctxFunc("Target")
	ctx = context.WithValue(ctx, ctxkey, "SignUp")

	defer cancel()
	binding.Validator = &defaultValidator{}
	err := c.ShouldBindWith(&supportInfo, binding.Form)
	if err == nil {
		ID, ok := c.Get("userID")

		if !ok {
			resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S1AUT", ErrTyp: "Invalid information", Err: fmt.Errorf("Cannot decode User ID from the request"), SuccessResp: successResp})
			c.JSON(http.StatusUnprocessableEntity, resp)
			c.Abort()
			return
		}
		fmt.Println("-----> Got ID", ID.(string))
		userType, ok := c.Get("userType")
		if !ok {
			resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S1AUT", ErrTyp: "Invalid information", Err: fmt.Errorf("Cannot decode User Type from the request"), SuccessResp: successResp})
			c.JSON(http.StatusUnprocessableEntity, resp)
			c.Abort()
			return
		}
		supportInfo.StakeholderID = ID.(string)
		supportInfo.StakeholderRole = userType.(string)
		err := supportInfo.Insert()
		if err != nil {
			resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S1AUT", ErrTyp: "Internal Server Error", Err: fmt.Errorf("Failed to Add support Due to %s", err.Error()), SuccessResp: successResp})
			c.JSON(http.StatusInternalServerError, resp)
			c.Abort()
			return
		}
		c.JSON(http.StatusOK, SucResp{"Thank you for contacting, we will call you shortly"})
		return
	}
	resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S1AUT", ErrTyp: "Required information not found ", Err: err, SuccessResp: successResp})
	c.JSON(http.StatusUnprocessableEntity, resp)
	c.Abort()
	return
}
