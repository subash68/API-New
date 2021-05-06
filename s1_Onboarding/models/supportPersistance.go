package models

import (
	"fmt"
	"time"
)

// SupportDataModel ...
type SupportDataModel struct {
	StakeholderID   string `form:"-" json:"-"`
	StakeholderRole string `form:"-" json:"-"`
	ContactEmail    string `form:"contactEmail" json:"contactEmail" binding:"required"`
	ContactPerson   string `form:"contactPerson" json:"contactPerson" binding:"required"`
	ContactPhone    string `form:"contactPhone" json:"contactPhone" binding:"required"`
	QueryOrIssue    string `form:"queryOrIssue" json:"queryOrIssue" binding:"required"`
	CreationDate    string `form:"-" json:"creationDate"`
	LastUpdatedDate string `form:"-" json:"lastUpdatedDat"`
}

// Insert ...
func (sd *SupportDataModel) Insert() error {
	sp, _ := RetriveSP("SUPPORT_INS")
	stmt, err := Db.Prepare(sp)
	if err != nil {
		return fmt.Errorf("Failed to Prepare Support query Due to : %v", err)
	}
	defer stmt.Close()
	currentTime := time.Now().Format(time.RFC3339)
	_, err = stmt.Exec(sd.StakeholderID, sd.StakeholderRole, sd.ContactEmail, sd.ContactPerson, sd.ContactPhone, sd.QueryOrIssue, currentTime, currentTime)

	if err != nil {
		return fmt.Errorf("Failed to Insert Support query Due to : %v", err)
	}
	return nil
}
