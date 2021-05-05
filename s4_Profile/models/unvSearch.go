package models

import (
	"database/sql"
	"fmt"
	"strings"
)

// UnvSearchModel ...
type UnvSearchModel struct {
	UniversityID string `json:"universityID"`
	Name         string `json:"universityName"`
	Locations    string `json:"locations"`
	NirfRanking  string `json:"ranking"`
	Accredations string `json:"accredations"`
}

// UnvSearchResults ...
type UnvSearchResults struct {
	Universities []UnvSearchModel `json:"universities"`
}

// SearchUniversities ...
func SearchUniversities(unvName string, hcID string, skills []string, locations []string, cutOff string) (UnvSearchResults, error) {
	filter := ""
	qrySP := ""
	vals := []interface{}{}
	var err error
	var unvSearchResults UnvSearchResults
	if hcID != "" {
		qrySP = "SRH_UNV_BY_HC"
		vals = append(vals, hcID)
	} else {
		if unvName != "" {
			filter += " u.University_Name LIKE '%" + unvName + "%'"
		}
		// if cutOff != "" {
		// 	qryHC, err = strconv.ParseFloat(cutOff, 64)
		// 	if err != nil {
		// 		return unvSearchResults, fmt.Errorf("Invalid cutoff value %v", err.Error())
		// 	}
		// 	filter += " AND h.MinimumCutoff>=" + cutOff

		// 	//fmt.Println("---->>>> ", qryHC, cutOff)
		// }
		// if len(industry) > 0 {
		// 	strInd := strings.Join(industry, "','")
		// 	filter += " AND c.CorporateIndustry IN ('" + strInd + "')"
		// }
		// if len(skills) > 0 {
		// 	strSkills := strings.Join(skills, "','")
		// 	filter += " AND a.SkillName IN ('" + strSkills + "')"
		// }
		if len(locations) > 0 {
			strLocations := strings.Join(locations, "','")
			if filter != "" {
				filter += " AND "
			}
			filter += "u.UniversityHQAddress_City IN ('" + strLocations + "')"
		}
		qrySP = "SRH_UNV_BY_NSLC"
		//vals = append(vals, hcID)
	}

	sp, _ := RetriveSP(qrySP)
	sp += filter
	fmt.Println(sp)
	rows, err := Db.Query(sp, vals...)
	if err != nil {
		return unvSearchResults, fmt.Errorf("Cannot get the Rows %v", err.Error())
	}
	defer rows.Close()
	for rows.Next() {
		var row UnvSearchModel
		err = rows.Scan(&row.UniversityID, &row.Name, &row.Locations, &row.NirfRanking, &row.Accredations)
		if err != nil {
			return unvSearchResults, fmt.Errorf("Cannot read the Rows %v", err.Error())
		}
		unvSearchResults.Universities = append(unvSearchResults.Universities, row)
	}
	return unvSearchResults, nil
}

// GetUnvByID ...
func GetUnvByID(ID string, subID string) (UniversityGetByIDModel, error) {
	sp, _ := RetriveSP("UNV_GET_PROFILE_BY_ID")
	fmt.Println("========================== UNV_GET_PROFILE_BY_ID==========", sp)
	row := Db.QueryRow(sp, ID)
	unvDB := UniversityGetByIDModel{}
	//var result map[string]interface{}

	// Unmarshal or Decode the JSON to the interface.

	err := row.Scan(&unvDB.StakeholderID, &unvDB.UniversityName, &unvDB.UniversityCollegeID, &unvDB.UniversityHQAddressCity, &unvDB.YearOfEstablishment, &unvDB.UniversityProfile, &unvDB.ProgramsOffered, &unvDB.Ranking, &unvDB.Accredations, &unvDB.StudentStrengthNullable, &unvDB.UnvInsightsAvailable)
	if err != nil {

		return unvDB, fmt.Errorf("Cannot scan ros due to : %v", err.Error())
	}
	subSP, _ := RetriveSP("UNV_SUB_DATA_IN_SRH")
	fmt.Println("========================== UNV_GET_PROFILE_BY_ID==========", sp)
	subrow, err := Db.Query(subSP, subID, ID)
	if err != nil && err != sql.ErrNoRows {
		return unvDB, fmt.Errorf("Cannot get the Rows %v", err.Error())
	} else if err == sql.ErrNoRows {

	} else {
		defer subrow.Close()
		for subrow.Next() {
			var newsub SubscriptionReq
			err = subrow.Scan(&newsub.Publisher, &newsub.DateOfSubscription, &newsub.PublishID, &newsub.TransactionID, &newsub.GeneralNote)
			newsub.GeneralNote = strings.Split(newsub.GeneralNote, " has been published")[0]
			if err != nil {
				return unvDB, fmt.Errorf("Cannot read the Rows %v", err.Error())
			}
			unvDB.Subscriptions = append(unvDB.Subscriptions, newsub)
		}
	}
	subSP, _ = RetriveSP("UNV_INSIGHTS_GET_ALL")
	fmt.Println("========================== UNV_GET_PROFILE_BY_ID==========", sp)
	subrow, err = Db.Query(subSP, subID, ID)
	if err != nil && err != sql.ErrNoRows {
		return unvDB, fmt.Errorf("Cannot get the Rows %v", err.Error())
	} else if err == sql.ErrNoRows {

	} else {
		defer subrow.Close()
		for subrow.Next() {
			var newsub SubscriptionReq
			err = subrow.Scan(&newsub.SubscriptionID, &newsub.Subscriber, &newsub.Publisher, &newsub.DateOfSubscription)
			newsub.GeneralNote = "University Information" // strings.Split(newsub.GeneralNote, " has been published")[0]
			if err != nil {
				return unvDB, fmt.Errorf("Cannot read the Rows %v", err.Error())
			}
			unvDB.Subscriptions = append(unvDB.Subscriptions, newsub)
		}
	}
	if unvDB.StakeholderID == "" {
		return unvDB, fmt.Errorf("User details not found for ID %s", ID)
	}
	if unvDB.StudentStrengthNullable.Valid {
		unvDB.StudentDbAvailable = true
		unvDB.StudentDbPublishID = unvDB.StudentStrengthNullable.String
	}

	return unvDB, nil
}
