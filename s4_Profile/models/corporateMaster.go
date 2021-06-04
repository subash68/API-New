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
	AttachmentName                      string    `json:"attachmentName,omitempty"`
	YearOfEstablishment                 int64     `json:"yearOfEstablishment" binding:"required" `
	DateOfJoining                       time.Time `json:"dateOfJoining,omitempty" `
	AccountStatus                       string    `json:"accountStatus,omitempty" `
	PrimaryPhoneVerified                bool      `json:"primaryPhoneVerified"`
	PrimaryEmailVerified                bool      `json:"primaryEmailVerified"`
	ProfilePicture                      []byte    `form:"-" json:"profilePicture"`
	AccountExpiryDate                   time.Time `form:"-" json:"accountExpiryDate"`
	PublishedFlag                       bool      `json:"publishedFlag"`
}

// CorpPushedDataReq ...
type CorpPushedDataReq struct {
	PublishID               string            `form:"-" json:"publishID"`
	DateOfPublish           string            `form:"-" json:"dateOfPublish"`
	HiringCriteriaPublished bool              `form:"hiringCriteriaPublished" json:"hiringCriteriaPublished"`
	JobsPublished           bool              `form:"jobsPublished" json:"jobsPublished"`
	ProfilePublished        bool              `form:"profilePublished" json:"profilePublished"`
	OtherPublished          bool              `form:"otherPublished" json:"otherPublished"`
	GeneralNote             string            `form:"-" json:"generalNote"`
	IsSubscribed            bool              `form:"isSubscribed" json:"isSubscribed"`
	PublishedData           string            `form:"-" json:"-"`
	Info                    map[string]string `form:"info" json:"info"`
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
	PublishedData                       []CorpPushedDataReq      `json:"publishedData"`
}

// HiringCriteriaDB ...
type HiringCriteriaDB struct {
	HiringCriteriaID            string    `form:"-" json:"hiringCriteriaID"`
	HiringCriteriaName          string    `form:"hiringCriteriaName" json:"hiringCriteriaName"`
	MinimumCutoffPercentage10th float64   `form:"minimumCutoffPercentage10th" json:"minimumCutoffPercentage10th"`
	MinimumCutoffPercentage12th float64   `form:"minimumCutoffPercentage12th" json:"minimumCutoffPercentage12th"`
	MinimumCutoffCGPAGrad       float64   `form:"minimumCutoffCGPAGrad" json:"minimumCutoffCGPAGrad"`
	MinimumCutoffPercentageGrad float64   `form:"minimumCutoffPercentageGrad" json:"minimumCutoffPercentageGrad"`
	EduGapsSchoolAllowed        bool      `form:"eduGapsSchoolAllowed" json:"eduGapsSchoolAllowed"`
	EduGaps11N12Allowed         bool      `form:"eduGaps11N12Allowed" json:"eduGaps11N12Allowed"`
	EduGaps12NGradAllowed       bool      `form:"eduGaps12NGradAllowed" json:"eduGaps12NGradAllowed"`
	EduGapsGradAllowed          bool      `form:"eduGapsGradAllowed" json:"eduGapsGradAllowed"`
	EduGapsGradNPGAllowed       bool      `form:"eduGapsGradNPGAllowed" json:"eduGapsGradNPGAllowed"`
	EduGapsSchool               int       `form:"eduGapsSchool" json:"eduGapsSchool"`
	EduGaps11N12                int       `form:"eduGaps11N12" json:"eduGaps11N12"`
	EduGaps12NGrad              int       `form:"eduGaps12NGrad" json:"eduGaps12NGrad"`
	EduGapsGrad                 int       `form:"eduGapsGrad" json:"eduGapsGrad"`
	EduGapsGradNPG              int       `form:"eduGapsGradNPG" json:"eduGapsGradNPG"`
	AllowActiveBacklogs         bool      `form:"allowActiveBacklogs" json:"allowActiveBacklogs"`
	NumberOfAllowedBacklogs     int       `form:"numberOfAllowedBacklogs" json:"numberOfAllowedBacklogs"`
	YearOfPassing               int       `form:"yearOfPassing" json:"yearOfPassing"`
	Remarks                     string    `form:"remarks" json:"remarks"`
	CreationDate                time.Time `form:"-" json:"creationDate" time_format="2006-12-01T21:23:34.409Z"`
	ProgramsInString            string    `json:"hcProgramsInString"`
}

// HcProgramsModel ...
type HcProgramsModel struct {
	ProgramName string `json:"programName"`
	ProgramID   string `json:"programID"`
	BranchName  string `json:"branchName"`
	BranchID    string `json:"branchID"`
}

// JobHcMappingDB ...
type JobHcMappingDB struct {
	StakeholderID   string    `form:"-" json:"stakeholderID,omitempty"`
	JobID           string    `form:"jobID" json:"jobID"`
	JobName         string    `form:"jobName" json:"jobName"`
	HcID            string    `form:"hiringCriteriaID" json:"hiringCriteriaID"`
	HcName          string    `form:"hiringCriteriaName" json:"hiringCriteriaName"`
	JobType         string    `form:"jobType" json:"jobType"`
	NoOfPositions   int       `form:"noOfPositions" json:"noOfPositions"`
	Location        string    `form:"location" json:"location"`
	SalaryMaxRange  string    `form:"salaryMaxRange" json:"salaryMaxRange"`
	SalaryMinRange  string    `form:"salaryMinRange" json:"salaryMinRange"`
	MonthOfHiring   time.Time `form:"monthOfHiring" json:"monthOfHiring" time_format="2006-12-01T21:23:34.409Z"`
	Remarks         string    `form:"remarks" json:"remarks"`
	Attachment      []byte    `form:"attachment" json:"attachment"`
	AttachmentName  string    `form:"attachmentName" json:"attachmentName"`
	Status          string    `form:"status" json:"status"`
	CreationDate    time.Time `form:"-" json:"creationDate" time_format="2006-12-01T21:23:34.409Z"`
	LastUpdatedDate time.Time `form:"-" json:"lastUpdatedDate" time_format="2006-12-01T21:23:34.409Z"`
	PublishedFlag   bool      `form:"-" json:"publishedFlag"`
	PublishID       string    `form:"-" json:"publishID"`
	SkillsInString  string    `form:"-" json:"skillsInString"`
}

// JobSkillsMapping ...
type JobSkillsMapping struct {
	ID            int       `form:"id" json:"id"`
	JobID         string    `form:"jobID" json:"jobID"`
	JobName       string    `form:"jobName" json:"jobName"`
	StakeholderID string    `form:"-" json:"stakeholder,omitempty"`
	SkillID       string    `form:"skillID" json:"skillID"`
	Skill         string    `form:"skill" json:"skill"`
	CreationDate  time.Time `form:"_" json:"creationDate"`
}
