package models

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/jaswanth-gorripati/PGK/s3_Publish/configuration"
)

// Corporate ...
const (
	Corporate         string = "Corporate"
	University        string = "University"
	Student           string = "Student"
	HideInfoString    string = "XXXXXXXXXXXXX"
	NftSubQueryString string = "NFT_PDH_GET_SUB_DATA"
)

// NftSubscriptionData ...
var (
	NftSubscriptionData nftSubscriptionData = nftSubscriptionData{}
	NonSubscriptionKeys []string            = []string{"hiringCriteriaName", "jobName", "programID", "universityName", "corporateName", "title", "information"}
)

type nftSubscriptionData struct{}

// NftSubDataRespModel  ...
type NftSubDataRespModel struct {
	IsSubscribed        bool                   `json:"isSubscribed"`
	DateOfPublish       string                 `json:"dateOfPublish`
	GeneralNote         string                 `json:"generalNote"`
	PublishedData       map[string]interface{} `json:"publishedData"`
	PublishedDataString string                 `json:"-"`
}

func (nsd *nftSubscriptionData) GetData(ID string, userRole string, publisherRole string, publishID string) (NftSubDataRespModel, error) {
	var nftData NftSubDataRespModel
	dbConfig := configuration.DbConfig()
	var subDbName, pubDbName string
	switch userRole {
	case Corporate:
		subDbName = dbConfig.CrpSubDBName
		break
	case University:
		subDbName = dbConfig.UnvSubDBName
		break
	case Student:
		subDbName = dbConfig.StuSubDBName
		break
	default:
		return nftData, fmt.Errorf("Invalid Publisher UserRole")
	}
	switch publisherRole {
	case Corporate:
		pubDbName = dbConfig.CorpPDHDbName
		break
	case University:
		pubDbName = dbConfig.UnvPDHDbName
		break
	// case Student:
	// 	subDbName =  dbConfig.StuPublishDBName
	// 	break
	default:
		return nftData, fmt.Errorf("Invalid Publisher UserRole")
	}

	sp, _ := RetriveSP(NftSubQueryString)

	sp = strings.ReplaceAll(sp, "//REPLACESUB", subDbName)
	sp = strings.ReplaceAll(sp, "//REPLACEPUB", pubDbName)
	fmt.Println("==========" + sp + "=============" + ID + "=====" + publishID)
	err := Db.QueryRow(sp, ID, publishID).Scan(&nftData.DateOfPublish, &nftData.GeneralNote, &nftData.PublishedDataString, &nftData.IsSubscribed)
	if err != nil && err == sql.ErrNoRows {
		nftData.GeneralNote = "No information found, Contact Admin"
		return nftData, nil
	} else if err != nil {
		return nftData, err
	}
	//nftData.PublishedDataString = strings.ReplaceAll(nftData.PublishedDataString, "\"", fmt.Sprintf("'"))
	fmt.Printf("\n%s\n", nftData.PublishedDataString)
	err = json.Unmarshal([]byte(nftData.PublishedDataString), &nftData.PublishedData)
	if err != nil {
		fmt.Println("========================JSON ERROR==========================")
		return nftData, err
	}

	//nftData.IsSubscribed = true
	if nftData.IsSubscribed == false {
		for key, value := range nftData.PublishedData {
			bool, _ := InArray(key, NonSubscriptionKeys)
			if bool {
				nftData.PublishedData[key] = value
			} else {
				nftData.PublishedData[key] = HideInfoString
			}
			fmt.Println(fmt.Sprintf("%s : %s", key, value))
		}
	}
	delete(nftData.PublishedData, "publishID")
	delete(nftData.PublishedData, "publishedFlag")
	return nftData, nil
}
