package models

// CorporateMasterDB ...
type CorporateMasterDB struct {
	StakeholderID                       string `form:"stakeholderID,omitempty" `
	CorporateName                       string `form:"corporateName" binding:"required" `
	CIN                                 string `form:"CIN" binding:"required" binding:"required"`
	CorporateHQAddressLine1             string `form:"corporateHQAddressLine1" binding:"required" `
	CorporateHQAddressLine2             string `form:"corporateHQAddressLine2,omitempty"`
	CorporateHQAddressLine3             string `form:"corporateHQAddressLine3,omitempty"`
	CorporateHQAddressCountry           string `form:"corporateHQAddressCountry" binding:"required" `
	CorporateHQAddressState             string `form:"corporateHQAddressState" binding:"required" `
	CorporateHQAddressCity              string `form:"corporateHQAddressCity" binding:"required" `
	CorporateHQAddressDistrict          string `form:"corporateHQAddressDistrict,omitempty"`
	CorporateHQAddressZipCode           string `form:"corporateHQAddressZipCode" binding:"required" `
	CorporateHQAddressPhone             string `form:"corporateHQAddressPhone" binding:"required,min=13,max=13" `
	CorporateHQAddressEmail             string `form:"corporateHQAddressEmail,omitempty"`
	CorporateLocalBranchAddressLine1    string `form:"corporateLocalBranchAddressLine1,omitempty" `
	CorporateLocalBranchAddressLine2    string `form:"corporateLocalBranchAddressLine2,omitempty"`
	CorporateLocalBranchAddressLine3    string `form:"corporateLocalBranchAddressLine3,omitempty" `
	CorporateLocalBranchAddressCountry  string `form:"corporateLocalBranchAddressCountry,omitempty" `
	CorporateLocalBranchAddressState    string `form:"corporateLocalBranchAddressState,omitempty" `
	CorporateLocalBranchAddressCity     string `form:"corporateLocalBranchAddressCity,omitempty" `
	CorporateLocalBranchAddressDistrict string `form:"corporateLocalBranchAddressDistrict,omitempty"`
	CorporateLocalBranchAddressZipCode  string `form:"corporateLocalBranchAddressZipCode,omitempty"  `
	CorporateLocalBranchAddressPhone    string `form:"corporateLocalBranchAddressPhone,omitempty" `
	CorporateLocalBranchAddressEmail    string `form:"corporateLocalBranchAddressEmail,omitempty" `
	PrimaryContactFirstName             string `form:"primaryContactFirstName" binding:"required" `
	PrimaryContactMiddleName            string `form:"primaryContactMiddleName,omitempty"`
	PrimaryContactLastName              string `form:"primaryContactLastName" binding:"required" `
	PrimaryContactDesignation           string `form:"primaryContactDesignation" binding:"required" `
	PrimaryContactPhone                 string `form:"primaryContactPhone,omitempty" binding:"required,min=13,max=13" `
	PrimaryContactEmail                 string `form:"primaryContactEmail" binding:"required,email" `
	SecondaryContactFirstName           string `form:"secondaryContactFirstName,omitempty" `
	SecondaryContactMiddleName          string `form:"secondaryContactMiddleName,omitempty"`
	SecondaryContactLastName            string `form:"secondaryContactLastName,omitempty"`
	SecondaryContactDesignation         string `form:"secondaryContactDesignation,omitempty" `
	SecondaryContactPhone               string `form:"secondaryContactPhone,omitempty" `
	SecondaryContactEmail               string `form:"secondaryContactEmail,omitempty" `
	CorporateType                       string `form:"corporateType" binding:"required" `
	CorporateCategory                   string `form:"corporateCategory" binding:"required" `
	CorporateIndustry                   string `form:"corporateIndustry" binding:"required" `
	CompanyProfile                      string `form:"companyProfile,omitempty"`
	Attachment                          []byte `form:"attachment,omitempty"`
	AttachmentName                      string `form:"attachmentName,omitempty"`
	YearOfEstablishment                 int64  `form:"yearOfEstablishment" binding:"required" `
	AccountStatus                       string `form:"accountStatus,omitempty" `
	Password                            string `form:"password" binding:"required,min=8,max=15"`
	PrimaryPhoneVerified                bool   `form:"primaryPhoneVerified,omitempty"`
	PrimaryEmailVerified                bool   `form:"primaryEmailVerified,omitempty"`
	ProfilePicture                      []byte `form:"-" json:"profilePicture"`
}
