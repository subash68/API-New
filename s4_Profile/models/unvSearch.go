package models

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/jaswanth-gorripati/PGK/s4_Profile/configuration"
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
func GetUnvByID(ID string, subID string, subUserType string) (UniversityGetByIDModel, error) {
	dbNames := configuration.DbConfig()
	var subDbName string
	if subUserType == "University" {
		subDbName = dbNames.UnvSubDBName
	} else if subUserType == "Student" {
		subDbName = dbNames.StuSubDBName
	} else if subUserType == "Corporate" {
		subDbName = dbNames.CrpSubDBName
	}
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
	subSP = strings.ReplaceAll(subSP, "//REPLCESUBDB", subDbName)
	fmt.Println("========================== UNV_SUB_DATA_IN_SRH==========", subSP)
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
	subSP, _ = RetriveSP("UNV_STU_DB_SUB_GET_ALL")
	fmt.Println("========================== UNV_GET_PROFILE_BY_ID==========", subSP)
	subrow, err = Db.Query(subSP, subID, ID)
	if err != nil && err != sql.ErrNoRows {
		return unvDB, fmt.Errorf("Cannot get the Rows %v", err.Error())
	} else if err == sql.ErrNoRows {

	} else {
		defer subrow.Close()
		for subrow.Next() {
			var newsub SubscriptionReq
			err = subrow.Scan(&newsub.SubscriptionID, &newsub.Publisher, &newsub.Subscriber, &newsub.DateOfSubscription, &newsub.SearchCriteria)
			newsub.GeneralNote = "Student Database" // strings.Split(newsub.GeneralNote, " has been published")[0]
			if err != nil {
				return unvDB, fmt.Errorf("Cannot read the Rows %v", err.Error())
			}
			unvDB.Subscriptions = append(unvDB.Subscriptions, newsub)
		}
	}

	subSP, _ = RetriveSP("UNV_INSIGHTS_GET_ALL")
	fmt.Println("========================== UNV_INSIGHTS_GET_ALL==========", subSP)
	subrow, err = Db.Query(subSP, subID, ID)
	if err != nil && err != sql.ErrNoRows {
		return unvDB, fmt.Errorf("Cannot get the Rows %v", err.Error())
	} else if err == sql.ErrNoRows {

	} else {
		defer subrow.Close()
		for subrow.Next() {
			var newsub SubscriptionReq
			err = subrow.Scan(&newsub.SubscriptionID, &newsub.Publisher, &newsub.Subscriber, &newsub.DateOfSubscription)
			newsub.GeneralNote = "University Information" // strings.Split(newsub.GeneralNote, " has been published")[0]
			if err != nil {
				return unvDB, fmt.Errorf("Cannot read the Rows %v", err.Error())
			}
			unvDB.Subscriptions = append(unvDB.Subscriptions, newsub)
		}
	}
	subSP, _ = RetriveSP("CORP_CD_GET_ALL")
	fmt.Println("========================== CORP_CD_GET_ALL==========", subSP, subID, ID)
	subrow, err = Db.Query(subSP, ID, subID)
	if err != nil && err != sql.ErrNoRows {
		return unvDB, fmt.Errorf("Cannot get the Rows %v", err.Error())
	} else if err == sql.ErrNoRows {

	} else {
		defer subrow.Close()
		for subrow.Next() {
			var newsub SubscriptionReq
			var cdReq, cdAr bool
			var rqDate, arDate time.Time
			var reqNftID, arNftID string
			err = subrow.Scan(&newsub.Subscriber, &newsub.Publisher, &newsub.CampusDriveID, &cdReq, &rqDate, &reqNftID, &cdAr, &arDate, &arNftID)
			newsub.GeneralNote = "Campus Hiring" // strings.Split(newsub.GeneralNote, " has been published")[0]
			if err != nil {
				return unvDB, fmt.Errorf("Cannot read the Rows %v", err.Error())
			}
			fmt.Printf("\n\n==== Campus details %+v , cdr %v , adAr %v===\n\n", newsub, cdReq, cdAr)
			if cdReq == true && cdAr == true && arNftID != "" {
				newsub.CampusDriveStatus = "Accepted"
				newsub.NftID = arNftID
				newsub.DateOfSubscription = rqDate
			} else if cdReq == true && cdAr == false && arNftID != "" {
				newsub.CampusDriveStatus = "Rejected"
				newsub.NftID = arNftID
				newsub.DateOfSubscription = rqDate
			} else if cdReq == true && cdAr == false && arNftID == "" {
				newsub.CampusDriveStatus = "Pending"
				newsub.NftID = reqNftID
				newsub.DateOfSubscription = rqDate
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
	pubSP, _ := RetriveSP("UNV_GET_PUBDATA_FOR_ID")
	pubSP = strings.ReplaceAll(pubSP, "//REPLCESUBDB", subDbName)
	fmt.Println("========================== UNV_GET_PUBDATA_FOR_ID==========", pubSP)
	pubrow, err := Db.Query(pubSP, subID, ID)
	if err != nil && err != sql.ErrNoRows {
		return unvDB, fmt.Errorf("Cannot get the Rows %v", err.Error())
	} else if err == sql.ErrNoRows {

	} else {
		defer pubrow.Close()
		for pubrow.Next() {
			var newpub UnvPublishReqModel
			err = pubrow.Scan(&newpub.PublishID, &newpub.DateOfPublish, &newpub.ProgramsPublished, &newpub.BranchesPublished, &newpub.StudentStrengthPublished, &newpub.AcredPublished, &newpub.COEsPublished, &newpub.RankingPublished, &newpub.OtherPublished, &newpub.ProfilePublished, &newpub.GeneralNote, &newpub.IsSubscribed, &newpub.PublishedData)
			newpub.GeneralNote = strings.Split(newpub.GeneralNote, " has been published")[0]
			newpub.Info = make(map[string]string)
			if newpub.GeneralNote == "Profile" {
				publishedTypes := ""
				if newpub.ProgramsPublished {
					publishedTypes += "Programs,"
				}
				if newpub.BranchesPublished {
					publishedTypes += "Branches,"
				}
				if newpub.StudentStrengthPublished {
					publishedTypes += "Student Strength,"
				}
				if newpub.AcredPublished {
					publishedTypes += "Accredations,"
				}
				if newpub.COEsPublished {
					publishedTypes += "COEs,"
				}
				if newpub.RankingPublished {
					publishedTypes += "Ranking,"
				}
				if newpub.ProfilePublished {
					publishedTypes += "Profile,"
				}
				newpub.Info["PublishedData"] = publishedTypes[:len(publishedTypes)-1]
			}
			if newpub.OtherPublished {

				newpub.parseOtherInfo()
			}
			if err != nil {
				return unvDB, fmt.Errorf("Cannot read the Rows %v", err.Error())
			}
			unvDB.PublishedData = append(unvDB.PublishedData, newpub)
		}
	}

	return unvDB, nil
}

func (upr *UnvPublishReqModel) parseOtherInfo() {
	oi := UnvOtherInformationModel{}
	err := json.Unmarshal([]byte(upr.PublishedData), &oi)
	if err != nil {
		fmt.Println("========= Error while parsing oi ===", err.Error())
	}
	fmt.Printf("\n%+v\n", oi)
	upr.Info["Title"] = oi.Title
	return

}
