// Package models ...
package models

import "time"

// StudentMasterDb ...
type StudentMasterDb struct {
	StakeholderID            string    `form:"-" json:"stakeholderID,omitempty"`
	FirstName                string    `form:"firstName" json:"firstName" binding:"required"`
	MiddleName               string    `form:"middleName" json:"middleName"`
	LastName                 string    `form:"lastName" json:"lastName" binding:"required"`
	PersonalEmail            string    `form:"personalEmail" json:"personalEmail" binding:"required,email"`
	PhoneNumber              string    `form:"phoneNumber" json:"phoneNumber" binding:"required,min=13,max=13"`
	AlternatePhoneNumber     string    `form:"alternatePhoneNumber" json:"alternatePhoneNumber"`
	Gender                   string    `form:"gender" json:"gender" binding:"required"`
	DateOfBirth              time.Time `form:"dateOfBirth" json:"dateOfBirth" binding:"required" time_format="2006-12-01T21:23:34.409Z"`
	AadharNumber             string    `form:"aadharNumber" json:"aadharNumber" binding:"required"`
	PermanentAddressLine1    string    `form:"permanentAddressLine1" json:"permanentAddressLine1" binding:"required"`
	PermanentAddressLine2    string    `form:"permanentAddressLine2" json:"permanentAddressLine2" binding:"required"`
	PermanentAddressLine3    string    `form:"permanentAddressLine3" json:"permanentAddressLine3"`
	PermanentAddressCountry  string    `form:"permanentAddressCountry" json:"permanentAddressCountry" binding:"required"`
	PermanentAddressState    string    `form:"permanentAddressState" json:"permanentAddressState" binding:"required"`
	PermanentAddressCity     string    `form:"permanentAddressCity" json:"permanentAddressCity" binding:"required"`
	PermanentAddressDistrict string    `form:"permanentAddressDistrict" json:"permanentAddressDistrict" binding:"required"`
	PermanentAddressZipcode  string    `form:"permanentAddressZipcode" json:"permanentAddressZipcode" binding:"required"`
	PermanentAddressPhone    string    `form:"permanentAddressPhone" json:"permanentAddressPhone" binding:"required"`
	PresentAddressLine1      string    `form:"presentAddressLine1" json:"presentAddressLine1" binding:"required"`
	PresentAddressLine2      string    `form:"presentAddressLine2" json:"presentAddressLine2" binding:"required"`
	PresentAddressLine3      string    `form:"presentAddressLine3" json:"presentAddressLine3"`
	PresentAddressCountry    string    `form:"presentAddressCountry" json:"presentAddressCountry" binding:"required"`
	PresentAddressState      string    `form:"presentAddressState" json:"presentAddressState" binding:"required"`
	PresentAddressCity       string    `form:"presentAddressCity" json:"presentAddressCity" binding:"required"`
	PresentAddressDistrict   string    `form:"presentAddressDistrict" json:"presentAddressDistrict" binding:"required"`
	PresentAddressZipcode    string    `form:"presentAddressZipcode" json:"presentAddressZipcode" binding:"required"`
	PresentAddressPhone      string    `form:"presentAddressPhone" json:"presentAddressPhone" binding:"required"`
	UniversityName           string    `form:"universityName" json:"universityName"`
	UniversityID             string    `form:"universityID" json:"universityID"`
	ProgramName              string    `form:"programName" json:"programName"`
	ProgramID                string    `form:"programID" json:"programID"`
	BranchName               string    `form:"branchName" json:"branchName"`
	BranchID                 string    `form:"branchID" json:"branchID"`
	CollegeID                string    `form:"collegeID" json:"collegeID" `
	CollegeEmailID           string    `form:"collegeEmailID" json:"collegeEmailID"`
	Password                 string    `form:"password" json:"password" binding:"required,min=8,max=15" binding:"required"`
	UniversityApprovedFlag   bool      `form:"-" json:"universityApprovedFlag`
	CreationDate             time.Time `form:"-" json:"creationDate"`
	LastUpdatedDate          time.Time `form:"-" json:"lastUpdatedDate"`
	AccountStatus            string    `form:"accountStatus" json:"accountStatus"`
	PrimaryPhoneVerified     bool      `form:"primaryPhoneVerified" json:"primaryPhoneVerified"`
	PrimaryEmailVerified     bool      `form:"primaryEmailVerified" json:"primaryEmailVerified"`
	ProfilePicture           []byte    `form:"-" json:"profilePicture"`
	Attachment               []byte    `form:"-" json:"Attachment"`
}
