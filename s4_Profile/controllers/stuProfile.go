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
	err := info.Scan(&stuDb.StakeholderID, &stuDb.FirstName, &stuDb.MiddleName, &stuDb.LastName, &stuDb.PersonalEmail, &stuDb.CollegeEmail, &stuDb.PhoneNumber, &stuDb.AlternatePhoneNumber, &stuDb.CollegeID, &stuDb.Gender, &stuDb.DateOfBirth, &stuDb.AadharNumber, &stuDb.PermanentAddressLine1, &stuDb.PermanentAddressLine2, &stuDb.PermanentAddressLine3, &stuDb.PermanentAddressCountry, &stuDb.PermanentAddressState, &stuDb.PermanentAddressCity, &stuDb.PermanentAddressDistrict, &stuDb.PermanentAddressZipcode, &stuDb.PermanentAddressPhone, &stuDb.PresentAddressLine1, &stuDb.PresentAddressLine2, &stuDb.PresentAddressLine3, &stuDb.PresentAddressCountry, &stuDb.PresentAddressState, &stuDb.PresentAddressCity, &stuDb.PresentAddressDistrict, &stuDb.PresentAddressZipcode, &stuDb.PresentAddressPhone, &stuDb.FathersGuardianFullName, &stuDb.FathersGuardianOccupation, &stuDb.FathersGuardianCompany, &stuDb.FathersGuardianPhoneNumber, &stuDb.FathersGuardianEmailID, &stuDb.MothersGuardianFullName, &stuDb.MothersGuardianOccupation, &stuDb.MothersGuardianCompany, &stuDb.MothersGuardianDesignation, &stuDb.MothersGuardianPhoneNumber, &stuDb.MothersGuardianEmailID, &stuDb.DateOfJoining, &stuDb.PrimaryEmailVerified, &stuDb.PrimaryPhoneVerified, &stuDb.AccountStatus, &stuDb.ProfilePicture, &stuDb.AccountExpiryDate, &stuDb.AboutMeNullable)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		c.Abort()
		return
	}
	if stuDb.AboutMeNullable.Valid {
		stuDb.AboutMe = stuDb.AboutMeNullable.String
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
	stuData.LanguagesArray = stuData.Languages.Languages
	stuData.CertificationsArray = stuData.Certifications.Certifications
	stuData.AssessmentsArray = stuData.Assessments.Assessments
	stuData.InternshipsArray = stuData.Internships.Internships
	c.JSON(http.StatusOK, stuData)
	return
}
