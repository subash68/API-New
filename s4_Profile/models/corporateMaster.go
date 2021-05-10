package models

import (
	"time"
)

// CorporateMasterDB ...
type CorporateMasterDB struct {
	StakeholderID                       string    `json:"stakeholderID,omitempty" `
	CorporateName                       string    `json:"corporateName" binding:"required" `
	CIN                                 string    `json:"CIN" binding:"required" `
	CorporateHQAddressLine1             string    `json:"corporateHQAddressLine1" binding:"required" `
	CorporateHQAddressLine2             string    `json:"corporateHQAddressLine2,omitempty"`
	CorporateHQAddressLine3             string    `json:"corporateHQAddressLine3,omitempty"`
	CorporateHQAddressCountry           string    `json:"corporateHQAddressCountry" binding:"required" `
	CorporateHQAddressState             string    `json:"corporateHQAddressState" binding:"required" `
	CorporateHQAddressCity              string    `json:"corporateHQAddressCity" binding:"required" `
	CorporateHQAddressDistrict          string    `json:"corporateHQAddressDistrict,omitempty"`
	CorporateHQAddressZipCode           string    `json:"corporateHQAddressZipCode" binding:"required" `
	CorporateHQAddressPhone             string    `json:"corporateHQAddressPhone" binding:"required,min=13,max=13" `
	CorporateHQAddressEmail             string    `json:"corporateHQAddressEmail,omitempty"`
	CorporateLocalBranchAddressLine1    string    `json:"corporateLocalBranchAddressLine1,omitempty" `
	CorporateLocalBranchAddressLine2    string    `json:"corporateLocalBranchAddressLine2,omitempty"`
	CorporateLocalBranchAddressLine3    string    `json:"corporateLocalBranchAddressLine3,omitempty" `
	CorporateLocalBranchAddressCountry  string    `json:"corporateLocalBranchAddressCountry,omitempty" `
	CorporateLocalBranchAddressState    string    `json:"corporateLocalBranchAddressState,omitempty" `
	CorporateLocalBranchAddressCity     string    `json:"corporateLocalBranchAddressCity,omitempty" `
	CorporateLocalBranchAddressDistrict string    `json:"corporateLocalBranchAddressDistrict,omitempty"`
	CorporateLocalBranchAddressZipCode  string    `json:"corporateLocalBranchAddressZipCode,omitempty"  `
	CorporateLocalBranchAddressPhone    string    `json:"corporateLocalBranchAddressPhone,omitempty" `
	CorporateLocalBranchAddressEmail    string    `json:"corporateLocalBranchAddressEmail,omitempty" `
	PrimaryContactFirstName             string    `json:"primaryContactFirstName" binding:"required" `
	PrimaryContactMiddleName            string    `json:"primaryContactMiddleName,omitempty"`
	PrimaryContactLastName              string    `json:"primaryContactLastName" binding:"required" `
	PrimaryContactDesignation           string    `json:"primaryContactDesignation" binding:"required" `
	PrimaryContactPhone                 string    `json:"primaryContactPhone,omitempty" binding:"required,min=13,max=13" `
	PrimaryContactEmail                 string    `json:"primaryContactEmail" binding:"required,email" `
	SecondaryContactFirstName           string    `json:"secondaryContactFirstName,omitempty" `
	SecondaryContactMiddleName          string    `json:"secondaryContactMiddleName,omitempty"`
	SecondaryContactLastName            string    `json:"secondaryContactLastName,omitempty"`
	SecondaryContactDesignation         string    `json:"secondaryContactDesignation,omitempty" `
	SecondaryContactPhone               string    `json:"secondaryContactPhone,omitempty" `
	SecondaryContactEmail               string    `json:"secondaryContactEmail,omitempty" `
	CorporateType                       string    `json:"corporateType" binding:"required" `
	CorporateCategory                   string    `json:"corporateCategory" binding:"required" `
	CorporateIndustry                   string    `json:"corporateIndustry" binding:"required" `
	CompanyProfile                      string    `json:"companyProfile,omitempty"`
	Attachment                          []byte    `json:"attachment,omitempty"`
	YearOfEstablishment                 int64     `json:"yearOfEstablishment" binding:"required" `
	DateOfJoining                       time.Time `json:"dateOfJoining,omitempty" `
	AccountStatus                       string    `json:"accountStatus,omitempty" `
	PrimaryPhoneVerified                bool      `json:"primaryPhoneVerified"`
	PrimaryEmailVerified                bool      `json:"primaryEmailVerified"`
	ProfilePicture                      []byte    `form:"-" json:"profilePicture"`
	AccountExpiryDate                   time.Time `form:"-" json:"accountExpiryDate"`
}

// CorporateByIDResp ....
type CorporateByIDResp struct {
	StakeholderID                       string                   `json:"stakeholderID" `
	CorporateName                       string                   `json:"corporateName"`
	CIN                                 string                   `json:"CIN" `
	CorporateHQAddressLine1             string                   `json:"corporateHQAddressLine1,omitempty" `
	CorporateHQAddressLine2             string                   `json:"corporateHQAddressLine2,omitempty"`
	CorporateHQAddressLine3             string                   `json:"corporateHQAddressLine3,omitempty"`
	CorporateHQAddressCountry           string                   `json:"corporateHQAddressCountry" binding:"required" `
	CorporateHQAddressState             string                   `json:"corporateHQAddressState" binding:"required" `
	CorporateHQAddressCity              string                   `json:"corporateHQAddressCity" binding:"required" `
	CorporateHQAddressDistrict          string                   `json:"corporateHQAddressDistrict,omitempty"`
	CorporateHQAddressZipCode           string                   `json:"corporateHQAddressZipCode" binding:"required" `
	CorporateLocalBranchAddressLine1    string                   `json:"corporateLocalBranchAddressLine1,omitempty" `
	CorporateLocalBranchAddressLine2    string                   `json:"corporateLocalBranchAddressLine2,omitempty"`
	CorporateLocalBranchAddressLine3    string                   `json:"corporateLocalBranchAddressLine3,omitempty" `
	CorporateLocalBranchAddressCountry  string                   `json:"corporateLocalBranchAddressCountry,omitempty" `
	CorporateLocalBranchAddressState    string                   `json:"corporateLocalBranchAddressState,omitempty" `
	CorporateLocalBranchAddressCity     string                   `json:"corporateLocalBranchAddressCity,omitempty" `
	CorporateLocalBranchAddressDistrict string                   `json:"corporateLocalBranchAddressDistrict,omitempty"`
	CorporateLocalBranchAddressZipCode  string                   `json:"corporateLocalBranchAddressZipCode,omitempty"  `
	CorporateType                       string                   `json:"corporateType" binding:"required" `
	CorporateCategory                   string                   `json:"corporateCategory" binding:"required" `
	CorporateIndustry                   string                   `json:"corporateIndustry,,omitempty" binding:"required" `
	CompanyProfile                      string                   `json:"companyProfile"`
	YearOfEstablishment                 int64                    `json:"yearOfEstablishment" binding:"required" `
	DateOfJoining                       time.Time                `json:"dateOfJoining,omitempty" `
	Jobs                                string                   `json:"jobs,omitempty" form:"jobs"`
	JobsAvailable                       []map[string]interface{} `json:"jobsAvailable"`
	Subscriptions                       []SubscriptionReq        `json:"subscriptions"`
}
