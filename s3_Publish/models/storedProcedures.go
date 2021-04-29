// Package models ...
package models

import (
	"fmt"

	"github.com/jaswanth-gorripati/PGK/s3_Publish/configuration"
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

func getCmd() string {
	return fmt.Sprint("'") + fmt.Sprint('"') + "noOfPositions" + fmt.Sprint('"') + ": " + fmt.Sprint('"') + "', jc.NoOfPositions,'" + fmt.Sprint('"') + "'), CONCAT('" + fmt.Sprint('"') + "dateOfHiring" + fmt.Sprint('"') + ": " + fmt.Sprint('"') + "', jc.DateOfHiring, '" + fmt.Sprint('"') + "'), CONCAT('" + fmt.Sprint('"') + "Location" + fmt.Sprint('"') + ": " + fmt.Sprint('"') + "', jc.Location, '" + fmt.Sprint('"') + "'),CONCAT('" + fmt.Sprint('"') + "skillName" + fmt.Sprint('"') + ": " + fmt.Sprint('"') + "', jc.SkillName, '" + fmt.Sprint('"') + "'),CONCAT('" + fmt.Sprint('"') + "salaryRange" + fmt.Sprint('"') + ": " + fmt.Sprint('"') + "', jc.SalaryRange, '" + fmt.Sprint('"') + "')"
}

// CreateSP Creates default stored procedures for Database
func CreateSP() {
	dbConfig := configuration.DbConfig()
	SP = map[string]string{
		"HC_INS_NEW":                    "INSERT INTO " + dbConfig.DbDatabaseName + "." + dbConfig.CorpHCDbName + " (HiringCriteria_ID,HiringCriteria_Name,Stakeholder_ID,Program_ID,Branch_ID,MinimumCutoffCategory,MinimumCutoff,EduGaps_School_Allowed,EduGaps_11N12_Allowed,EduGaps_Grad_Allowed,EduGaps_PG_Allowed,ActiveBacklogsAllowed,TotalNumberOfBacklogsAllowed,YearOfPassing,Remarks) VALUES",
		"HC_Rowcount":                   "SELECT COUNT(*) FROM " + dbConfig.DbDatabaseName + "." + dbConfig.CorpHCDbName + " WHERE Stakeholder_ID= ?",
		"HC_Get_Last_ID":                "SELECT HiringCriteria_ID FROM " + dbConfig.DbDatabaseName + "." + dbConfig.CorpHCDbName + " where Stakeholder_ID=? ORDER BY HiringCriteria_ID DESC LIMIT 1",
		"HC_GET_BY_ID":                  "SELECT HiringCriteria_ID,HiringCriteria_Name,Program_ID,(select BranchName as b from CollabToHire.LUT_BranchCatalog as b where b.Branch_ID=hc.Branch_ID),MinimumCutoffCategory,MinimumCutoff,EduGaps_School_Allowed,EduGaps_11N12_Allowed,EduGaps_Grad_Allowed,EduGaps_PG_Allowed,ActiveBacklogsAllowed,TotalNumberOfBacklogsAllowed,YearOfPassing,Remarks,CreationDate,PublishFlag,PublishID FROM " + dbConfig.DbDatabaseName + "." + dbConfig.CorpHCDbName + " as hc WHERE HiringCriteria_ID= ? GROUP BY HiringCriteria_ID,HiringCriteria_Name,Program_ID,Branch_ID,MinimumCutoffCategory,MinimumCutoff,EduGaps_School_Allowed,EduGaps_11N12_Allowed,EduGaps_Grad_Allowed,EduGaps_PG_Allowed,ActiveBacklogsAllowed,TotalNumberOfBacklogsAllowed,YearOfPassing,Remarks,CreationDate,PublishFlag,PublishID LIMIT 1",
		"HC_GET_ALL":                    "SELECT HiringCriteria_ID,HiringCriteria_Name,Program_ID,(select BranchName as b from CollabToHire.LUT_BranchCatalog as b where b.Branch_ID=hc.Branch_ID),MinimumCutoffCategory,MinimumCutoff,EduGaps_School_Allowed,EduGaps_11N12_Allowed,EduGaps_Grad_Allowed,EduGaps_PG_Allowed,ActiveBacklogsAllowed,TotalNumberOfBacklogsAllowed,YearOfPassing,Remarks,CreationDate,PublishFlag,PublishID FROM " + dbConfig.DbDatabaseName + "." + dbConfig.CorpHCDbName + " as hc WHERE Stakeholder_ID= ?  GROUP BY HiringCriteria_ID,HiringCriteria_Name,Program_ID,Branch_ID,MinimumCutoffCategory,MinimumCutoff,EduGaps_School_Allowed,EduGaps_11N12_Allowed,EduGaps_Grad_Allowed,EduGaps_PG_Allowed,ActiveBacklogsAllowed,TotalNumberOfBacklogsAllowed,YearOfPassing,Remarks,CreationDate,PublishFlag,PublishID ORDER BY CreationDate",
		"HC_GET_ALL_PUB":                "SELECT HiringCriteria_ID,HiringCriteria_Name,Program_ID,(select BranchName as b from CollabToHire.LUT_BranchCatalog as b where b.Branch_ID=hc.Branch_ID),MinimumCutoffCategory,MinimumCutoff,EduGaps_School_Allowed,EduGaps_11N12_Allowed,EduGaps_Grad_Allowed,EduGaps_PG_Allowed,ActiveBacklogsAllowed,TotalNumberOfBacklogsAllowed,YearOfPassing,Remarks,CreationDate,PublishFlag,PublishID FROM " + dbConfig.DbDatabaseName + "." + dbConfig.CorpHCDbName + " as hc WHERE Stakeholder_ID= ? AND PublishFlag=1  GROUP BY HiringCriteria_ID,HiringCriteria_Name,Program_ID,Branch_ID,MinimumCutoffCategory,MinimumCutoff,EduGaps_School_Allowed,EduGaps_11N12_Allowed,EduGaps_Grad_Allowed,EduGaps_PG_Allowed,ActiveBacklogsAllowed,TotalNumberOfBacklogsAllowed,YearOfPassing,Remarks,CreationDate,PublishFlag,PublishID ORDER BY CreationDate",
		"HC_UPDATE_BY_ID":               "UPDATE " + dbConfig.DbDatabaseName + "." + dbConfig.CorpHCDbName + " SET ",
		"HC_UPDATE_WHERE":               " WHERE HiringCriteria_ID= ? AND Stakeholder_ID= ?",
		"HC_DELETE_BY_ID":               "DELETE FROM " + dbConfig.DbDatabaseName + "." + dbConfig.CorpHCDbName + " WHERE HiringCriteria_ID= ? AND Stakeholder_ID= ?",
		"JOB_HC_MAP_INS":                "INSERT INTO " + dbConfig.DbDatabaseName + "." + dbConfig.JobHcDbName + " (Job_ID,Stakeholder_ID,HiringCriteria_ID,HiringCriteria_Name,JobName) VALUES(?,?,?,?,?)",
		"JOB_HC_MAP_UPD":                "INSERT INTO " + dbConfig.DbDatabaseName + "." + dbConfig.JobHcDbName + " (Job_ID,Stakeholder_ID,HiringCriteria_ID,HiringCriteria_Name,JobName) VALUES(?,?,?,?,?)",
		"JOB_HC_Last_ID":                "SELECT Job_ID FROM " + dbConfig.DbDatabaseName + "." + dbConfig.JobHcDbName + " WHERE Stakeholder_ID= ? ORDER BY Job_ID DESC LIMIT 1",
		"JOB_HC_GET_BY_ID":              "SELECT Job_ID,Stakeholder_ID,HiringCriteria_ID,HiringCriteria_Name,JobName,CreationDate,ifnull(PublishFlag,0),ifnull(Publish_ID,'') FROM " + dbConfig.DbDatabaseName + "." + dbConfig.JobHcDbName + " WHERE Job_ID= ? GROUP BY Job_ID,Stakeholder_ID,HiringCriteria_ID,HiringCriteria_Name,JobName,CreationDate,PublishFlag,Publish_ID ORDER BY CreationDate LIMIT 1",
		"JOB_HC_GET_CRT_DATE":           "SELECT CreationDate FROM " + dbConfig.DbDatabaseName + "." + dbConfig.JobHcDbName + " WHERE Job_ID= ? GROUP BY CreationDate ORDER BY CreationDate LIMIT 1",
		"JOB_HC_GETALL_BY_ID":           "SELECT a.Job_ID,a.HiringCriteria_ID,a.HiringCriteria_Name,a.JobName,a.CreationDate,ifnull(a.PublishFlag,0),ifnull(a.Publish_ID,'') FROM " + dbConfig.DbDatabaseName + "." + dbConfig.JobHcDbName + " as a  WHERE Stakeholder_ID= ? GROUP BY Job_ID,HiringCriteria_ID,HiringCriteria_Name,JobName,CreationDate,PublishFlag,Publish_ID ORDER BY CreationDate",
		"JOB_UPDATE_BY_ID":              "UPDATE " + dbConfig.DbDatabaseName + "." + dbConfig.JobHcDbName + " SET ",
		"JOB_UPDATE_WHERE":              " WHERE Job_ID= ? AND Stakeholder_ID= ?",
		"JOB_PUBLISH_BY_ID":             "UPDATE " + dbConfig.DbDatabaseName + "." + dbConfig.JobHcDbName + " SET Publish_ID=?, PublishFlag=1, LastUpdatedDate=?  WHERE Job_ID= ? AND Stakeholder_ID= ? ",
		"JOB_UPD_HC_MAP":                "UPDATE " + dbConfig.DbDatabaseName + "." + dbConfig.JobHcDbName + " SET HiringCriteria_ID= ?,HiringCriteria_Name= ? WHERE Job_ID= ? AND Stakeholder_ID= ?",
		"JS_DELETE_BY_ID":               "DELETE FROM " + dbConfig.DbDatabaseName + "." + dbConfig.JobHcDbName + " WHERE Job_ID= ? AND Stakeholder_ID= ?",
		"JOB_SKill_MAP_INS":             "INSERT INTO " + dbConfig.DbDatabaseName + "." + dbConfig.JobSkillDbName + " (Job_ID,JobName,Stakeholder_ID,Skill_ID,SkillName,NoOfPositions,Location,SalaryRange,DateOfHiring,Status,Remarks,AttachFile) VALUES",
		"JOB_SKill_MAP_UPD":             "UPDATE " + dbConfig.DbDatabaseName + "." + dbConfig.JobSkillDbName + " SET ",
		"JOB_SKill_MAP_WHR":             " WHERE id= ? AND Stakeholder_ID= ?",
		"JOB_SKill_GET_BY_ID":           "SELECT id,Job_ID,JobName,Skill_ID,SkillName,NoOfPositions,Location,SalaryRange,DateOfHiring,Status,Remarks,AttachFile,CreationDate FROM " + dbConfig.DbDatabaseName + "." + dbConfig.JobSkillDbName + " WHERE Job_ID= ? GROUP BY id,Job_ID,JobName,Skill_ID,SkillName,NoOfPositions,Location,SalaryRange,DateOfHiring,Status,Remarks,AttachFile,CreationDate ORDER BY CreationDate",
		"JS_SM_DELETE_BY_ID":            "DELETE FROM " + dbConfig.DbDatabaseName + "." + dbConfig.JobSkillDbName + " WHERE id=? AND Job_ID= ? AND Stakeholder_ID= ?",
		"JS_SM_DELETE_All":              "DELETE FROM " + dbConfig.DbDatabaseName + "." + dbConfig.JobSkillDbName + " WHERE Job_ID= ? AND Stakeholder_ID= ?",
		"PJ_INS_NEW":                    "INSERT INTO " + dbConfig.DbDatabaseName + "." + dbConfig.CorpPJDbName + " (Publish_ID,Job_ID,JobName,Stakeholder_ID) VALUES",
		"PJ_GET_BY_ID":                  "SELECT Publish_ID,Job_ID,JobName,Stakeholder_ID,CreationDate FROM " + dbConfig.DbDatabaseName + "." + dbConfig.JobHcDbName + " WHERE Publish_ID= ?  GROUP BY Publish_ID,Job_ID,JobName,Stakeholder_ID,CreationDate LIMIT 1",
		"PJ_GET_ALL":                    "SELECT Publish_ID,Job_ID,JobName,Stakeholder_ID,CreationDate  FROM " + dbConfig.DbDatabaseName + "." + dbConfig.JobHcDbName + " WHERE Stakeholder_ID= ? AND PublishFlag=1 GROUP BY Publish_ID,Job_ID,JobName,Stakeholder_ID,CreationDate ORDER BY Publish_ID",
		"PJ_UPDATE_BY_ID":               "UPDATE " + dbConfig.DbDatabaseName + "." + dbConfig.CorpPJDbName + " SET ",
		"PJ_UPDATE_WHERE":               " WHERE Publish_ID= ? AND Stakeholder_ID= ?",
		"PJ_DELETE_BY_ID":               "UPDATE " + dbConfig.DbDatabaseName + "." + dbConfig.JobHcDbName + " SET Publish_ID=NULL,PublishFlag=0  WHERE Publish_ID=? AND Stakeholder_ID=? ",
		"PJ_Get_Last_ID":                "SELECT Publish_ID FROM " + dbConfig.DbDatabaseName + "." + dbConfig.CorpPJDbName + " where Stakeholder_ID=? ORDER BY Publish_ID DESC LIMIT 1",
		"PDH_INS_NEW":                   "INSERT INTO " + dbConfig.DbDatabaseName + "." + dbConfig.CorpPDHDbName + " (Stakeholder_ID,Publish_ID,DateOfPublish,HiringCriteria_Published,Jobs_Published,Profile_Published,Other_Published,GeneralNote,CreationDate,LastUpdatedDate,PublishData_JSON) VALUES ",
		"CRP_PDH_GET_ALL":               "SELECT Publish_ID,DateOfPublish,HiringCriteria_Published,Jobs_Published,Profile_Published,Other_Published,GeneralNote,CreationDate,LastUpdatedDate,PublishData_JSON FROM " + dbConfig.DbDatabaseName + "." + dbConfig.CorpPDHDbName + " where Stakeholder_ID= ? GROUP BY Publish_ID,DateOfPublish,HiringCriteria_Published,Jobs_Published,Profile_Published,Other_Published,GeneralNote,CreationDate,LastUpdatedDate,PublishData_JSON ORDER BY CreationDate DESC ",
		"PDH_Get_Last_ID":               "SELECT Publish_ID FROM " + dbConfig.DbDatabaseName + "." + dbConfig.CorpPDHDbName + " where Stakeholder_ID=? ORDER BY Publish_ID DESC LIMIT 1",
		"OI_INS_NEW":                    "INSERT INTO " + dbConfig.DbDatabaseName + "." + dbConfig.CorpOIDbName + " (Stakeholder_ID,Title,Information,Attachment,CreationDate,LastUpdatedDate) VALUES(?,?,?,?,?,?) ",
		"OI_UPDATE_BY_TITLE":            "UPDATE " + dbConfig.DbDatabaseName + "." + dbConfig.CorpOIDbName + " SET ",
		"OI_UPDATE_WHERE":               " WHERE id= ? AND Stakeholder_ID= ?",
		"OI_GET_ALL":                    "SELECT id,Title,Information,Attachment,CreationDate,LastUpdatedDate,PublishFlag,ifnull(Publish_ID,'')  FROM " + dbConfig.DbDatabaseName + "." + dbConfig.CorpOIDbName + " WHERE Stakeholder_ID= ?  GROUP BY id,Title,Information,Attachment,CreationDate,LastUpdatedDate,PublishFlag,Publish_ID ORDER BY id DESC",
		"OI_GET_BY_ID":                  "SELECT Title,Information,CreationDate  FROM " + dbConfig.DbDatabaseName + "." + dbConfig.CorpOIDbName + " WHERE Stakeholder_ID= ? AND id =?  GROUP BY Title,Information,CreationDate",
		"OI_GET_ALL_PUB":                "SELECT id,Title,Information,Attachment,CreationDate,LastUpdatedDate,PublishFlag,Publish_ID  FROM " + dbConfig.DbDatabaseName + "." + dbConfig.CorpOIDbName + " WHERE Stakeholder_ID= ? AND PublishFlag=1  GROUP BY id,Title,Information,Attachment,CreationDate,LastUpdatedDate,PublishFlag,Publish_ID ORDER BY id DESC",
		"UNV_GET_PROFILE":               "SELECT Stakeholder_ID,University_Name,University_College_ID,UniversityHQAddress_Line1,UniversityHQAddress_Line2,UniversityHQAddress_Line3,UniversityHQAddress_Country,UniversityHQAddress_State,UniversityHQAddress_City,UniversityHQAddress_District,UniversityHQAddress_Zipcode,UniversityHQAddress_Phone,UniversityHQAddress_Email,UniversityLocal_BranchAddress_Line1,UniversityLocal_BranchAddress_Line2,UniversityLocal_BranchAddress_Line3,UniversityLocal_BranchAddress_Country,UniversityLocal_BranchAddress_State,UniversityLocal_BranchAddress_City,UniversityLocal_BranchAddress_District,UniversityLocal_BranchAddress_Zipcode,UniversityLocal_BranchAddress_Phone,UniversityLocal_BranchAddress_Email,PrimaryContact_FirstName,PrimaryContact_MiddleName,PrimaryContact_LastName,PrimaryContact_Designation,PrimaryContact_Phone,PrimaryContact_Email,UniversitySector,UniversityProfile,YearOfEstablishment,Attachment,DateOfJoiningPlatform FROM " + dbConfig.DbDatabaseName + "." + dbConfig.UnvMasterDbName + " WHERE Stakeholder_ID= ? GROUP BY Stakeholder_ID,University_Name,University_College_ID,UniversityHQAddress_Line1,UniversityHQAddress_Line2,UniversityHQAddress_Line3,UniversityHQAddress_Country,UniversityHQAddress_State,UniversityHQAddress_City,UniversityHQAddress_District,UniversityHQAddress_Zipcode,UniversityHQAddress_Phone,UniversityHQAddress_Email,UniversityLocal_BranchAddress_Line1,UniversityLocal_BranchAddress_Line2,UniversityLocal_BranchAddress_Line3,UniversityLocal_BranchAddress_Country,UniversityLocal_BranchAddress_State,UniversityLocal_BranchAddress_City,UniversityLocal_BranchAddress_District,UniversityLocal_BranchAddress_Zipcode,UniversityLocal_BranchAddress_Phone,UniversityLocal_BranchAddress_Email,PrimaryContact_FirstName,PrimaryContact_MiddleName,PrimaryContact_LastName,PrimaryContact_Designation,PrimaryContact_Phone,PrimaryContact_Email,UniversitySector,UniversityProfile,YearOfEstablishment,Attachment,DateOfJoiningPlatform LIMIT 1",
		"UNV_GET_Name":                  "SELECT University_Name FROM " + dbConfig.DbDatabaseName + "." + dbConfig.UnvMasterDbName + " WHERE Stakeholder_ID= ?  LIMIT 1",
		"UNV_PDH_INS_NEW":               "INSERT INTO " + dbConfig.DbDatabaseName + "." + dbConfig.UnvPDHDbName + " (Stakeholder_ID,Publish_ID,UniversityName,DateOfPublish,Programs_Published,BranchesPublished,StudentStrength_Published,Accreditations_Published,COEs_Published,Rankings_Published,Other_Published,Profile_Published,InfoPublished,GeneralNote,CreationDate,LastUpdatedDate,PublishData_JSON) VALUES ",
		"UNV_PDH_GET_BY_ID":             "SELECT Stakeholder_ID,Publish_ID,DateOfPublish,Programs_Published,BranchesPublished,StudentStrength_Published,Accreditations_Published,COEs_Published,Rankings_Published,Other_Published,Profile_Published,InfoPublished,GeneralNote,CreationDate,LastUpdatedDate,PublishData_JSON FROM " + dbConfig.DbDatabaseName + "." + dbConfig.UnvPDHDbName + " where Stakeholder_ID=?  GROUP BY Stakeholder_ID,Publish_ID,DateOfPublish,Programs_Published,BranchesPublished,StudentStrength_Published,Accreditations_Published,COEs_Published,Rankings_Published,Other_Published,Profile_Published,InfoPublished,GeneralNote,CreationDate,LastUpdatedDate,PublishData_JSON ORDER BY  CreationDate DESC",
		"UNV_PDH_Get_Last_ID":           "SELECT Publish_ID FROM " + dbConfig.DbDatabaseName + "." + dbConfig.UnvPDHDbName + " where Stakeholder_ID=? ORDER BY Publish_ID DESC LIMIT 1",
		"UNV_OI_INS_NEW":                "INSERT INTO " + dbConfig.DbDatabaseName + "." + dbConfig.UnvOIDbName + " (Stakeholder_ID,Title,Information,Attachment,PublishFlag,Publish_ID,CreationDate,LastUpdatedDate) VALUES(?,?,?,?,?,?,?,?) ",
		"UNV_OI_GET_ALL":                "SELECT id,Title,Information,Publish_ID,Attachment,CreationDate,LastUpdatedDate,PublishFlag  FROM " + dbConfig.DbDatabaseName + "." + dbConfig.UnvOIDbName + " WHERE Stakeholder_ID= ?  GROUP BY id,Title,Information,Publish_ID,Attachment,CreationDate,LastUpdatedDate,PublishFlag ORDER BY id DESC",
		"UNV_OI_GET_ALL_PUB":            "SELECT id,Title,Information,Publish_ID,Attachment,CreationDate,LastUpdatedDate,PublishFlag  FROM " + dbConfig.DbDatabaseName + "." + dbConfig.UnvOIDbName + " WHERE Stakeholder_ID= ? AND PublishFlag=1  GROUP BY id,Title,Information,Publish_ID,Attachment,CreationDate,LastUpdatedDate,PublishFlag ORDER BY id DESC",
		"UNV_Add_Program":               "INSERT INTO " + dbConfig.DbDatabaseName + "." + dbConfig.UnvProgDbName + " (Stakeholder_ID,Program_ID,ProgramName,ProgramType,Start_Date,End_Date,EnabledFlag,LastUpdatedDate) VALUES",
		"UNV_GET_Program":               "SELECT id,Program_ID,ProgramName,ProgramType,Start_Date,End_Date,EnabledFlag,CreationDate,LastUpdatedDate FROM " + dbConfig.DbDatabaseName + "." + dbConfig.UnvProgDbName + " WHERE Stakeholder_ID=?  GROUP BY id,Program_ID,ProgramName,ProgramType,Start_Date,End_Date,EnabledFlag,CreationDate,LastUpdatedDate ORDER BY CreationDate ",
		"UNV_Program_DELETE_BY_ID":      "DELETE FROM " + dbConfig.DbDatabaseName + "." + dbConfig.UnvProgDbName + " WHERE Stakeholder_ID= ? ",
		"Program_UNV_PRP_DEL":           "DELETE FROM " + dbConfig.DbDatabaseName + "." + dbConfig.UnvProgDbName + " WHERE id=? AND Stakeholder_ID= ? ",
		"UNV_Add_Coes":                  "INSERT INTO " + dbConfig.DbDatabaseName + "." + dbConfig.UnvCoesDbName + " (Stakeholder_ID,COEName,COEType,COEDescription,InternallyManagedFlag,OutsourcedVendor_Name,OutsourcedVendor_Stakeholder_ID,AttachFile,Start_Date,End_Date,EnabledFlag) VALUES",
		"UNV_GET_Coes":                  "SELECT id,COEName,COEType,COEDescription,InternallyManagedFlag,OutsourcedVendor_Name,OutsourcedVendor_Stakeholder_ID,AttachFile,Start_Date,End_Date,EnabledFlag,CreationDate,LastUpdatedDate FROM " + dbConfig.DbDatabaseName + "." + dbConfig.UnvCoesDbName + " WHERE Stakeholder_ID=? GROUP BY id,COEName,COEType,COEDescription,InternallyManagedFlag,OutsourcedVendor_Name,OutsourcedVendor_Stakeholder_ID,AttachFile,Start_Date,End_Date,EnabledFlag,CreationDate,LastUpdatedDate ORDER BY CreationDate",
		"UNV_Add_Accredations":          "INSERT INTO " + dbConfig.DbDatabaseName + "." + dbConfig.UnvAccrDbName + " (Stakeholder_ID,AccreditationName,AccreditationType,AccreditationDescription,IssuingAuthority,AttachFile,Start_Date,End_Date,EnabledFlag,LastUpdatedDate) VALUES",
		"UNV_GET_Accredations":          "SELECT id,AccreditationName,AccreditationType,AccreditationDescription,IssuingAuthority,AttachFile,Start_Date,End_Date,EnabledFlag,CreationDate,LastUpdatedDate FROM " + dbConfig.DbDatabaseName + "." + dbConfig.UnvAccrDbName + " WHERE Stakeholder_ID= ?  GROUP BY Stakeholder_ID,AccreditationName,AccreditationType,AccreditationDescription,IssuingAuthority,AttachFile,Start_Date,End_Date,EnabledFlag,CreationDate,LastUpdatedDate ORDER BY CreationDate",
		"UNV_Accredations_DELETE_BY_ID": "DELETE FROM " + dbConfig.DbDatabaseName + "." + dbConfig.UnvAccrDbName + " WHERE Stakeholder_ID= ? ",
		"UNV_Add_Branches":              "INSERT INTO " + dbConfig.DbDatabaseName + "." + dbConfig.UnvBranchDbName + " (Stakeholder_ID,Program_ID,ProgramName,Branch_ID,BranchName,Start_Date,End_Date,EnabledFlag,NoOfPassingStudents,MonthYearOfPassing,LastUpdatedDate) VALUES",
		"UNV_GET_Branches":              "SELECT id,Program_ID,ProgramName,Branch_ID,BranchName,Start_Date,End_Date,EnabledFlag,NoOfPassingStudents,MonthYearOfPassing,CreationDate,LastUpdatedDate FROM " + dbConfig.DbDatabaseName + "." + dbConfig.UnvBranchDbName + " WHERE Stakeholder_ID=?  GROUP BY id,Program_ID,ProgramName,Branch_ID,BranchName,Start_Date,End_Date,EnabledFlag,NoOfPassingStudents,MonthYearOfPassing,CreationDate,LastUpdatedDate ORDER BY CreationDate ",
		"UNV_Branches_DELETE_BY_ID":     "DELETE FROM " + dbConfig.DbDatabaseName + "." + dbConfig.UnvBranchDbName + " WHERE Stakeholder_ID= ? ",
		"Branch_UNV_PRP_DEL":            "DELETE FROM " + dbConfig.DbDatabaseName + "." + dbConfig.UnvBranchDbName + " WHERE id=? AND Stakeholder_ID= ? ",
		"UNV_Add_Ranking":               "INSERT INTO " + dbConfig.DbDatabaseName + "." + dbConfig.UnvRankDbName + " (Stakeholder_ID,Ranking,IssuingAuthority,AttachFile) VALUES",
		"UNV_GET_Ranking":               "SELECT id,Ranking,IssuingAuthority,AttachFile,CreationDate,LastUpdatedDate FROM " + dbConfig.DbDatabaseName + "." + dbConfig.UnvRankDbName + " WHERE Stakeholder_ID=? GROUP BY id,Ranking,IssuingAuthority,AttachFile,CreationDate,LastUpdatedDate ORDER BY CreationDate",
		"UNV_Add_SpecialOfferings":      "INSERT INTO " + dbConfig.DbDatabaseName + "." + dbConfig.UnvSplOfrDbName + " (Stakeholder_ID,SpecialOfferingName,SpecialOfferingType,SpecialOfferingDescription,InternallyManagedFlag,OutsourcedVendor_Name,OutsourcedVendor_ContachInfo,OutsourcedVendor_Stakeholder_ID,Univ_SplOfferscol,AttachFile,Start_Date,End_Date,EnabledFlag) VALUES",
		"UNV_GET_SpecialOfferings":      "SELECT id,SpecialOfferingName,SpecialOfferingType,SpecialOfferingDescription,InternallyManagedFlag,OutsourcedVendor_Name,OutsourcedVendor_ContachInfo,OutsourcedVendor_Stakeholder_ID,Univ_SplOfferscol,AttachFile,Start_Date,End_Date,EnabledFlag,CreationDate,LastUpdatedDate  FROM " + dbConfig.DbDatabaseName + "." + dbConfig.UnvSplOfrDbName + " WHERE Stakeholder_ID=? GROUP BY id,SpecialOfferingName,SpecialOfferingType,SpecialOfferingDescription,InternallyManagedFlag,OutsourcedVendor_Name,OutsourcedVendor_ContachInfo,OutsourcedVendor_Stakeholder_ID,Univ_SplOfferscol,AttachFile,Start_Date,End_Date,EnabledFlag,CreationDate,LastUpdatedDate ORDER BY CreationDate",
		"UNV_Add_Tieups":                "INSERT INTO " + dbConfig.DbDatabaseName + "." + dbConfig.UnvTieupDbName + " (Stakeholder_ID,TieupName,TieupType,TieupDescription,TiedupWith_Name,TiedupWith_ContactInfo,TiedupWith_Stakeholder_ID,AttachFile,Start_Date,End_Date,EnabledFlag) VALUES",
		"UNV_GET_Tieups":                "SELECT id,TieupName,TieupType,TieupDescription,TiedupWith_Name,TiedupWith_ContactInfo,TiedupWith_Stakeholder_ID,AttachFile,Start_Date,End_Date,EnabledFlag,CreationDate,LastUpdatedDate FROM " + dbConfig.DbDatabaseName + "." + dbConfig.UnvTieupDbName + " WHERE Stakeholder_ID=? GROUP BY id,TieupName,TieupType,TieupDescription,TiedupWith_Name,TiedupWith_ContactInfo,TiedupWith_Stakeholder_ID,AttachFile,Start_Date,End_Date,EnabledFlag,CreationDate,LastUpdatedDate ORDER BY CreationDate",
		"UNV_VRF_SUB":                   "select IF(COUNT(*),'true','false') from " + dbConfig.DbDatabaseName + "." + dbConfig.UnvSubDBName + " WHERE Subscriber_Stakeholder_ID = ? AND Publish_ID =?",
		"CRP_VRF_SUB":                   "select IF(COUNT(*),'true','false') from " + dbConfig.DbDatabaseName + "." + dbConfig.CrpSubDBName + " WHERE Subscriber_Stakeholder_ID = ? AND Publish_ID =?",
		"STU_VRF_SUB":                   "select IF(COUNT(*),'true','false') from " + dbConfig.DbDatabaseName + "." + dbConfig.StuSubDBName + " WHERE Subscriber_Stakeholder_ID = ? AND Publish_ID =?",
		"CRP_PDH_GET_PID":               "SELECT HiringCriteria_Published,Jobs_Published,Profile_Published,Other_Published,PublishData_JSON FROM " + dbConfig.DbDatabaseName + "." + dbConfig.CorpPDHDbName + " WHERE Publish_ID=? limit 1 ",
		"UNV_PDH_GET_PID":               "SELECT Programs_Published,BranchesPublished,StudentStrength_Published,Accreditations_Published,COEs_Published,Rankings_Published,Other_Published,Profile_Published,PublishData_JSON FROM " + dbConfig.DbDatabaseName + "." + dbConfig.UnvPDHDbName + " WHERE Publish_ID=? limit 1 ",
		"CRP_GET_JOB_BY_PID":            "SELECT b.Job_ID, b.JobName,c.Corporate_Name,(select ProgramName from CollabToHire.LUT_ProgramCatalog where Program_ID=hc.Program_ID limit 1) as ProgramName,(select BranchName from CollabToHire.LUT_BranchCatalog where Branch_ID=hc.Branch_ID limit 1) as BranchName,hc.MinimumCutoffCategory, hc.MinimumCutoff, hc.ActiveBacklogsAllowed, hc.TotalNumberOfBacklogsAllowed, hc.EduGaps_11N12_Allowed, hc.EduGaps_Grad_Allowed, hc.EduGaps_School_Allowed, hc.EduGaps_PG_Allowed, hc.YearOfPassing, hc.Remarks,(SELECT CONCAT('[{', result, '}]') AS final FROM (SELECT GROUP_CONCAT(DISTINCT CONCAT_WS(',',CONCAT(" + getCmd() + ") SEPARATOR '},{') AS result FROM CollabToHire.Corp_JobsToSkills_Mapping AS jc where jc.Job_ID = b.Job_ID) as jcd) as skills FROM CollabToHire.Corp_CreateJob AS b,CollabToHire.Corp_HiringCriteria as hc,CollabToHire.Corporate_Master as c WHERE b.Publish_ID = ? AND hc.HiringCriteria_ID = b.HiringCriteria_ID AND c.Stakeholder_ID=b.Stakeholder_ID Limit 1",
		"CRP_GET_HC_BY_PID":             "SELECT hc.HiringCriteria_ID,hc.HiringCriteria_Name,c.Corporate_Name,(select ProgramName from CollabToHire.LUT_ProgramCatalog where Program_ID=hc.Program_ID limit 1) as ProgramName,(select BranchName from CollabToHire.LUT_BranchCatalog where Branch_ID=hc.Branch_ID limit 1) as BranchName,hc.MinimumCutoffCategory, hc.MinimumCutoff, hc.ActiveBacklogsAllowed, hc.TotalNumberOfBacklogsAllowed, hc.EduGaps_11N12_Allowed, hc.EduGaps_Grad_Allowed, hc.EduGaps_School_Allowed, hc.EduGaps_PG_Allowed, hc.YearOfPassing, hc.Remarks FROM CollabToHire.Corp_HiringCriteria as hc,CollabToHire.Corporate_Master as c WHERE hc.PublishID = ? AND c.Stakeholder_ID=hc.Stakeholder_ID Limit 1",
		"CRP_GET_OI_BY_PID":             "Select Title,Information,Attachment FROM " + dbConfig.DbDatabaseName + "." + dbConfig.CorpOIDbName + " where Publish_ID=?",
		"UNV_GET_OI_BY_PID":             "Select Title,Information,Attachment FROM " + dbConfig.DbDatabaseName + "." + dbConfig.UnvOIDbName + " where Publish_ID=?",
		"STU_PDH_INS_NEW":               "INSERT INTO " + dbConfig.DbDatabaseName + "." + dbConfig.StuPublishDBName + " (Stakeholder_ID,Publish_ID,StudentName,DateOfPublish,ContactInfoPublished,EducationPublished,LanguagesPublished,CertificationsPublished,AssessmentsPublished,InternshipPublished,OtherInformationPublished,GeneralNote,CreationDate,LastUpdatedDate,PublishData_JSON) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)",
		"STU_GET_Name":                  "SELECT group_concat(Student_FirstName,' ',Student_LastName,' ',Student_MiddleName) FROM " + dbConfig.DbDatabaseName + "." + dbConfig.UnvMasterDbName + " WHERE Stakeholder_ID= ?  LIMIT 1",
		"STU_PDH_GET_BY_ID":             "SELECT Stakeholder_ID,Publish_ID,StudentName,DateOfPublish,ContactInfoPublished,EducationPublished,LanguagesPublished,CertificationsPublished,AssessmentsPublished,InternshipPublished,OtherInformationPublished,GeneralNote,CreationDate,LastUpdatedDate,PublishData_JSON FROM " + dbConfig.DbDatabaseName + "." + dbConfig.StuPublishDBName + " where Stakeholder_ID=?  GROUP BY Stakeholder_ID,Publish_ID,StudentName,DateOfPublish,ContactInfoPublished,EducationPublished,LanguagesPublished,CertificationsPublished,AssessmentsPublished,InternshipPublished,OtherInformationPublished,GeneralNote,CreationDate,LastUpdatedDate,PublishData_JSON ORDER BY  CreationDate DESC",
		"STU_PDH_GET_ALL":               "SELECT Publish_ID,StudentName,DateOfPublish,ContactInfoPublished,EducationPublished,LanguagesPublished,CertificationsPublished,AssessmentsPublished,InternshipPublished,OtherInformationPublished,GeneralNote,CreationDate,LastUpdatedDate,PublishData_JSON FROM " + dbConfig.DbDatabaseName + "." + dbConfig.StuPublishDBName + " where Stakeholder_ID=?  GROUP BY Publish_ID,StudentName,DateOfPublish,ContactInfoPublished,EducationPublished,LanguagesPublished,CertificationsPublished,AssessmentsPublished,InternshipPublished,OtherInformationPublished,GeneralNote,CreationDate,LastUpdatedDate,PublishData_JSON ORDER BY  CreationDate DESC",
		"STU_PDH_GET_PID":               "SELECT ContactInfoPublished,EducationPublished,LanguagesPublished,CertificationsPublished,AssessmentsPublished,InternshipPublished,OtherInformationPublished,PublishData_JSON FROM " + dbConfig.DbDatabaseName + "." + dbConfig.StuPublishDBName + " where Publish_ID=?  GROUP BY ContactInfoPublished,EducationPublished,LanguagesPublished,CertificationsPublished,AssessmentsPublished,InternshipPublished,OtherInformationPublished,PublishData_JSON LIMIT 1",
		"STU_PDH_Get_Last_ID":           "SELECT Publish_ID FROM " + dbConfig.DbDatabaseName + "." + dbConfig.StuPublishDBName + " where Stakeholder_ID=? ORDER BY Publish_ID DESC LIMIT 1",
		"STU_OI_INS_NEW":                "INSERT INTO " + dbConfig.DbDatabaseName + "." + dbConfig.StuOiDBName + " (Stakeholder_ID,Title,Information,Attachment,PublishedFlag,Publish_ID,CreationDate,LastUpdatedDate) VALUES(?,?,?,?,?,?,?,?) ",
		"STU_OI_GET_ALL":                "SELECT id,Title,Information,Publish_ID,Attachment,CreationDate,LastUpdatedDate,PublishedFlag  FROM " + dbConfig.DbDatabaseName + "." + dbConfig.StuOiDBName + " WHERE Stakeholder_ID= ?  GROUP BY id,Title,Information,Publish_ID,Attachment,CreationDate,LastUpdatedDate,PublishedFlag ORDER BY id DESC",
		"STU_GET_OI_BY_PID":             "SELECT Title,Information,Attachment FROM " + dbConfig.DbDatabaseName + "." + dbConfig.StuOiDBName + " where Publish_ID=?",
	}
}

// "JC_INS_NEW":      "INSERT INTO " + dbConfig.DbDatabaseName + "." + dbConfig.CorpJCDbName + " (Job_ID,Job_Name,CorporateID,NoOfPositions,Location,SalaryType,Salary,DateOfHiring,HC_ID,Status,Remarks,Attachment)  VALUES(?,?,?,?,?,?,?,?,?,?,?,?)",
// "JS_INS_NEW":      "INSERT INTO " + dbConfig.DbDatabaseName + "." + dbConfig.CorpJSDbName + " (Job_ID,Skill)  VALUES ",
// "JC_GET_BY_ID":    "SELECT Job_ID,Job_Name,CorporateID,NoOfPositions,Location,SalaryType,Salary,DateOfHiring,HC_ID,Status,Remarks,Attachment,CreationDate FROM " + dbConfig.DbDatabaseName + "." + dbConfig.CorpJCDbName + " WHERE Job_ID= ?  GROUP BY Job_ID,Job_Name,CorporateID,NoOfPositions,Location,SalaryType,Salary,DateOfHiring,HC_ID,Status,Remarks,Attachment,CreationDate LIMIT 1",
// "JC_GET_ALL":      "SELECT Job_ID,Job_Name,CorporateID,NoOfPositions,Location,SalaryType,Salary,DateOfHiring,HC_ID,Status,Remarks,Attachment,CreationDate FROM " + dbConfig.DbDatabaseName + "." + dbConfig.CorpJCDbName + " WHERE CorporateID= ?  GROUP BY Job_ID,Job_Name,CorporateID,NoOfPositions,Location,SalaryType,Salary,DateOfHiring,HC_ID,Status,Remarks,Attachment,CreationDate ORDER BY CreationDate",
// "GET_JOB_DETAILS": "SELECT a.*,(SELECT GROUP_CONCAT(b.Skill) FROM " + dbConfig.DbDatabaseName + "." + dbConfig.CorpJSDbName + " AS b WHERE a.Job_ID=b.Job_ID) FROM " + dbConfig.DbDatabaseName + "." + dbConfig.CorpJCDbName + " AS a WHERE JOb_ID=?",
// "JC_Get_Last_ID":  "SELECT Job_ID FROM " + dbConfig.DbDatabaseName + "." + dbConfig.CorpJCDbName + " where CorporateID=? ORDER BY CreationDate DESC LIMIT 1",
// "JC_UPDATE_BY_ID": "UPDATE " + dbConfig.DbDatabaseName + "." + dbConfig.CorpJCDbName + " SET ",
// "JC_UPDATE_WHERE": " WHERE Job_ID= ? AND CorporateID= ?",
// "JS_DELETE_BY_ID": "DELETE FROM " + dbConfig.DbDatabaseName + "." + dbConfig.CorpJSDbName + " WHERE Job_ID= ?",
// "JC_DELETE_BY_ID": "DELETE FROM " + dbConfig.DbDatabaseName + "." + dbConfig.CorpJCDbName + " WHERE Job_ID= ? AND CorporateID= ?",
//