// Package models ...
package models

import "time"

// StudentMasterDb ...
type StudentMasterDb struct {
	StakeholderID              string    `form:"stakeholderID" json:"stakeholderID,omitempty"`
	FirstName                  string    `form:"firstName" json:"firstName" binding:"required"`
	MiddleName                 string    `form:"middleName" json:"middleName"`
	LastName                   string    `form:"lastName" json:"lastName" binding:"required"`
	PersonalEmail              string    `form:"personalEmail" json:"personalEmail"`
	CollegeEmail               string    `form:"collegeEmail" json:"collegeEmail" binding:"required,email"`
	PhoneNumber                string    `form:"phoneNuber" json:"phoneNuber" binding:"required"`
	AlternatePhoneNumber       string    `form:"alternatePhoneNuber" json:"alternatePhoneNuber"`
	CollegeID                  string    `form:"collegeID" json:"collegeID" binding:"required"`
	Gender                     string    `form:"gender" json:"gender" binding:"required"`
	DateOfBirth                string    `form:"dateOfBirth" json:"dateOfBirth" binding:"required"`
	AadharNumber               string    `form:"aadharNumber" json:"aadharNumber" binding:"required"`
	PermanentAddressLine1      string    `form:"permanentAddressLine1" json:"permanentAddressLine1" binding:"required"`
	PermanentAddressLine2      string    `form:"permanentAddressLine2" json:"permanentAddressLine2" binding:"required"`
	PermanentAddressLine3      string    `form:"permanentAddressLine3" json:"permanentAddressLine3"`
	PermanentAddressCountry    string    `form:"permanentAddressCountry" json:"permanentAddressCountry" binding:"required"`
	PermanentAddressState      string    `form:"permanentAddressState" json:"permanentAddressState" binding:"required"`
	PermanentAddressCity       string    `form:"permanentAddressCity" json:"permanentAddressCity" binding:"required"`
	PermanentAddressDistrict   string    `form:"permanentAddressDistrict" json:"permanentAddressDistrict" binding:"required"`
	PermanentAddressZipcode    string    `form:"permanentAddressZipcode" json:"permanentAddressZipcode" binding:"required"`
	PermanentAddressPhone      string    `form:"permanentAddressPhone" json:"permanentAddressPhone" binding:"required"`
	PresentAddressLine1        string    `form:"presentAddressLine1" json:"presentAddressLine1" binding:"required"`
	PresentAddressLine2        string    `form:"presentAddressLine2" json:"presentAddressLine2" binding:"required"`
	PresentAddressLine3        string    `form:"presentAddressLine3" json:"presentAddressLine3"`
	PresentAddressCountry      string    `form:"presentAddressCountry" json:"presentAddressCountry" binding:"required"`
	PresentAddressState        string    `form:"presentAddressState" json:"presentAddressState" binding:"required"`
	PresentAddressCity         string    `form:"presentAddressCity" json:"presentAddressCity" binding:"required"`
	PresentAddressDistrict     string    `form:"presentAddressDistrict" json:"presentAddressDistrict" binding:"required"`
	PresentAddressZipcode      string    `form:"presentAddressZipcode" json:"presentAddressZipcode" binding:"required"`
	PresentAddressPhone        string    `form:"presentAddressPhone" json:"presentAddressPhone" binding:"required"`
	FathersGuardianFullName    string    `form:"fathersGuardianFullName" json:"fathersGuardianFullName" binding:"required"`
	FathersGuardianOccupation  string    `form:"fathersGuardianOccupation" json:"fathersGuardianOccupation" binding:"required"`
	FathersGuardianCompany     string    `form:"fathersGuardianCompany" json:"fathersGuardianCompany" binding:"required"`
	FathersGuardianPhoneNumber string    `form:"fathersGuardianPhoneNumber" json:"fathersGuardianPhoneNumber" binding:"required"`
	FathersGuardianEmailID     string    `form:"fathersGuardianEmailID" json:"fathersGuardianEmailID" binding:"required,email"`
	MothersGuardianFullName    string    `form:"mothersGuardianFullName" json:"mothersGuardianFullName"`
	MothersGuardianOccupation  string    `form:"mothersGuardianOccupation" json:"mothersGuardianOccupation"`
	MothersGuardianCompany     string    `form:"mothersGuardianCompany" json:"mothersGuardianCompany"`
	MothersGuardianDesignation string    `form:"mothersGuardianDesignation" json:"mothersGuardianDesignation"`
	MothersGuardianPhoneNumber string    `form:"mothersGuardianPhoneNumber" json:"mothersGuardianPhoneNumber" `
	MothersGuardianEmailID     string    `form:"mothersGuardianEmailID" json:"mothersGuardianEmailID" `
	AccountStatus              string    `form:"accountStatus" json:"accountStatus"`
	Password                   string    `form:"password" json:"password" binding:"required,min=8,max=15" binding:"required"`
	PrimaryPhoneVerified       bool      `form:"primaryPhoneVerified" json:"primaryPhoneVerified"`
	PrimaryEmailVerified       bool      `form:"primaryEmailVerified" json:"primaryEmailVerified"`
	DateOfJoining              time.Time `json:"dateOfJoining,omitempty"`
}
