package models

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

// CorpSearchModel ...
type CorpSearchModel struct {
	CorporateID       string `json:"corporateID"`
	Name              string `json:"corporateName"`
	Industry          string `json:"corporateIndustry"`
	Locations         string `json:"locations"`
	AvgHiringCriteria string `json:"avgHiringCriteria"`
}

// CrpSearchResults ...
type CrpSearchResults struct {
	Corporates []CorpSearchModel `json:"corporates"`
}

// SeacrhCorporate ...
func SeacrhCorporate(corporateName string, industry []string, skills []string, locations []string, cutOff string) (CrpSearchResults, error) {
	filter := ""
	qryHC := 0.0
	var err error
	var crpSearchResults CrpSearchResults
	if corporateName != "" {
		filter += " AND c.Corporate_Name LIKE '%" + corporateName + "%'"
	}
	if cutOff != "" {
		qryHC, err = strconv.ParseFloat(cutOff, 64)
		if err != nil {
			return crpSearchResults, fmt.Errorf("Invalid cutoff value %v", err.Error())
		}
		filter += " AND h.MinimumCutoff>=" + cutOff

		//fmt.Println("---->>>> ", qryHC, cutOff)
	}
	if len(industry) > 0 {
		strInd := strings.Join(industry, "','")
		filter += " AND c.CorporateIndustry IN ('" + strInd + "')"
	}
	if len(skills) > 0 {
		strSkills := strings.Join(skills, "','")
		filter += " AND a.SkillName IN ('" + strSkills + "')"
	}
	if len(locations) > 0 {
		strLocations := strings.Join(locations, "','")
		filter += " AND a.Location IN ('" + strLocations + "')"
	}

	sp, _ := RetriveSP("SRH_CRP_BY_SKILL_MAPPING")
	sp += filter
	fmt.Println(sp, qryHC, cutOff, cutOff != "")
	rows, err := Db.Query(sp+filter, qryHC)
	if err != nil {
		return crpSearchResults, fmt.Errorf("Cannot get the Rows %v", err.Error())
	}
	defer rows.Close()
	for rows.Next() {
		var row CorpSearchModel
		err = rows.Scan(&row.CorporateID, &row.Locations, &row.AvgHiringCriteria, &row.Name, &row.Industry)
		if err != nil {
			return crpSearchResults, fmt.Errorf("Cannot read the Rows %v", err.Error())
		}
		crpSearchResults.Corporates = append(crpSearchResults.Corporates, row)
	}
	return crpSearchResults, nil
}

// GetCorpByID ...
func GetCorpByID(ID string, count int, shID string) (CorporateByIDResp, error) {
	sp, _ := RetriveSP("CORP_GET_PROFILE_BY_ID")
	fmt.Println("========================== CORP_GET_PROFILE_BY_ID==========", sp)
	row := Db.QueryRow(sp, shID, count, ID)
	corpDb := CorporateByIDResp{}
	//var result map[string]interface{}

	// Unmarshal or Decode the JSON to the interface.

	err := row.Scan(&corpDb.StakeholderID, &corpDb.CorporateName, &corpDb.CIN, &corpDb.CorporateHQAddressLine1, &corpDb.CorporateHQAddressLine2, &corpDb.CorporateHQAddressLine3, &corpDb.CorporateHQAddressCountry, &corpDb.CorporateHQAddressState, &corpDb.CorporateHQAddressCity, &corpDb.CorporateHQAddressDistrict, &corpDb.CorporateHQAddressZipCode, &corpDb.CorporateLocalBranchAddressLine1, &corpDb.CorporateLocalBranchAddressLine2, &corpDb.CorporateLocalBranchAddressLine3, &corpDb.CorporateLocalBranchAddressCountry, &corpDb.CorporateLocalBranchAddressState, &corpDb.CorporateLocalBranchAddressCity, &corpDb.CorporateLocalBranchAddressDistrict, &corpDb.CorporateLocalBranchAddressZipCode, &corpDb.CorporateType, &corpDb.CorporateCategory, &corpDb.CorporateIndustry, &corpDb.CompanyProfile, &corpDb.YearOfEstablishment, &corpDb.DateOfJoining, &corpDb.Jobs)
	if err != nil {

		return corpDb, fmt.Errorf("Cannot scan ros due to : %v", err.Error())
	}
	fmt.Println("==================", corpDb.Jobs, corpDb.Jobs == "", corpDb, " ===================")
	if corpDb.StakeholderID == "" {
		return corpDb, fmt.Errorf("User details not found for ID %s", ID)
	}

	if corpDb.Jobs != "" {
		corpDb.Jobs = strings.ReplaceAll(corpDb.Jobs, "34", fmt.Sprint("\""))
		//corpDb.Jobs = strings.ReplaceAll(corpDb.Jobs, "34", fmt.Sprint('"'))
		err = json.Unmarshal([]byte(corpDb.Jobs), &corpDb.JobsAvailable)
		if err != nil {

			return corpDb, fmt.Errorf("Cannot  ros due to : %v ---> : %v", corpDb.Jobs, err)
		}
	} else {
		corpDb.JobsAvailable = []map[string]interface{}{}
	}
	corpDb.Jobs = ""
	return corpDb, nil
}
