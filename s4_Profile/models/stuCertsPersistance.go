package models

import (
	"database/sql"
	"fmt"
	"time"
)

// InsertCert ....
func (sc *StudentCertsModel) InsertCert() error {
	// Preparing Database insert
	slInsertCmd, _ := RetriveSP("STU_CERTS_INS")

	stmt, err := Db.Prepare(slInsertCmd)
	if err != nil {
		return fmt.Errorf("Cannot prepare -- %v  -- insert due to %v", slInsertCmd, err.Error())
	}
	currentTime := time.Now()
	_, err = stmt.Exec(sc.StakeholderID, sc.Name, sc.IssuingAuthority, sc.StartDate, sc.EndDate, sc.Attachment, sc.AttachmentName, true, currentTime, currentTime)
	if err != nil {
		return fmt.Errorf("Failed to insert in database -- %v -- insert due to %v", slInsertCmd, err.Error())
	}
	return nil
}

// GetAllCerts ....
func (sc *StudentAllCertsModel) GetAllCerts() error {
	// Preparing Database insert
	slInsertCmd, _ := RetriveSP("STU_CERTS_GETALL")

	slRows, err := Db.Query(slInsertCmd, sc.StakeholderID)
	if err != nil && err != sql.ErrNoRows {
		return fmt.Errorf("Cannot get the Rows %v", err.Error())
	} else if err == sql.ErrNoRows {
		return nil
	}
	defer slRows.Close()
	for slRows.Next() {
		var newSl StudentCertsModel
		err = slRows.Scan(&newSl.ID, &newSl.Name, &newSl.IssuingAuthority, &newSl.StartDate, &newSl.EndDate, &newSl.Attachment, &newSl.AttachmentName, &newSl.EnabledFlag, &newSl.CreationDate, &newSl.LastUpdatedDate)
		if err != nil {
			return fmt.Errorf("Cannot read the Rows %v", err.Error())
		}
		if newSl.EndDate == "0001-01-01T00:00:00Z" {
			newSl.EndDate = ""
		}
		sc.Certifications = append(sc.Certifications, newSl)
	}

	return nil
}

// UpdateCert ....
func (sc *StudentCertsModel) UpdateCert() error {
	// Preparing Database insert
	slInsertCmd, _ := RetriveSP("STU_CERTS_UPD")

	stmt, err := Db.Prepare(slInsertCmd)
	if err != nil {
		return fmt.Errorf("Cannot prepare -- %v -- insert due to %v", slInsertCmd, err.Error())
	}

	_, err = stmt.Exec(sc.Name, sc.IssuingAuthority, sc.StartDate, sc.EndDate, sc.Attachment, sc.AttachmentName, time.Now(), sc.ID, sc.StakeholderID)
	if err != nil {
		return fmt.Errorf("Failed to update in database -- %v  -- insert due to %v", slInsertCmd, err.Error())
	}
	return nil
}

// DeleteCert ....
func (sc *StudentCertsModel) DeleteCert() error {
	// Preparing Database insert
	slInsertCmd, _ := RetriveSP("STU_CERTS_DLT")

	stmt, err := Db.Prepare(slInsertCmd)
	if err != nil {
		return fmt.Errorf("Cannot prepare -- %v -- insert due to %v", slInsertCmd, err.Error())
	}
	_, err = stmt.Exec(sc.ID, sc.StakeholderID)
	if err != nil {
		return fmt.Errorf("Failed to update in database -- %v  -- insert due to %v", slInsertCmd, err.Error())
	}
	return nil
}
