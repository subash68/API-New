// Package routers provide the routes for the application api
package routers

// Imports
import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jaswanth-gorripati/PGK/s1_Onboarding/controllers"
	"github.com/jaswanth-gorripati/PGK/s1_Onboarding/middleware"
)

// InitialzeRoutes : initalizing routes to the Campus recruit application API
func InitialzeRoutes() *gin.Engine {

	// Setting Release mode in GIN
	gin.SetMode(gin.ReleaseMode)

	// Declaring and assigning router as gin default
	router := gin.Default()

	// Adding logger to the console, Prints the request URL details
	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {

		// printing URL parameters
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			// Client IP
			param.ClientIP,

			// Date and time of the URL request
			param.TimeStamp.Format(time.RFC1123),

			// Method (GET / POST / PUT / PATCH )
			param.Method,

			// URL Path
			param.Path,

			// Requested Protocol (http / https)
			param.Request.Proto,

			// Status code
			param.StatusCode,

			// Latency of the client
			param.Latency,

			// User agent of the client
			param.Request.UserAgent(),

			// Error message
			param.ErrorMessage,
		)
	}))

	// Allow all origins for dev
	// router.Use(cors.Default())
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "HEAD", "OPTIONS", "POST", "PUT", "PATCH", "DELETE"},
		AllowHeaders:     []string{"Origin", "X-Requested-With", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// Catching if any errors happens in the api call
	router.Use(gin.Recovery())

	// Test Route URL

	onboard := router.Group("/o")
	onboard.GET("/", func(c *gin.Context) {
		c.Header("Title", "Campus Hiring")
		c.JSON(http.StatusOK, "Campus Hiring API is working")
	})
	onboard.POST("/signUp/:stakeholder", controllers.Signup)
	onboard.POST("/verifyOTP", controllers.CommonOTPVerifier)
	onboard.POST("/verifyMobile", controllers.VerifyMobile)
	onboard.POST("/changeAccStatus", middleware.AuthorizeRequest(), middleware.AuthorizeInternalRequest("INTERNAL", []string{"S7"}), controllers.ActivateAccount)
	onboard.POST("/verifyEmail", controllers.VerifyEmail)
	onboard.POST("/resendOtp", controllers.ResendOTP)
	onboard.POST("/login", controllers.Login)
	onboard.POST("/sendOTP", controllers.SendChangePasswordOTP)
	onboard.POST("/resetPassword", controllers.ChangePassword)
	onboard.POST("/logout", middleware.AuthorizeRequest(), controllers.Logout)
	onboard.POST("/addSupport", middleware.AuthorizeRequest(), controllers.AddSupport)
	// users := router.Group("/u")
	// users.Use(middleware.AuthorizeRequest())
	// users.GET("/getPayment", controllers.GetPayment)
	//router.GET("/vrfAccount", controllers.VerifyAccount)

	// Starting the server with address specified
	//router.Run("0.0.0.0:8080")
	return router
}
