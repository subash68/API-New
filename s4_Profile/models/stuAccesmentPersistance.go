package models

import (
	"database/sql"
	"fmt"
	"time"
)

// InsertAssessment ....
func (sa *StudentAssessmentModel) InsertAssessment() error {
	// Preparing Database insert
	slInsertCmd, _ := RetriveSP("STU_ASSESSMENT_INS")

	stmt, err := Db.Prepare(slInsertCmd)
	if err != nil {
		return fmt.Errorf("Cannot prepare -- %v  -- insert due to %v", slInsertCmd, err.Error())
	}
	currentTime := time.Now()
	_, err = stmt.Exec(sa.StakeholderID, sa.Name, sa.Score, sa.IssuingAuthority, sa.AssessmentDate, sa.Description, sa.Attachment, true, currentTime, currentTime)
	if err != nil {
		return fmt.Errorf("Failed to insert in database -- %v -- insert due to %v", slInsertCmd, err.Error())
	}
	return nil
}

// GetAllAssessment ....
func (sa *StudentAllAssessmentModel) GetAllAssessment() error {
	// Preparing Database insert
	slInsertCmd, _ := RetriveSP("STU_ASSESSMENT_GETALL")

	slRows, err := Db.Query(slInsertCmd, sa.StakeholderID)
	if err != nil && err != sql.ErrNoRows {
		return fmt.Errorf("Cannot get the Rows %v", err.Error())
	} else if err == sql.ErrNoRows {
		return nil
	}
	defer slRows.Close()
	for slRows.Next() {
		var newSl StudentAssessmentModel
		err = slRows.Scan(&newSl.ID, &newSl.Name, &newSl.Score, &newSl.IssuingAuthority, &newSl.AssessmentDate, &newSl.Description, &newSl.Attachment, &newSl.EnabledFlag, &newSl.CreationDate, &newSl.LastUpdatedDate)
		if err != nil {
			return fmt.Errorf("Cannot read the Rows %v", err.Error())
		}

		sa.Assessments = append(sa.Assessments, newSl)
	}

	return nil
}

// UpdateAssessment ....
func (sa *StudentAssessmentModel) UpdateAssessment() error {
	// Preparing Database insert
	slInsertCmd, _ := RetriveSP("STU_ASSESSMENT_UPD")

	stmt, err := Db.Prepare(slInsertCmd)
	if err != nil {
		return fmt.Errorf("Cannot prepare -- %v -- insert due to %v", slInsertCmd, err.Error())
	}

	_, err = stmt.Exec(sa.Name, sa.Score, sa.IssuingAuthority, sa.AssessmentDate, sa.Description, sa.Attachment, time.Now(), sa.ID, sa.StakeholderID)
	if err != nil {
		return fmt.Errorf("Failed to update in database -- %v  -- insert due to %v", slInsertCmd, err.Error())
	}
	return nil
}

// DeleteAssessment ....
func (sa *StudentAssessmentModel) DeleteAssessment() error {
	// Preparing Database insert
	slInsertCmd, _ := RetriveSP("STU_ASSESSMENT_DLT")

	stmt, err := Db.Prepare(slInsertCmd)
	if err != nil {
		return fmt.Errorf("Cannot prepare -- %v -- insert due to %v", slInsertCmd, err.Error())
	}
	_, err = stmt.Exec(sa.ID, sa.StakeholderID)
	if err != nil {
		return fmt.Errorf("Failed to update in database -- %v  -- insert due to %v", slInsertCmd, err.Error())
	}
	return nil
}
