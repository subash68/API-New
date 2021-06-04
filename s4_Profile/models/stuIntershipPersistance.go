package models

import (
	"database/sql"
	"fmt"
	"time"
)

// InsertInternship ....
func (si *StudentInternshipModel) InsertInternship() error {
	// Preparing Database insert
	slInsertCmd, _ := RetriveSP("STU_INTERSHIP_INS")

	stmt, err := Db.Prepare(slInsertCmd)
	if err != nil {
		return fmt.Errorf("Cannot prepare -- %v  -- insert due to %v", slInsertCmd, err.Error())
	}
	currentTime := time.Now()
	_, err = stmt.Exec(si.StakeholderID, si.Name, si.OrganizationName, si.FieldOfWork, si.OrganizationCity, si.StartDate, si.EndDate, si.Description, si.Attachment, si.AttachmentName, true, currentTime, currentTime)
	if err != nil {
		return fmt.Errorf("Failed to insert in database -- %v -- insert due to %v", slInsertCmd, err.Error())
	}
	return nil
}

// GetAllInternships ....
func (si *StudentAllInternshipModel) GetAllInternships() error {
	// Preparing Database insert
	slInsertCmd, _ := RetriveSP("STU_INTERSHIP_GETALL")

	slRows, err := Db.Query(slInsertCmd, si.StakeholderID)
	if err != nil && err != sql.ErrNoRows {
		return fmt.Errorf("Cannot get the Rows %v", err.Error())
	} else if err == sql.ErrNoRows {
		return nil
	}
	defer slRows.Close()
	for slRows.Next() {
		var newSl StudentInternshipModel
		err = slRows.Scan(&newSl.ID, &newSl.Name, &newSl.OrganizationName, &newSl.FieldOfWork, &newSl.OrganizationCity, &newSl.StartDate, &newSl.EndDate, &newSl.Description, &newSl.Attachment, &newSl.AttachmentName, &newSl.EnabledFlag, &newSl.CreationDate, &newSl.LastUpdatedDate, &newSl.SentforVerification, &newSl.DateSentforVerification, &newSl.Verified, &newSl.DateVerified, &newSl.SentbackforRevalidation, &newSl.DateSentBackForRevalidation, &newSl.ValidatorRemarks, &newSl.VerificationType, &newSl.VerifiedByStakeholderID, &newSl.VerifiedByEmailID)
		if err != nil {
			return fmt.Errorf("Cannot read the Rows %v", err.Error())
		}
		// if newSl.EndDate == "0001-01-01T00:00:00Z" {
		// 	newSl.EndDate = ""
		// }
		si.Internships = append(si.Internships, newSl)
	}

	return nil
}

// UpdateInternship ....
func (si *StudentInternshipModel) UpdateInternship() error {
	// Preparing Database insert
	slInsertCmd, _ := RetriveSP("STU_INTERSHIP_UPD")

	stmt, err := Db.Prepare(slInsertCmd)
	if err != nil {
		return fmt.Errorf("Cannot prepare -- %v -- insert due to %v", slInsertCmd, err.Error())
	}

	_, err = stmt.Exec(si.Name, si.OrganizationName, si.FieldOfWork, si.OrganizationCity, si.StartDate, si.EndDate, si.Description, si.Attachment, si.AttachmentName, time.Now(), si.ID, si.StakeholderID)
	if err != nil {
		return fmt.Errorf("Failed to update in database -- %v  -- insert due to %v", slInsertCmd, err.Error())
	}
	return nil
}

// DeleteInternship ....
func (si *StudentInternshipModel) DeleteInternship() error {
	// Preparing Database insert
	slInsertCmd, _ := RetriveSP("STU_INTERSHIP_DLT")

	stmt, err := Db.Prepare(slInsertCmd)
	if err != nil {
		return fmt.Errorf("Cannot prepare -- %v -- insert due to %v", slInsertCmd, err.Error())
	}
	_, err = stmt.Exec(si.ID, si.StakeholderID)
	if err != nil {
		return fmt.Errorf("Failed to update in database -- %v  -- insert due to %v", slInsertCmd, err.Error())
	}
	return nil
}
