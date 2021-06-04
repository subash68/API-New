package models

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/jaswanth-gorripati/PGK/s4_Profile/configuration"
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
		filter += " AND h.MinimumCutoffCGPA_Grad>=" + cutOff

		//fmt.Println("---->>>> ", qryHC, cutOff)
	}
	if len(industry) > 0 {
		strInd := strings.Join(industry, "','")
		filter += " AND c.CorporateIndustry IN ('" + strInd + "')"
	}
	if len(skills) > 0 {
		strSkills := strings.Join(skills, "','")
		filter += " AND sk.SkillName IN ('" + strSkills + "')"
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
func GetCorpByID(ID string, count int, shID string, userType string) (CorporateByIDResp, error) {
	dbNames := configuration.DbConfig()
	subDbName := ""
	if userType == "University" {
		subDbName = dbNames.UnvSubDBName
	} else if userType == "Student" {
		subDbName = dbNames.StuSubDBName
	} else if userType == "Corporate" {
		subDbName = dbNames.CrpSubDBName
	}
	sp, _ := RetriveSP("CORP_GET_PROFILE_BY_ID")
	sp = strings.ReplaceAll(sp, "//RPLCSUB", subDbName)
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
	subSP, _ := RetriveSP("CORP_HCI_GET_ALL_SUB")
	fmt.Println("========================== CORP_HCI_GET_ALL_SUB==========", subSP)
	subrow, err := Db.Query(subSP, ID, shID)
	if err != nil && err != sql.ErrNoRows {
		return corpDb, fmt.Errorf("Cannot get the Rows %v", err.Error())
	} else if err == sql.ErrNoRows {

	} else {
		defer subrow.Close()
		for subrow.Next() {
			var newsub SubscriptionReq
			err = subrow.Scan(&newsub.SubscriptionID, &newsub.Publisher, &newsub.Subscriber, &newsub.DateOfSubscription)
			newsub.GeneralNote = "Hiring Insights" // strings.Split(newsub.GeneralNote, " has been published")[0]
			if err != nil {
				return corpDb, fmt.Errorf("Cannot read the Rows %v", err.Error())
			}
			corpDb.Subscriptions = append(corpDb.Subscriptions, newsub)
		}
	}
	subSP, _ = RetriveSP("UNV_CD_GET_ALL")
	fmt.Println("========================== UNV_CD_GET_ALL==========", subSP, shID, ID)
	subrow, err = Db.Query(subSP, ID, shID)
	if err != nil && err != sql.ErrNoRows {
		return corpDb, fmt.Errorf("Cannot get the Rows %v", err.Error())
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
				return corpDb, fmt.Errorf("Cannot read the Rows %v", err.Error())
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
			corpDb.Subscriptions = append(corpDb.Subscriptions, newsub)
		}
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

	subSP, _ = RetriveSP("CORP_GET_PUB_DATA_FOR_ID")
	subSP = strings.ReplaceAll(subSP, "//RPLCSUB", subDbName)
	fmt.Println("========================== CORP_GET_PUB_DATA_FOR_ID==========", subSP)
	subrow, err = Db.Query(subSP, shID, ID)
	if err != nil && err != sql.ErrNoRows {
		return corpDb, fmt.Errorf("Cannot get the Rows %v", err.Error())
	} else if err == sql.ErrNoRows {

	} else {
		defer subrow.Close()
		for subrow.Next() {
			var newsub CorpPushedDataReq
			err = subrow.Scan(&newsub.PublishID, &newsub.DateOfPublish, &newsub.HiringCriteriaPublished, &newsub.JobsPublished, &newsub.ProfilePublished, &newsub.OtherPublished, &newsub.GeneralNote, &newsub.IsSubscribed, &newsub.PublishedData)
			if err != nil {
				return corpDb, fmt.Errorf("Cannot read the Rows %v", err.Error())
			}
			newsub.GeneralNote = strings.Split(newsub.GeneralNote, " has been published")[0]
			newsub.Info = make(map[string]string)
			newsub.parseDataInfo()
			corpDb.PublishedData = append(corpDb.PublishedData, newsub)
		}
	}

	corpDb.Jobs = ""
	return corpDb, nil
}

func (cpd *CorpPushedDataReq) parseDataInfo() {
	//fmt.Printf("PublishId : %v \n PublishData %v\nhc: %v\n jp: %v \n,oi: %v \n", cpd.PublishID, cpd.PublishedData, cpd.HiringCriteriaPublished, cpd.JobsPublished, cpd.OtherPublished)
	if cpd.PublishedData == "" || cpd.PublishedData == "[]" {
		fmt.Println("======= No data found for id %s =======", cpd.PublishID)
		return

	}
	var err error
	if cpd.HiringCriteriaPublished {
		hcInfo := HiringCriteriaDB{}
		err = json.Unmarshal([]byte(cpd.PublishedData), &hcInfo)
		if err == nil {
			cpd.Info["programs"] = hcInfo.ProgramsInString
			return
		}
	} else if cpd.JobsPublished {
		jobInfo := JobHcMappingDB{}
		err = json.Unmarshal([]byte(cpd.PublishedData), &jobInfo)
		if err == nil {
			cpd.Info["hiringCriteriaID"] = jobInfo.HcID
			cpd.Info["hiringCriteriaName"] = jobInfo.HcName
			return
		}
	} else if cpd.OtherPublished {
		oiInfo := UnvOtherInformationModel{}
		err = json.Unmarshal([]byte(cpd.PublishedData), &oiInfo)
		fmt.Printf("`````````````````other info %v \n  %+v``````````", cpd.PublishedData, oiInfo)
		if err == nil {
			cpd.Info["Title"] = oiInfo.Title
			return
		}
	}
	if err != nil {
		fmt.Printf("Error while un marshaling data Published data === > %v \n Error: %v\n", cpd.PublishedData, err.Error())
	}
	return
}
