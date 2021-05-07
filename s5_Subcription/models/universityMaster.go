// Package models ...
package models

import "time"

// UniversityMasterDb ...
type UniversityMasterDb struct {
	StakeholderID                        string    `form:"stakeholderID"`
	UniversityName                       string    `form:"universityName"`
	UniversityCollegeID                  string    `form:"universityCollegeID,omitempty"`
	UniversityHQAddressLine1             string    `form:"universityHQAddressLine1,omitempty"`
	UniversityHQAddressLine2             string    `form:"universityHQAddressLine2,omitempty"`
	UniversityHQAddressLine3             string    `form:"universityHQAddressLine3,omitempty"`
	UniversityHQAddressCountry           string    `form:"universityHQAddressCountry,omitempty"`
	UniversityHQAddressState             string    `form:"universityHQAddressState,omitempty"`
	UniversityHQAddressCity              string    `form:"universityHQAddressCity,omitempty"`
	UniversityHQAddressDistrict          string    `form:"universityHQAddressDistrict,omitempty"`
	UniversityHQAddressZipcode           string    `form:"universityHQAddressZipcode,omitempty"`
	UniversityHQAddressPhone             string    `form:"universityHQAddressPhone,omitempty"`
	UniversityHQAddressemail             string    `form:"universityHQAddressemail,omitempty"`
	UniversityLocalBranchAddressLine1    string    `form:"universityLocalBranchAddressLine1,omitempty"`
	UniversityLocalBranchAddressLine2    string    `form:"universityLocalBranchAddressLine2,omitempty"`
	UniversityLocalBranchAddressLine3    string    `form:"universityLocalBranchAddressLine3,omitempty"`
	UniversityLocalBranchAddressCountry  string    `form:"universityLocalBranchAddressCountry,omitempty"`
	UniversityLocalBranchAddressState    string    `form:"universityLocalBranchAddressState,omitempty"`
	UniversityLocalBranchAddressCity     string    `form:"universityLocalBranchAddressCity,omitempty"`
	UniversityLocalBranchAddressDistrict string    `form:"universityLocalBranchAddressDistrict,omitempty"`
	UniversityLocalBranchAddressZipcode  string    `form:"universityLocalBranchAddressZipcode,omitempty"`
	UniversityLocalBranchAddressPhone    string    `form:"universityLocalBranchAddressPhone,omitempty"`
	UniversityLocalBranchAddressemail    string    `form:"universityLocalBranchAddressemail,omitempty"`
	PrimaryContactFirstName              string    `form:"primaryContactFirstName,omitempty"`
	PrimaryContactMiddleName             string    `form:"primaryContactMiddleName,omitempty"`
	PrimaryContactLastName               string    `form:"primaryContactLastName,omitempty"`
	PrimaryContactDesignation            string    `form:"primaryContactDesignation,omitempty"`
	PrimaryContactPhone                  string    `form:"primaryContactPhone" binding:"required"`
	PrimaryContactEmail                  string    `form:"primaryContactEmail" binding:"required,email"`
	SecondaryContactFirstName            string    `form:"secondaryContactFirstName,omitempty"`
	SecondaryContactMiddleName           string    `form:"secondaryContactMiddleName,omitempty"`
	SecondaryContactLastName             string    `form:"secondaryContactLastName,omitempty"`
	SecondaryContactDesignation          string    `form:"secondaryContactDesignation,omitempty"`
	SecondaryContactPhone                string    `form:"secondaryContactPhone,omitempty"`
	SecondaryContactEmail                string    `form:"secondaryContactEmail,omitempty"`
	UniversitySector                     string    `form:"universitySector" binding:"required"`
	YearOfEstablishment                  int64     `form:"yearOfEstablishment" binding:"required"`
	UniversityProfile                    string    `form:"universityProfile,omitempty"`
	Attachment                           []byte    `form:"attachment,omitempty"`
	DateOfJoining                        time.Time `form:"dateOfJoining,omitempty"`
	AccountStatus                        string    `form:"accountStatus,omitempty"`
	PrimaryPhoneVerified                 bool      `form:"primaryPhoneVerified"`
	PrimaryEmailVerified                 bool      `form:"primaryEmailVerified"`
}

// UnvInsightsModel ...
type UnvInsightsModel struct {
	SubscriptionID            string    `form:"-" json:"subscriptionID"`
	SubscriberStakeholderID   string    `form:"universityID" json:"universityID" binding:"required"`
	SubscribedStakeholderID   string    `form:"-" json:"subscribedStakeholderID`
	AverageCGPA               float64   `form:"-" json:"averageCGPA"`
	AveragePercentage         float64   `form:"-" json:"averagePercentage"`
	HighestCGPA               float64   `form:"-" json:"highestCGPA"`
	HighestPercentage         float64   `form:"-" json:"highestPercentage"`
	HighestPackageReceived    string    `form:"-" json:"highestPackage"`
	AveragePackageReceived    string    `form:"-" json:"averagePackage"`
	UniversityConversionRatio float64   `form:"-" json:"universityConvertionRatio"`
	TentativeMonthOfPassing   string    `form:"-" json:"tentativeMonthOfPassing"`
	Top5Recruiters            []string  `form:"-" json:"top5Recruiters"`
	Top5Skills                []string  `form:"-" json:"top5Skills"`
	SubscribedDate            time.Time `form:"-" json:"subscribedTime" time_format="2006-12-01T21:23:34.409Z"`
	CreationDate              time.Time `form:"-" json:"creationDate" time_format="2006-12-01T21:23:34.409Z"`
	LastUpdatedDate           time.Time `form:"-" json:"lastUpdatedTime" time_format="2006-12-01T21:23:34.409Z"`
}

// UnvInsightSubsReqModel ...
type UnvInsightSubsReqModel struct {
	SubscriberStakeholderID string  `form:"universityID" json:"universityID" binding:"required"`
	PaidTokensUsed          float64 `form:"paidTokensUsed" json:"paidTokensUsed"`
	BonusTokensUsed         float64 `form:"bonusTokensUsed" json:"bonusTokensUsed"`
	TransactionID           string  `form:"transactionID" json:"transactionID,omitempty"`
}

// StuInfoFromUnvDatabaseModel ...
type StuInfoFromUnvDatabaseModel struct {
	StakeholderID string  `form:"studentID" json:"studentID"`
	Name          string  `form:"studentName" json:"studentName"`
	CollegeID     string  `form:"collegeID" json:"collegeID"`
	ProgramName   string  `form:"programName" json:"programName"`
	ProgramID     string  `form:"programID" json:"programID"`
	BranchName    string  `form:"branchName" json:"branchName"`
	BranchID      string  `form:"branchID" json:"branchID"`
	AvgCgpa       float64 `form:"avgCGPA" json:"avgCGPA"`
	AvgPercentage float64 `form:"avgPercentage" json:"avgPercentage"`
}

// UnvStuDataModel ...
type UnvStuDataModel struct {
	SubscriptionID          string                        `form:"-" json:"subscriptionID"`
	SubscriberStakeholderID string                        `form:"universityID" json:"universityID" binding:"required"`
	SubscribedStakeholderID string                        `form:"-" json:"subscribedStakeholderID`
	SubscriptionValidityTag bool                          `form:"-" json:"subscriptionValidityTag"`
	StudentDataExists       bool                          `form:"-" json:"studentDataExists"`
	StudentsData            []StuInfoFromUnvDatabaseModel `form:"studentsData" json:"studentsData"`
	SubscribedDate          time.Time                     `form:"-" json:"subscribedTime" time_format="2006-12-01T21:23:34.409Z"`
	CreationDate            time.Time                     `form:"-" json:"creationDate" time_format="2006-12-01T21:23:34.409Z"`
	LastUpdatedDate         time.Time                     `form:"-" json:"lastUpdatedTime" time_format="2006-12-01T21:23:34.409Z"`
}

// UnvStuDataQueryDataModel ...
type UnvStuDataQueryDataModel struct {
	SubscriberStakeholderID string   `form:"universityID" json:"universityID" binding:"required"`
	SubscriptionID          string   `form:"subscriptionID" json:"subscriptionID"  binding:"required"`
	HiringCriteriaID        string   `form:"hiringCriteriaID" json:"hiringCriteriaID"`
	ProgramID               string   `form:"programID" json:"programID"`
	BranchID                string   `form:"branchID" json:"branchID"`
	MonthOfHiring           string   `form:"mothOfHiring" json:"mothOfHiring"`
	Skills                  []string `form:"skills" json:"skills"`
}

// UnvCampusPlacementInviteModel ...
type UnvCampusPlacementInviteModel struct {
	CorpID                 string    `form:"-" json:"initiatorStakeholderID"`
	CorporateName          string    `form:"-" json:"corporateName"`
	UniversityID           string    `form:"universityID" json:"universityID"`
	UniversityName         string    `form:"universityName" json:"universityName"`
	CampusDriveID          string    `form:"-" json:"campusDriveID"`
	CampusDriverRequested  bool      `form:"campusDriveRequested" json:"campusDriveRequested"`
	RequestedDate          time.Time `form:"-" json:"requestedDate" time_format="2006-12-01T21:23:34.409Z"`
	RequestedNftID         string    `form:"-" json:"requestedNotificationID"`
	Accepted               bool      `form:"-" json:"accepted"`
	AcceptedOrRejectedDate time.Time `form:"-" json:"acceptedOrRejectedDate" time_format="2006-12-01T21:23:34.409Z"`
	AccOrRejectNftID       string    `form:"-" json:"acceptOrRejectNftID"`
	ReasonToReject         string    `form:"reasonToReject" json:"reasonToReject"`
}

// UnvSubCDModel ...
type UnvSubCDModel struct {
	CorporateID   string `form:"corporateID" json:"corporateID" binding:"required"`
	CampusDriveID string `form:"-" json:"campusDriveID" binding:"required"`
	EmailFrom     string `form:"emailFrom" json:"emailFrom" binding:"required,email"`
	EmailTo       string `form:"emailTo" json:"emailTo" binding:"required,email"`
	EmailSubject  string `form:"emailSubject" json:"emailSubject" binding:"required"`
	EmailBody     string `form:"emailBody" json:"emailBody" binding:"required"`
}
