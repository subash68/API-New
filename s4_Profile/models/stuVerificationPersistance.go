package models

import (
	"fmt"
	"strings"
	"time"

	"github.com/jaswanth-gorripati/PGK/s4_Profile/configuration"
)

// PvcPending
const (
	PvcPending    string = "PENDING"
	PvcRevalidate string = "REVALIDATE"
	PvcVerified   string = "VERIFIED"
)

// RequestProfileVerification ...
func (sm *StudentMasterDb) RequestProfileVerification() (string, error) {
	var spv StudentProfileVerificationDataModel
	dbError := spv.GetVrfProfileData(sm.StakeholderID)
	if dbError.Err != nil {
		return "", dbError.Err
	}
	if spv.ContactInfo.UniversityID == "" {
		return "", fmt.Errorf("Current University details are empty")
	}
	if spv.ContactInfo.PermanentAddressZipcode == "" {
		return "", fmt.Errorf("Permanent Address cannot be empty")
	}
	if spv.Academics.Tenth.Percentage == "" && spv.Academics.Twelfth.Percentage == "" {
		return "", fmt.Errorf("Academics cannot be empty")
	}
	dbConfig := configuration.DbConfig()
	tbName := dbConfig.DbDatabaseName
	dbNames := tbName + "." + dbConfig.StuMasterDbName + "," +
		tbName + "." + dbConfig.StuAcademics + "," +
		tbName + "." + dbConfig.StuSemDbName + "," +
		tbName + "." + dbConfig.StuCertsDbName + "," +
		tbName + "." + dbConfig.StuAssessmentDbName + "," +
		tbName + "." + dbConfig.StuIntershipsDbName + "," +
		tbName + "." + dbConfig.StuAwardsDbName + "," +
		tbName + "." + dbConfig.StuEventsDbName + "," +
		tbName + "." + dbConfig.StuExtraCurDbName + "," +
		tbName + "." + dbConfig.StuPatentsDbName + "," +
		tbName + "." + dbConfig.StuProjectsDbName + "," +
		tbName + "." + dbConfig.StuPublicationsDbName + "," +
		tbName + "." + dbConfig.StuScholarshipsDbName + "," +
		tbName + "." + dbConfig.StuTestScoresDbName + "," +
		tbName + "." + dbConfig.StuVolunteerExpDbName
	stmts := addUpdateVrfStmt(dbNames)
	fmt.Printf("=======> Sending verification query %v ", stmts)
	err := execUpdateVrfStmt(stmts, sm.StakeholderID)
	if err != nil {
		return "", err
	}
	return spv.ContactInfo.UniversityID, nil
}

func addUpdateVrfStmt(dbNames string) []string {
	stmt := []string{}
	sp, _ := RetriveSP("STU_SEND_FOR_VRF")
	allDbName := strings.Split(dbNames, ",")
	for _, v := range allDbName {
		if v != "" {
			newSp := strings.ReplaceAll(sp, "//UPDATE_DB_NAMES", v)
			stmt = append(stmt, newSp)
		}
	}
	return stmt
}

func execUpdateVrfStmt(stmt []string, sh string) error {
	currentTime := time.Now().Format(time.RFC3339)
	for _, v := range stmt {
		if v != "" {
			fmt.Printf("\nPreparing %v , %s, %s\n", v, currentTime, sh)
			stmt, err := Db.Prepare(v)
			if err != nil {
				fmt.Printf("\nFailed Preparing %v\n", v)
				return err
			}
			_, err = stmt.Exec(currentTime, sh)
			if err != nil {
				fmt.Printf("\nFailed Executing %v\n", v)
				return err
			}
		}
	}
	return nil
}

// GetVrfProfileData ...
func (spv *StudentProfileVerificationDataModel) GetVrfProfileData(ID string) DbModelError {
	var dbError DbModelError
	info, dbError := GetProfile(ID, "STU_GET_PROFILE")
	if dbError.ErrTyp != "000" {
		return dbError
	}

	stuDb := StudentMasterDb{}
	err := info.Scan(&stuDb.StakeholderID, &stuDb.FirstName, &stuDb.MiddleName, &stuDb.LastName, &stuDb.PersonalEmail, &stuDb.CollegeEmail, &stuDb.PhoneNumber, &stuDb.AlternatePhoneNumber, &stuDb.CollegeID, &stuDb.Gender, &stuDb.DateOfBirth, &stuDb.AadharNumber, &stuDb.PermanentAddressLine1, &stuDb.PermanentAddressLine2, &stuDb.PermanentAddressLine3, &stuDb.PermanentAddressCountry, &stuDb.PermanentAddressState, &stuDb.PermanentAddressCity, &stuDb.PermanentAddressDistrict, &stuDb.PermanentAddressZipcode, &stuDb.PermanentAddressPhone, &stuDb.PresentAddressLine1, &stuDb.PresentAddressLine2, &stuDb.PresentAddressLine3, &stuDb.PresentAddressCountry, &stuDb.PresentAddressState, &stuDb.PresentAddressCity, &stuDb.PresentAddressDistrict, &stuDb.PresentAddressZipcode, &stuDb.PresentAddressPhone, &stuDb.FathersGuardianFullName, &stuDb.FathersGuardianOccupation, &stuDb.FathersGuardianCompany, &stuDb.FathersGuardianPhoneNumber, &stuDb.FathersGuardianEmailID, &stuDb.MothersGuardianFullName, &stuDb.MothersGuardianOccupation, &stuDb.MothersGuardianCompany, &stuDb.MothersGuardianDesignation, &stuDb.MothersGuardianPhoneNumber, &stuDb.MothersGuardianEmailID, &stuDb.DateOfJoining, &stuDb.PrimaryEmailVerified, &stuDb.PrimaryPhoneVerified, &stuDb.AccountStatus, &stuDb.ProfilePicture, &stuDb.AccountExpiryDate, &stuDb.UniversityName, &stuDb.UniversityID, &stuDb.Attachment, &stuDb.AttachmentName, &stuDb.DateOfJoining, &stuDb.SentforVerification, &stuDb.DateSentforVerification, &stuDb.Verified, &stuDb.DateVerified, &stuDb.SentbackforRevalidation, &stuDb.DateSentBackForRevalidation, &stuDb.ValidatorRemarks, &stuDb.VerificationType, &stuDb.VerifiedByStakeholderID, &stuDb.VerifiedByEmailID)
	if err != nil {
		dbError.ErrTyp = "500"
		dbError.Err = err
		return dbError
	}
	spv.Profile = stuDb
	spv.GetContactFromProfile()
	spv.Academics, err = GetAcademics(ID)
	if err != nil {
		dbError.ErrTyp = "500"
		dbError.Err = err
		return dbError
	}
	dbError.ErrTyp = "000"
	return dbError
}

// GetAllStudentProfileMetadata ...
func GetAllStudentProfileMetadata(ID string, verificationStatus string) (sap []StudentAllProfiles, dbError DbModelError) {
	sp, _ := RetriveSP("STU_GET_ALL_PROFILES")
	fmt.Println(sp, ID, verificationStatus)
	rows, err := Db.Query(sp, ID, verificationStatus)
	fmt.Println("===== rows", rows, err)
	defer rows.Close()
	for rows.Next() {
		var newSl StudentAllProfiles
		var gradProgram, gradBranch, gradYear, pgProgram, pgBranch, pgYear string
		err = rows.Scan(&newSl.StudentPlatformID, &newSl.StudentName, &newSl.UniversityID, &gradProgram, &gradBranch, &gradYear, &pgProgram, &pgBranch, &pgYear)
		if err != nil {
			dbError.ErrTyp = "500"
			dbError.Err = err
			return sap, dbError
		}
		if newSl.StudentPlatformID != "" {
			if pgProgram != "" && pgBranch != "" && pgYear != "" {
				newSl.Program = pgProgram
				newSl.BranchName = pgBranch
				newSl.Year = switchToText(pgYear, true)
			} else {
				newSl.Program = gradProgram
				newSl.BranchName = gradBranch
				newSl.Year = switchToText(gradYear, false)
			}

			sap = append(sap, newSl)
		}
	}
	dbError.ErrTyp = "000"
	return sap, dbError

}

// ValidateStudentProfile ...
func (sm *StudentMasterDb) ValidateStudentProfile(status bool, ID string) error {
	sp, _ := RetriveSP("STU_VALIDATE_PROFILE")
	stmt, err := Db.Prepare(sp)
	if err != nil {
		return err
	}
	pvs := PvcRevalidate
	if status {
		pvs = PvcVerified
	}
	_, err = stmt.Exec(pvs, status, sm.StakeholderID, ID)
	if err != nil {
		return err
	}
	return nil
}

func switchToText(id string, isPg bool) string {
	switch id {
	case "1":
		return "First"
		break
	case "2":
		if isPg {
			return "Final"
		}
		return "Second"
		break
	case "3":
		return "Third"
		break
	case "4":
		return "Final"
		break
	case "5":
		return "Final"
		break
	}
	return ""
}
