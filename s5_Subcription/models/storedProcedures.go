// Package models ...
package models

import (
	"github.com/jaswanth-gorripati/PGK/s5_Subcription/configuration"
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
		"CORP_GET_PROFILE":            "SELECT Stakeholder_ID,Corporate_Name,Corporate_CIN,CorporateHQAddress_Line1,CorporateHQAddress_Line2,CorporateHQAddress_Line3,CorporateHQAddress_Country,CorporateHQAddress_State,CorporateHQAddress_City,CorporateHQAddress_District,CorporateHQAddress_ZipCode,CorporateHQAddress_Phone,CorporateHQAddress_Email,CorporateLocal_BranchAddress_Line1,CorporateLocal_BranchAddress_Line2,CorporateLocal_BranchAddress_Line3,CorporateLocal_BranchAddress_Country,CorporateLocal_BranchAddress_State,CorporateLocal_BranchAddress_City,CorporateLocal_BranchAddress_District,CorporateLocal_BranchAddress_ZipCode,CorporateLocal_BranchAddress_Phone,CorporateLocal_BranchAddress_Email,PrimaryContact_FirstName,PrimaryContact_MiddleName,PrimaryContact_LastName,PrimaryContact_Designation,PrimaryContact_Phone,PrimaryContact_Email,SecondaryContact_FirstName,SecondaryContact_MiddleName,SecondaryContact_LastName,SecondaryContact_Designation,SecondaryContact_Phone,SecondaryContact_Email,CorporateType,CorporateCategory,CorporateIndustry,CompanyProfile,Attachment,YearOfEstablishment,DateOfJoiningPlatform,AccountStatus,Primary_Email_Verified,Primary_Phone_Verified FROM " + dbConfig.DbDatabaseName + "." + dbConfig.CorpMasterDbName + " WHERE Stakeholder_ID= ? GROUP BY Stakeholder_ID,Corporate_Name,Corporate_CIN,CorporateHQAddress_Line1,CorporateHQAddress_Line2,CorporateHQAddress_Line3,CorporateHQAddress_Country,CorporateHQAddress_State,CorporateHQAddress_City,CorporateHQAddress_District,CorporateHQAddress_ZipCode,CorporateHQAddress_Phone,CorporateHQAddress_Email,CorporateLocal_BranchAddress_Line1,CorporateLocal_BranchAddress_Line2,CorporateLocal_BranchAddress_Line3,CorporateLocal_BranchAddress_Country,CorporateLocal_BranchAddress_State,CorporateLocal_BranchAddress_City,CorporateLocal_BranchAddress_District,CorporateLocal_BranchAddress_ZipCode,CorporateLocal_BranchAddress_Phone,CorporateLocal_BranchAddress_Email,PrimaryContact_FirstName,PrimaryContact_MiddleName,PrimaryContact_LastName,PrimaryContact_Designation,PrimaryContact_Phone,PrimaryContact_Email,SecondaryContact_FirstName,SecondaryContact_MiddleName,SecondaryContact_LastName,SecondaryContact_Designation,SecondaryContact_Phone,SecondaryContact_Email,CorporateType,CorporateCategory,CorporateIndustry,CompanyProfile,Attachment,YearOfEstablishment,DateOfJoiningPlatform,AccountStatus,Primary_Email_Verified,Primary_Phone_Verified LIMIT 1",
		"CORP_UPDATE_BY_ID":           "UPDATE " + dbConfig.DbDatabaseName + "." + dbConfig.CorpMasterDbName + " SET ",
		"UNV_GET_PROFILE":             "SELECT Stakeholder_ID,University_Name,University_College_ID,UniversityHQAddress_Line1,UniversityHQAddress_Line2,UniversityHQAddress_Line3,UniversityHQAddress_Country,UniversityHQAddress_State,UniversityHQAddress_City,UniversityHQAddress_District,UniversityHQAddress_Zipcode,UniversityHQAddress_Phone,UniversityHQAddress_Email,UniversityLocal_BranchAddress_Line1,UniversityLocal_BranchAddress_Line2,UniversityLocal_BranchAddress_Line3,UniversityLocal_BranchAddress_Country,UniversityLocal_BranchAddress_State,UniversityLocal_BranchAddress_City,UniversityLocal_BranchAddress_District,UniversityLocal_BranchAddress_Zipcode,UniversityLocal_BranchAddress_Phone,UniversityLocal_BranchAddress_Email,PrimaryContact_FirstName,PrimaryContact_MiddleName,PrimaryContact_LastName,PrimaryContact_Designation,PrimaryContact_Phone,PrimaryContact_Email,SecondaryContact_FirstName,SecondaryContact_MiddleName,SecondaryContact_LastName,SecondaryContact_Designation,SecondaryContact_Phone,SecondaryContact_Email,UniversitySector,UniversityProfile,YearOfEstablishment,Attachment,DateOfJoiningPlatform,Primary_Email_Verified,Primary_Phone_Verified,AccountStatus FROM " + dbConfig.DbDatabaseName + "." + dbConfig.UnvMasterDbName + " WHERE Stakeholder_ID= ? GROUP BY Stakeholder_ID,University_Name,University_College_ID,UniversityHQAddress_Line1,UniversityHQAddress_Line2,UniversityHQAddress_Line3,UniversityHQAddress_Country,UniversityHQAddress_State,UniversityHQAddress_City,UniversityHQAddress_District,UniversityHQAddress_Zipcode,UniversityHQAddress_Phone,UniversityHQAddress_Email,UniversityLocal_BranchAddress_Line1,UniversityLocal_BranchAddress_Line2,UniversityLocal_BranchAddress_Line3,UniversityLocal_BranchAddress_Country,UniversityLocal_BranchAddress_State,UniversityLocal_BranchAddress_City,UniversityLocal_BranchAddress_District,UniversityLocal_BranchAddress_Zipcode,UniversityLocal_BranchAddress_Phone,UniversityLocal_BranchAddress_Email,PrimaryContact_FirstName,PrimaryContact_MiddleName,PrimaryContact_LastName,PrimaryContact_Designation,PrimaryContact_Phone,PrimaryContact_Email,SecondaryContact_FirstName,SecondaryContact_MiddleName,SecondaryContact_LastName,SecondaryContact_Designation,SecondaryContact_Phone,SecondaryContact_Email,UniversitySector,UniversityProfile,YearOfEstablishment,Attachment,DateOfJoiningPlatform,Primary_Email_Verified,Primary_Phone_Verified,AccountStatus LIMIT 1",
		"UNV_UPDATE_BY_ID":            "UPDATE " + dbConfig.DbDatabaseName + "." + dbConfig.UnvMasterDbName + " SET ",
		"STU_GET_PROFILE":             "SELECT Stakeholder_ID,Student_FirstName,Student_MiddleName,Student_LastName,Student_PersonalEmailID,Student_CollegeEmailID,Student_PhoneNumber,Student_AlternateContactNumber,Student_CollegeID,Student_Gender,Student_DateOfBirth,Student_AadharNumber,StudentPermanantAddress_Line1,StudentPermanantAddress_Line2,StudentPermanantAddress_Line3,StudentPermanantAddress_Country,StudentPermanantAddress_State,StudentPermanantAddress_City,StudentPermanantAddress_District,StudentPermanantAddress_Zipcode,StudentPermanantAddress_Phone,StudentPresentAddress_Line1,StudentPresentAddress_Line2,StudentPresentAddress_Line3,StudentPresentAddress_Country,StudentPresentAddress_State,StudentPresentAddress_City,StudentPresentAddress_District,StudentPresentAddress_Zipcode,StudentPresentAddress_Phone,Father_Guardian_FullName,Father_Guardian_Occupation,Father_Guardian_Company,Father_Guardian_PhoneNumber,Father_Guardian_EmailID,Mother_Guardian_FullName,Mother_Guardian_Occupation,Mother_Guardian_Comany,Mother_Guardian_Designation,Mother_Guardian_PhoneNumber,Mother_Guardian_EmailID,CreationDate,Primary_Email_Verified,Primary_Phone_Verified,AccountStatus FROM " + dbConfig.DbDatabaseName + "." + dbConfig.StuMasterDbName + " WHERE Stakeholder_ID= ? GROUP BY Stakeholder_ID,Student_FirstName,Student_MiddleName,Student_LastName,Student_PersonalEmailID,Student_CollegeEmailID,Student_PhoneNumber,Student_AlternateContactNumber,Student_CollegeID,Student_Gender,Student_DateOfBirth,Student_AadharNumber,StudentPermanantAddress_Line1,StudentPermanantAddress_Line2,StudentPermanantAddress_Line3,StudentPermanantAddress_Country,StudentPermanantAddress_State,StudentPermanantAddress_City,StudentPermanantAddress_District,StudentPermanantAddress_Zipcode,StudentPermanantAddress_Phone,StudentPresentAddress_Line1,StudentPresentAddress_Line2,StudentPresentAddress_Line3,StudentPresentAddress_Country,StudentPresentAddress_State,StudentPresentAddress_City,StudentPresentAddress_District,StudentPresentAddress_Zipcode,StudentPresentAddress_Phone,Father_Guardian_FullName,Father_Guardian_Occupation,Father_Guardian_Company,Father_Guardian_PhoneNumber,Father_Guardian_EmailID,Mother_Guardian_FullName,Mother_Guardian_Occupation,Mother_Guardian_Comany,Mother_Guardian_Designation,Mother_Guardian_PhoneNumber,Mother_Guardian_EmailID,CreationDate,Primary_Email_Verified,Primary_Phone_Verified,AccountStatus LIMIT 1",
		"STU_UPDATE_BY_ID":            "UPDATE " + dbConfig.DbDatabaseName + "." + dbConfig.StuMasterDbName + " SET ",
		"PROFILE_UPDATE_WHERE":        " WHERE Stakeholder_ID= ?",
		"UPLOAD_PROFILE_PIC":          "INSERT INTO " + dbConfig.DbDatabaseName + "." + dbConfig.ProfilePicDbName + " (Stakeholder_ID,Profile_Pic) VALUES(?,?) ON DUPLICATE KEY UPDATE Profile_Pic= ?",
		"GET_PRF_PIC":                 "SELECT Profile_Pic FROM " + dbConfig.DbDatabaseName + "." + dbConfig.ProfilePicDbName + " WHERE Stakeholder_ID= ? GROUP BY Profile_Pic LIMIT 1",
		"GET_LUT_CRP_TPY":             "SELECT CorporateType,CorporateTypeCode,OneLtrCode FROM " + dbConfig.DbDatabaseName + "." + dbConfig.LutCrpTyp + " GROUP BY CorporateType,CorporateTypeCode,OneLtrCode",
		"GET_LUT_CRP_CAT":             "SELECT CorporateCategoryName,CorporateCategoryCode,OneLtrCode FROM " + dbConfig.DbDatabaseName + "." + dbConfig.LutCrpCat + " GROUP BY CorporateCategoryName,CorporateCategoryCode,OneLtrCode",
		"GET_LUT_CRP_IND":             "SELECT CorporateIndutryName,CorporateIndustryCode FROM " + dbConfig.DbDatabaseName + "." + dbConfig.LutCrpInd + " GROUP BY CorporateIndutryName,CorporateIndustryCode",
		"GET_LUT_UNV_CAT":             "SELECT UniversityTypeName,UniversityTypeCode,OneLtrCode FROM " + dbConfig.DbDatabaseName + "." + dbConfig.LutUnvCat + " GROUP BY UniversityTypeName,UniversityTypeCode,OneLtrCode",
		"GET_LUT_SKILLS":              "SELECT Skill_ID,SkillName FROM " + dbConfig.DbDatabaseName + "." + dbConfig.LutSkills + " GROUP BY Skill_ID,SkillName",
		"GET_LUT_PROGRAMS":            "SELECT Program_ID,ProgramName FROM " + dbConfig.DbDatabaseName + "." + dbConfig.LutPrograms + " GROUP BY Program_ID,ProgramName",
		"GET_LUT_DEPART":              "SELECT Branch_ID,Program_ID,BranchName FROM " + dbConfig.DbDatabaseName + "." + dbConfig.LutDepartments + " GROUP BY Branch_ID,Program_ID,BranchName",
		"UNV_SUB_INS":                 "INSERT INTO " + dbConfig.DbDatabaseName + "." + dbConfig.UnvSubDBName + " (Subscriber_Stakeholder_ID,Publisher_Stakeholder_ID,DateOfSubscription,Publish_ID,Transaction_ID) VALUES",
		"UNV_GET_SH_PUB":              "SELECT Stakeholder_ID FROM " + dbConfig.DbDatabaseName + "." + dbConfig.UnvPubDBName + " WHERE Publish_ID=? ",
		"UNV_GET_ALL_SUBS":            "SELECT a.Publisher_Stakeholder_ID,a.DateOfSubscription,a.Publish_ID,a.Transaction_ID,c.Corporate_Name,d.GeneralNote FROM " + dbConfig.DbDatabaseName + "." + dbConfig.UnvSubDBName + " as a," + dbConfig.DbDatabaseName + "." + dbConfig.CorpMasterDbName + " AS c," + dbConfig.DbDatabaseName + "." + dbConfig.CrpPubDBName + " AS d where c.Stakeholder_ID = a.Publisher_Stakeholder_ID AND d.Publish_ID=a.Publish_ID AND  a.Subscriber_Stakeholder_ID=?",
		"CRP_SUB_INS":                 "INSERT INTO " + dbConfig.DbDatabaseName + "." + dbConfig.CrpSubDBName + " (Subscriber_Stakeholder_ID,Publisher_Stakeholder_ID,DateOfSubscription,Publish_ID,Transaction_ID) VALUES",
		"CRP_GET_SH_PUB":              "SELECT Stakeholder_ID FROM " + dbConfig.DbDatabaseName + "." + dbConfig.CrpPubDBName + " WHERE Publish_ID=? ",
		"CRP_GET_ALL_SUBS":            "SELECT a.Publisher_Stakeholder_ID,a.DateOfSubscription,a.Publish_ID,a.Transaction_ID,c.University_Name,d.GeneralNote FROM " + dbConfig.DbDatabaseName + "." + dbConfig.CrpSubDBName + " as a," + dbConfig.DbDatabaseName + "." + dbConfig.UnvMasterDbName + " AS c," + dbConfig.DbDatabaseName + "." + dbConfig.UnvPubDBName + " AS d where c.Stakeholder_ID = a.Publisher_Stakeholder_ID AND d.Publish_ID=a.Publish_ID AND  a.Subscriber_Stakeholder_ID=?",
		"STU_GET_ALL_SUBS":            "SELECT a.Publisher_Stakeholder_ID,a.DateOfSubscription,a.Publish_ID,a.Transaction_ID,c.Corporate_Name,d.GeneralNote FROM " + dbConfig.DbDatabaseName + "." + dbConfig.StuSubDBName + " as a," + dbConfig.DbDatabaseName + "." + dbConfig.CorpMasterDbName + " AS c," + dbConfig.DbDatabaseName + "." + dbConfig.CrpPubDBName + " AS d where c.Stakeholder_ID = a.Publisher_Stakeholder_ID AND d.Publish_ID=a.Publish_ID AND  a.Subscriber_Stakeholder_ID=?",
		"STU_SUB_INS":                 "INSERT INTO " + dbConfig.DbDatabaseName + "." + dbConfig.StuSubDBName + " (Subscriber_Stakeholder_ID,Publisher_Stakeholder_ID,DateOfSubscription,Publish_ID,Transaction_ID) VALUES",
		"STU_GET_SH_PUB":              "SELECT Stakeholder_ID FROM " + dbConfig.DbDatabaseName + "." + dbConfig.StuPubDBName + " WHERE Publish_ID=? ",
		"CRP_CAMPUS_DRIVE_REQ":        "INSERT INTO " + dbConfig.DbDatabaseName + "." + dbConfig.CorpCDDbName + " (Initiator_Stakeholder_ID,Receiver_Stakeholder_ID,CampusDrive_ID,CampusDrive_Requested,CampusDrive_Requested_Date,CampusDrive_Requested_NotificationID) VALUES(?,?,?,?,?,?)",
		"UNV_CAMPUS_DRIVE_REQ":        "INSERT INTO " + dbConfig.DbDatabaseName + "." + dbConfig.UnvCPDbName + " (Initiator_Stakeholder_ID,Receiver_Stakeholder_ID,CampusDrive_ID,CampusDrive_Requested,CampusDrive_Requested_Date,CampusDrive_Requested_NotificationID) VALUES(?,?,?,?,?,?)",
		"CRP_CD_INVITE_UPDATE":        "UPDATE " + dbConfig.DbDatabaseName + "." + dbConfig.CorpCDDbName + " SET CampusDrive_AcceptedorRejectedbyUniv=?,CampusDrive_AcceptedorRejectedbyUniv_Date=?,CampusDrive_AcceptedorRejectedbyUniv_NotificationID=?,CampusDrice_AcceptedorRejectedbyUniv_Reason=? WHERE Receiver_Stakeholder_ID=? AND CampusDrive_ID=?",
		"UNV_CD_INVITE_UPDATE":        "UPDATE " + dbConfig.DbDatabaseName + "." + dbConfig.UnvCPDbName + " SET CampusDrive_AcceptedorRejectedbyCorp=?,CampusDrive_AcceptedorRejectedbyCorp_Date=?,CampusDrive_AcceptedorRejectedbyCorp_NotificationID=?,CampusDrice_AcceptedorRejectedbyCorp_Reason=? WHERE Receiver_Stakeholder_ID=? AND CampusDrive_ID=?",
		"UNV_INSIGHTS_INS":            "INSERT INTO " + dbConfig.DbDatabaseName + "." + dbConfig.UnvInsightsDbName + " (Subscription_ID,Subscriber_Stakeholder_ID,Subscribed_Stakeholder_ID,AverageCGPA,AveragePercentage,HighestCGPA,HighestPercentage,HighestPackageReceived,AveragePackageReceived,UniversityConversionRatio,TentativeMonthofPassing,Top5Recruiters,Top5Skills,SubscribedDate,CreationDate,LastUpdatedDate) VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)",
		"UNV_INSIGHTS_Get_Last_ID":    "SELECT Subscription_ID FROM " + dbConfig.DbDatabaseName + "." + dbConfig.UnvInsightsDbName + " where Subscribed_Stakeholder_ID=? ORDER BY Subscription_ID DESC LIMIT 1",
		"UNV_INSIGHTS_GET_ALL":        "SELECT Subscription_ID,Subscriber_Stakeholder_ID,Subscribed_Stakeholder_ID,AverageCGPA,AveragePercentage,HighestCGPA,HighestPercentage,HighestPackageReceived,AveragePackageReceived,UniversityConversionRatio,TentativeMonthofPassing,Top5Recruiters,Top5Skills,SubscribedDate,CreationDate,LastUpdatedDate FROM " + dbConfig.DbDatabaseName + "." + dbConfig.UnvInsightsDbName + " Where Subscribed_Stakeholder_ID=? GROUP BY Subscription_ID,Subscriber_Stakeholder_ID,Subscribed_Stakeholder_ID,AverageCGPA,AveragePercentage,HighestCGPA,HighestPercentage,HighestPackageReceived,AveragePackageReceived,UniversityConversionRatio,TentativeMonthofPassing,Top5Recruiters,Top5Skills,SubscribedDate,CreationDate,LastUpdatedDate Order by SubscribedDate DESC",
		"UNV_INSIGHTS_GET_SUB":        "SELECT Subscription_ID,Subscriber_Stakeholder_ID,Subscribed_Stakeholder_ID,AverageCGPA,AveragePercentage,HighestCGPA,HighestPercentage,HighestPackageReceived,AveragePackageReceived,UniversityConversionRatio,TentativeMonthofPassing,Top5Recruiters,Top5Skills,SubscribedDate,CreationDate,LastUpdatedDate FROM " + dbConfig.DbDatabaseName + "." + dbConfig.UnvInsightsDbName + " Where Subscribed_Stakeholder_ID=? AND Subscription_ID=? GROUP BY Subscription_ID,Subscriber_Stakeholder_ID,Subscribed_Stakeholder_ID,AverageCGPA,AveragePercentage,HighestCGPA,HighestPercentage,HighestPackageReceived,AveragePackageReceived,UniversityConversionRatio,TentativeMonthofPassing,Top5Recruiters,Top5Skills,SubscribedDate,CreationDate,LastUpdatedDate Order by SubscribedDate DESC",
		"UNV_STU_DB_SUB_INIT":         "INSERT INTO " + dbConfig.DbDatabaseName + "." + dbConfig.UnvStuDataDbName + " (Subscription_ID,Subscriber_Stakeholder_ID,Subscribed_Stakeholder_ID,Student_CollegeRollNo,Student_ProgramName,Student_Program_ID,Student_BranchName,Student_BranchID,AverageCGPA,AveragePercentage,Student_Stakeholder_ID,SubscribedDate,Subscription_ValidityFlag,CreationDate,LastUpdatedDate) VALUES",
		"UNV_STU_DB_SUB_GET_ALL":      "SELECT Subscription_ID,Subscriber_Stakeholder_ID,Subscribed_Stakeholder_ID,Student_CollegeRollNo,Student_ProgramName,Student_Program_ID,Student_BranchName,Student_BranchID,ifnull(AverageCGPA,70),ifnull(AveragePercentage,80),Student_Stakeholder_ID,SubscribedDate,Subscription_ValidityFlag,CreationDate,LastUpdatedDate FROM " + dbConfig.DbDatabaseName + "." + dbConfig.UnvStuDataDbName + " WHERE Subscription_ID =? AND Subscribed_Stakeholder_ID=? ",
		"UNV_STU_DB_VAL_INS":          "SELECT Subscription_ID,Subscriber_Stakeholder_ID,Subscribed_Stakeholder_ID,Student_CollegeRollNo,count(*) FROM " + dbConfig.DbDatabaseName + "." + dbConfig.UnvStuDataDbName + " WHERE Subscription_ID=? AND Subscribed_Stakeholder_ID=? AND Subscriber_Stakeholder_ID=? ",
		"UNV_STU_DB_DEL_BFR_UPDATE":   "DELETE FROM " + dbConfig.DbDatabaseName + "." + dbConfig.UnvStuDataDbName + " WHERE Subscription_ID=? AND Subscribed_Stakeholder_ID=? AND Subscriber_Stakeholder_ID=? ",
		"UNV_STU_DB_Get_Last_ID":      "SELECT Subscription_ID FROM " + dbConfig.DbDatabaseName + "." + dbConfig.UnvStuDataDbName + " where Subscribed_Stakeholder_ID=? ORDER BY Subscription_ID DESC LIMIT 1",
		"UNV_GET_STU_DATA_FOR_QUERY":  "SELECT a.Stakeholder_ID,ifnull(((Tenth_Percentage+Twelfth_Percentage)/2),70) as avgPercent,ifnull(Grad_ProgramName,'Master of Technology') as Program,ifnull(Grad_Program_ID,'MTech') ProgramID,ifnull(Grad_BranchName,'Bachelor of Technology (Food Science)')as BranchName,ifnull(Grad_Branch_ID,'1') as BranchID,ifnull(Grad_CollegeRollNumber,(SELECT Student_CollegeID FROM CollabToHire.Student_Master where Stakeholder_ID=a.Stakeholder_ID)) as CollegeID FROM CollabToHire.Stud_Grades as a,CollabToHire.Student_Master as b Where a.Stakeholder_ID=b.Stakeholder_ID AND b.University_Stakeholder_ID=?",
		"UNV_INSIGHTS_GET_ALL_SUB":    "SELECT a.Subscription_ID,a.Subscriber_Stakeholder_ID,a.Subscribed_Stakeholder_ID,a.SubscribedDate,b.University_Name,CONCAT(b.UniversityHQAddress_City,',',b.UniversityLocal_BranchAddress_City,',') FROM " + dbConfig.DbDatabaseName + "." + dbConfig.UnvInsightsDbName + " as a," + dbConfig.DbDatabaseName + "." + dbConfig.UnvMasterDbName + " as b Where Subscribed_Stakeholder_ID=?  Order by SubscribedDate DESC",
		"UNV_STU_DB_SUB_GET_ALL_SUB":  "SELECT DISTINCT Subscription_ID,Subscriber_Stakeholder_ID,Subscribed_Stakeholder_ID,SubscribedDate,b.University_Name,CONCAT(b.UniversityHQAddress_City,',',b.UniversityLocal_BranchAddress_City,',') FROM " + dbConfig.DbDatabaseName + "." + dbConfig.UnvStuDataDbName + " as a," + dbConfig.DbDatabaseName + "." + dbConfig.UnvMasterDbName + " as b WHERE  Subscribed_Stakeholder_ID=?   Order by SubscribedDate DESC ",
		"CORP_CD_INIT":                "INSERT INTO " + dbConfig.DbDatabaseName + "." + dbConfig.CorpCDDbName + " (Initiator_Stakeholder_ID,Receiver_Stakeholder_ID,CampusDrive_ID,CampusDrive_Requested,CampusDrive_Requested_Date,CampusDrive_Requested_NotificationID,CreationDate,LastUpdatedDate) VALUES(?,?,?,?,?,?,?,?)",
		"CORP_CD_SUB_UPDATE":          "UPDATE " + dbConfig.DbDatabaseName + "." + dbConfig.CorpCDDbName + " SET CampusDrive_Requested=?,CampusDrive_Requested_Date=?,CampusDrive_Requested_NotificationID=?,LastUpdatedDate=? WHERE Initiator_Stakeholder_ID=? AND CampusDrive_ID=?",
		"CORP_CD_UNV_RESP":            "UPDATE " + dbConfig.DbDatabaseName + "." + dbConfig.CorpCDDbName + " SET CampusDrive_AcceptedorRejectedbyUniv=?,CampusDrive_AcceptedorRejectedbyUniv_Date=?,CampusDrive_AcceptedorRejectedbyUniv_NotificationID=?,CampusDrice_AcceptedorRejectedbyUniv_Reason=?,LastUpdatedDate=? WHERE Receiver_Stakeholder_ID=? AND CampusDrive_ID=?",
		"CORP_CD_GET_BY_ID":           "SELECT Initiator_Stakeholder_ID,Receiver_Stakeholder_ID,CampusDrive_ID,CampusDrive_Requested,CampusDrive_Requested_Date,CampusDrive_Requested_NotificationID,CampusDrive_AcceptedorRejectedbyUniv,CampusDrive_AcceptedorRejectedbyUniv_Date,CampusDrive_AcceptedorRejectedbyUniv_NotificationID,CampusDrice_AcceptedorRejectedbyUniv_Reason,CreationDate,LastUpdatedDate FROM " + dbConfig.DbDatabaseName + "." + dbConfig.CorpCDDbName + " WHERE CampusDrive_ID=? AND (Initiator_Stakeholder_ID=? OR Receiver_Stakeholder_ID=? )",
		"CORP_CD_GET_INITIATOR_FR_CD": "SELECT Initiator_Stakeholder_ID,Receiver_Stakeholder_ID FROM " + dbConfig.DbDatabaseName + "." + dbConfig.CorpCDDbName + " WHERE CampusDrive_ID=? AND (Initiator_Stakeholder_ID=? OR Receiver_Stakeholder_ID=? )",
		"CORP_CD_Get_Last_ID":         "SELECT Initiator_Stakeholder_ID FROM " + dbConfig.DbDatabaseName + "." + dbConfig.CorpCDDbName + " where Initiator_Stakeholder_ID=? ORDER BY CampusDrive_ID DESC LIMIT 1",
		"UNV_CD_INIT":                 "INSERT INTO " + dbConfig.DbDatabaseName + "." + dbConfig.UnvCPDbName + " (Initiator_Stakeholder_ID,Receiver_Stakeholder_ID,CampusDrive_ID,CampusDrive_Requested,CampusDrive_Requested_Date,CampusDrive_Requested_NotificationID,CreationDate,LastUpdatedDate) VALUES(?,?,?,?,?,?,?,?)",
		"UNV_CD_Get_Last_ID":          "SELECT Initiator_Stakeholder_ID FROM " + dbConfig.DbDatabaseName + "." + dbConfig.UnvCPDbName + " where Initiator_Stakeholder_ID=? ORDER BY CampusDrive_ID DESC LIMIT 1",
		"UNV_CD_SUB_UPDATE":           "UPDATE " + dbConfig.DbDatabaseName + "." + dbConfig.UnvCPDbName + " SET CampusDrive_Requested=?,CampusDrive_Requested_Date=?,CampusDrive_Requested_NotificationID=?,LastUpdatedDate=? WHERE Initiator_Stakeholder_ID=? AND CampusDrive_ID=?",
		"UNV_CD_UNV_RESP":             "UPDATE " + dbConfig.DbDatabaseName + "." + dbConfig.UnvCPDbName + " SET CampusDrive_AcceptedorRejectedbyCorp=?,CampusDrive_AcceptedorRejectedbyCorp_Date=?,CampusDrive_AcceptedorRejectedbyCorp_NotificationID=?,CampusDrice_AcceptedorRejectedbyCorp_Reason=?,LastUpdatedDate=? WHERE Receiver_Stakeholder_ID=? AND CampusDrive_ID=?",
		"UNV_CD_GET_BY_ID":            "SELECT Initiator_Stakeholder_ID,Receiver_Stakeholder_ID,CampusDrive_ID,CampusDrive_Requested,CampusDrive_Requested_Date,CampusDrive_Requested_NotificationID,CampusDrive_AcceptedorRejectedbyCorp,CampusDrive_AcceptedorRejectedbyCorp_Date,CampusDrive_AcceptedorRejectedbyCorp_NotificationID,CampusDrice_AcceptedorRejectedbyCorp_Reason,CreationDate,LastUpdatedDate FROM " + dbConfig.DbDatabaseName + "." + dbConfig.UnvCPDbName + " WHERE CampusDrive_ID=? AND (Initiator_Stakeholder_ID=? OR Receiver_Stakeholder_ID=? )",
		"UNV_CD_GET_INITIATOR_FR_CD":  "SELECT Initiator_Stakeholder_ID,Receiver_Stakeholder_ID FROM " + dbConfig.DbDatabaseName + "." + dbConfig.UnvCPDbName + " WHERE CampusDrive_ID=? AND (Initiator_Stakeholder_ID=? OR Receiver_Stakeholder_ID=? )",
	}
}

// `CampmusDrivePlanned_Date`
//
