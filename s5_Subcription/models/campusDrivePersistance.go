package models

import (
	"database/sql"
	"fmt"
	"time"
)

// Corporate ...
const (
	Corporate       string = "Corporate"
	University      string = "University"
	LastCrpID       string = "CORP_CD_Get_Last_ID"
	LastUnvID       string = "UNV_CD_Get_Last_ID"
	CrpCode         string = "CCDI"
	UnvCode         string = "UCPI"
	CrpInsCmd       string = "CORP_CD_INIT"
	UnvInsCmd       string = "UNV_CD_INIT"
	CrpInviteCmd    string = "CORP_CD_SUB_UPDATE"
	UnvInviteCmd    string = "UNV_CD_SUB_UPDATE"
	CrpRespondCmd   string = "CORP_CD_UNV_RESP"
	UnvRespondCmd   string = "UNV_CD_UNV_RESP"
	CrpGetByIDCmd   string = "CORP_CD_GET_BY_ID"
	UnvGetByIDCmd   string = "UNV_CD_GET_BY_ID"
	CrpGetIRByIDCmd string = "CORP_CD_GET_INITIATOR_FR_CD"
	UnvGetIRByIDCmd string = "UNV_CD_GET_INITIATOR_FR_CD"
)

// SubscribeToInviteForCD ...
func (cdm *CampusDriveDataModel) SubscribeToInviteForCD(lastQueryCmd string, code string, insCmd string) error {
	// var lastQueryCmd, code, insCmd string
	// //var initiatorEmail,receiverEmail string
	// switch userType {
	// case Corporate:
	// 	lastQueryCmd, code, insCmd = LastCrpID, CrpCode, CrpInsCmd
	// 	break
	// case University:
	// 	lastQueryCmd, code, insCmd = LastUnvID, UnvCode, UnvInsCmd
	// 	break
	// default:
	// 	return fmt.Errorf("Invalid user type %s", userType)
	// }
	// var err error
	// cdm.CampusDriveID, err = CreateSudID(cdm.InitiatorID, lastQueryCmd, code)
	// if err != nil {
	// 	return err
	// }
	currentTime := time.Now().Format(time.RFC3339)
	newUISubIns, _ := RetriveSP(insCmd)
	subInsStmt, err := Db.Prepare(newUISubIns)
	if err != nil {
		return fmt.Errorf("Cannot prepare Campus drive Subscription insert due to %v %v", newUISubIns, err.Error())
	}

	_, err = subInsStmt.Exec(cdm.InitiatorID, cdm.ReceiverID, cdm.CampusDriveID, cdm.CampusDriveRequested, currentTime, cdm.RequestedNftID, currentTime, currentTime)
	if err != nil {
		return fmt.Errorf("Cannot Insert Campus drive Subscription  due to %v %v", newUISubIns, err.Error())
	}

	return nil
}

// SendInvitationToReceiver ...
func (cdm *CampusDriveDataModel) SendInvitationToReceiver(userType string) error {
	var updateCmd string
	switch userType {
	case Corporate:
		updateCmd = CrpInviteCmd
		break
	case University:
		updateCmd = UnvInviteCmd
		break
	default:
		return fmt.Errorf("Invalid user type %s", userType)
	}
	currentTime := time.Now().Format(time.RFC3339)
	newUISubIns, _ := RetriveSP(updateCmd)
	subInsStmt, err := Db.Prepare(newUISubIns)
	if err != nil {
		return fmt.Errorf("Cannot prepare Campus drive Invitation Update due to %v %v", newUISubIns, err.Error())
	}

	_, err = subInsStmt.Exec(true, currentTime, cdm.RequestedNftID, currentTime, cdm.InitiatorID, cdm.CampusDriveID)
	if err != nil {
		return fmt.Errorf("Cannot Update Campus drive Invitation   due to %v %v", newUISubIns, err.Error())
	}
	return nil
}

// Respond ...
func (cdm *CampusDriveDataModel) Respond(userType string) error {
	var updateCmd string
	switch userType {
	case Corporate:
		updateCmd = UnvRespondCmd
		break
	case University:
		updateCmd = CrpRespondCmd
		break
	default:
		return fmt.Errorf("Invalid user type %s", userType)
	}
	currentTime := time.Now().Format(time.RFC3339)
	newUISubIns, _ := RetriveSP(updateCmd)
	subInsStmt, err := Db.Prepare(newUISubIns)
	if err != nil {
		return fmt.Errorf("Cannot prepare Campus drive Invitation Response due to %v %v", newUISubIns, err.Error())
	}

	_, err = subInsStmt.Exec(cdm.Accepted, currentTime, cdm.AccOrRejectNftID, cdm.ReasonToReject, currentTime, cdm.ReceiverID, cdm.CampusDriveID)
	if err != nil {
		return fmt.Errorf("Cannot Respond Campus drive Invitation due to %v %v", newUISubIns, err.Error())
	}
	return nil
}

// GetByID ...
func (cdm *CampusDriveDataModel) GetByID(userType string) error {
	var updateCmd string
	switch userType {
	case Corporate:
		updateCmd = CrpGetByIDCmd
		break
	case University:
		updateCmd = UnvGetByIDCmd
		break
	default:
		return fmt.Errorf("Invalid user type %s", userType)
	}
	newUISubIns, _ := RetriveSP(updateCmd)
	err := Db.QueryRow(newUISubIns, cdm.CampusDriveID, cdm.InitiatorID, cdm.InitiatorID).Scan(&cdm.InitiatorID, &cdm.ReceiverID, &cdm.CampusDriveID, &cdm.CampusDriveRequested, &cdm.RequestedDate, &cdm.RequestedNftID, &cdm.Accepted, &cdm.AcceptedOrRejectedDate, &cdm.AccOrRejectNftID, &cdm.ReasonToReject, &cdm.CreationDate, &cdm.LastUpdatedDate)
	if err != nil && err != sql.ErrNoRows {
		return fmt.Errorf("Cannot Scan rows of Campus drive Invitation due to %v %v", newUISubIns, err.Error())
	}
	return nil
}

// GetIRByID ...
func (cdm *CampusDriveDataModel) GetIRByID(userType string, ID string, isResponse bool) (string, string, error) {
	var updateCmd string
	switch userType {
	case Corporate:
		if isResponse {
			updateCmd = UnvGetIRByIDCmd
		} else {
			updateCmd = CrpGetIRByIDCmd
		}
		break
	case University:
		if isResponse {
			updateCmd = CrpGetIRByIDCmd
		} else {
			updateCmd = UnvGetIRByIDCmd
		}
		break
	default:
		return "", "", fmt.Errorf("Invalid user type %s", userType)
	}
	newUISubIns, _ := RetriveSP(updateCmd)
	var i, r string
	err := Db.QueryRow(newUISubIns, cdm.CampusDriveID, ID, ID).Scan(&i, &r)
	if err != nil && err != sql.ErrNoRows {
		return "", "", fmt.Errorf("Cannot Scan rows of Campus drive Invitation due to %v %v", newUISubIns, err.Error())
	}
	return i, r, nil
}

// GetContentByNftID ...
func GetContentByNftID(nftID string, ID string) (string, error) {
	newUISubIns, _ := RetriveSP("NFT_GET_BY_ID")
	var content string
	err := Db.QueryRow(newUISubIns, nftID, ID).Scan(&content)
	if err != nil {
		return "", err
	}
	return content, nil
}
