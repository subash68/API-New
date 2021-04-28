// Package routers provide the routes for the application api
package routers

// Imports
import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jaswanth-gorripati/PGK/s4_Profile/controllers"
	"github.com/jaswanth-gorripati/PGK/s4_Profile/middleware"
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
	router.MaxMultipartMemory = 1 << 20 // Max 1mb files

	// Test Route URL
	router.GET("/", func(c *gin.Context) {
		c.Header("Title", "Campus Hiring")
		c.JSON(http.StatusOK, "Campus Hiring API is working")
	})

	user := router.Group("/u")
	// user.Use(middleware.AuthorizeRequest())

	profile := user.Group("/profile")
	profile.Use(middleware.AuthorizeRequest())
	//profile.Use(middleware.RestrictContentType("multipart/form-data"))

	profilePicture := user.Group("/profilePic")
	profilePicture.Use(middleware.AuthorizeRequest())

	lut := user.Group("/lut")

	// middleware.RestrictContentType("multipart/form-data"),
	profile.GET("/", controllers.GetProfile)
	profile.PATCH("/", controllers.UpdateProfile)
	profilePicture.POST("/", controllers.UploadProfilePic)
	profilePicture.GET("/", controllers.GetProfilePic)
	lut.GET("/", controllers.GetLutData)

	// Search profiles

	corporates := user.Group("/crp")
	corporates.Use(middleware.AuthorizeRequest())

	corporates.GET("/search", controllers.SearchCrp)
	corporates.GET("/search/:corporateID", controllers.GetCrpByID)

	unv := user.Group("/unv")
	unv.Use(middleware.AuthorizeRequest())

	unv.GET("/search", controllers.SearchUnv)
	unv.GET("/search/:unvID", controllers.GetUnvByID)

	stu := user.Group("/stu")
	stu.Use(middleware.AuthorizeRequest())
	stu.POST("/academics", controllers.AddAcademics)
	stu.GET("/academics", controllers.GetAcademics)

	// Student Languages
	stu.POST("/language", controllers.AddLanguages)
	stu.GET("/language", controllers.GetAllLanguages)
	stu.PATCH("/language/:id", controllers.UpdateLanguage)
	stu.DELETE("/language/:id", controllers.DeleteLanguage)

	// Student Certifications
	stu.POST("/certs", controllers.AddCert)
	stu.GET("/certs", controllers.GetAllCerts)
	stu.PATCH("/certs/:id", controllers.UpdateCert)
	stu.DELETE("/certs/:id", controllers.DeleteCert)

	// Student Assessment
	stu.POST("/assessment", controllers.AddAssessment)
	stu.GET("/assessment", controllers.GetAllAssessments)
	stu.PATCH("/assessment/:id", controllers.UpdateAssessment)
	stu.DELETE("/assessment/:id", controllers.DeleteAssessment)

	// Student Internship
	stu.POST("/internship", controllers.AddInternship)
	stu.GET("/internship", controllers.GetAllInternships)
	stu.PATCH("/internship/:id", controllers.UpdateInternship)
	stu.DELETE("/internship/:id", controllers.DeleteInternship)

	// Student Awards
	stu.POST("/awards", controllers.StudentAwards.AddAwards)
	stu.GET("/awards", controllers.StudentAwards.GetAllAwards)
	stu.PATCH("/awards/:id", controllers.StudentAwards.UpdateAwards)
	stu.DELETE("/awards/:id", controllers.StudentAwards.DeleteAwards)

	// Student Competition
	stu.POST("/competition", controllers.StudentCompetitions.AddCompetitions)
	stu.GET("/competition", controllers.StudentCompetitions.GetAllCompetitions)
	stu.PATCH("/competition/:id", controllers.StudentCompetitions.UpdateCompetitions)
	stu.DELETE("/competition/:id", controllers.StudentCompetitions.DeleteCompetitions)

	// Student ConferenceWorkshop
	stu.POST("/confWorkshop", controllers.StudentConfWorkshop.AddConfWorkshop)
	stu.GET("/confWorkshop", controllers.StudentConfWorkshop.GetAllConfWorkshop)
	stu.PATCH("/confWorkshop/:id", controllers.StudentConfWorkshop.UpdateConfWorkshop)
	stu.DELETE("/confWorkshop/:id", controllers.StudentConfWorkshop.DeleteConfWorkshop)

	// Student ExtraCurricular
	stu.POST("/extraCurricular", controllers.StudentExtraCurricular.AddExtraCurricular)
	stu.GET("/extraCurricular", controllers.StudentExtraCurricular.GetAllExtraCurricular)
	stu.PATCH("/extraCurricular/:id", controllers.StudentExtraCurricular.UpdateExtraCurricular)
	stu.DELETE("/extraCurricular/:id", controllers.StudentExtraCurricular.DeleteExtraCurricular)

	// Student Patents
	stu.POST("/patents", controllers.StudentPatents.AddPatents)
	stu.GET("/patents", controllers.StudentPatents.GetAllPatents)
	stu.PATCH("/patents/:id", controllers.StudentPatents.UpdatePatents)
	stu.DELETE("/patents/:id", controllers.StudentPatents.DeletePatents)

	// Student Projects
	stu.POST("/projects", controllers.StudentProjects.AddProjects)
	stu.GET("/projects", controllers.StudentProjects.GetAllProjects)
	stu.PATCH("/projects/:id", controllers.StudentProjects.UpdateProjects)
	stu.DELETE("/projects/:id", controllers.StudentProjects.DeleteProjects)

	// Student Publications
	stu.POST("/publications", controllers.StudentPublications.AddPublications)
	stu.GET("/publications", controllers.StudentPublications.GetAllPublications)
	stu.PATCH("/publications/:id", controllers.StudentPublications.UpdatePublications)
	stu.DELETE("/publications/:id", controllers.StudentPublications.DeletePublications)

	// Student Scholarship
	stu.POST("/scholarship", controllers.StudentScholarships.AddScholarships)
	stu.GET("/scholarship", controllers.StudentScholarships.GetAllScholarships)
	stu.PATCH("/scholarship/:id", controllers.StudentScholarships.UpdateScholarships)
	stu.DELETE("/scholarship/:id", controllers.StudentScholarships.DeleteScholarships)

	// Student complete profile
	stu.GET("/completeProfile", controllers.GetStudentProfile)

	return router
}
