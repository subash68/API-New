package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jaswanth-gorripati/PGK/s4_Profile/models"
)

// GetStudentProfile ...
func GetStudentProfile(c *gin.Context) {
	ctx, ID, _, successResp := getFuncReq(c, "Get Student Profile")
	var stuData models.StudentCompleteProfileDataModel
	info, dbError := models.GetProfile(ID, "STU_GET_PROFILE")
	if dbError.ErrTyp != "000" {
		resp := ErrCheck(ctx, dbError)
		c.Error(dbError.Err)
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	stuDb := models.StudentMasterDb{}
	err := info.Scan(&stuDb.StakeholderID, &stuDb.FirstName, &stuDb.MiddleName, &stuDb.LastName, &stuDb.PersonalEmail, &stuDb.CollegeEmail, &stuDb.PhoneNumber, &stuDb.AlternatePhoneNumber, &stuDb.CollegeID, &stuDb.Gender, &stuDb.DateOfBirth, &stuDb.AadharNumber, &stuDb.PermanentAddressLine1, &stuDb.PermanentAddressLine2, &stuDb.PermanentAddressLine3, &stuDb.PermanentAddressCountry, &stuDb.PermanentAddressState, &stuDb.PermanentAddressCity, &stuDb.PermanentAddressDistrict, &stuDb.PermanentAddressZipcode, &stuDb.PermanentAddressPhone, &stuDb.PresentAddressLine1, &stuDb.PresentAddressLine2, &stuDb.PresentAddressLine3, &stuDb.PresentAddressCountry, &stuDb.PresentAddressState, &stuDb.PresentAddressCity, &stuDb.PresentAddressDistrict, &stuDb.PresentAddressZipcode, &stuDb.PresentAddressPhone, &stuDb.FathersGuardianFullName, &stuDb.FathersGuardianOccupation, &stuDb.FathersGuardianCompany, &stuDb.FathersGuardianPhoneNumber, &stuDb.FathersGuardianEmailID, &stuDb.MothersGuardianFullName, &stuDb.MothersGuardianOccupation, &stuDb.MothersGuardianCompany, &stuDb.MothersGuardianDesignation, &stuDb.MothersGuardianPhoneNumber, &stuDb.MothersGuardianEmailID, &stuDb.DateOfJoining, &stuDb.PrimaryEmailVerified, &stuDb.PrimaryPhoneVerified, &stuDb.AccountStatus, &stuDb.ProfilePicture, &stuDb.AccountExpiryDate, &stuDb.UniversityName, &stuDb.UniversityID, &stuDb.Attachment, &stuDb.AttachmentName, &stuDb.DateOfJoining)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		c.Abort()
		return
	}

	stuData.Profile = stuDb
	stuData.GetContactFromProfile()
	stuData.Academics, err = models.GetAcademics(ID)
	if err != nil {
		resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Failed to Process request", Err: err, SuccessResp: successResp})
		c.JSON(http.StatusInternalServerError, resp)
		c.Abort()
		return
	}
	stuData.Languages.StakeholderID = ID
	err = stuData.Languages.GetAllLanguages()
	if err != nil {
		resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Failed to Process request", Err: err, SuccessResp: successResp})
		c.JSON(http.StatusInternalServerError, resp)
		c.Abort()
		return
	}
	stuData.Certifications.StakeholderID = ID
	err = stuData.Certifications.GetAllCerts()
	if err != nil {
		resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Failed to Process request", Err: err, SuccessResp: successResp})
		c.JSON(http.StatusInternalServerError, resp)
		c.Abort()
		return
	}
	stuData.Assessments.StakeholderID = ID
	err = stuData.Assessments.GetAllAssessment()
	if err != nil {
		resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Failed to Process request", Err: err, SuccessResp: successResp})
		c.JSON(http.StatusInternalServerError, resp)
		c.Abort()
		return
	}
	stuData.Internships.StakeholderID = ID
	err = stuData.Internships.GetAllInternships()
	if err != nil {
		resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Failed to Process request", Err: err, SuccessResp: successResp})
		c.JSON(http.StatusInternalServerError, resp)
		c.Abort()
		return
	}

	// GetAwards ...
	sa, err := getAwards(ID)
	if err != nil {
		resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Failed to Process request", Err: err, SuccessResp: successResp})
		c.JSON(http.StatusInternalServerError, resp)
		c.Abort()
		return
	}

	sac, err := getCompetitions(ID)
	if err != nil {
		resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Failed to Read rows", Err: err, SuccessResp: successResp})
		c.JSON(http.StatusInternalServerError, resp)
		c.Abort()
	}

	sae, err := getAllExtraCurricular(ID)
	if err != nil {
		resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Failed to Read rows", Err: err, SuccessResp: successResp})
		c.JSON(http.StatusInternalServerError, resp)
		c.Abort()
		return
	}

	sap, err := getPatents(ID)
	if err != nil {
		resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Failed to Get Patents", Err: err, SuccessResp: successResp})
		c.JSON(http.StatusInternalServerError, resp)
		c.Abort()
		return
	}

	sapj, err := getProjects(ID)
	if err != nil {
		resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Failed to Get Projects", Err: err, SuccessResp: successResp})
		c.JSON(http.StatusInternalServerError, resp)
		c.Abort()
		return
	}

	sapb, err := getPublications(ID)
	if err != nil {
		resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Failed to Read rows", Err: err, SuccessResp: successResp})
		c.JSON(http.StatusInternalServerError, resp)
		c.Abort()
		return
	}

	sasc, err := getScholarships(ID)
	if err != nil {
		resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Failed to Get Scholarships", Err: err, SuccessResp: successResp})
		c.JSON(http.StatusInternalServerError, resp)
		c.Abort()
		return
	}

	sasa, err := getSocialAccounts(ID)
	if err != nil {
		resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Failed to Get SocialAccount", Err: err, SuccessResp: successResp})
		c.JSON(http.StatusInternalServerError, resp)
		c.Abort()
		return
	}

	ts, err := getTechSkills(ID)
	if err != nil {
		resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Failed to Get TechSkills", Err: err, SuccessResp: successResp})
		c.JSON(http.StatusInternalServerError, resp)
		c.Abort()
		return
	}

	saTs, err := getTestScores(ID)
	if err != nil {
		resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Failed to Read rows", Err: err, SuccessResp: successResp})
		c.JSON(http.StatusInternalServerError, resp)
		c.Abort()
		return
	}

	sav, err := getAllVolunteerExperience(ID)
	if err != nil {
		resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Failed to Read rows", Err: err, SuccessResp: successResp})
		c.JSON(http.StatusInternalServerError, resp)
		c.Abort()
		return
	}

	saEvent, err := getEvents(ID)
	if err != nil {
		resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Failed to Get Events", Err: err, SuccessResp: successResp})
		c.JSON(http.StatusInternalServerError, resp)
		c.Abort()
		return
	}

	stuData.LanguagesArray = stuData.Languages.Languages
	stuData.CertificationsArray = stuData.Certifications.Certifications
	stuData.AssessmentsArray = stuData.Assessments.Assessments
	stuData.InternshipsArray = stuData.Internships.Internships
	stuData.AwardsArray = sa.Awards
	stuData.CompetitionsArray = sac.Competitions
	stuData.ExtraCurricularArray = sae.ExtraCurricular
	stuData.PatentsArray = sap.Patents
	stuData.ProjectsArray = sapj.Projects
	stuData.PublicationsArray = sapb.Publications
	stuData.ScholarshipsArray = sasc.Scholarships
	stuData.SocialAccountArray = sasa.SocialAccounts
	stuData.TechSkillsArray = ts.TechSkills
	stuData.TestScoresArray = saTs.TestScores
	stuData.VolunteerExperienceArray = sav.VolunteerExperience
	stuData.EventsArray = saEvent.Events
	c.JSON(http.StatusOK, stuData)
	return
}
