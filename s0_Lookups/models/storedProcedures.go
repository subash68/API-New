// Package models ...
package models

import (
	"github.com/jaswanth-gorripati/PGK/s0_Lookups/configuration"
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
		"LUT_10Boards_GET":             "SELECT BoardName,CertificateName,Board_ID FROM " + dbConfig.DbDatabaseName + "." + dbConfig.LUT10Boards + "",
		"LUT_12Boards_GET":             "SELECT BoardName,CertificateName,Board_ID FROM " + dbConfig.DbDatabaseName + "." + dbConfig.LUT12Boards + "",
		"LUT_AccountStatus":            "SELECT AccountStatus,AccountStatusCode FROM " + dbConfig.DbDatabaseName + "." + dbConfig.LUTAccountStatus + "",
		"LUT_Branches":                 "SELECT Branch_ID,BranchName,Duration,Program_ID,ProgramType FROM " + dbConfig.DbDatabaseName + "." + dbConfig.LUTBranchCatalog + "",
		"LUT_CorporateCategory":        "SELECT CorporateCategoryName,CorporateCategoryCode,OneLtrCode FROM " + dbConfig.DbDatabaseName + "." + dbConfig.LUTCorporateCategory + "",
		"LUT_CorporateIndustry":        "SELECT CorporateIndustryName,CorporateIndustryCode FROM " + dbConfig.DbDatabaseName + "." + dbConfig.LUTCorporateIndustry + "",
		"LUT_CorporateType":            "SELECT CorporateTypeCode,CorporateType,OneLtrCode FROM " + dbConfig.DbDatabaseName + "." + dbConfig.LUTCorporateType + "",
		"LUT_JobType":                  "SELECT Job_Type,Job_Type_Code FROM " + dbConfig.DbDatabaseName + "." + dbConfig.LUTJob + "",
		"LUT_Lang_Prof":                "SELECT LanguageProficiency,LanguageProficiency_ID FROM " + dbConfig.DbDatabaseName + "." + dbConfig.LUTLanguageProficiency + "",
		"LUT_Token_MOI":                "SELECT ModeOfIssueOfToken,ModeOfIssueOfToken_ID FROM " + dbConfig.DbDatabaseName + "." + dbConfig.LUTModeOfIssueOfToken + "",
		"LUT_Nft_Type":                 "SELECT NotificationType,NotificationType_ID FROM " + dbConfig.DbDatabaseName + "." + dbConfig.LUTNotificationType + "",
		"LUT_Payment_Mode":             "SELECT PaymentMode,PaymentMode_ID FROM " + dbConfig.DbDatabaseName + "." + dbConfig.LUTPaymentMode + "",
		"LUT_Program_Catalog":          "SELECT ProgramCode,ProgramName,ProgramType FROM " + dbConfig.DbDatabaseName + "." + dbConfig.LUTProgramCatalog + "",
		"LUT_Program_Types":            "SELECT ProgramType,ProgramTypeCode FROM " + dbConfig.DbDatabaseName + "." + dbConfig.LUTProgramType + "",
		"LUT_Skill_Prof":               "SELECT SkillProficiency,SkillProficiencyCode FROM " + dbConfig.DbDatabaseName + "." + dbConfig.LUTSkillProficiency + "",
		"LUT_Skills":                   "SELECT Skill_ID,SkillName,ifnull(DisableFlag,false) FROM " + dbConfig.DbDatabaseName + "." + dbConfig.LUTSkillsMaster + "",
		"LUT_SortBy":                   "SELECT SortBy,SortByCode FROM " + dbConfig.DbDatabaseName + "." + dbConfig.LUTSortBy + "",
		"LUT_Stakeholders":             "SELECT StakeholderTypeName,StakeholderTypeCode,OneLtrCode FROM " + dbConfig.DbDatabaseName + "." + dbConfig.LUTStakeholderType + "",
		"LUT_StudentEventResults":      "SELECT EventResult,EventResultCode FROM " + dbConfig.DbDatabaseName + "." + dbConfig.LUTStudentEventResult + "",
		"LUT_StudentEvent":             "SELECT EventType,EventTypeCode FROM " + dbConfig.DbDatabaseName + "." + dbConfig.LUTStudentEventType + "",
		"LUT_Student_Vrf_Status":       "SELECT VerificationStatus,VerificationStatusCode FROM " + dbConfig.DbDatabaseName + "." + dbConfig.LUTStudentProfileVerificationStatus + "",
		"LUT_Student_Vrf_Type":         "SELECT VerificationType,VerificationTypeCode FROM " + dbConfig.DbDatabaseName + "." + dbConfig.LUTStudentVerificationType + "",
		"LUT_Subscription_Type":        "SELECT SubscriptionType,SubscriptionTypeCode FROM " + dbConfig.DbDatabaseName + "." + dbConfig.LUTSubscriptionType + "",
		"LUT_Token_Events":             "SELECT Event,EventCode,StakeholderType FROM " + dbConfig.DbDatabaseName + "." + dbConfig.LUTTokenEventsDefinition + "",
		"LUT_University_Accreditation": "SELECT AccreditationType,AccreditationTypeCode FROM " + dbConfig.DbDatabaseName + "." + dbConfig.LUTUniversityAccreditationType + "",
		"LUT_University_Catalog":       "SELECT UniversityCode,UniversityName,UniversityAddress,UniversityType,University_State,DateofEstablishment FROM " + dbConfig.DbDatabaseName + "." + dbConfig.LUTUniversityCatalog + "",
		"LUT_University_Coes":          "SELECT COEType,COETypeCode FROM " + dbConfig.DbDatabaseName + "." + dbConfig.LUTUniversityCOEType + "",
		"LUT_University_Spl_Type":      "SELECT SplOfferingType,SplOfferingTypeCode FROM " + dbConfig.DbDatabaseName + "." + dbConfig.LUTUniversitySplOffersType + "",
		"LUT_University_Tie_up":        "SELECT TieUpType,TieUpTypeCode FROM " + dbConfig.DbDatabaseName + "." + dbConfig.LUTUniversityTieUpType + "",
		"LUT_University_Types":         "SELECT UniversityTypeName,UniversityTypeCode,OneLtrCode FROM " + dbConfig.DbDatabaseName + "." + dbConfig.LUTUniversityType + "",
	}
}
