package models

import (
	"database/sql"
	"fmt"
)

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
func (ald *AllLutData) GetAllLutData(reqLut []string) error {
	err := CheckPing()
	if err != nil {
		return fmt.Errorf("Internal")
	}
	for _, v := range reqLut {

		log.Debugf("\nGetting Lut details of '%s'\n", v)

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

		default:
			return fmt.Errorf("Invalid search query : %s, Expecting %v", v, queryList)
		}
		if err != nil {
			log.Errorf("Failed to get data, Error: %v", err)
			return fmt.Errorf("Internal")
		}
	}
	return nil
}
