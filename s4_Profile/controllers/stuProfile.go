package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jaswanth-gorripati/PGK/s4_Profile/models"
)

// GetStudentProfile ...
func GetStudentProfile(c *gin.Context) {
	_, ID, _, _ := getFuncReq(c, "Get Student Profile")
	stuData, err := returnCompleteStuData(ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, stuData)
	return
}

func returnCompleteStuData(ID string) (models.StudentCompleteProfileDataModel, error) {
	var stuData models.StudentCompleteProfileDataModel
	info, dbError := models.GetProfile(ID, "STU_GET_PROFILE")
	if dbError.ErrTyp != "000" {
		return stuData, dbError.Err
	}

	stuDb := models.StudentMasterDb{}
	err := info.Scan(&stuDb.StakeholderID, &stuDb.FirstName, &stuDb.MiddleName, &stuDb.LastName, &stuDb.PersonalEmail, &stuDb.CollegeEmail, &stuDb.PhoneNumber, &stuDb.AlternatePhoneNumber, &stuDb.CollegeID, &stuDb.Gender, &stuDb.DateOfBirth, &stuDb.AadharNumber, &stuDb.PermanentAddressLine1, &stuDb.PermanentAddressLine2, &stuDb.PermanentAddressLine3, &stuDb.PermanentAddressCountry, &stuDb.PermanentAddressState, &stuDb.PermanentAddressCity, &stuDb.PermanentAddressDistrict, &stuDb.PermanentAddressZipcode, &stuDb.PermanentAddressPhone, &stuDb.PresentAddressLine1, &stuDb.PresentAddressLine2, &stuDb.PresentAddressLine3, &stuDb.PresentAddressCountry, &stuDb.PresentAddressState, &stuDb.PresentAddressCity, &stuDb.PresentAddressDistrict, &stuDb.PresentAddressZipcode, &stuDb.PresentAddressPhone, &stuDb.FathersGuardianFullName, &stuDb.FathersGuardianOccupation, &stuDb.FathersGuardianCompany, &stuDb.FathersGuardianPhoneNumber, &stuDb.FathersGuardianEmailID, &stuDb.MothersGuardianFullName, &stuDb.MothersGuardianOccupation, &stuDb.MothersGuardianCompany, &stuDb.MothersGuardianDesignation, &stuDb.MothersGuardianPhoneNumber, &stuDb.MothersGuardianEmailID, &stuDb.DateOfJoining, &stuDb.PrimaryEmailVerified, &stuDb.PrimaryPhoneVerified, &stuDb.AccountStatus, &stuDb.ProfilePicture, &stuDb.AccountExpiryDate, &stuDb.UniversityName, &stuDb.UniversityID, &stuDb.Attachment, &stuDb.AttachmentName, &stuDb.DateOfJoining, &stuDb.SentforVerification, &stuDb.DateSentforVerification, &stuDb.Verified, &stuDb.DateVerified, &stuDb.SentbackforRevalidation, &stuDb.DateSentBackForRevalidation, &stuDb.ValidatorRemarks, &stuDb.VerificationType, &stuDb.VerifiedByStakeholderID, &stuDb.VerifiedByEmailID)
	if err != nil {
		return stuData, err
	}

	stuData.Profile = stuDb
	stuData.GetContactFromProfile()
	stuData.Academics, err = models.GetAcademics(ID)
	if err != nil {
		return stuData, err
	}

	stuData.Certifications.StakeholderID = ID
	err = stuData.Certifications.GetAllCerts()
	if err != nil {
		return stuData, err
	}
	stuData.Assessments.StakeholderID = ID
	err = stuData.Assessments.GetAllAssessment()
	if err != nil {
		return stuData, err
	}
	stuData.Internships.StakeholderID = ID
	err = stuData.Internships.GetAllInternships()
	if err != nil {
		return stuData, err
	}

	// GetAwards ...
	sa, err := getAwards(ID)
	if err != nil {
		return stuData, err
	}

	// sac, err := getCompetitions(ID)
	// if err != nil {
	// 	resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S3PJ", ErrTyp: "Failed to Read rows", Err: err, SuccessResp: successResp})
	// 	c.JSON(http.StatusInternalServerError, resp)
	// 	c.Abort()
	// }

	sae, err := getAllExtraCurricular(ID)
	if err != nil {
		return stuData, err
	}

	sap, err := getPatents(ID)
	if err != nil {
		return stuData, err
	}

	sapj, err := getProjects(ID)
	if err != nil {
		return stuData, err
	}

	sapb, err := getPublications(ID)
	if err != nil {
		return stuData, err
	}

	sasc, err := getScholarships(ID)
	if err != nil {
		return stuData, err
	}

	saTs, err := getTestScores(ID)
	if err != nil {
		return stuData, err
	}

	sav, err := getAllVolunteerExperience(ID)
	if err != nil {
		return stuData, err
	}

	saEvent, err := getEvents(ID)
	if err != nil {
		return stuData, err
	}

	stuData.CertificationsArray = stuData.Certifications.Certifications
	stuData.AssessmentsArray = stuData.Assessments.Assessments
	stuData.InternshipsArray = stuData.Internships.Internships
	stuData.AwardsArray = sa.Awards
	// stuData.CompetitionsArray = sac.Competitions
	stuData.ExtraCurricularArray = sae.ExtraCurricular
	stuData.PatentsArray = sap.Patents
	stuData.ProjectsArray = sapj.Projects
	stuData.PublicationsArray = sapb.Publications
	stuData.ScholarshipsArray = sasc.Scholarships
	stuData.TestScoresArray = saTs.TestScores
	stuData.VolunteerExperienceArray = sav.VolunteerExperience
	stuData.EventsArray = saEvent.Events

	return stuData, nil

}
