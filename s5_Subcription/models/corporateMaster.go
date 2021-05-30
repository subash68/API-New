package models

import "time"

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
}

// CampusDriveDataModel ...
type CampusDriveDataModel struct {
	InitiatorID            string    `form:"-" json:"initiatorStakeholderID"`
	InitiatorName          string    `form:"-" json:"initiatorName"`
	ReceiverID             string    `form:"-" json:"receiverID"`
	ReceiverName           string    `form:"-" json:"receiverName"`
	CampusDriveID          string    `form:"-" json:"campusDriveID"`
	CampusDriveRequested   bool      `form:"-" json:"campusDriveRequested"`
	RequestedDate          time.Time `form:"-" json:"requestedDate" time_format="2006-12-01T21:23:34.409Z"`
	RequestedNftID         string    `form:"-" json:"requestedNotificationID"`
	Accepted               bool      `form:"-" json:"accepted"`
	AcceptedOrRejectedDate time.Time `form:"-" json:"acceptedOrRejectedDate" time_format="2006-12-01T21:23:34.409Z"`
	AccOrRejectNftID       string    `form:"-" json:"acceptOrRejectNftID"`
	ReasonToReject         string    `form:"-" json:"reasonToReject"`
	CreationDate           time.Time `form:"-" json:"creationDate" time_format="2006-12-01T21:23:34.409Z"`
	LastUpdatedDate        time.Time `form:"-" json:"lastUpdatedTime" time_format="2006-12-01T21:23:34.409Z"`
}

// CDSubscriptionReq ...
type CDSubscriptionReq struct {
	InitiatorID          string    `form:"-" json:"initiatorID"`
	InitiatorName        string    `form:"-" json:"initiatorName,omitempty"`
	InitiatorLocation    string    `json:"initiatorLocation,omitempty"`
	ReceiverID           string    `form:"-" json:"receiverID"`
	ReceiverName         string    `form:"-" json:"receiverName,omitempty"`
	ReceiverLocation     string    `json:"receiverLocation,omitempty"`
	CampusDriveID        string    `form:"-" json:"campusDriveID"`
	CampusDriveRequested bool      `form:"-" json:"campusDriveRequested"`
	RequestedDate        time.Time `form:"-" json:"requestedDate" time_format="2006-12-01T21:23:34.409Z"`
	CampusDriveStatus    string    `form:"-" json:"campusDriveStatus"`
	NftID                string    `form:"-" json:"nftID"`
	TokensUsed           float64   `form:"-" json:"tokensUsed,omitempty"`
	TransactionID        string    `form:"-" json:"transactionID,omitempty"`
}

//AllCdInvites ...
type AllCdInvites struct {
	CampusInviteSent     []CDSubscriptionReq `json:"campusInviteSent"`
	CampusInviteReceived []CDSubscriptionReq `json:"campusInviteReceived"`
}

// CampusDriveRespondDataModel ...
type CampusDriveRespondDataModel struct {
	CampusDriveID  string `form:"campusDriveID" json:"campusDriveID" binding:"required"`
	NftID          string `form:"nftID" json:"nftID" binding:"required"`
	Accepted       bool   `form:"accepted" json:"accepted"`
	ReasonToReject string `form:"reasonToReject" json:"reasonToReject"`
}

// UnvCDDataModel ...
type UnvCDDataModel struct {
	Name         string `json:"universityName,omitempty"`
	Location     string `json:"location,omitempty"`
	YearOfEst    int    `json:"yearOfEstablishment,omitempty"`
	Programs     string `json:"program,omitempty"`
	Ranking      string `json:"ranking,omitempty"`
	Accredations string `json:"accredations,omitempty"`
}

// CorpCDDataModel ...
type CorpCDDataModel struct {
	Name        string `json:"corporateName,omitempty"`
	Location    string `json:"location,omitempty"`
	Category    int    `json:"category,omitempty"`
	CIN         string `json:"cin,omitempty"`
	CorporateID string `json:"corporateID,omitempty"`
}

// CampusDriveInviteEmailModel ...
type CampusDriveInviteEmailModel struct {
	CampusDriveID     string          `form:"campusDriveID" json:"campusDriveID" binding:"required"`
	EmailFrom         string          `form:"emailFrom" json:"emailFrom" binding:"required,email"`
	EmailTo           string          `form:"emailTo" json:"emailTo" binding:"required,email"`
	EmailSubject      string          `form:"emailSubject" json:"emailSubject" binding:"required"`
	EmailBody         string          `form:"emailBody" json:"emailBody" binding:"required"`
	UniversityDetails UnvCDDataModel  `form:"-" json:"universityDetails,omitempty"`
	CorporateDetails  CorpCDDataModel `form:"-" json:"corporateDetails,omitempty"`
}

// CampusDriveDirectInviteEmailModel ...
type CampusDriveDirectInviteEmailModel struct {
	CampusDriveID  string `form:"campusDriveID" json:"campusDriveID" binding:"required"`
	EmailFrom      string `form:"emailFrom" json:"emailFrom" binding:"required,email"`
	EmailTo        string `form:"emailTo" json:"emailTo" binding:"required,email"`
	EmailSubject   string `form:"emailSubject" json:"emailSubject" binding:"required"`
	EmailBody      string `form:"emailBody" json:"emailBody" binding:"required"`
	StudentIDs     string `form:"studentIDs" json:"studentIDs"  binding:"required"`
	SubscriptionID string `form:"subscriptionID" json:"subscriptionID" binding:"required"`
}

// CampusDriveSubInitModel ...
type CampusDriveSubInitModel struct {
	ReceiverID      string  `form:"receiverID" json:"receiverID" binding:"required"`
	PaidTokensUsed  float64 `form:"paidTokensUsed" json:"paidTokensUsed"`
	BonusTokensUsed float64 `form:"bonusTokensUsed" json:"bonusTokensUsed"`
	TransactionID   string  `form:"transactionID" json:"transactionID,omitempty"`
}

// CrpHiringInsightsModel ...
type CrpHiringInsightsModel struct {
	SubscriptionID           string    `form:"-" json:"subscriptionID"`
	SubscriberStakeholderID  string    `form:"corporateID" json:"corporateID" binding:"required"`
	SubscribedStakeholderID  string    `form:"-" json:"subscribedStakeholderID"`
	SubscriptionValidityFlag bool      `form:"-" json:"subscriptionValidityFlag,omitempty"`
	Top5LocationsRecruited   []string  `form:"-" json:"top5LocationsRecruited"`
	Top5SkillsRecruited      []string  `form:"-" json:"top5SkillsRecruited"`
	Top5ProgramsRecruited    []string  `form:"-" json:"top5ProgramsRecruited"`
	TentativeMonthOfHiring   string    `form:"-" json:"tentativeMonthOfHiring"`
	AverageCutoffCGPA        float64   `form:"-" json:"averageCutoffCGPA"`
	AverageCutoffPercentage  float64   `form:"-" json:"averageCutoffPercentage"`
	AverageSalary            string    `form:"-" json:"averageSalary"`
	SubscribedDate           time.Time `form:"-" json:"subscribedTime" time_format="2006-12-01T21:23:34.409Z"`
	CreationDate             time.Time `form:"-" json:"creationDate" time_format="2006-12-01T21:23:34.409Z"`
	LastUpdatedDate          time.Time `form:"-" json:"lastUpdatedTime" time_format="2006-12-01T21:23:34.409Z"`
}

// CorpHiringInsightsSubsReqModel ...
type CorpHiringInsightsSubsReqModel struct {
	SubscriberStakeholderID string  `form:"corporateID" json:"corporateID" binding:"required"`
	PaidTokensUsed          float64 `form:"paidTokensUsed" json:"paidTokensUsed"`
	BonusTokensUsed         float64 `form:"bonusTokensUsed" json:"bonusTokensUsed"`
	TransactionID           string  `form:"transactionID" json:"transactionID,omitempty"`
}
