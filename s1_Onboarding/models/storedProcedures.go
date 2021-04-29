package models

import (
	"github.com/jaswanth-gorripati/PGK/s1_Onboarding/configuration"
)

// SP is a Mapping of procedure name to Procedure query
var SP map[string]string

// RetriveSP takes in the required Proceder name and returns the procedure
func RetriveSP(procedureName string) (string, bool) {
	if SP[procedureName] == "" {
		return "", false
	}
	return SP[procedureName], true
}

// CreateSP Creates default stored procedures for Database
func CreateSP() {
	dbConfig := configuration.DbConfig()
	SP = map[string]string{
		"CORP_EXISTS_WITH_EMAIL": "SELECT Stakeholder_ID,PrimaryContact_Email,IF(COUNT(*),'true','false') FROM " + dbConfig.DbDatabaseName + "." + dbConfig.CorpMasterDbName + " WHERE PrimaryContact_Email= ? GROUP BY Stakeholder_ID,PrimaryContact_Email LIMIT 1",
		"CORP_INS_NEW_USR":       "INSERT INTO " + dbConfig.DbDatabaseName + "." + dbConfig.CorpMasterDbName + " (Stakeholder_ID,Corporate_Name,Corporate_CIN,CorporateHQAddress_Line1,CorporateHQAddress_Line2,CorporateHQAddress_Line3,CorporateHQAddress_Country,CorporateHQAddress_State,CorporateHQAddress_City,CorporateHQAddress_District,CorporateHQAddress_ZipCode,CorporateHQAddress_Phone,CorporateHQAddress_Email,CorporateLocal_BranchAddress_Line1,CorporateLocal_BranchAddress_Line2,CorporateLocal_BranchAddress_Line3,CorporateLocal_BranchAddress_Country,CorporateLocal_BranchAddress_State,CorporateLocal_BranchAddress_City,CorporateLocal_BranchAddress_District,CorporateLocal_BranchAddress_ZipCode,CorporateLocal_BranchAddress_Phone,CorporateLocal_BranchAddress_Email,PrimaryContact_FirstName,PrimaryContact_MiddleName,PrimaryContact_LastName,PrimaryContact_Designation,PrimaryContact_Phone,PrimaryContact_Email,SecondaryContact_FirstName,SecondaryContact_MiddleName,SecondaryContact_LastName,SecondaryContact_Designation,SecondaryContact_Phone,SecondaryContact_Email,CorporateType,CorporateCategory,CorporateIndustry,CompanyProfile,Attachment,YearOfEstablishment,AccountStatus,UserPassword,AccountExpiryDate)  VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)",
		"CORP_ACC_STATUS_UPD":    "UPDATE " + dbConfig.DbDatabaseName + "." + dbConfig.CorpMasterDbName + " SET Primary_Email_Verified= ?,Primary_Phone_Verified= ?,AccountStatus= ? WHERE Stakeholder_ID= ?",
		"CORP_ACC_ACTIVATION":    "UPDATE " + dbConfig.DbDatabaseName + "." + dbConfig.CorpMasterDbName + " SET AccountStatus= ?,AccountExpiryDate= ? WHERE AccountStatus= ? AND Stakeholder_ID= ?",
		"CORP_CNG_PASS":          "UPDATE " + dbConfig.DbDatabaseName + "." + dbConfig.CorpMasterDbName + " SET UserPassword= ? WHERE Stakeholder_ID= ?",
		"CORP_MBL_VRF_QRY":       "SELECT Primary_Email_Verified,Primary_Phone_Verified,AccountStatus,IF(COUNT(*),'true','false') FROM " + dbConfig.DbDatabaseName + "." + dbConfig.CorpMasterDbName + " WHERE Stakeholder_ID= ? GROUP BY Primary_Email_Verified,Primary_Phone_Verified,AccountStatus LIMIT 1",
		"CORP_VRF_ME_QRY":        "SELECT PrimaryContact_Phone,PrimaryContact_Email,IF(COUNT(*),'true','false') FROM " + dbConfig.DbDatabaseName + "." + dbConfig.CorpMasterDbName + " WHERE Stakeholder_ID= ? GROUP BY PrimaryContact_Phone,PrimaryContact_Email LIMIT 1",
		"CORP_GET_PID":           "SELECT Stakeholder_ID,IF(COUNT(*),'true','false') FROM " + dbConfig.DbDatabaseName + "." + dbConfig.CorpMasterDbName + " WHERE PrimaryContact_Phone= ? OR PrimaryContact_Email= ? GROUP BY Stakeholder_ID",
		"CORP_LOGIN":             "SELECT Stakeholder_ID,AccountStatus,UserPassword,IF(COUNT(*),'true','false') FROM " + dbConfig.DbDatabaseName + "." + dbConfig.CorpMasterDbName + " WHERE PrimaryContact_Email= ? OR Stakeholder_ID= ? OR PrimaryContact_Phone= ? GROUP BY Stakeholder_ID,AccountStatus,UserPassword LIMIT 1",
		"CORP_ROW_CNT":           "SELECT COUNT(*) FROM " + dbConfig.DbDatabaseName + "." + dbConfig.CorpMasterDbName + "",
		"UNV_EXISTS_WITH_EMAIL":  "SELECT Stakeholder_ID,PrimaryContact_Email,IF(COUNT(*),'true','false') FROM " + dbConfig.DbDatabaseName + "." + dbConfig.UnvMasterDbName + " WHERE PrimaryContact_Email= ? GROUP BY Stakeholder_ID,PrimaryContact_Email LIMIT 1",
		"UNV_INS_NEW_USR":        "INSERT INTO " + dbConfig.DbDatabaseName + "." + dbConfig.UnvMasterDbName + " (Stakeholder_ID,University_Name,University_College_ID,UniversityHQAddress_Line1,UniversityHQAddress_Line2,UniversityHQAddress_Line3,UniversityHQAddress_Country,UniversityHQAddress_State,UniversityHQAddress_City,UniversityHQAddress_District,UniversityHQAddress_Zipcode,UniversityHQAddress_Phone,UniversityHQAddress_Email,UniversityLocal_BranchAddress_Line1,UniversityLocal_BranchAddress_Line2,UniversityLocal_BranchAddress_Line3,UniversityLocal_BranchAddress_Country,UniversityLocal_BranchAddress_State,UniversityLocal_BranchAddress_City,UniversityLocal_BranchAddress_District,UniversityLocal_BranchAddress_Zipcode,UniversityLocal_BranchAddress_Phone,UniversityLocal_BranchAddress_Email,PrimaryContact_FirstName,PrimaryContact_MiddleName,PrimaryContact_LastName,PrimaryContact_Designation,PrimaryContact_Phone,PrimaryContact_Email,SecondaryContact_FirstName,SecondaryContact_MiddleName,SecondaryContact_LastName,SecondaryContact_Designation,SecondaryContact_Phone,SecondaryContact_Email,UniversitySector,UniversityProfile,YearOfEstablishment,Attachment,AccountStatus,UserPassword,AccountExpiryDate)  VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)",
		"UNV_ROW_CNT":            "SELECT COUNT(*) FROM " + dbConfig.DbDatabaseName + "." + dbConfig.UnvMasterDbName + "",
		"UNV_ACC_STATUS_UPD":     "UPDATE " + dbConfig.DbDatabaseName + "." + dbConfig.UnvMasterDbName + " SET Primary_Email_Verified= ?,Primary_Phone_Verified= ?,AccountStatus= ? WHERE Stakeholder_ID= ?",
		"UNV_ACC_ACTIVATION":     "UPDATE " + dbConfig.DbDatabaseName + "." + dbConfig.UnvMasterDbName + " SET AccountStatus= ?,AccountExpiryDate= ? WHERE AccountStatus= ? AND Stakeholder_ID= ?",
		"UNV_CNG_PASS":           "UPDATE " + dbConfig.DbDatabaseName + "." + dbConfig.UnvMasterDbName + " SET UserPassword= ? WHERE Stakeholder_ID= ?",
		"UNV_MBL_VRF_QRY":        "SELECT Primary_Email_Verified,Primary_Phone_Verified,AccountStatus,IF(COUNT(*),'true','false') FROM " + dbConfig.DbDatabaseName + "." + dbConfig.UnvMasterDbName + " WHERE Stakeholder_ID= ? GROUP BY Primary_Email_Verified,Primary_Phone_Verified,AccountStatus LIMIT 1",
		"UNV_VRF_ME_QRY":         "SELECT PrimaryContact_Phone,PrimaryContact_Email,IF(COUNT(*),'true','false') FROM " + dbConfig.DbDatabaseName + "." + dbConfig.UnvMasterDbName + " WHERE Stakeholder_ID= ? GROUP BY PrimaryContact_Phone,PrimaryContact_Email LIMIT 1",
		"UNV_GET_PID":            "SELECT Stakeholder_ID,IF(COUNT(*),'true','false') FROM " + dbConfig.DbDatabaseName + "." + dbConfig.UnvMasterDbName + " WHERE PrimaryContact_Phone=? OR PrimaryContact_Email=? GROUP BY Stakeholder_ID",
		"UNV_LOGIN":              "SELECT Stakeholder_ID,AccountStatus,UserPassword,IF(COUNT(*),'true','false') FROM " + dbConfig.DbDatabaseName + "." + dbConfig.UnvMasterDbName + " WHERE PrimaryContact_Email= ? OR Stakeholder_ID= ? OR PrimaryContact_Phone= ? GROUP BY Stakeholder_ID,AccountStatus,UserPassword LIMIT 1",
		"STU_EXISTS_WITH_EMAIL":  "SELECT Stakeholder_ID,Student_PersonalEmailID,IF(COUNT(*),'true','false') FROM " + dbConfig.DbDatabaseName + "." + dbConfig.StuMasterDbName + " WHERE Student_PersonalEmailID= ? GROUP BY Stakeholder_ID,Student_PersonalEmailID LIMIT 1",
		"STU_INS_NEW_USR":        "INSERT INTO " + dbConfig.DbDatabaseName + "." + dbConfig.StuMasterDbName + " (Stakeholder_ID,Student_FirstName,Student_MiddleName,Student_LastName,Student_PersonalEmailID,Student_PhoneNumber,Student_AlternateContactNumber,Student_Gender,Student_DateOfBirth,Student_AadharNumber,StudentPermanantAddress_Line1,StudentPermanantAddress_Line2,StudentPermanantAddress_Line3,StudentPermanantAddress_Country,StudentPermanantAddress_State,StudentPermanantAddress_City,StudentPermanantAddress_District,StudentPermanantAddress_Zipcode,StudentPermanantAddress_Phone,StudentPresentAddress_Line1,StudentPresentAddress_Line2,StudentPresentAddress_Line3,StudentPresentAddress_Country,StudentPresentAddress_State,StudentPresentAddress_City,StudentPresentAddress_District,StudentPresentAddress_Zipcode,StudentPresentAddress_Phone,University_Name,University_Stakeholder_ID,ProgramName,Program_ID,BranchName,Branch_ID,Student_CollegeID,Student_CollegeEmailID,UserPassword,UniversityApprovedFlag,CreationDate,LastUpdatedDate,AccountStatus,Primary_Phone_Verified,Primary_Email_Verified,AccountExpiryDate,Attachment)  VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)",
		"STU_ROW_CNT":            "SELECT COUNT(*) FROM " + dbConfig.DbDatabaseName + "." + dbConfig.StuMasterDbName + "",
		"STU_ACC_STATUS_UPD":     "UPDATE " + dbConfig.DbDatabaseName + "." + dbConfig.StuMasterDbName + " SET Primary_Email_Verified= ?,Primary_Phone_Verified= ?,AccountStatus= ? WHERE Stakeholder_ID= ?",
		"STU_ACC_ACTIVATION":     "UPDATE " + dbConfig.DbDatabaseName + "." + dbConfig.StuMasterDbName + " SET AccountStatus= ?,AccountExpiryDate= ? WHERE AccountStatus= ? AND Stakeholder_ID= ?",
		"STU_CNG_PASS":           "UPDATE " + dbConfig.DbDatabaseName + "." + dbConfig.StuMasterDbName + " SET UserPassword= ? WHERE Stakeholder_ID= ?",
		"STU_MBL_VRF_QRY":        "SELECT Primary_Email_Verified,Primary_Phone_Verified,AccountStatus,IF(COUNT(*),'true','false') FROM " + dbConfig.DbDatabaseName + "." + dbConfig.StuMasterDbName + " WHERE Stakeholder_ID= ? GROUP BY Primary_Email_Verified,Primary_Phone_Verified,AccountStatus LIMIT 1",
		"STU_VRF_ME_QRY":         "SELECT Student_PhoneNumber,Student_PersonalEmailID,IF(COUNT(*),'true','false') FROM " + dbConfig.DbDatabaseName + "." + dbConfig.StuMasterDbName + " WHERE Stakeholder_ID= ? GROUP BY Student_PhoneNumber,Student_PersonalEmailID LIMIT 1",
		"STU_GET_PID":            "SELECT Stakeholder_ID,IF(COUNT(*),'true','false') FROM " + dbConfig.DbDatabaseName + "." + dbConfig.StuMasterDbName + " WHERE Student_PhoneNumber= ? OR Student_PersonalEmailID= ? GROUP BY Stakeholder_ID",
		"STU_LOGIN":              "SELECT Stakeholder_ID,AccountStatus,UserPassword,IF(COUNT(*),'true','false') FROM " + dbConfig.DbDatabaseName + "." + dbConfig.StuMasterDbName + " WHERE Student_PersonalEmailID= ? OR Stakeholder_ID= ? OR Student_PhoneNumber= ? GROUP BY Stakeholder_ID,AccountStatus,UserPassword LIMIT 1",
		"LUT_GET_CRP_TYPE":       "SELECT OneLtrCode,IF(COUNT(*),'true','false') FROM " + dbConfig.DbDatabaseName + ".LUT_CorporateType WHERE CorporateType=? GROUP BY OneLtrCode LIMIT 1",
		"LUT_GET_CRP_CAT":        "SELECT OneLtrCode,IF(COUNT(*),'true','false') FROM " + dbConfig.DbDatabaseName + ".LUT_CorporateCategory WHERE CorporateCategoryName=? GROUP BY OneLtrCode LIMIT 1",
		"LUT_GET_UNV_CAT":        "SELECT OneLtrCode,IF(COUNT(*),'true','false') FROM " + dbConfig.DbDatabaseName + ".LUT_UniversityType WHERE UniversityTypeName=? GROUP BY OneLtrCode LIMIT 1",
	}
}