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
func GetAllStudentProfileMetadata(ID string, verificationStatus bool) (sap []StudentAllProfiles, dbError DbModelError) {
	sp, _ := RetriveSP("STU_SRH_VRF_DYNAMIC")
	if verificationStatus {
		sp = strings.ReplaceAll(sp, "//REPLACE_VRF_CHECKING", "VerificationStatus_Verified=1 ")
	} else {
		sp = strings.ReplaceAll(sp, "//REPLACE_VRF_CHECKING", "VerificationStatus_SentforVerification=1 ")
	}
	// TODO STUDENT FILTERS
	filters := ""

	sp = strings.ReplaceAll(sp, "//REPLACE_FILTER", filters)

	fmt.Println(sp, ID, verificationStatus)
	rows, err := Db.Query(sp, ID)
	fmt.Println("===== rows", rows, err)
	if err != nil {
		dbError.ErrTyp = "500"
		dbError.Err = err
		return sap, dbError
	}
	defer rows.Close()
	for rows.Next() {
		var newSl StudentAllProfiles
		var gradProgram, gradBranch, gradYear, pgProgram, pgBranch, pgYear string
		err = rows.Scan(&newSl.StudentPlatformID, &newSl.StudentFirstName, &newSl.StudentMiddleName, &newSl.StudentLastName, &newSl.UniversityID, &gradProgram, &gradBranch, &gradYear, &pgProgram, &pgBranch, &pgYear)
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
func (sm *StuVrfDataModel) ValidateStudentProfile(ID string) error {
	var unvData UnvVrfData
	unvEmailSp, _ := RetriveSP("UNV_GET_EMAIL_BY_ID")
	err := Db.QueryRow(unvEmailSp, ID).Scan(&unvData.Email)
	if err != nil {
		return err
	}
	unvData.ID = ID
	unvData.currentTime = time.Now().Format(time.RFC3339)
	unvData.VrfType = "U"
	dbConfig := configuration.DbConfig()
	queries := []string{}
	if sm.Account.Verified == true || sm.Account.Remarks != "" {
		table := dbConfig.DbDatabaseName + "." + dbConfig.StuMasterDbName
		query := ""
		if sm.Account.Verified {
			query = "UPDATE " + table + " SET VerificationStatus_SentforVerification=false,VerificationStatus_Verified=true,Date_Verified='" + unvData.currentTime + "',VerificationType='" + unvData.VrfType + "',VerifiedBy_Stakeholder_ID='" + unvData.ID + "',VerifiedBy_Email_ID='" + unvData.Email + "' WHERE VerificationStatus_SentforVerification=true AND Stakeholder_ID='" + sm.StudentID + "' "
		} else {
			query = "UPDATE " + table + " SET VerificationStatus_SentforVerification=false,VerificationStatus_SentbackforRevalidation=true,Date_SentbackforRevalidation='" + unvData.currentTime + "',Validator_Remarks='" + sm.Account.Remarks + "' WHERE VerificationStatus_SentforVerification=true AND Stakeholder_ID='" + sm.StudentID + "'"
		}
		queries = append(queries, query)
	}

	if sm.BelowGrad.Verified == true || sm.BelowGrad.Remarks != "" {
		table := dbConfig.DbDatabaseName + "." + dbConfig.StuAcademics
		query := ""
		if sm.BelowGrad.Verified {
			query = "UPDATE " + table + " SET VerificationStatus_SentforVerification=false,VerificationStatus_Verified=true,Date_Verified='" + unvData.currentTime + "',VerificationType='" + unvData.VrfType + "',VerifiedBy_Stakeholder_ID='" + unvData.ID + "',VerifiedBy_Email_ID='" + unvData.Email + "' WHERE VerificationStatus_SentforVerification=true AND Stakeholder_ID='" + sm.StudentID + "' "
		} else {
			query = "UPDATE " + table + " SET VerificationStatus_SentforVerification=false,VerificationStatus_SentbackforRevalidation=true,Date_SentbackforRevalidation='" + unvData.currentTime + "',Validator_Remarks='" + sm.Account.Remarks + "' WHERE VerificationStatus_SentforVerification=true AND Stakeholder_ID='" + sm.StudentID + "' "
		}
		queries = append(queries, query)
	}
	for _, v := range sm.Semister {
		if v.Verified == true || v.Remarks != "" {
			table := dbConfig.DbDatabaseName + "." + dbConfig.StuSemDbName
			query := ConstructVrfUpdQry(table, unvData, v, sm.StudentID)
			queries = append(queries, query)
		}
	}

	for _, v := range sm.Certifications {
		if v.Verified == true || v.Remarks != "" {
			table := dbConfig.DbDatabaseName + "." + dbConfig.StuCertsDbName
			query := ConstructVrfUpdQry(table, unvData, v, sm.StudentID)
			queries = append(queries, query)
		}
	}

	for _, v := range sm.Assessments {
		if v.Verified == true || v.Remarks != "" {
			table := dbConfig.DbDatabaseName + "." + dbConfig.StuAssessmentDbName
			query := ConstructVrfUpdQry(table, unvData, v, sm.StudentID)
			queries = append(queries, query)
		}
	}

	for _, v := range sm.Internships {
		if v.Verified == true || v.Remarks != "" {
			table := dbConfig.DbDatabaseName + "." + dbConfig.StuIntershipsDbName
			query := ConstructVrfUpdQry(table, unvData, v, sm.StudentID)
			queries = append(queries, query)
		}
	}

	for _, v := range sm.Awards {
		if v.Verified == true || v.Remarks != "" {
			table := dbConfig.DbDatabaseName + "." + dbConfig.StuAwardsDbName
			query := ConstructVrfUpdQry(table, unvData, v, sm.StudentID)
			queries = append(queries, query)
		}
	}
	for _, v := range sm.Events {
		if v.Verified == true || v.Remarks != "" {
			table := dbConfig.DbDatabaseName + "." + dbConfig.StuEventsDbName
			query := ConstructVrfUpdQry(table, unvData, v, sm.StudentID)
			queries = append(queries, query)
		}
	}
	for _, v := range sm.ExtraCurricular {
		if v.Verified == true || v.Remarks != "" {
			table := dbConfig.DbDatabaseName + "." + dbConfig.StuExtraCurDbName
			query := ConstructVrfUpdQry(table, unvData, v, sm.StudentID)
			queries = append(queries, query)
		}
	}
	for _, v := range sm.Patents {
		if v.Verified == true || v.Remarks != "" {
			table := dbConfig.DbDatabaseName + "." + dbConfig.StuPatentsDbName
			query := ConstructVrfUpdQry(table, unvData, v, sm.StudentID)
			queries = append(queries, query)
		}
	}
	for _, v := range sm.Projects {
		if v.Verified == true || v.Remarks != "" {
			table := dbConfig.DbDatabaseName + "." + dbConfig.StuProjectsDbName
			query := ConstructVrfUpdQry(table, unvData, v, sm.StudentID)
			queries = append(queries, query)
		}
	}
	for _, v := range sm.Publications {
		if v.Verified == true || v.Remarks != "" {
			table := dbConfig.DbDatabaseName + "." + dbConfig.StuPublicationsDbName
			query := ConstructVrfUpdQry(table, unvData, v, sm.StudentID)
			queries = append(queries, query)
		}
	}
	for _, v := range sm.Scholarships {
		if v.Verified == true || v.Remarks != "" {
			table := dbConfig.DbDatabaseName + "." + dbConfig.StuScholarshipsDbName
			query := ConstructVrfUpdQry(table, unvData, v, sm.StudentID)
			queries = append(queries, query)
		}
	}
	for _, v := range sm.TestScores {
		if v.Verified == true || v.Remarks != "" {
			table := dbConfig.DbDatabaseName + "." + dbConfig.StuTestScoresDbName
			query := ConstructVrfUpdQry(table, unvData, v, sm.StudentID)
			queries = append(queries, query)
		}
	}
	for _, v := range sm.VolunteerExperience {
		if v.Verified == true || v.Remarks != "" {
			table := dbConfig.DbDatabaseName + "." + dbConfig.StuVolunteerExpDbName
			query := ConstructVrfUpdQry(table, unvData, v, sm.StudentID)
			queries = append(queries, query)
		}
	}
	err = execVrfUpdQueries(queries)
	if err != nil {
		return err
	}

	return nil
}

// ConstructVrfUpdQry ...
func ConstructVrfUpdQry(table string, unvData UnvVrfData, vd VrfDataModel, stuID string) string {
	var query string
	if vd.Verified {
		query = "UPDATE " + table + " SET VerificationStatus_SentforVerification=false,VerificationStatus_Verified=true,Date_Verified='" + unvData.currentTime + "',VerificationType='" + unvData.VrfType + "',VerifiedBy_Stakeholder_ID='" + unvData.ID + "',VerifiedBy_Email_ID='" + unvData.Email + "' WHERE VerificationStatus_SentforVerification=true AND Stakeholder_ID='" + stuID + "' AND id=" + vd.ID + " "
	} else {
		query = "UPDATE " + table + " SET VerificationStatus_SentforVerification=false,VerificationStatus_SentbackforRevalidation=true,Date_SentbackforRevalidation='" + unvData.currentTime + "',Validator_Remarks='" + vd.Remarks + "' WHERE VerificationStatus_SentforVerification=true AND Stakeholder_ID='" + stuID + "' AND id=" + vd.ID + " "
	}
	return query
}

func execVrfUpdQueries(queries []string) error {
	for _, v := range queries {
		if v != "" {
			fmt.Printf("\nPreparing %v , %s, %s\n")
			stmt, err := Db.Prepare(v)
			if err != nil {
				fmt.Printf("\nFailed Preparing %v\n", v)
				return err
			}
			_, err = stmt.Exec()
			if err != nil {
				fmt.Printf("\nFailed Executing %v\n", v)
				return err
			}
		}
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

func dynamicSearchStuProfile() string {
	return ""
}
