// Package models ...
package models

import "time"

// UniversityMasterDb ...
type UniversityMasterDb struct {
	StakeholderID                        string    `form:"stakeholderID" json:"stakeholderID"`
	UniversityID                         string    `form:"universityID" json:"universityID"`
	UniversityName                       string    `form:"universityName" json:"universityName"`
	UniversityCollegeID                  string    `form:"universityCollegeID,omitempty" json:"universityCollegeID"`
	UniversityCollageName                string    `form:"universityCollegeName" json:"universityCollegeName"`
	UniversityHQAddressLine1             string    `form:"universityHQAddressLine1,omitempty" json:"universityHQAddressLine1"`
	UniversityHQAddressLine2             string    `form:"universityHQAddressLine2,omitempty" json:"universityHQAddressLine2"`
	UniversityHQAddressLine3             string    `form:"universityHQAddressLine3,omitempty" json:"universityHQAddressLine3"`
	UniversityHQAddressCountry           string    `form:"universityHQAddressCountry,omitempty" json:"universityHQAddressCountry"`
	UniversityHQAddressState             string    `form:"universityHQAddressState,omitempty" json:"universityHQAddressState"`
	UniversityHQAddressCity              string    `form:"universityHQAddressCity,omitempty" json:"universityHQAddressCity"`
	UniversityHQAddressDistrict          string    `form:"universityHQAddressDistrict,omitempty" json:"universityHQAddressDistrict"`
	UniversityHQAddressZipcode           string    `form:"universityHQAddressZipcode,omitempty" json:"universityHQAddressZipcode"`
	UniversityHQAddressPhone             string    `form:"universityHQAddressPhone,omitempty" json:"universityHQAddressPhone"`
	UniversityHQAddressemail             string    `form:"universityHQAddressemail,omitempty" json:"universityHQAddressemail"`
	UniversityLocalBranchAddressLine1    string    `form:"universityLocalBranchAddressLine1,omitempty" json:"universityLocalBranchAddressLine1"`
	UniversityLocalBranchAddressLine2    string    `form:"universityLocalBranchAddressLine2,omitempty" json:"universityLocalBranchAddressLine2"`
	UniversityLocalBranchAddressLine3    string    `form:"universityLocalBranchAddressLine3,omitempty" json:"universityLocalBranchAddressLine3"`
	UniversityLocalBranchAddressCountry  string    `form:"universityLocalBranchAddressCountry,omitempty" json:"universityLocalBranchAddressCountry"`
	UniversityLocalBranchAddressState    string    `form:"universityLocalBranchAddressState,omitempty" json:"universityLocalBranchAddressState"`
	UniversityLocalBranchAddressCity     string    `form:"universityLocalBranchAddressCity,omitempty" json:"universityLocalBranchAddressCity"`
	UniversityLocalBranchAddressDistrict string    `form:"universityLocalBranchAddressDistrict,omitempty" json:"universityLocalBranchAddressDistrict"`
	UniversityLocalBranchAddressZipcode  string    `form:"universityLocalBranchAddressZipcode,omitempty" json:"universityLocalBranchAddressZipcode"`
	UniversityLocalBranchAddressPhone    string    `form:"universityLocalBranchAddressPhone,omitempty" json:"universityLocalBranchAddressPhone"`
	UniversityLocalBranchAddressemail    string    `form:"universityLocalBranchAddressemail,omitempty" json:"universityLocalBranchAddressemail"`
	PrimaryContactFirstName              string    `form:"primaryContactFirstName,omitempty" json:"primaryContactFirstName"`
	PrimaryContactMiddleName             string    `form:"primaryContactMiddleName,omitempty" json:"primaryContactMiddleName"`
	PrimaryContactLastName               string    `form:"primaryContactLastName,omitempty" json:"primaryContactLastName"`
	PrimaryContactDesignation            string    `form:"primaryContactDesignation,omitempty" json:"primaryContactDesignation"`
	PrimaryContactPhone                  string    `form:"primaryContactPhone" binding:"required" json:"primaryContactPhone"`
	PrimaryContactEmail                  string    `form:"primaryContactEmail" binding:"required,email" json:"primaryContactEmail"`
	SecondaryContactFirstName            string    `form:"secondaryContactFirstName,omitempty" json:"secondaryContactFirstName"`
	SecondaryContactMiddleName           string    `form:"secondaryContactMiddleName,omitempty" json:"secondaryContactMiddleName"`
	SecondaryContactLastName             string    `form:"secondaryContactLastName,omitempty" json:"secondaryContactLastName"`
	SecondaryContactDesignation          string    `form:"secondaryContactDesignation,omitempty" json:"secondaryContactDesignation"`
	SecondaryContactPhone                string    `form:"secondaryContactPhone,omitempty" json:"secondaryContactPhone"`
	SecondaryContactEmail                string    `form:"secondaryContactEmail,omitempty" json:"secondaryContactEmail"`
	UniversitySector                     string    `form:"universitySector" binding:"required" json:"universitySector"`
	YearOfEstablishment                  int64     `form:"yearOfEstablishment" binding:"required" json:"yearOfEstablishment"`
	UniversityProfile                    string    `form:"universityProfile,omitempty" json:"universityProfile"`
	Attachment                           []byte    `form:"attachment,omitempty" json:"attachment"`
	AttachmentName                       string    `json:"attachmentName,omitempty"`
	DateOfJoining                        time.Time `form:"dateOfJoining,omitempty" json:"dateOfJoining"`
	AccountStatus                        string    `form:"accountStatus,omitempty" json:"accountStatus"`
	PrimaryPhoneVerified                 bool      `form:"primaryPhoneVerified" json:"primaryPhoneVerified"`
	PrimaryEmailVerified                 bool      `form:"primaryEmailVerified" json:"primaryEmailVerified"`
	ProfilePicture                       []byte    `form:"-" json:"profilePicture" json:"profilePicture"`
	AccountExpiryDate                    time.Time `form:"-" json:"accountExpiryDate" json:"accountExpiryDate"`
	PublishedFlag                        bool      `json:"publishedFlag"`
}

// UniversityGetByIDModel ...
type UniversityGetByIDModel struct {
	StakeholderID           string               `json:"stakeholderID"`
	UniversityName          string               `json:"universityName"`
	UniversityCollegeID     string               `json:"universityCollegeID"`
	UniversityHQAddressCity string               `json:"universityHQAddressCity,omitempty"`
	YearOfEstablishment     int64                `json:"yearOfEstablishment"`
	UniversityProfile       string               `json:"universityProfile"`
	ProgramsOffered         string               `json:"programsOffered"`
	Ranking                 string               `json:"ranking"`
	Accredations            string               `json:"accredations"`
	StudentStrengthNullable NullString           `form:"-" json:"-"`
	StudentDbAvailable      bool                 `form:"-" json:"studentDbAvailable"`
	StudentDbPublishID      string               `form:"-" json:"studentDbPublishID,omitempty"`
	UnvInsightsAvailable    bool                 `form:"-" json:"universityInsight"`
	Subscriptions           []SubscriptionReq    `json:"subscriptions"`
	PublishedData           []UnvPublishReqModel `json:"publishedData"`
}

// SubscriptionReq ...
type SubscriptionReq struct {
	Subscriber         string    `form:"-" json:"subscriber,omitempty"`
	Publisher          string    `form:"-" json:"publisher,omitempty"`
	PublisherType      string    `form:"-" json:"publisherType,omitempty"`
	SubscriptionType   string    `form:"-" json:"subscriptionType,omitempty"`
	DateOfSubscription time.Time `form:"-" json:"dateOfSubscription"`
	PublishID          string    `form:"publishId" json:"publishId"`
	TokensUsed         float64   `form:"tokensUsed" json:"tokensUsed,omitempty"`
	TransactionID      string    `form:"transactionID" json:"transactionID,omitempty"`
	CorporateName      string    `form:"-" json:"corporateName,omitempty"`
	GeneralNote        string    `form:"-" json:"generalNote"`
	SubscriptionID     string    `form:"-" json:"subscriptionID,omitempty"`
	CampusDriveID      string    `form:"-" json:"campusDriveID,omitempty"`
	CampusDriveStatus  string    `form:"-" json:"campusDriveStatus,omitempty"`
	NftID              string    `form:"-" json:"nftID,omitempty"`
	SearchCriteria     string    `form:"-" json:"searchCriteria,omitempty"`
}

// UnvPublishReqModel ...
type UnvPublishReqModel struct {
	PublishID                string            `form:"-" json:"publishID"`
	DateOfPublish            time.Time         `form:"-" json:"dateOfPublish"`
	ProgramsPublished        bool              `form:"programsPublished" json:"programsPublished"`
	BranchesPublished        bool              `form:"branchesPublished" json:"branchesPublished"`
	StudentStrengthPublished bool              `form:"studentStrengthPublished" json:"studentStrengthPublished"`
	AcredPublished           bool              `form:"acredPublished" json:"acredPublished"`
	COEsPublished            bool              `form:"coesPublished" json:"coesPublished"`
	RankingPublished         bool              `form:"rankingPublished" json:"rankingPublished"`
	OtherPublished           bool              `form:"otherPublished" json:"otherPublished"`
	ProfilePublished         bool              `form:"profilePublished" json:"profilePublished"`
	InfoPublished            bool              `form:"infoPublished" json:"infoPublished"`
	GeneralNote              string            `form:"-" json:"generalNote"`
	IsSubscribed             bool              `form:"isSubscribed" json:"isSubscribed"`
	PublishedData            string            `form:"-" json:"-"`
	Info                     map[string]string `form:"info" json:"info"`
}

// UnvOtherInformationModel ...
type UnvOtherInformationModel struct {
	StakeholderID   string    `form:"-" json:"-"`
	Title           string    `form:"title" json:"title" binding:"required"`
	Information     string    `form:"information" json:"information" binding:"required"`
	Attachment      string    `form:"-" json:"attachment,omitEmpty"`
	AttachmentName  string    `form:"-" json:"attachmentName,omitEmpty"`
	PublishID       string    `form:"-" json:"publishID"`
	CreationDate    time.Time `form:"-" json:"creationDate,omitEmpty"`
	LastUpdatedDate time.Time `form:"-" json:"lastUpdatedDate,omitEmpty"`
}

// UnvVrfData ...
type UnvVrfData struct {
	ID          string
	Email       string
	currentTime string
	VrfType     string
}
