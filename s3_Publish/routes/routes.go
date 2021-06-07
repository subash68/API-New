// Package routers provide the routes for the application api
package routers

// Imports
import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jaswanth-gorripati/PGK/s3_Publish/controllers"
	"github.com/jaswanth-gorripati/PGK/s3_Publish/middleware"
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
	//
	router.MaxMultipartMemory = 2 << 20 // Max 2mb files

	// Test Route URL
	router.GET("/", func(c *gin.Context) {
		c.Header("Title", "Campus Hiring")
		c.JSON(http.StatusOK, "Campus Hiring API is working")
	})

	publish := router.Group("/p")

	subData := publish.Group("/subData")
	subData.Use(middleware.AuthorizeRequest(""))

	subData.GET("/publishedData/:publishID", controllers.GetCrpSubscribedData)
	subData.GET("/nftData", controllers.NftDataController.GetNftData)

	corporate := publish.Group("/crp")
	corporate.Use(middleware.AuthorizeRequest("Corporate"))

	student := publish.Group("/stu")
	student.Use(middleware.AuthorizeRequest("Student"))
	stuPublish := student.Group("/publish")

	hiringCriteria := corporate.Group("/hiringCriteria")
	//hiringCriteria.Use()
	jobsCreation := corporate.Group("/createJob")
	//jobsCreation.Use(middleware.RestrictContentType())
	publishJob := corporate.Group("/publishJob")

	otherInfo := corporate.Group("/oi")

	university := publish.Group("/unv")
	university.Use(middleware.AuthorizeRequest("University"))

	proposal := university.Group("/proposal")
	unvPublish := university.Group("/publish")

	// unvPublish := university.Group("/publish")

	// HIRING CRITERIA
	hiringCriteria.POST("/", controllers.AddHiringCriteria)
	hiringCriteria.POST("/publish", controllers.PublishHiringCriteria)
	hiringCriteria.GET("/published", controllers.GetAllPublishedHC)
	hiringCriteria.GET("/getByID/:hcID", controllers.GetHiringCriteriaByID)
	hiringCriteria.GET("/all", controllers.GetAllHiringCriteria)
	hiringCriteria.PATCH("/:hcID", controllers.UpdateHiringCriteria)
	hiringCriteria.DELETE("/:hcID", controllers.DeleteHiringCriteria)

	// CREATE JOB
	//  middleware.RestrictContentType("multipart/form-data"),
	jobsCreation.POST("/", controllers.AddJobsCreation)
	jobsCreation.POST("/addSkills/:jobID", controllers.AddSkills)
	jobsCreation.PATCH("/skill/:jobID/:id", controllers.UpdateJobSkill)
	jobsCreation.GET("/getByID/:jobID", controllers.GetJobsCreationByID)
	jobsCreation.GET("/all", controllers.GetAllJobsCreated)
	jobsCreation.PATCH("/job/:jobID", controllers.UpdateJobsCreation)
	jobsCreation.PATCH("/mapHC/:jobID", controllers.MapJobToHC)
	jobsCreation.DELETE("/job/:jobID", controllers.DeleteJobsCreation)
	jobsCreation.DELETE("/skill/:jobID/:id", controllers.DeleteJobSkill)

	// PUBLISH JOB
	publishJob.POST("/", controllers.AddPublishedJobs)
	publishJob.GET("/getByID/:pjID", controllers.GetPublishedJobsByID)
	publishJob.GET("/all", controllers.GetAllPublishedJobs)
	publishJob.PATCH("/:pjID", controllers.UpdatePublishedJobs)
	publishJob.DELETE("/:pjID", controllers.DeletePublishedJobs)

	// PUBLISH CORPORATE PROFILE
	corporate.POST("/publish/profile", controllers.ProfilePublish)

	//OTHER INFORMATION
	otherInfo.POST("/", controllers.AddOtherInfo)
	otherInfo.GET("/", controllers.GetAllOI)
	otherInfo.GET("/published", controllers.GetAllPublishedOI)
	otherInfo.POST("/publish/:id", controllers.PublishOI)

	corporate.GET("/publish/all", controllers.GetAllCrpPublishedData)
	corporate.GET("/publish/id/:publishID", controllers.GetCrpPublishedDataByID)

	// University publish
	proposal.POST("/", controllers.AddNewProposal)
	proposal.GET("/", controllers.GetUnvProposal)
	proposal.DELETE("/program/:programID", controllers.DeleteProgramByID)
	proposal.DELETE("/branch/:branchID", controllers.DeleteBranchByID)
	proposal.DELETE("/accredations/:accredationsID", controllers.DeleteAccredationsByID)
	proposal.DELETE("/coes/:coesID", controllers.DeleteCoesByID)
	proposal.DELETE("/ranking/:rankingID", controllers.DeleteRankingByID)
	proposal.DELETE("/tieups/:tieupsID", controllers.DeleteTieupsByID)
	proposal.DELETE("/specialOfferings/:specialOfferingsID", controllers.DeleteSpecialOfferingsByID)
	proposal.DELETE("/otherInfo/:otherInfoID", controllers.DeleteOtherInfoByID)
	unvPublish.POST("/profile", controllers.PublishProfile)
	unvPublish.POST("/oi", controllers.PublishUnvOI)
	unvPublish.GET("/oi", controllers.GetPublishedUnvOI)
	unvPublish.GET("/all", controllers.GetPublishedUnvData)
	unvPublish.GET("/getByID/:publishID", controllers.GetUnvPublishedDataByID)

	stuPublish.POST("/profile", controllers.PublishStudentProfile)
	stuPublish.POST("/oi", controllers.PublishStudentOI)
	stuPublish.GET("/oi", controllers.GetAllStudentOI)
	stuPublish.GET("/all", controllers.GetStudentPublishedData)
	stuPublish.GET("/getByID/:publishID", controllers.GetStudentPublishedDataByID)

	return router
}
