// Package models ...
package models

import (
	"fmt"

	"github.com/jaswanth-gorripati/PGK/s4_Profile/configuration"
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
	return fmt.Sprint("'") + fmt.Sprint('"') + "jobID" + fmt.Sprint('"') + ": " + fmt.Sprint('"') + "',d.Job_ID,'" + fmt.Sprint('"') + "'),CONCAT('" + fmt.Sprint('"') + "isSubscribed" + fmt.Sprint('"') + ": " + fmt.Sprint('"') + "',pj.isSubscribed,'" + fmt.Sprint('"') + "'),CONCAT('" + fmt.Sprint('"') + "publishID" + fmt.Sprint('"') + ": " + fmt.Sprint('"') + "',pj.Publish_ID,'" + fmt.Sprint('"') + "'),CONCAT('" + fmt.Sprint('"') + "jobName" + fmt.Sprint('"') + ": " + fmt.Sprint('"') + "',d.JobName,'" + fmt.Sprint('"') + "'),CONCAT('" + fmt.Sprint('"') + "noOfPositions" + fmt.Sprint('"') + ": ',d.NoOfPositions),CONCAT('" + fmt.Sprint('"') + "dateOfHiring" + fmt.Sprint('"') + ": " + fmt.Sprint('"') + "',d.DateOfHiring,'" + fmt.Sprint('"') + "'),CONCAT('" + fmt.Sprint('"') + "location" + fmt.Sprint('"') + ": " + fmt.Sprint('"') + "',d.Location,'" + fmt.Sprint('"') + "')"
}

// CreateSP Creates default stored procedures for Database
func CreateSP() {
	dbConfig := configuration.DbConfig()
	SP = map[string]string{
		"CORP_GET_PROFILE":            "SELECT Stakeholder_ID,Corporate_Name,Corporate_CIN,CorporateHQAddress_Line1,CorporateHQAddress_Line2,CorporateHQAddress_Line3,CorporateHQAddress_Country,CorporateHQAddress_State,CorporateHQAddress_City,CorporateHQAddress_District,CorporateHQAddress_ZipCode,CorporateHQAddress_Phone,CorporateHQAddress_Email,CorporateLocal_BranchAddress_Line1,CorporateLocal_BranchAddress_Line2,CorporateLocal_BranchAddress_Line3,CorporateLocal_BranchAddress_Country,CorporateLocal_BranchAddress_State,CorporateLocal_BranchAddress_City,CorporateLocal_BranchAddress_District,CorporateLocal_BranchAddress_ZipCode,CorporateLocal_BranchAddress_Phone,CorporateLocal_BranchAddress_Email,PrimaryContact_FirstName,PrimaryContact_MiddleName,PrimaryContact_LastName,PrimaryContact_Designation,PrimaryContact_Phone,PrimaryContact_Email,SecondaryContact_FirstName,SecondaryContact_MiddleName,SecondaryContact_LastName,SecondaryContact_Designation,SecondaryContact_Phone,SecondaryContact_Email,CorporateType,CorporateCategory,CorporateIndustry,CompanyProfile,Attachment,YearOfEstablishment,DateOfJoiningPlatform,AccountStatus,Primary_Email_Verified,Primary_Phone_Verified,ProfilePicture,AccountExpiryDate FROM " + dbConfig.DbDatabaseName + "." + dbConfig.CorpMasterDbName + " WHERE Stakeholder_ID= ? GROUP BY Stakeholder_ID,Corporate_Name,Corporate_CIN,CorporateHQAddress_Line1,CorporateHQAddress_Line2,CorporateHQAddress_Line3,CorporateHQAddress_Country,CorporateHQAddress_State,CorporateHQAddress_City,CorporateHQAddress_District,CorporateHQAddress_ZipCode,CorporateHQAddress_Phone,CorporateHQAddress_Email,CorporateLocal_BranchAddress_Line1,CorporateLocal_BranchAddress_Line2,CorporateLocal_BranchAddress_Line3,CorporateLocal_BranchAddress_Country,CorporateLocal_BranchAddress_State,CorporateLocal_BranchAddress_City,CorporateLocal_BranchAddress_District,CorporateLocal_BranchAddress_ZipCode,CorporateLocal_BranchAddress_Phone,CorporateLocal_BranchAddress_Email,PrimaryContact_FirstName,PrimaryContact_MiddleName,PrimaryContact_LastName,PrimaryContact_Designation,PrimaryContact_Phone,PrimaryContact_Email,SecondaryContact_FirstName,SecondaryContact_MiddleName,SecondaryContact_LastName,SecondaryContact_Designation,SecondaryContact_Phone,SecondaryContact_Email,CorporateType,CorporateCategory,CorporateIndustry,CompanyProfile,Attachment,YearOfEstablishment,DateOfJoiningPlatform,AccountStatus,Primary_Email_Verified,Primary_Phone_Verified,ProfilePicture,AccountExpiryDate LIMIT 1",
		"CORP_UPDATE_BY_ID":           "UPDATE " + dbConfig.DbDatabaseName + "." + dbConfig.CorpMasterDbName + " SET ",
		"UNV_GET_PROFILE":             "SELECT Stakeholder_ID,University_Name,University_College_ID,UniversityHQAddress_Line1,UniversityHQAddress_Line2,UniversityHQAddress_Line3,UniversityHQAddress_Country,UniversityHQAddress_State,UniversityHQAddress_City,UniversityHQAddress_District,UniversityHQAddress_Zipcode,UniversityHQAddress_Phone,UniversityHQAddress_Email,UniversityLocal_BranchAddress_Line1,UniversityLocal_BranchAddress_Line2,UniversityLocal_BranchAddress_Line3,UniversityLocal_BranchAddress_Country,UniversityLocal_BranchAddress_State,UniversityLocal_BranchAddress_City,UniversityLocal_BranchAddress_District,UniversityLocal_BranchAddress_Zipcode,UniversityLocal_BranchAddress_Phone,UniversityLocal_BranchAddress_Email,PrimaryContact_FirstName,PrimaryContact_MiddleName,PrimaryContact_LastName,PrimaryContact_Designation,PrimaryContact_Phone,PrimaryContact_Email,SecondaryContact_FirstName,SecondaryContact_MiddleName,SecondaryContact_LastName,SecondaryContact_Designation,SecondaryContact_Phone,SecondaryContact_Email,UniversitySector,UniversityProfile,YearOfEstablishment,Attachment,DateOfJoiningPlatform,Primary_Email_Verified,Primary_Phone_Verified,AccountStatus,ProfilePicture,AccountExpiryDate FROM " + dbConfig.DbDatabaseName + "." + dbConfig.UnvMasterDbName + " WHERE Stakeholder_ID= ? GROUP BY Stakeholder_ID,University_Name,University_College_ID,UniversityHQAddress_Line1,UniversityHQAddress_Line2,UniversityHQAddress_Line3,UniversityHQAddress_Country,UniversityHQAddress_State,UniversityHQAddress_City,UniversityHQAddress_District,UniversityHQAddress_Zipcode,UniversityHQAddress_Phone,UniversityHQAddress_Email,UniversityLocal_BranchAddress_Line1,UniversityLocal_BranchAddress_Line2,UniversityLocal_BranchAddress_Line3,UniversityLocal_BranchAddress_Country,UniversityLocal_BranchAddress_State,UniversityLocal_BranchAddress_City,UniversityLocal_BranchAddress_District,UniversityLocal_BranchAddress_Zipcode,UniversityLocal_BranchAddress_Phone,UniversityLocal_BranchAddress_Email,PrimaryContact_FirstName,PrimaryContact_MiddleName,PrimaryContact_LastName,PrimaryContact_Designation,PrimaryContact_Phone,PrimaryContact_Email,SecondaryContact_FirstName,SecondaryContact_MiddleName,SecondaryContact_LastName,SecondaryContact_Designation,SecondaryContact_Phone,SecondaryContact_Email,UniversitySector,UniversityProfile,YearOfEstablishment,Attachment,DateOfJoiningPlatform,Primary_Email_Verified,Primary_Phone_Verified,AccountStatus,ProfilePicture,AccountExpiryDate LIMIT 1",
		"UNV_UPDATE_BY_ID":            "UPDATE " + dbConfig.DbDatabaseName + "." + dbConfig.UnvMasterDbName + " SET ",
		"STU_GET_PROFILE":             "SELECT Stakeholder_ID,Student_FirstName,Student_MiddleName,Student_LastName,Student_PersonalEmailID,Student_CollegeEmailID,Student_PhoneNumber,Student_AlternateContactNumber,Student_CollegeID,Student_Gender,Student_DateOfBirth,Student_AadharNumber,StudentPermanantAddress_Line1,StudentPermanantAddress_Line2,StudentPermanantAddress_Line3,StudentPermanantAddress_Country,StudentPermanantAddress_State,StudentPermanantAddress_City,StudentPermanantAddress_District,StudentPermanantAddress_Zipcode,StudentPermanantAddress_Phone,StudentPresentAddress_Line1,StudentPresentAddress_Line2,StudentPresentAddress_Line3,StudentPresentAddress_Country,StudentPresentAddress_State,StudentPresentAddress_City,StudentPresentAddress_District,StudentPresentAddress_Zipcode,StudentPresentAddress_Phone,Father_Guardian_FullName,Father_Guardian_Occupation,Father_Guardian_Company,Father_Guardian_PhoneNumber,Father_Guardian_EmailID,Mother_Guardian_FullName,Mother_Guardian_Occupation,Mother_Guardian_Comany,Mother_Guardian_Designation,Mother_Guardian_PhoneNumber,Mother_Guardian_EmailID,CreationDate,Primary_Email_Verified,Primary_Phone_Verified,AccountStatus,ProfilePicture,AccountExpiryDate,Student_AboutMe,ifnull(University_Name,''),ifnull(University_Stakeholder_ID,'') FROM " + dbConfig.DbDatabaseName + "." + dbConfig.StuMasterDbName + " WHERE Stakeholder_ID= ? GROUP BY Stakeholder_ID,Student_FirstName,Student_MiddleName,Student_LastName,Student_PersonalEmailID,Student_CollegeEmailID,Student_PhoneNumber,Student_AlternateContactNumber,Student_CollegeID,Student_Gender,Student_DateOfBirth,Student_AadharNumber,StudentPermanantAddress_Line1,StudentPermanantAddress_Line2,StudentPermanantAddress_Line3,StudentPermanantAddress_Country,StudentPermanantAddress_State,StudentPermanantAddress_City,StudentPermanantAddress_District,StudentPermanantAddress_Zipcode,StudentPermanantAddress_Phone,StudentPresentAddress_Line1,StudentPresentAddress_Line2,StudentPresentAddress_Line3,StudentPresentAddress_Country,StudentPresentAddress_State,StudentPresentAddress_City,StudentPresentAddress_District,StudentPresentAddress_Zipcode,StudentPresentAddress_Phone,Father_Guardian_FullName,Father_Guardian_Occupation,Father_Guardian_Company,Father_Guardian_PhoneNumber,Father_Guardian_EmailID,Mother_Guardian_FullName,Mother_Guardian_Occupation,Mother_Guardian_Comany,Mother_Guardian_Designation,Mother_Guardian_PhoneNumber,Mother_Guardian_EmailID,CreationDate,Primary_Email_Verified,Primary_Phone_Verified,AccountStatus,ProfilePicture,AccountExpiryDate,Student_AboutMe,University_Name,University_Stakeholder_ID LIMIT 1",
		"STU_UPDATE_BY_ID":            "UPDATE " + dbConfig.DbDatabaseName + "." + dbConfig.StuMasterDbName + " SET ",
		"PROFILE_UPDATE_WHERE":        " WHERE Stakeholder_ID= ?",
		"CORP_UPLOAD_PROFILE_PIC":     "UPDATE " + dbConfig.DbDatabaseName + "." + dbConfig.CorpMasterDbName + " SET ProfilePicture= ? WHERE Stakeholder_ID=?",
		"UNV_UPLOAD_PROFILE_PIC":      "UPDATE " + dbConfig.DbDatabaseName + "." + dbConfig.UnvMasterDbName + " SET ProfilePicture= ? WHERE Stakeholder_ID=?",
		"STU_UPLOAD_PROFILE_PIC":      "UPDATE " + dbConfig.DbDatabaseName + "." + dbConfig.StuMasterDbName + " SET ProfilePicture= ? WHERE Stakeholder_ID=?",
		"CORP_GET_PRF_PIC":            "SELECT ProfilePicture FROM " + dbConfig.DbDatabaseName + "." + dbConfig.CorpMasterDbName + " WHERE Stakeholder_ID= ? GROUP BY ProfilePicture LIMIT 1",
		"UNV_GET_PRF_PIC":             "SELECT ProfilePicture FROM " + dbConfig.DbDatabaseName + "." + dbConfig.UnvMasterDbName + " WHERE Stakeholder_ID= ? GROUP BY ProfilePicture LIMIT 1",
		"STU_GET_PRF_PIC":             "SELECT ProfilePicture FROM " + dbConfig.DbDatabaseName + "." + dbConfig.StuMasterDbName + " WHERE Stakeholder_ID= ? GROUP BY ProfilePicture LIMIT 1",
		"GET_LUT_CRP_TPY":             "SELECT CorporateType,CorporateTypeCode,OneLtrCode FROM " + dbConfig.DbDatabaseName + "." + dbConfig.LutCrpTyp + " GROUP BY CorporateType,CorporateTypeCode,OneLtrCode",
		"GET_LUT_CRP_CAT":             "SELECT CorporateCategoryName,CorporateCategoryCode,OneLtrCode FROM " + dbConfig.DbDatabaseName + "." + dbConfig.LutCrpCat + " GROUP BY CorporateCategoryName,CorporateCategoryCode,OneLtrCode",
		"GET_LUT_CRP_IND":             "SELECT CorporateIndutryName,CorporateIndustryCode FROM " + dbConfig.DbDatabaseName + "." + dbConfig.LutCrpInd + " GROUP BY CorporateIndutryName,CorporateIndustryCode",
		"GET_LUT_UNV_CAT":             "SELECT UniversityTypeName,UniversityTypeCode,OneLtrCode FROM " + dbConfig.DbDatabaseName + "." + dbConfig.LutUnvCat + " GROUP BY UniversityTypeName,UniversityTypeCode,OneLtrCode",
		"GET_LUT_SKILLS":              "SELECT Skill_ID,SkillName FROM " + dbConfig.DbDatabaseName + "." + dbConfig.LutSkills + " GROUP BY Skill_ID,SkillName",
		"GET_LUT_PROGRAMS":            "SELECT Program_ID,ProgramName FROM " + dbConfig.DbDatabaseName + "." + dbConfig.LutPrograms + " GROUP BY Program_ID,ProgramName",
		"GET_LUT_DEPART":              "SELECT Branch_ID,Program_ID,BranchName FROM " + dbConfig.DbDatabaseName + "." + dbConfig.LutDepartments + " GROUP BY Branch_ID,Program_ID,BranchName",
		"SRH_CRP_BY_SKILL_MAPPING":    "SELECT distinct a.Stakeholder_ID,(SELECT group_concat(distinct d.Location) as Locations from CollabToHire.Corp_JobsToSkills_Mapping AS d Where d.Stakeholder_ID=a.Stakeholder_ID),(SELECT AVG(b.MinimumCutoff) as AvgCutoff FROM CollabToHire.Corp_HiringCriteria AS b where b.Stakeholder_ID=a.Stakeholder_ID AND b.MinimumCutoff>=?),c.Corporate_Name,c.CorporateIndustry FROM CollabToHire.Corp_JobsToSkills_Mapping as a,CollabToHire.Corporate_Master as c,CollabToHire.Corp_HiringCriteria as h where c.Stakeholder_ID=a.Stakeholder_ID AND h.Stakeholder_ID=a.Stakeholder_ID",
		"CORP_GET_PROFILE_BY_ID":      "SELECT a.Stakeholder_ID,a.Corporate_Name,a.Corporate_CIN,a.CorporateHQAddress_Line1,a.CorporateHQAddress_Line2,a.CorporateHQAddress_Line3,a.CorporateHQAddress_Country,a.CorporateHQAddress_State,a.CorporateHQAddress_City,a.CorporateHQAddress_District,a.CorporateHQAddress_ZipCode,a.CorporateLocal_BranchAddress_Line1,a.CorporateLocal_BranchAddress_Line2,a.CorporateLocal_BranchAddress_Line3,a.CorporateLocal_BranchAddress_Country,a.CorporateLocal_BranchAddress_State,a.CorporateLocal_BranchAddress_City,a.CorporateLocal_BranchAddress_District,a.CorporateLocal_BranchAddress_ZipCode,a.CorporateType,a.CorporateCategory,a.CorporateIndustry,a.CompanyProfile,a.YearOfEstablishment,a.DateOfJoiningPlatform,(SELECT CONCAT('[{', result, '}]') as final FROM (SELECT GROUP_CONCAT(distinct CONCAT_WS(',',CONCAT(" + getCmd() + ") separator '},{') as result  from CollabToHire.Corp_JobsToSkills_Mapping AS d INNER JOIN (SELECT q.Job_ID,q.Publish_ID,(select IF(COUNT(*),'true','false') from CollabToHire.//RPLCSUB WHERE Subscriber_Stakeholder_ID = ? AND Publish_ID = q.Publish_ID) as isSubscribed FROM CollabToHire.Corp_CreateJob as q where Stakeholder_ID=a.Stakeholder_ID AND PublishFlag=1 Order BY LastUpdatedDate DESC LIMIT ?) as pj ON  d.Job_ID=pj.Job_ID order by d.CreationDate desc limit 1)as  jhg) AS Jobs  From CollabToHire.Corporate_Master as a where a.Stakeholder_ID=?",
		"UNV_GET_PROFILE_BY_ID":       "SELECT u.Stakeholder_ID, u.University_Name, University_College_ID, u.UniversityHQAddress_City, u.YearOfEstablishment, u.UniversityProfile, ifnull((select GROUP_CONCAT(DISTINCT ProgramName) From CollabToHire.Univ_ProgramsOffered where Stakeholder_ID=u.Stakeholder_ID),'') as Programs, ifnull((select json_object('issuingAuthority',IssuingAuthority,'rank',Ranking ) from CollabToHire.Univ_Ranking where Stakeholder_ID=u.Stakeholder_ID limit 1),'') as Ranking, ifnull((select json_object('issuingAuthority',IssuingAuthority,'name',AccreditationName,'type',AccreditationType ) from CollabToHire.Univ_Accreditations where Stakeholder_ID=u.Stakeholder_ID limit 1),'') as Accredations,(select Publish_ID FROM CollabToHire.Univ_PublishHistory where StudentStrength_Published=true AND Stakeholder_ID=u.Stakeholder_ID Order by DateOfPublish DESC LIMIT 1) as StuStrenthPublishID,(select if(count(*),true,false) FROM CollabToHire.Univ_PublishHistory where Stakeholder_ID=u.Stakeholder_ID LIMIT 1) as UnvInsightExists FROM CollabToHire.University_Master as u where u.Stakeholder_ID=?",
		"SRH_UNV_BY_HC":               "SELECT  u.Stakeholder_ID, u.University_Name, u.UniversityHQAddress_City, (SELECT JSON_OBJECT('issuingAuthority',IssuingAuthority,'rank',Ranking) FROM CollabToHire.Univ_Ranking WHERE Stakeholder_ID = u.Stakeholder_ID LIMIT 1) AS Ranking, (SELECT JSON_OBJECT('issuingAuthority',IssuingAuthority,'name',AccreditationName,'type',AccreditationType) FROM CollabToHire.Univ_Accreditations WHERE Stakeholder_ID = u.Stakeholder_ID  LIMIT 1) AS Accredation  FROM CollabToHire.University_Master as u  INNER JOIN (SELECT pg.Stakeholder_ID from CollabToHire.Univ_ProgramWiseBranches as pg,CollabToHire.Corp_HiringCriteria as hc  where HiringCriteria_ID=? AND pg.Program_ID=hc.Program_ID AND pg.Branch_ID=hc.Branch_ID) as phc ON u.Stakeholder_ID=phc.Stakeholder_ID",
		"SRH_UNV_BY_NSLC":             "SELECT  u.Stakeholder_ID, u.University_Name, u.UniversityHQAddress_City, (SELECT JSON_OBJECT('issuingAuthority',IssuingAuthority,'rank',Ranking,'exists',if(count(*),true,false)) FROM CollabToHire.Univ_Ranking WHERE Stakeholder_ID = u.Stakeholder_ID LIMIT 1) AS Ranking, (SELECT JSON_OBJECT('issuingAuthority',IssuingAuthority,'name',AccreditationName,'type',AccreditationType,'exists',if(count(*),true,false)) FROM CollabToHire.Univ_Accreditations WHERE Stakeholder_ID = u.Stakeholder_ID  LIMIT 1) AS Accredation FROM CollabToHire.University_Master as u where ",
		"STU_INS_ACADEMICS":           "INSERT INTO " + dbConfig.DbDatabaseName + "." + dbConfig.StuAcademics + " (Stakeholder_ID,",
		"STU_GET_ACADEMICS":           "SELECT Tenth_NameOfSchool,Tenth_LocationOfSchool,Tenth_MonthYearOfPassing,Tenth_SchoolingBoard,Tenth_Percentage,Tenth_AttachFile,Twelfth_NameOfInstitute,Twelfth_LocationOfInstitute,Twelfth_MonthYearOfPassing,Twelfth_InstituteBoard,Twelfth_Percentage,Twelfth_AttachFile,ifnull(Grad_Stakeholder_ID_Univ,''),ifnull(Grad_CollegeRollNumber,''),ifnull(Grad_ExpectedYearOfPassing,''),ifnull(Grad_Program_ID,''),ifnull(Grad_ProgramName,''),ifnull(Grad_Branch_ID,''),ifnull(Grad_BranchName,''),ifnull(Grad_FinalCGPA,''),ifnull(Grad_FinalPercentage,''),ifnull(Grad_ActiveBacklogs_Number,0),ifnull(Grad_TotalNumberOfBacklogs,0),ifnull(PostGrad_Unviversity_Stakeholder_ID,''),ifnull(PostGrad_CollegeRollNumber,''),ifnull(PostGrad_ExpectedYearOfPassing,''),ifnull(PostGrad_Program_ID,''),ifnull(PostGrad_ProgramName,''),ifnull(PostGrad_Branch_ID,''),ifnull(PostGrad_BranchName,''),ifnull(PostGrad_FinalCGPA,''),ifnull(PostGrad_FinalPercentage,'') FROM " + dbConfig.DbDatabaseName + "." + dbConfig.StuAcademics + " WHERE Stakeholder_ID=? ",
		"STU_TENTH_INS":               "Tenth_NameOfSchool,Tenth_LocationOfSchool,Tenth_MonthYearOfPassing,Tenth_SchoolingBoard,Tenth_Percentage,Tenth_AttachFile,",
		"STU_TENTH_UPD":               "Tenth_NameOfSchool=?,Tenth_LocationOfSchool=?,Tenth_MonthYearOfPassing=?,Tenth_SchoolingBoard=?,Tenth_Percentage=?,Tenth_AttachFile=?,",
		"STU_Twelfth_INS":             "Twelfth_NameOfInstitute,Twelfth_LocationOfInstitute,Twelfth_MonthYearOfPassing,Twelfth_InstituteBoard,Twelfth_Percentage,Twelfth_AttachFile,",
		"STU_Twelfth_UPD":             "Twelfth_NameOfInstitute=?,Twelfth_LocationOfInstitute=?,Twelfth_MonthYearOfPassing=?,Twelfth_InstituteBoard=?,Twelfth_Percentage=?,Twelfth_AttachFile=?,",
		"STU_SEM_INS":                 "INSERT INTO " + dbConfig.DbDatabaseName + "." + dbConfig.StuSemDbName + " (Student_Stakeholder_ID,University_Stakeholder_ID,IsGrad,IsPG,Semester,Student_CollegeRollNo,ProgramName,Program_ID,BranchName,Branch_ID,CGPA,Percentage,AttachFile,EnabledFlag,CreationDate,LastUpdatedDate,SemesterCompletionDate) VALUES",
		"STU_SEM_GET_ALL":             "SELECT id,Student_Stakeholder_ID,University_Stakeholder_ID,IsGrad,IsPG,Semester,Student_CollegeRollNo,ProgramName,Program_ID,BranchName,Branch_ID,CGPA,Percentage,AttachFile,EnabledFlag,CreationDate,LastUpdatedDate,SemesterCompletionDate FROM " + dbConfig.DbDatabaseName + "." + dbConfig.StuSemDbName + " WHERE Student_Stakeholder_ID=? ",
		"STU_GRAD_INS":                "Grad_Stakeholder_ID_Univ,Grad_CollegeRollNumber,Grad_ExpectedYearOfPassing,Grad_Program_ID,Grad_ProgramName,Grad_Branch_ID,Grad_BranchName,Grad_FinalCGPA,Grad_FinalPercentage,Grad_ActiveBacklogs_Number,Grad_TotalNumberOfBacklogs,",
		"STU_GRAD_UPD":                "Grad_Stakeholder_ID_Univ=?,Grad_CollegeRollNumber=?,Grad_ExpectedYearOfPassing=?,Grad_Program_ID=?,Grad_ProgramName=?,Grad_Branch_ID=?,Grad_BranchName=?,Grad_FinalCGPA=?,Grad_FinalPercentage=?,Grad_ActiveBacklogs_Number=?,Grad_TotalNumberOfBacklogs=?,",
		"STU_PG_INS":                  "PostGrad_Unviversity_Stakeholder_ID,PostGrad_CollegeRollNumber,PostGrad_ExpectedYearOfPassing,PostGrad_Program_ID,PostGrad_ProgramName,PostGrad_Branch_ID,PostGrad_BranchName,PostGrad_FinalCGPA,PostGrad_FinalPercentage,",
		"STU_PG_UPD":                  "PostGrad_Unviversity_Stakeholder_ID=?,PostGrad_CollegeRollNumber=?,PostGrad_ExpectedYearOfPassing=?,PostGrad_Program_ID=?,PostGrad_ProgramName=?,PostGrad_Branch_ID=?,PostGrad_BranchName=?,PostGrad_FinalCGPA=?,PostGrad_FinalPercentage=?,",
		"STU_LANG_INS":                "INSERT INTO " + dbConfig.DbDatabaseName + "." + dbConfig.StuLangKnown + " (Stakeholder_ID,LanguageName,SpeakProficiency,IsMotherTongue,ReadProficiency,WriteProficiency,EnabledFlag) VALUES",
		"STU_LANG_UPD":                "Update " + dbConfig.DbDatabaseName + "." + dbConfig.StuLangKnown + " SET LanguageName=?,SpeakProficiency=?,IsMotherTongue=?,ReadProficiency=?,WriteProficiency=?,LastUpdatedDate=? WHERE id=? AND Stakeholder_ID=?",
		"STU_LANG_DLT":                "DELETE from " + dbConfig.DbDatabaseName + "." + dbConfig.StuLangKnown + " WHERE id=? AND Stakeholder_ID=?",
		"STU_LANG_GETALL":             "SELECT id,LanguageName,SpeakProficiency,IsMotherTongue,ReadProficiency,WriteProficiency,EnabledFlag,CreationDate,LastUpdatedDate from " + dbConfig.DbDatabaseName + "." + dbConfig.StuLangKnown + " WHERE Stakeholder_ID=?",
		"STU_CERTS_INS":               "INSERT INTO " + dbConfig.DbDatabaseName + "." + dbConfig.StuCertsDbName + " (Stakeholder_ID,CourseName,IssuingAuthority,Start_Date,End_Date,AttachFile,Result,Description,EnabledFlag,CreationDate,LastUpdatedDate) VALUES(?,?,?,?,?,?,?,?,?,?,?) ",
		"STU_CERTS_UPD":               "Update " + dbConfig.DbDatabaseName + "." + dbConfig.StuCertsDbName + " SET CourseName=?,IssuingAuthority=?,Start_Date=?,End_Date=?,AttachFile=?,Result=?,Description=?,LastUpdatedDate=? WHERE id=? AND Stakeholder_ID=?",
		"STU_CERTS_DLT":               "DELETE from " + dbConfig.DbDatabaseName + "." + dbConfig.StuCertsDbName + " WHERE id=? AND Stakeholder_ID=?",
		"STU_CERTS_GETALL":            "SELECT id,CourseName,IssuingAuthority,Start_Date,End_Date,AttachFile,Result,Description,EnabledFlag,CreationDate,LastUpdatedDate from " + dbConfig.DbDatabaseName + "." + dbConfig.StuCertsDbName + " WHERE Stakeholder_ID=?",
		"STU_ASSESSMENT_INS":          "INSERT INTO " + dbConfig.DbDatabaseName + "." + dbConfig.StuAssessmentDbName + " (Stakeholder_ID,AssessmentName,AssessmentScore,IssuingAuthority,Assessment_Date,Description,AttachFile,EnabledFlag,CreationDate,LastUpdatedDate) VALUES(?,?,?,?,?,?,?,?,?,?) ",
		"STU_ASSESSMENT_UPD":          "Update " + dbConfig.DbDatabaseName + "." + dbConfig.StuAssessmentDbName + " SET AssessmentName=?,AssessmentScore=?,IssuingAuthority=?,Assessment_Date=?,Description=?,AttachFile=?,LastUpdatedDate=? WHERE id=? AND Stakeholder_ID=?",
		"STU_ASSESSMENT_DLT":          "DELETE from " + dbConfig.DbDatabaseName + "." + dbConfig.StuAssessmentDbName + " WHERE id=? AND Stakeholder_ID=?",
		"STU_ASSESSMENT_GETALL":       "SELECT id,AssessmentName,AssessmentScore,IssuingAuthority,Assessment_Date,Description,AttachFile,EnabledFlag,CreationDate,LastUpdatedDate from " + dbConfig.DbDatabaseName + "." + dbConfig.StuAssessmentDbName + " WHERE Stakeholder_ID=?",
		"STU_INTERSHIP_INS":           "INSERT INTO " + dbConfig.DbDatabaseName + "." + dbConfig.StuIntershipsDbName + " (Stakeholder_ID,InternshipName,NameOfOrganisation,FieldOfWork,CityOfOrganisation,Start_Date,End_Date,Description,AttachFile,EnabledFlag,CreationDate,LastUpdatedDate) VALUES(?,?,?,?,?,?,?,?,?,?,?,?) ",
		"STU_INTERSHIP_UPD":           "Update " + dbConfig.DbDatabaseName + "." + dbConfig.StuIntershipsDbName + " SET InternshipName=?,NameOfOrganisation=?,FieldOfWork=?,CityOfOrganisation=?,Start_Date=?,End_Date=?,Description=?,AttachFile=?,LastUpdatedDate=? WHERE id=? AND Stakeholder_ID=?",
		"STU_INTERSHIP_DLT":           "DELETE from " + dbConfig.DbDatabaseName + "." + dbConfig.StuIntershipsDbName + " WHERE id=? AND Stakeholder_ID=?",
		"STU_INTERSHIP_GETALL":        "SELECT id,InternshipName,NameOfOrganisation,FieldOfWork,CityOfOrganisation,Start_Date,End_Date,Description,AttachFile,EnabledFlag,CreationDate,LastUpdatedDate from " + dbConfig.DbDatabaseName + "." + dbConfig.StuIntershipsDbName + " WHERE Stakeholder_ID=?",
		"UNV_SUB_DATA_IN_SRH":         "SELECT a.Publisher_Stakeholder_ID,a.DateOfSubscription,a.Publish_ID,a.Transaction_ID,d.GeneralNote FROM CollabToHire.Corp_SubscriptionHistory as a,CollabToHire.Univ_PublishHistory AS d where d.Publish_ID=a.Publish_ID AND  a.Subscriber_Stakeholder_ID=? AND a.Publisher_Stakeholder_ID=?",
		"STU_Awards_INS":              "INSERT INTO " + dbConfig.DbDatabaseName + "." + dbConfig.StuAwardsDbName + " (Stakeholder_ID,AwardRecognitionName,AwardRecognition_Date,IssuingAuthority,AttachFile,EnabledFlag,CreationDate,LastUpdatedDate) VALUES(?,?,?,?,?,?,?,?) ",
		"STU_Awards_UPD":              "Update " + dbConfig.DbDatabaseName + "." + dbConfig.StuAwardsDbName + " SET AwardRecognitionName=?,AwardRecognition_Date=?,IssuingAuthority=?,AttachFile=?,LastUpdatedDate=? WHERE id=? AND Stakeholder_ID=?",
		"STU_Awards_DLT":              "DELETE from " + dbConfig.DbDatabaseName + "." + dbConfig.StuAwardsDbName + " WHERE id=? AND Stakeholder_ID=?",
		"STU_Awards_GETALL":           "SELECT id,AwardRecognitionName,AwardRecognition_Date,IssuingAuthority,AttachFile,EnabledFlag,CreationDate,LastUpdatedDate from " + dbConfig.DbDatabaseName + "." + dbConfig.StuAwardsDbName + " WHERE Stakeholder_ID=?",
		"STU_Competitions_INS":        "INSERT INTO " + dbConfig.DbDatabaseName + "." + dbConfig.StuCompetitionDbName + " (Stakeholder_ID,CompetitionName,Competition_Date,CompetitionRank,AttachFile,EnabledFlag,CreationDate,LastUpdatedDate) VALUES(?,?,?,?,?,?,?,?) ",
		"STU_Competitions_UPD":        "Update " + dbConfig.DbDatabaseName + "." + dbConfig.StuCompetitionDbName + " SET CompetitionName=?,Competition_Date=?,CompetitionRank=?,AttachFile=?,LastUpdatedDate=? WHERE id=? AND Stakeholder_ID=?",
		"STU_Competitions_DLT":        "DELETE from " + dbConfig.DbDatabaseName + "." + dbConfig.StuCompetitionDbName + " WHERE id=? AND Stakeholder_ID=?",
		"STU_Competitions_GETALL":     "SELECT id,CompetitionName,Competition_Date,CompetitionRank,AttachFile,EnabledFlag,CreationDate,LastUpdatedDate from " + dbConfig.DbDatabaseName + "." + dbConfig.StuCompetitionDbName + " WHERE Stakeholder_ID=?",
		"STU_CONF_WORKSHOP_INS":       "INSERT INTO " + dbConfig.DbDatabaseName + "." + dbConfig.StuConfWorkshopsDbName + " (Stakeholder_ID,ConferenceWorkshopName,ConferenceWorkshop_Date,AttachFile,EnabledFlag,CreationDate,LastUpdatedDate) VALUES(?,?,?,?,?,?,?) ",
		"STU_CONF_WORKSHOP_UPD":       "Update " + dbConfig.DbDatabaseName + "." + dbConfig.StuConfWorkshopsDbName + " SET ConferenceWorkshopName=?,ConferenceWorkshop_Date=?,AttachFile=?,LastUpdatedDate=? WHERE id=? AND Stakeholder_ID=?",
		"STU_CONF_WORKSHOP_DLT":       "DELETE from " + dbConfig.DbDatabaseName + "." + dbConfig.StuConfWorkshopsDbName + " WHERE id=? AND Stakeholder_ID=?",
		"STU_CONF_WORKSHOP_GETALL":    "SELECT id,ConferenceWorkshopName,ConferenceWorkshop_Date,AttachFile,EnabledFlag,CreationDate,LastUpdatedDate from " + dbConfig.DbDatabaseName + "." + dbConfig.StuConfWorkshopsDbName + " WHERE Stakeholder_ID=?",
		"STU_EXTRA_CURRICULAR_INS":    "INSERT INTO " + dbConfig.DbDatabaseName + "." + dbConfig.StuExtraCurDbName + " (Stakeholder_ID,ExtraCurricularName,AttachFile,EnabledFlag,CreationDate,LastUpdatedDate) VALUES(?,?,?,?,?,?) ",
		"STU_EXTRA_CURRICULAR_UPD":    "Update " + dbConfig.DbDatabaseName + "." + dbConfig.StuExtraCurDbName + " SET ExtraCurricularName=?,AttachFile=?,LastUpdatedDate=? WHERE id=? AND Stakeholder_ID=?",
		"STU_EXTRA_CURRICULAR_DLT":    "DELETE from " + dbConfig.DbDatabaseName + "." + dbConfig.StuExtraCurDbName + " WHERE id=? AND Stakeholder_ID=?",
		"STU_EXTRA_CURRICULAR_GETALL": "SELECT id,ExtraCurricularName,AttachFile,EnabledFlag,CreationDate,LastUpdatedDate from " + dbConfig.DbDatabaseName + "." + dbConfig.StuExtraCurDbName + " WHERE Stakeholder_ID=?",
		"STU_PATENTS_INS":             "INSERT INTO " + dbConfig.DbDatabaseName + "." + dbConfig.StuPatentsDbName + " (Stakeholder_ID,PatentName,Patent_Type,Patent_Number,Patent_Status,AttachFile,EnabledFlag,CreationDate,LastUpdatedDate) VALUES(?,?,?,?,?,?,?,?,?) ",
		"STU_PATENTS_UPD":             "Update " + dbConfig.DbDatabaseName + "." + dbConfig.StuPatentsDbName + " SET PatentName=?,Patent_Type=?,Patent_Number=?,Patent_Status=?,AttachFile=?,LastUpdatedDate=? WHERE id=? AND Stakeholder_ID=?",
		"STU_PATENTS_DLT":             "DELETE from " + dbConfig.DbDatabaseName + "." + dbConfig.StuPatentsDbName + " WHERE id=? AND Stakeholder_ID=?",
		"STU_PATENTS_GETALL":          "SELECT id,PatentName,Patent_Type,Patent_Number,Patent_Status,AttachFile,EnabledFlag,CreationDate,LastUpdatedDate from " + dbConfig.DbDatabaseName + "." + dbConfig.StuPatentsDbName + " WHERE Stakeholder_ID=?",
		"STU_PROJECTS_INS":            "INSERT INTO " + dbConfig.DbDatabaseName + "." + dbConfig.StuProjectsDbName + " (Stakeholder_ID,ProjectName,ProjectAbstract,GuideName,Guide_EmailID,Start_Date,End_Date,AttachFile,EnabledFlag,CreationDate,LastUpdatedDate) VALUES(?,?,?,?,?,?,?,?,?,?,?) ",
		"STU_PROJECTS_UPD":            "Update " + dbConfig.DbDatabaseName + "." + dbConfig.StuProjectsDbName + " SET ProjectName=?,ProjectAbstract=?,GuideName=?,Guide_EmailID=?,Start_Date=?,End_Date=?,AttachFile=?,LastUpdatedDate=? WHERE id=? AND Stakeholder_ID=?",
		"STU_PROJECTS_DLT":            "DELETE from " + dbConfig.DbDatabaseName + "." + dbConfig.StuProjectsDbName + " WHERE id=? AND Stakeholder_ID=?",
		"STU_PROJECTS_GETALL":         "SELECT id,ProjectName,ProjectAbstract,ifnull(GuideName,''),ifnull(Guide_EmailID,''),Start_Date,End_Date,ifnull(AttachFile,''),EnabledFlag,CreationDate,LastUpdatedDate from " + dbConfig.DbDatabaseName + "." + dbConfig.StuProjectsDbName + " WHERE Stakeholder_ID=?",
		"STU_PUBLICATIONS_INS":        "INSERT INTO " + dbConfig.DbDatabaseName + "." + dbConfig.StuPublicationsDbName + " (Stakeholder_ID,PublicationName,PublishingAuthority,Guide_Name,Guide_EmailID,Start_Date,End_Date,AttachFile,EnabledFlag,CreationDate,LastUpdatedDate) VALUES(?,?,?,?,?,?,?,?,?,?,?) ",
		"STU_PUBLICATIONS_UPD":        "Update " + dbConfig.DbDatabaseName + "." + dbConfig.StuPublicationsDbName + " SET PublicationName=?,PublishingAuthority=?,Guide_Name=?,Guide_EmailID=?,Start_Date=?,End_Date=?,AttachFile=?,LastUpdatedDate=? WHERE id=? AND Stakeholder_ID=?",
		"STU_PUBLICATIONS_DLT":        "DELETE from " + dbConfig.DbDatabaseName + "." + dbConfig.StuPublicationsDbName + " WHERE id=? AND Stakeholder_ID=?",
		"STU_PUBLICATIONS_GETALL":     "SELECT id,PublicationName,PublishingAuthority,Guide_Name,Guide_EmailID,Start_Date,End_Date,AttachFile,EnabledFlag,CreationDate,LastUpdatedDate from " + dbConfig.DbDatabaseName + "." + dbConfig.StuPublicationsDbName + " WHERE Stakeholder_ID=?",
		"STU_SCHOLARSHIPS_INS":        "INSERT INTO " + dbConfig.DbDatabaseName + "." + dbConfig.StuScholarshipsDbName + " (Stakeholder_ID,Scholarship_Name,ScholarshipGivenBy,Scholarship_Date,AttachFile,EnabledFlag,CreationDate,LastUpdatedDate) VALUES(?,?,?,?,?,?,?,?) ",
		"STU_SCHOLARSHIPS_UPD":        "Update " + dbConfig.DbDatabaseName + "." + dbConfig.StuScholarshipsDbName + " SET Scholarship_Name=?,ScholarshipGivenBy=?,Scholarship_Date=?,AttachFile=?,LastUpdatedDate=? WHERE id=? AND Stakeholder_ID=?",
		"STU_SCHOLARSHIPS_DLT":        "DELETE from " + dbConfig.DbDatabaseName + "." + dbConfig.StuScholarshipsDbName + " WHERE id=? AND Stakeholder_ID=?",
		"STU_SCHOLARSHIPS_GETALL":     "SELECT id,Scholarship_Name,ScholarshipGivenBy,Scholarship_Date,AttachFile,EnabledFlag,CreationDate,LastUpdatedDate from " + dbConfig.DbDatabaseName + "." + dbConfig.StuScholarshipsDbName + " WHERE Stakeholder_ID=?",
		"STU_SOCIAL_ACCOUNTS_INS":     "INSERT INTO " + dbConfig.DbDatabaseName + "." + dbConfig.StuSocialAccountsDbName + " (Stakeholder_ID,SocialAccount_UserID,EnabledFlag,CreationDate,LastUpdatedDate) VALUES(?,?,?,?,?) ",
		"STU_SOCIAL_ACCOUNTS_UPD":     "Update " + dbConfig.DbDatabaseName + "." + dbConfig.StuSocialAccountsDbName + " SET SocialAccount_UserID=?,LastUpdatedDate=? WHERE id=? AND Stakeholder_ID=?",
		"STU_SOCIAL_ACCOUNTS_DLT":     "DELETE from " + dbConfig.DbDatabaseName + "." + dbConfig.StuSocialAccountsDbName + " WHERE id=? AND Stakeholder_ID=?",
		"STU_SOCIAL_ACCOUNTS_GETALL":  "SELECT id,SocialAccount_UserID,EnabledFlag,CreationDate,LastUpdatedDate from " + dbConfig.DbDatabaseName + "." + dbConfig.StuSocialAccountsDbName + " WHERE Stakeholder_ID=?",
		"STU_TECH_SKILLS_INS":         "INSERT INTO " + dbConfig.DbDatabaseName + "." + dbConfig.StuTechSkillsDbName + " (Stakeholder_ID,TechnicalSkill_ID,TechnicalSkillName,EnabledFlag,CreationDate,LastUpdateDate) VALUES(?,?,?,?,?,?) ",
		"STU_TECH_SKILLS_UPD":         "Update " + dbConfig.DbDatabaseName + "." + dbConfig.StuTechSkillsDbName + " SET TechnicalSkill_ID=?,TechnicalSkillName=?,LastUpdateDate=? WHERE id=? AND Stakeholder_ID=?",
		"STU_TECH_SKILLS_DLT":         "DELETE from " + dbConfig.DbDatabaseName + "." + dbConfig.StuTechSkillsDbName + " WHERE id=? AND Stakeholder_ID=?",
		"STU_TECH_SKILLS_GETALL":      "SELECT id,TechnicalSkill_ID,TechnicalSkillName,EnabledFlag,CreationDate,LastUpdateDate from " + dbConfig.DbDatabaseName + "." + dbConfig.StuTechSkillsDbName + " WHERE Stakeholder_ID=?",
		"STU_TEST_SCORES_INS":         "INSERT INTO " + dbConfig.DbDatabaseName + "." + dbConfig.StuTestScoresDbName + " (Stakeholder_ID,TestScoreName,TestScore_Date,Test_Score,TestScore_TotalScore,AttachFile,EnabledFlag,CreationDate,LastUpdatedDate) VALUES(?,?,?,?,?,?,?,?,?) ",
		"STU_TEST_SCORES_UPD":         "Update " + dbConfig.DbDatabaseName + "." + dbConfig.StuTestScoresDbName + " SET TestScoreName=?,TestScore_Date=?,Test_Score=?,TestScore_TotalScore=?,AttachFile=?,LastUpdatedDate=? WHERE id=? AND Stakeholder_ID=?",
		"STU_TEST_SCORES_DLT":         "DELETE from " + dbConfig.DbDatabaseName + "." + dbConfig.StuTestScoresDbName + " WHERE id=? AND Stakeholder_ID=?",
		"STU_TEST_SCORES_GETALL":      "SELECT id,TestScoreName,TestScore_Date,Test_Score,TestScore_TotalScore,AttachFile,EnabledFlag,CreationDate,LastUpdatedDate from " + dbConfig.DbDatabaseName + "." + dbConfig.StuTestScoresDbName + " WHERE Stakeholder_ID=?",
		"STU_VOLUNTEER_INS":           "INSERT INTO " + dbConfig.DbDatabaseName + "." + dbConfig.StuVolunteerExpDbName + " (Stakeholder_ID,VolunteerExperience_Name,VolunteerExperience_Organisation,VolunteerExperience_Location,Start_Date,End_Date,AttachFile,EnabledFlag,CreationDate,LastUpdatedDate) VALUES(?,?,?,?,?,?,?,?,?,?) ",
		"STU_VOLUNTEER_UPD":           "Update " + dbConfig.DbDatabaseName + "." + dbConfig.StuVolunteerExpDbName + " SET VolunteerExperience_Name=?,VolunteerExperience_Organisation=?,VolunteerExperience_Location=?,Start_Date=?,End_Date=?,AttachFile=?,LastUpdatedDate=? WHERE id=? AND Stakeholder_ID=?",
		"STU_VOLUNTEER_DLT":           "DELETE from " + dbConfig.DbDatabaseName + "." + dbConfig.StuVolunteerExpDbName + " WHERE id=? AND Stakeholder_ID=?",
		"STU_VOLUNTEER_GETALL":        "SELECT id,VolunteerExperience_Name,VolunteerExperience_Organisation,VolunteerExperience_Location,Start_Date,End_Date,AttachFile,EnabledFlag,CreationDate,LastUpdatedDate from " + dbConfig.DbDatabaseName + "." + dbConfig.StuVolunteerExpDbName + " WHERE Stakeholder_ID=?",
		"STU_REQ_PROFILE_VRF":         "UPDATE " + dbConfig.DbDatabaseName + "." + dbConfig.StuMasterDbName + " SET UniversityVerificationStatus='PENDING',UniversityApprovedFlag=false WHERE Stakeholder_ID= ?",
		"STU_GET_ALL_PROFILES":        "SELECT ifnull(sm.Stakeholder_ID,''),ifnull(GROUP_CONCAT(sm.Student_FirstName,' ',sm.Student_MiddleName,' ',sm.Student_LastName),''),ifnull(sm.Student_CollegeID,''),ifnull(sg.Grad_Program_ID,''),ifnull(sg.Grad_BranchName,''),ifnull(sg.Grad_ExpectedYearOfPassing,''),ifnull(sg.PostGrad_Program_ID,''),ifnull(sg.PostGrad_BranchName,''),ifnull(sg.PostGrad_ExpectedYearOfPassing,'')  from " + dbConfig.DbDatabaseName + "." + dbConfig.StuMasterDbName + " as sm, " + dbConfig.DbDatabaseName + "." + dbConfig.StuAcademics + " as sg WHERE sm.University_Stakeholder_ID= ? AND sg.Stakeholder_ID=sm.Stakeholder_ID AND sm.UniversityVerificationStatus=?",
		"STU_VALIDATE_PROFILE":        "UPDATE " + dbConfig.DbDatabaseName + "." + dbConfig.StuMasterDbName + " SET UniversityVerificationStatus=?,UniversityApprovedFlag=? WHERE Stakeholder_ID= ? AND University_Stakeholder_ID= ?",
		"UNV_INSIGHTS_GET_ALL":        "SELECT Subscription_ID,Subscriber_Stakeholder_ID,Subscribed_Stakeholder_ID,SubscribedDate FROM " + dbConfig.DbDatabaseName + "." + dbConfig.UnvInsightsDbName + " Where Subscribed_Stakeholder_ID=? AND Subscriber_Stakeholder_ID=? GROUP BY Subscription_ID,Subscriber_Stakeholder_ID,Subscribed_Stakeholder_ID,SubscribedDate Order by SubscribedDate DESC",
		"UNV_STU_DB_SUB_GET_ALL":      "SELECT DISTINCT Subscription_ID,Subscriber_Stakeholder_ID,Subscribed_Stakeholder_ID,SubscribedDate,SearchCriteria FROM " + dbConfig.DbDatabaseName + "." + dbConfig.UnvStuDataDbName + " WHERE  Subscribed_Stakeholder_ID=? AND Subscriber_Stakeholder_ID=? GROUP BY Subscription_ID,Subscriber_Stakeholder_ID,Subscribed_Stakeholder_ID,SubscribedDate Order by SubscribedDate DESC ",
		"CORP_CD_GET_ALL":             "SELECT Initiator_Stakeholder_ID,Receiver_Stakeholder_ID,CampusDrive_ID,CampusDrive_Requested,CampusDrive_Requested_Date,ifnull(CampusDrive_Requested_NotificationID,''),ifnull(CampusDrive_AcceptedorRejectedbyUniv,false),ifnull(CampusDrive_AcceptedorRejectedbyUniv_Date,CURRENT_TIMESTAMP),ifnull(CampusDrive_AcceptedorRejectedbyUniv_NotificationID,'') FROM " + dbConfig.DbDatabaseName + "." + dbConfig.CorpCDDbName + " WHERE Receiver_Stakeholder_ID=? AND Initiator_Stakeholder_ID=? ",
		"CORP_HCI_GET_ALL_SUB":        "SELECT a.Subscription_ID,a.Subscriber_Stakeholder_ID,a.Subscribed_Stakeholder_ID,a.SubscribedDate FROM " + dbConfig.DbDatabaseName + "." + dbConfig.CorpHcInsightsDbName + " as a Where Subscriber_Stakeholder_ID=? AND Subscribed_Stakeholder_ID=?  Order by SubscribedDate DESC",
		"UNV_CD_GET_ALL":              "SELECT Initiator_Stakeholder_ID,Receiver_Stakeholder_ID,CampusDrive_ID,CampusDrive_Requested,CampusDrive_Requested_Date,ifnull(CampusDrive_Requested_NotificationID,''),ifnull(CampusDrive_AcceptedorRejectedbyCorp,false),ifnull(CampusDrive_AcceptedorRejectedbyCorp_Date,CURRENT_TIMESTAMP),ifnull(CampusDrive_AcceptedorRejectedbyCorp_NotificationID,'') FROM " + dbConfig.DbDatabaseName + "." + dbConfig.UnvCPDbName + " as a WHERE Receiver_Stakeholder_ID=? AND Initiator_Stakeholder_ID=? ",
	}
}
