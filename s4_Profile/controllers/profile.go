// Package controllers ...
package controllers

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jaswanth-gorripati/PGK/s4_Profile/models"
)

// ProfileRespModel ...
type ProfileRespModel struct {
	Message    string `json:"message,omitempty"`
	ProfilePic []byte `json:"profilePic,omitempty"`
}

// GetProfile ...
func GetProfile(c *gin.Context) {
	ctx, ID, userType, successResp := getFuncReq(c, "Get Profile")
	cmd := ""
	switch userType {
	case "Corporate":
		cmd = "CORP_GET_PROFILE"
		break
	case "University":
		cmd = "UNV_GET_PROFILE"
		break
	case "Student":
		cmd = "STU_GET_PROFILE"
		break
	default:
		resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S4Profile", ErrTyp: "Invalid Stakehodler type", Err: fmt.Errorf("" + userType + " is invaild,  Expecting Corporate,University or Student"), SuccessResp: successResp})
		c.JSON(http.StatusUnprocessableEntity, resp)
		return
	}
	info, dbError := models.GetProfile(ID, cmd)
	if dbError.ErrTyp != "000" {
		resp := ErrCheck(ctx, dbError)
		c.Error(dbError.Err)
		c.JSON(http.StatusInternalServerError, resp)
		return
	}
	switch userType {
	case "Corporate":
		corpDb := models.CorporateMasterDB{}
		err := info.Scan(&corpDb.StakeholderID, &corpDb.CorporateName, &corpDb.CIN, &corpDb.CorporateHQAddressLine1, &corpDb.CorporateHQAddressLine2, &corpDb.CorporateHQAddressLine3, &corpDb.CorporateHQAddressCountry, &corpDb.CorporateHQAddressState, &corpDb.CorporateHQAddressCity, &corpDb.CorporateHQAddressDistrict, &corpDb.CorporateHQAddressZipCode, &corpDb.CorporateHQAddressPhone, &corpDb.CorporateHQAddressEmail, &corpDb.CorporateLocalBranchAddressLine1, &corpDb.CorporateLocalBranchAddressLine2, &corpDb.CorporateLocalBranchAddressLine3, &corpDb.CorporateLocalBranchAddressCountry, &corpDb.CorporateLocalBranchAddressState, &corpDb.CorporateLocalBranchAddressCity, &corpDb.CorporateLocalBranchAddressDistrict, &corpDb.CorporateLocalBranchAddressZipCode, &corpDb.CorporateLocalBranchAddressPhone, &corpDb.CorporateLocalBranchAddressEmail, &corpDb.PrimaryContactFirstName, &corpDb.PrimaryContactMiddleName, &corpDb.PrimaryContactLastName, &corpDb.PrimaryContactDesignation, &corpDb.PrimaryContactPhone, &corpDb.PrimaryContactEmail, &corpDb.SecondaryContactFirstName, &corpDb.SecondaryContactMiddleName, &corpDb.SecondaryContactLastName, &corpDb.SecondaryContactDesignation, &corpDb.SecondaryContactPhone, &corpDb.SecondaryContactEmail, &corpDb.CorporateType, &corpDb.CorporateCategory, &corpDb.CorporateIndustry, &corpDb.CompanyProfile, &corpDb.Attachment, &corpDb.YearOfEstablishment, &corpDb.DateOfJoining, &corpDb.AccountStatus, &corpDb.PrimaryEmailVerified, &corpDb.PrimaryPhoneVerified, &corpDb.ProfilePicture, &corpDb.AccountExpiryDate, &corpDb.AttachmentName, &corpDb.PublishedFlag)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			c.Abort()
			return
		}
		c.JSON(http.StatusOK, corpDb)
		return
		break
	case "University":
		unvDb := models.UniversityMasterDb{}
		err := info.Scan(&unvDb.StakeholderID, &unvDb.UniversityName, &unvDb.UniversityCollegeID, &unvDb.UniversityHQAddressLine1, &unvDb.UniversityHQAddressLine2, &unvDb.UniversityHQAddressLine3, &unvDb.UniversityHQAddressCountry, &unvDb.UniversityHQAddressState, &unvDb.UniversityHQAddressCity, &unvDb.UniversityHQAddressDistrict, &unvDb.UniversityHQAddressZipcode, &unvDb.UniversityHQAddressPhone, &unvDb.UniversityHQAddressemail, &unvDb.UniversityLocalBranchAddressLine1, &unvDb.UniversityLocalBranchAddressLine2, &unvDb.UniversityLocalBranchAddressLine3, &unvDb.UniversityLocalBranchAddressCountry, &unvDb.UniversityLocalBranchAddressState, &unvDb.UniversityLocalBranchAddressCity, &unvDb.UniversityLocalBranchAddressDistrict, &unvDb.UniversityLocalBranchAddressZipcode, &unvDb.UniversityLocalBranchAddressPhone, &unvDb.UniversityLocalBranchAddressemail, &unvDb.PrimaryContactFirstName, &unvDb.PrimaryContactMiddleName, &unvDb.PrimaryContactLastName, &unvDb.PrimaryContactDesignation, &unvDb.PrimaryContactPhone, &unvDb.PrimaryContactEmail, &unvDb.SecondaryContactFirstName, &unvDb.SecondaryContactMiddleName, &unvDb.SecondaryContactLastName, &unvDb.SecondaryContactDesignation, &unvDb.SecondaryContactPhone, &unvDb.SecondaryContactEmail, &unvDb.UniversitySector, &unvDb.UniversityProfile, &unvDb.YearOfEstablishment, &unvDb.Attachment, &unvDb.DateOfJoining, &unvDb.PrimaryEmailVerified, &unvDb.PrimaryPhoneVerified, &unvDb.AccountStatus, &unvDb.ProfilePicture, &unvDb.AccountExpiryDate, &unvDb.AttachmentName, &unvDb.PublishedFlag)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			c.Abort()
			return
		}
		c.JSON(http.StatusOK, unvDb)
		c.Abort()
		return
		break
	case "Student":
		stuDb := models.StudentMasterDb{}
		err := info.Scan(&stuDb.StakeholderID, &stuDb.FirstName, &stuDb.MiddleName, &stuDb.LastName, &stuDb.PersonalEmail, &stuDb.CollegeEmail, &stuDb.PhoneNumber, &stuDb.AlternatePhoneNumber, &stuDb.CollegeID, &stuDb.Gender, &stuDb.DateOfBirth, &stuDb.AadharNumber, &stuDb.PermanentAddressLine1, &stuDb.PermanentAddressLine2, &stuDb.PermanentAddressLine3, &stuDb.PermanentAddressCountry, &stuDb.PermanentAddressState, &stuDb.PermanentAddressCity, &stuDb.PermanentAddressDistrict, &stuDb.PermanentAddressZipcode, &stuDb.PermanentAddressPhone, &stuDb.PresentAddressLine1, &stuDb.PresentAddressLine2, &stuDb.PresentAddressLine3, &stuDb.PresentAddressCountry, &stuDb.PresentAddressState, &stuDb.PresentAddressCity, &stuDb.PresentAddressDistrict, &stuDb.PresentAddressZipcode, &stuDb.PresentAddressPhone, &stuDb.FathersGuardianFullName, &stuDb.FathersGuardianOccupation, &stuDb.FathersGuardianCompany, &stuDb.FathersGuardianPhoneNumber, &stuDb.FathersGuardianEmailID, &stuDb.MothersGuardianFullName, &stuDb.MothersGuardianOccupation, &stuDb.MothersGuardianCompany, &stuDb.MothersGuardianDesignation, &stuDb.MothersGuardianPhoneNumber, &stuDb.MothersGuardianEmailID, &stuDb.DateOfJoining, &stuDb.PrimaryEmailVerified, &stuDb.PrimaryPhoneVerified, &stuDb.AccountStatus, &stuDb.ProfilePicture, &stuDb.AccountExpiryDate, &stuDb.AboutMeNullable, &stuDb.UniversityName, &stuDb.UniversityID, &stuDb.Attachment, &stuDb.AttachmentName)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			c.Abort()
			return
		}
		if stuDb.AboutMeNullable.Valid {
			stuDb.AboutMe = stuDb.AboutMeNullable.String
		}
		c.JSON(http.StatusOK, stuDb)
		c.Abort()
		return
	}
	c.JSON(http.StatusInternalServerError, "Cannot process request")
	return
}

// UpdateProfile ...
func UpdateProfile(c *gin.Context) {
	ctx, ID, userType, successResp := getFuncReq(c, "Update Profile")
	var err error
	for key, value := range c.Request.PostForm {
		fmt.Println(key, value)
	}
	c.Request.ParseMultipartForm(1000)

	updateReq := c.Request.PostForm
	attachmentUpdate := false
	var attachment []byte
	var attachmentName string
	ppUpdate := false
	var ppfile []byte

	form, _ := c.MultipartForm()
	files := form.File["attachment"]
	for _, file := range files {
		fileContent, _ := file.Open()
		attachment, err = ioutil.ReadAll(fileContent)
		attachmentName = file.Filename
		if err != nil {
			c.JSON(http.StatusUnprocessableEntity, err.Error())
			c.Abort()
		}
		attachmentUpdate = true
	}
	ppfiles := form.File["profilePicture"]
	for _, file := range ppfiles {
		fileContent, _ := file.Open()
		ppfile, err = ioutil.ReadAll(fileContent)
		if err != nil {
			c.JSON(http.StatusUnprocessableEntity, err.Error())
			c.Abort()
		}
		ppUpdate = true
	}

	fmt.Printf("\n form data := %+v \n", updateReq)

	sp := ""
	switch userType {
	case "Corporate":
		sp = "CORP_UPDATE_BY_ID"
		break
	case "University":
		sp = "UNV_UPDATE_BY_ID"
		break
	case "Student":
		sp = "STU_UPDATE_BY_ID"
		break
	default:
		resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S4Profile", ErrTyp: "Invalid Stakehodler type", Err: fmt.Errorf("" + userType + " is invaild,  Expecting Corporate,University or Student"), SuccessResp: successResp})
		c.JSON(http.StatusUnprocessableEntity, resp)
		return
	}

	dbError := models.UpdateProfileData(updateReq, sp, "PROFILE_UPDATE_WHERE", ID, attachmentUpdate, attachment, attachmentName, ppUpdate, ppfile)
	if dbError.ErrTyp != "000" {
		resp := ErrCheck(ctx, dbError)
		c.Error(dbError.Err)
		c.JSON(http.StatusInternalServerError, resp)
		return
	}
	c.JSON(http.StatusOK, ProfileRespModel{Message: "Successfully updated"})
	return

}

// GetProfilePic ...
func GetProfilePic(c *gin.Context) {
	ctx, ID, userType, _ := getFuncReq(c, "Get Profile Picture")
	pic, dbError := models.GetProfilePic(ID, userType, "GET_PRF_PIC")
	if dbError.ErrTyp != "000" {
		resp := ErrCheck(ctx, dbError)
		c.Error(dbError.Err)
		c.JSON(http.StatusInternalServerError, resp)
		return
	}
	c.JSON(http.StatusOK, ProfileRespModel{ProfilePic: pic})
	return

}

// UploadProfilePic ...
func UploadProfilePic(c *gin.Context) {
	ctx, ID, userType, successResp := getFuncReq(c, "Upload Profile Picture")
	form, _ := c.MultipartForm()
	var err error
	var profilePic []byte
	files := form.File["profilePic"]
	if len(files) == 0 {
		resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S4Profile003", ErrTyp: "File Not found", Err: fmt.Errorf("Cannot file File from the request, Send file in profilePic"), SuccessResp: successResp})
		c.JSON(http.StatusBadRequest, resp)
		return
	}
	if len(files) > 1 {
		resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S4Profile003", ErrTyp: "Cannot accept Multiple files", Err: fmt.Errorf("Expected Single file"), SuccessResp: successResp})
		c.JSON(http.StatusBadRequest, resp)
		return
	}
	for _, file := range files {
		fileContent, _ := file.Open()
		profilePic, err = ioutil.ReadAll(fileContent)
		if err != nil {
			resp := ErrCheck(ctx, models.DbModelError{ErrCode: "S4Profile003", ErrTyp: "Invalid File", Err: fmt.Errorf("Cannot decode File from the request"), SuccessResp: successResp})
			c.JSON(http.StatusBadRequest, resp)
			return
		}
	}
	dbError := models.UpdateProfilePic(profilePic, ID, userType, "UPLOAD_PROFILE_PIC")
	if dbError.ErrTyp != "000" {
		resp := ErrCheck(ctx, dbError)
		c.Error(dbError.Err)
		c.JSON(http.StatusInternalServerError, resp)
		return
	}
	c.JSON(http.StatusOK, ProfileRespModel{Message: "Profile Picture Updated"})
	return

}
