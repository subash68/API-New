// Package middleware ...
package middleware

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jaswanth-gorripati/PGK/s8_Notifications/configuration"
)

// TokenClaims ...
type TokenClaims struct {
	TokenType string `json:"tokenType"`
	UserType  string `json:"userType"`
	UserID    string `json:"userId"`
}

// AuthorizeRequest ...
func AuthorizeRequest() gin.HandlerFunc {
	return func(c *gin.Context) {
		clientToken := c.Request.Header.Get("Authorization")
		if clientToken == "" {
			c.JSON(http.StatusUnauthorized, "No Authorization header provided")
			c.Abort()
			return
		}

		extractedToken := strings.Split(clientToken, "Bearer ")
		fmt.Printf("\ntoken: %+v\n", extractedToken)
		if len(extractedToken) == 2 {
			clientToken = strings.TrimSpace(extractedToken[1])
		} else {
			c.JSON(400, "Incorrect Format of Authorization Token")
			c.Abort()
			return
		}
		authService := configuration.AuthConfig()
		endpoint := "http://" + authService.Host + ":" + strconv.Itoa(authService.Port) + "/a/verify"
		data := url.Values{}
		data.Set("token", clientToken)

		client := &http.Client{}
		r, err := http.NewRequest("POST", endpoint, strings.NewReader(data.Encode())) // URL-encoded payload
		if err != nil {
			c.JSON(http.StatusInternalServerError, "Cannot create token authorization request "+err.Error())
			c.Abort()
			return
		}
		r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		r.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

		res, err := client.Do(r)
		if err != nil && strings.Contains(err.Error(), "refused") {
			fmt.Printf("\n %+v \n", err)
			c.JSON(http.StatusInternalServerError, "Authorization service is not active")
			c.Abort()
			return
		} else if err != nil {
			c.JSON(http.StatusUnauthorized, "Invalid token "+err.Error())
			c.Abort()
			return
		}
		defer res.Body.Close()
		body, err := ioutil.ReadAll(res.Body)
		var tokenClaims TokenClaims
		err = json.Unmarshal(body, &tokenClaims)
		if err != nil {
			c.JSON(http.StatusUnauthorized, "Invalid token ")
			c.Abort()
			return
		}
		// if tokenClaims.UserType != stakeholderType {
		// 	c.JSON(http.StatusUnauthorized, "Unauthorized request")
		// 	c.Abort()
		// 	return
		// }
		c.Set("token", clientToken)
		c.Set("tokenType", tokenClaims.TokenType)
		c.Set("userType", tokenClaims.UserType)
		c.Set("userID", tokenClaims.UserID)
		c.Next()
		return
	}
}

// RestrictContentType ...
func RestrictContentType(contentType string) gin.HandlerFunc {
	return func(c *gin.Context) {
		ct := c.GetHeader("content-type")
		ct = strings.Split(ct, ";")[0]
		fmt.Println("content type :", ct)
		if ct != contentType {
			c.JSON(http.StatusNotAcceptable, "Content type "+contentType+" is required")
			c.Abort()
			return
		}
		c.Next()
		return
	}
}

// AuthorizeInternalRequest ...
func AuthorizeInternalRequest(tknType string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// userType, ok := c.Get("userType")
		// if !ok {
		// 	c.JSON(http.StatusUnauthorized, "Unauthorized request, Only internal calls allowed")
		// 	c.Abort()
		// 	return
		// }

		tokenType, ok := c.Get("tokenType")
		if !ok {
			c.JSON(http.StatusUnauthorized, "Unauthorized request, Only internal calls allowed")
			c.Abort()
			return
		}
		// sort.Sort(sort.StringSlice(usrTyps))
		// fmt.Println("====================>>", userType.(string), tokenType.(string), sort.StringSlice(usrTyps).Search(userType.(string)) >= 0, tokenType.(string) == tknType, " <<========================")
		//if sort.StringSlice(usrTyps).Search(userType.(string)) < 0 || tokenType.(string) != tknType {
		if tokenType.(string) != tknType {
			c.JSON(http.StatusUnauthorized, "Unauthorized request, Only internal calls allowed")
			c.Abort()
			return
		}
		c.Next()
		return
	}
}
