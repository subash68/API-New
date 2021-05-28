package models

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/gin-contrib/cache/persistence"
)

var cacheDuration time.Duration = 15 * time.Minute

var cacheInMemoryStore *persistence.InMemoryStore = persistence.NewInMemoryStore(cacheDuration)

func queryRows(query string) (*sql.Rows, error) {
	sp, _ := RetriveSP(query)
	log.Debugf("Query : %s", sp)
	rows, err := Db.Query(sp)
	if err != nil {
		log.Fatalf("\n%s, Error: %v\n", QueryFetchError, err)
		return rows, fmt.Errorf("Failed to Get Data")
	}
	return rows, nil
}

// GetAllLutData ...
func (ald *AllLutData) GetAllLutData(reqLut []string, isIgnoreCache bool) error {
	err := CheckPing()
	if err != nil {
		return fmt.Errorf("Internal")
	}
	for _, v := range reqLut {

		log.Debugf("\nGetting Lut details of '%s'\n", v)
		existsInCache := false
		if !isIgnoreCache {
			existsInCache = getQueryFromCacheStore(ald, v)
		}
		if !existsInCache || isIgnoreCache {
			log.Debugf("\nGetting Lut details of '%s' from not cache \n", v)
			switch v {

			// 10Boards
			case queryList[0]:
				rows, err := queryRows("LUT_10Boards_GET")
				err = ald.ScanTNB(rows, err)
				break

			// 12Boards
			case queryList[1]:
				rows, err := queryRows("LUT_12Boards_GET")
				err = ald.ScanTWB(rows, err)
				break

			// Account Status
			case queryList[2]:
				rows, err := queryRows("LUT_AccountStatus")
				err = ald.ScanAS(rows, err)
				break

			// Branch catalog
			case queryList[3]:
				rows, err := queryRows("LUT_Branches")
				err = ald.ScanBranchCatalog(rows, err)
				break

			// Corporate Category
			case queryList[4]:
				rows, err := queryRows("LUT_CorporateCategory")
				err = ald.ScanCoporateCategory(rows, err)
				break

			// Corporate Industry
			case queryList[5]:
				rows, err := queryRows("LUT_CorporateIndustry")
				err = ald.ScanCoporateIndustry(rows, err)
				break

			// Corporate Type
			case queryList[6]:
				rows, err := queryRows("LUT_CorporateType")
				err = ald.ScanCoporateType(rows, err)
				break

			// Job Type
			case queryList[7]:
				rows, err := queryRows("LUT_JobType")
				err = ald.ScanJobType(rows, err)
				break

			// Language Proficiency
			case queryList[8]:
				rows, err := queryRows("LUT_Lang_Prof")
				err = ald.ScanLangProf(rows, err)
				break

			// Mode Of Token Issue
			case queryList[9]:
				rows, err := queryRows("LUT_Token_MOI")
				err = ald.ScanModeOfIssue(rows, err)
				break

			// Notification Type
			case queryList[10]:
				rows, err := queryRows("LUT_Nft_Type")
				err = ald.ScanNftType(rows, err)
				break

			// Payment Mode
			case queryList[11]:
				rows, err := queryRows("LUT_Payment_Mode")
				err = ald.ScanPaymentMode(rows, err)
				break

			// Program Catalog
			case queryList[12]:
				rows, err := queryRows("LUT_Program_Catalog")
				err = ald.ScanProgramCatalog(rows, err)
				break

			// Program Type
			case queryList[13]:
				rows, err := queryRows("LUT_Program_Types")
				err = ald.ScanProgramType(rows, err)
				break

			// Skill Proficiency
			case queryList[14]:
				rows, err := queryRows("LUT_Skill_Prof")
				err = ald.ScanSkillProficiency(rows, err)
				break

			// Skills
			case queryList[15]:
				rows, err := queryRows("LUT_Skills")
				err = ald.ScanSkills(rows, err)
				break

			// Sort by
			case queryList[16]:
				rows, err := queryRows("LUT_SortBy")
				err = ald.ScanSortBy(rows, err)
				break

			// Stakeholder types
			case queryList[17]:
				rows, err := queryRows("LUT_Stakeholders")
				err = ald.ScanStakeholderTypes(rows, err)
				break

			// Student Event Result
			case queryList[18]:
				rows, err := queryRows("LUT_StudentEventResults")
				err = ald.ScanStudentEventResult(rows, err)
				break

			// Student Event
			case queryList[19]:
				rows, err := queryRows("LUT_StudentEvent")
				err = ald.ScanStudentEvent(rows, err)
				break

			// Student Verification Status
			case queryList[20]:
				rows, err := queryRows("LUT_Student_Vrf_Status")
				err = ald.ScanStudentVrfStatus(rows, err)
				break

			// Student Verification Type
			case queryList[21]:
				rows, err := queryRows("LUT_Student_Vrf_Type")
				err = ald.ScanStudentVrfType(rows, err)
				break

			// Subscription Type
			case queryList[22]:
				rows, err := queryRows("LUT_Subscription_Type")
				err = ald.ScanSubscriptionType(rows, err)
				break

			// Token Event
			case queryList[23]:
				rows, err := queryRows("LUT_Token_Events")
				err = ald.ScanTokenEvent(rows, err)
				break

			// University Accreditation
			case queryList[24]:
				rows, err := queryRows("LUT_University_Accreditation")
				err = ald.ScanUniversityAccreditation(rows, err)
				break

			// University Catalog
			case queryList[25]:
				rows, err := queryRows("LUT_University_Catalog")
				err = ald.ScanUniversityCatalog(rows, err)
				break

			// University Coe
			case queryList[26]:
				rows, err := queryRows("LUT_University_Coes")
				err = ald.ScanUniversityCoe(rows, err)
				break

			// University Special offering Type
			case queryList[27]:
				rows, err := queryRows("LUT_University_Spl_Type")
				err = ald.ScanUniversitySplOffType(rows, err)
				break

			// University Tie Up Type
			case queryList[28]:
				rows, err := queryRows("LUT_University_Tie_up")
				err = ald.ScanUniversityTieUpType(rows, err)
				break

			// University Type
			case queryList[29]:
				rows, err := queryRows("LUT_University_Types")
				err = ald.ScanUniversityType(rows, err)
				break

			default:
				return fmt.Errorf("Invalid search query : %s, Expecting %v", v, queryList)
			}

			if err != nil {
				log.Errorf("Failed to get data, Error: %v", err)
				return fmt.Errorf("Internal")
			}
		}
	}
	return nil
}

func getQueryFromCacheStore(ald *AllLutData, key string) bool {
	dataInStore, exists := cacheInMemoryStore.Cache.Get(key)
	if !exists {
		return false
	}
	switch key {

	// 10Boards
	case queryList[0]:
		_, ok := dataInStore.([]Lut10BoardsModel)
		if !ok {
			return ok
		}
		ald.TenthBoards = dataInStore.([]Lut10BoardsModel)
		return ok

	// 12Boards
	case queryList[1]:
		_, ok := dataInStore.([]Lut10BoardsModel)
		if !ok {
			return ok
		}
		ald.TwelfthBoards = dataInStore.([]Lut10BoardsModel)
		return ok

	// Account Status
	case queryList[2]:
		_, ok := dataInStore.([]LutAccountStatusModel)
		if !ok {
			return ok
		}
		ald.AccountStatus = dataInStore.([]LutAccountStatusModel)
		return ok

	// Branch catalog
	case queryList[3]:
		_, ok := dataInStore.([]LutBranchCatalogModel)
		if !ok {
			return ok
		}
		ald.BranchCatalog = dataInStore.([]LutBranchCatalogModel)
		return ok

	// Corporate Category
	case queryList[4]:
		_, ok := dataInStore.([]LutCorporateCategoryModel)
		if !ok {
			return ok
		}
		ald.CoporateCategory = dataInStore.([]LutCorporateCategoryModel)
		return ok

	// Corporate Industry
	case queryList[5]:
		_, ok := dataInStore.([]LutCorporateIndustryModel)
		if !ok {
			return ok
		}
		ald.CoporateIndustry = dataInStore.([]LutCorporateIndustryModel)
		return ok

	// Corporate Type
	case queryList[6]:
		_, ok := dataInStore.([]LutCorporateTypeModel)
		if !ok {
			return ok
		}
		ald.CorporateType = dataInStore.([]LutCorporateTypeModel)
		return ok

	// Job Type
	case queryList[7]:
		_, ok := dataInStore.([]LutJobTypeModel)
		if !ok {
			return ok
		}
		ald.JobType = dataInStore.([]LutJobTypeModel)
		return ok

	// Language Proficiency
	case queryList[8]:
		_, ok := dataInStore.([]LutLangProficiencyModel)
		if !ok {
			return ok
		}
		ald.LanguageProficiency = dataInStore.([]LutLangProficiencyModel)
		return ok

	// Mode Of Token Issue
	case queryList[9]:
		_, ok := dataInStore.([]LutModeOfTokenIssueModel)
		if !ok {
			return ok
		}
		ald.ModeOfTokenIssue = dataInStore.([]LutModeOfTokenIssueModel)
		return ok

	// Notification Type
	case queryList[10]:
		_, ok := dataInStore.([]LutNotificationTypeModel)
		if !ok {
			return ok
		}
		ald.NotificationType = dataInStore.([]LutNotificationTypeModel)
		return ok

	// Payment Mode
	case queryList[11]:
		_, ok := dataInStore.([]LutPaymentModeModel)
		if !ok {
			return ok
		}
		ald.PaymentMode = dataInStore.([]LutPaymentModeModel)
		return ok

	// Program Catalog
	case queryList[12]:
		_, ok := dataInStore.([]LutProgramCatalogModel)
		if !ok {
			return ok
		}
		ald.ProgramCatalog = dataInStore.([]LutProgramCatalogModel)
		return ok

	// Program Type
	case queryList[13]:
		_, ok := dataInStore.([]LutProgramTypeModel)
		if !ok {
			return ok
		}
		ald.ProgramType = dataInStore.([]LutProgramTypeModel)
		return ok

	// Skill Proficiency
	case queryList[14]:
		_, ok := dataInStore.([]LutSkillProficiencyModel)
		if !ok {
			return ok
		}
		ald.SkillProficiency = dataInStore.([]LutSkillProficiencyModel)
		return ok

	// Skills
	case queryList[15]:
		_, ok := dataInStore.([]LutSkillsModel)
		if !ok {
			return ok
		}
		ald.Skills = dataInStore.([]LutSkillsModel)
		return ok

	// Sort by
	case queryList[16]:
		_, ok := dataInStore.([]LutSortByModel)
		if !ok {
			return ok
		}
		ald.SortBy = dataInStore.([]LutSortByModel)
		return ok

	// Stakeholder types
	case queryList[17]:
		_, ok := dataInStore.([]LutStakeholdersModel)
		if !ok {
			return ok
		}
		ald.StakeholderType = dataInStore.([]LutStakeholdersModel)
		return ok

	// Student Event Result
	case queryList[18]:
		_, ok := dataInStore.([]LutStudentEventResultModel)
		if !ok {
			return ok
		}
		ald.StudentEventResult = dataInStore.([]LutStudentEventResultModel)
		return ok

	// Student Event
	case queryList[19]:
		_, ok := dataInStore.([]LutStudentEventModel)
		if !ok {
			return ok
		}
		ald.StudentEvent = dataInStore.([]LutStudentEventModel)
		return ok

	// Student Verification Status
	case queryList[20]:
		_, ok := dataInStore.([]LutStudentVerfStatusModel)
		if !ok {
			return ok
		}
		ald.StudentVerificationStatus = dataInStore.([]LutStudentVerfStatusModel)
		return ok

	// Student Verification Type
	case queryList[21]:
		_, ok := dataInStore.([]LutStudentVerfTypeModel)
		if !ok {
			return ok
		}
		ald.StudentVerificationType = dataInStore.([]LutStudentVerfTypeModel)
		return ok

	// Subscription Type
	case queryList[22]:
		_, ok := dataInStore.([]LutSubscriptionType)
		if !ok {
			return ok
		}
		ald.SubscriptionType = dataInStore.([]LutSubscriptionType)
		return ok

	// Token Event
	case queryList[23]:
		_, ok := dataInStore.([]LutTokenEvent)
		if !ok {
			return ok
		}
		ald.TokenEvent = dataInStore.([]LutTokenEvent)
		return ok

	// University Accreditation
	case queryList[24]:
		_, ok := dataInStore.([]LutUniversityAccreditationModel)
		if !ok {
			return ok
		}
		ald.UniversityAccreditation = dataInStore.([]LutUniversityAccreditationModel)
		return ok

	// University Catalog
	case queryList[25]:
		_, ok := dataInStore.([]LutUniversityCatalogModel)
		if !ok {
			return ok
		}
		ald.UniversityCatalog = dataInStore.([]LutUniversityCatalogModel)
		log.Debugf("University catalog from Cache")
		return ok

	// University Coe
	case queryList[26]:
		_, ok := dataInStore.([]LutUniversityCoeModel)
		if !ok {
			return ok
		}
		ald.UniversityCoe = dataInStore.([]LutUniversityCoeModel)
		return ok

	// University Special offering Type
	case queryList[27]:
		_, ok := dataInStore.([]LutUniversitySplOffTypeModel)
		if !ok {
			return ok
		}
		ald.UniversitySpecialOfferingType = dataInStore.([]LutUniversitySplOffTypeModel)
		return ok

	// University Tie Up Type
	case queryList[28]:
		_, ok := dataInStore.([]LutUniversityTieUpTypeModel)
		if !ok {
			return ok
		}
		ald.UniversityTieUpType = dataInStore.([]LutUniversityTieUpTypeModel)
		return ok

	// University Type
	case queryList[29]:
		_, ok := dataInStore.([]LutUniversityTypeModel)
		if !ok {
			return ok
		}
		ald.UniversityType = dataInStore.([]LutUniversityTypeModel)
		return ok

	default:
		return false
	}
	return false
}
