package models

import (
	"database/sql"
	"fmt"
	"strconv"
	"time"
)

// NftPersistance ...
var (
	NftPersistance nftPersistance = nftPersistance{}
)

type nftPersistance struct{}

// Corporate ...
const (
	Corporate  string = "Corporate"
	University string = "University"
	Student    string = "Student"
)

// AddNotification ...
func (np *nftPersistance) AddNotification(newNft NotificationsModel) (string, error) {
	var err error
	senderRole := ""
	switch newNft.SenderUserRole {
	case Corporate:
		senderRole = "CRP"
		break
	case University:
		senderRole = "UNV"
		break
	case Student:
		senderRole = "STU"
		break
	default:
		return "", fmt.Errorf("Invalid Publisher UserRole")
	}

	newNft.NotificationID, err = CreateNftID(newNft.SenderID, senderRole)
	if err != nil {
		return "", err
	}
	nftInsCmd, _ := RetriveSP("NFT_INS")

	stmt, err := Db.Prepare(nftInsCmd)
	if err != nil {
		fmt.Printf("Cannot prepare query -- %v  -- due to %v", nftInsCmd, err.Error())
		return "", fmt.Errorf("Cannot prepare query -- %v  -- due to %v", nftInsCmd, err.Error())
	}
	currentTime := time.Now()

	fmt.Printf("\n========= Add Notification query : %+v ====\n ======\n", newNft)
	//  isGeneric := false
	//  if newNft.GenericMessage == "true" {
	// 	isGeneric = true
	//  }
	_, err = stmt.Exec(newNft.NotificationID, newNft.SenderID, newNft.SenderUserRole, newNft.ReceiverID,
		currentTime, newNft.NotificationType, newNft.Content, newNft.AttachFile, newNft.RedirectedURL,
		newNft.PublishID, newNft.PublishFlag, currentTime, currentTime, newNft.NotificationTypeID, newNft.GenericMessage)
	if err != nil {
		fmt.Printf("Failed to insert into database -- %v -- insert due to %v", nftInsCmd, err.Error())
		return "", fmt.Errorf("Failed to insert into database -- %v -- insert due to %v", nftInsCmd, err.Error())
	}
	return newNft.NotificationID, nil
}

// GetAllNotifications ...
func (np *nftPersistance) GetAllNotifications(ID string, filters string, page int, size int, userType string) ([]NotificationsModel, error) {
	// Preparing Database insert
	nftGetAllCmd, _ := RetriveSP("NFT_GET_ALL")
	nftGroupCmd, _ := RetriveSP("NFT_GROUP_COND")

	var allNfts []NotificationsModel
	nftGetAllCmd = nftGetAllCmd + filters + nftGroupCmd
	fmt.Println(nftGetAllCmd, ID, ID, ((page - 1) * size), size)

	nftRows, err := Db.Query(nftGetAllCmd, ID, ID, userType, ((page - 1) * size), size)
	if err != nil && err != sql.ErrNoRows {
		return allNfts, fmt.Errorf("Cannot get the Rows %v", err.Error())
	} else if err == sql.ErrNoRows {
		return allNfts, nil
	}
	defer nftRows.Close()
	for nftRows.Next() {
		var newNFT NotificationsModel
		err = nftRows.Scan(&newNFT.NotificationID, &newNFT.SenderID, &newNFT.SenderUserRole, &newNFT.ReceiverID,
			&newNFT.DateofNotification, &newNFT.NotificationType, &newNFT.Content, &newNFT.AttachFile, &newNFT.RedirectedURL,
			&newNFT.PublishID, &newNFT.PublishFlag, &newNFT.CreationDate, &newNFT.LastUpdatedDate)
		if err != nil {
			return allNfts, err
		}
		if newNFT.SenderID != "" {
			switch newNFT.SenderUserRole {
			case "Corporate":
				newNFT.SenderName = getName("GET_CORPORATE_NAME", newNFT.SenderID)
				break
			case "University":
				newNFT.SenderName = getName("GET_UNIVERSITY_NAME", newNFT.SenderID)
				break
			case "Student":
				newNFT.SenderName = getName("GET_STUDENT_NAME", newNFT.SenderID)
				break
			}
		}
		allNfts = append(allNfts, newNFT)
	}

	return allNfts, nil
}

// GetNotificationByID ...
func (np *nftPersistance) GetNotificationByID(ID string, nftID string) (NotificationsModel, error) {
	// Preparing Database insert
	nftGetAllCmd, _ := RetriveSP("NFT_GET_BY_ID")

	var allNfts NotificationsModel

	err := Db.QueryRow(nftGetAllCmd, nftID, ID, ID).Scan(&allNfts.NotificationID, &allNfts.SenderID, &allNfts.SenderUserRole, &allNfts.ReceiverID,
		&allNfts.DateofNotification, &allNfts.NotificationType, &allNfts.Content, &allNfts.AttachFile, &allNfts.RedirectedURL,
		&allNfts.PublishID, &allNfts.PublishFlag, &allNfts.CreationDate, &allNfts.LastUpdatedDate)
	if err != nil && err != sql.ErrNoRows {
		return allNfts, fmt.Errorf("Cannot get the Rows %v", err.Error())
	}
	if err != nil && err == sql.ErrNoRows {
		return allNfts, fmt.Errorf("Invalid / Unauthorized Notification ID")
	}

	return allNfts, nil
}

// CreateNftID ...
func CreateNftID(senderID string, senderRole string) (string, error) {
	rowSP, _ := RetriveSP("NFT_Get_Last_ID")
	lastID := ""
	err := Db.QueryRow(rowSP, senderID).Scan(&lastID)

	if err != nil && err != sql.ErrNoRows {
		return "", fmt.Errorf("Failed to create Notification ID ", err)
	}
	if err == sql.ErrNoRows {
		lastID = "0000000000000"
	}
	corporateNum, _ := strconv.Atoi(senderID[7:])
	countNum, _ := strconv.Atoi(lastID[len(lastID)-7:])
	fmt.Println("--------------------> ", lastID, countNum)

	return "NFT" + senderRole + strconv.Itoa(corporateNum) + (fmt.Sprintf("%07d", (countNum + 1))), nil
}

func getName(query string, ID string) (name string) {
	nameQryCmd, _ := RetriveSP(query)
	err := Db.QueryRow(nameQryCmd, ID).Scan(&name)
	fmt.Println(err)
	return name
}
