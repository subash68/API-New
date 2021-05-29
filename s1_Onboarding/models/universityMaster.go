// Pakcage models
package models

// UniversityMasterDb ...
type UniversityMasterDb struct {
	StakeholderID                        string `form:"stakeholderID"`
	UniversityName                       string `form:"universityName" binding:"required"`
	University                           string `form:"universityCollegeID" binding:"required"`
	UniversityHQAddressLine1             string `form:"universityHQAddressLine1" binding:"required"`
	UniversityHQAddressLine2             string `form:"universityHQAddressLine2" binding:"required"`
	UniversityHQAddressLine3             string `form:"universityHQAddressLine3"`
	UniversityHQAddressCountry           string `form:"universityHQAddressCountry" binding:"required"`
	UniversityHQAddressState             string `form:"universityHQAddressState" binding:"required"`
	UniversityHQAddressCity              string `form:"universityHQAddressCity" binding:"required"`
	UniversityHQAddressDistrict          string `form:"universityHQAddressDistrict" binding:"required"`
	UniversityHQAddressZipcode           string `form:"universityHQAddressZipcode" binding:"required"`
	UniversityHQAddressPhone             string `form:"universityHQAddressPhone" binding:"required"`
	UniversityHQAddressemail             string `form:"universityHQAddressemail" binding:"required,email"`
	UniversityLocalBranchAddressLine1    string `form:"universityLocalBranchAddressLine1,omitempty"`
	UniversityLocalBranchAddressLine2    string `form:"universityLocalBranchAddressLine2,omitempty"`
	UniversityLocalBranchAddressLine3    string `form:"universityLocalBranchAddressLine3,omitempty"`
	UniversityLocalBranchAddressCountry  string `form:"universityLocalBranchAddressCountry,omitempty"`
	UniversityLocalBranchAddressState    string `form:"universityLocalBranchAddressState,omitempty"`
	UniversityLocalBranchAddressCity     string `form:"universityLocalBranchAddressCity,omitempty"`
	UniversityLocalBranchAddressDistrict string `form:"universityLocalBranchAddressDistrict,omitempty"`
	UniversityLocalBranchAddressZipcode  string `form:"universityLocalBranchAddressZipcode,omitempty"`
	UniversityLocalBranchAddressPhone    string `form:"universityLocalBranchAddressPhone,omitempty"`
	UniversityLocalBranchAddressemail    string `form:"universityLocalBranchAddressemail,omitempty"`
	PrimaryContactFirstName              string `form:"primaryContactFirstName" binding:"required"`
	PrimaryContactMiddleName             string `form:"primaryContactMiddleName,omitempty"`
	PrimaryContactLastName               string `form:"primaryContactLastName" binding:"required"`
	PrimaryContactDesignation            string `form:"primaryContactDesignation" binding:"required"`
	PrimaryContactPhone                  string `form:"primaryContactPhone" binding:"required,min=13,max=13"`
	PrimaryContactEmail                  string `form:"primaryContactEmail" binding:"required,email"`
	SecondaryContactFirstName            string `form:"secondaryContactFirstName,omitempty"`
	SecondaryContactMiddleName           string `form:"secondaryContactMiddleName,omitempty"`
	SecondaryContactLastName             string `form:"secondaryContactLastName,omitempty"`
	SecondaryContactDesignation          string `form:"secondaryContactDesignation,omitempty"`
	SecondaryContactPhone                string `form:"secondaryContactPhone,omitempty"`
	SecondaryContactEmail                string `form:"secondaryContactEmail,omitempty"`
	UniversitySector                     string `form:"universitySector" binding:"required"`
	YearOfEstablishment                  int64  `form:"yearOfEstablishment" binding:"required"`
	UniversityProfile                    string `form:"universityProfile,omitempty"`
	Attachment                           []byte `form:"attachment,omitempty"`
	AttachmentName                       string `form:"attachmentName,omitempty"`
	AccountStatus                        string `form:"accountStatus,omitempty"`
	Password                             string `form:"password" binding:"required,min=8,max=15"`
	PrimaryPhoneVerified                 bool   `form:"primaryPhoneVerified,omitempty"`
	PrimaryEmailVerified                 bool   `form:"primaryEmailVerified,omitempty"`
	ProfilePicture                       []byte `form:"-" json:"profilePicture"`
}
