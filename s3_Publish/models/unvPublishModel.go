// Package models ...
package models

import "time"

// UnvProgramWiseBranchDBModel ...
type UnvProgramWiseBranchDBModel struct {
	ID                  int       `form:"-" json:"id"`
	StakeholderID       string    `form:"-" json:"stakeholder,omitempty"`
	ProgramID           string    `form:"programID" json:"programID" binding:"required"`
	ProgramName         string    `form:"programName" json:"programName" binding:"required"`
	BranchID            string    `form:"branchID" json:"branchID" binding:"required"`
	BranchName          string    `form:"branchName" json:"branchName" binding:"required"`
	StartDate           string    `form:"startDate" json:"startDate" binding:"required"`
	EndDate             string    `form:"endDate" json:"endDate"`
	EnablingFlag        string    `form:"enablingFlag" json:"enablingFlag" binding:"required"`
	NoOfPassingStudents int64     `form:"noOfPassingStudents" json:"noOfPassingStudents" binding:"required"`
	MonthYearOfPassing  string    `form:"monthYearOfPassing" json:"monthYearOfPassing" binding:"required"`
	CreationDate        time.Time `form:"-" json:"creationDate,omitempty" `
	LastUpdatedDate     time.Time `form:"-" json:"lastUpdatedDate,omitempty"`
}

// UnvYearWiseRanking ...
type UnvYearWiseRanking struct {
	ID               int       `form:"-" json:"id"`
	StakeholderID    string    `form:"-" json:"stakeholder,omitempty"`
	Rank             string    `form:"rank" json:"rank" binding:"required"`
	IssuingAuthority string    `form:"issuingAuthority" json:"issuingAuthority" binding:"required"`
	RankingFile      string    `form:"-" json:"rankingFile"`
	CreationDate     time.Time `form:"-" json:"creationDate,omitempty"`
	LastUpdatedDate  time.Time `form:"-" json:"lastUpdatedDate,omitempty"`
}

// UnvAccredationsDBModel ...
type UnvAccredationsDBModel struct {
	ID                     int       `form:"-" json:"id"`
	StakeholderID          string    `form:"-" json:"stakeholder,omitempty"`
	AccredationName        string    `form:"accredationName" binding:"required"`
	AccredationType        string    `form:"accredationType" binding:"required"`
	AccredationDescription string    `form:"accredationDescription" binding:"required"`
	IssuingAuthority       string    `form:"issuingAuthority" json:"issuingAuthority" binding:"required"`
	AccredationFile        string    `form:"accredationFile" json:"accredationfile"`
	StartDate              string    `form:"startDate" json:"startDate" binding:"required"`
	EndDate                string    `form:"endDate" json:"endDate" binding:"required"`
	EnablingFlag           string    `form:"enablingFlag" json:"enablingFlag" binding:"required"`
	CreationDate           time.Time `form:"-" json:"creationDate,omitempty"`
	LastUpdatedDate        time.Time `form:"-" json:"lastUpdatedDate,omitempty"`
}

// UnvTieupsDBModel ...
type UnvTieupsDBModel struct {
	ID                     int       `form:"-" json:"id"`
	StakeholderID          string    `form:"-" json:"stakeholder,omitempty"`
	TieupType              string    `form:"typeupType" json:"typeupType" binding:"required"`
	TieupName              string    `form:"tieupName" json:"tieupName" binding:"required"`
	TieupDescription       string    `form:"tieupDescription" json:"tieupDescription" binding:"required"`
	TieupWithName          string    `form:"tieupWithName" json:"tieupWithName"`
	TieupWithContact       string    `form:"tieupWithContact" json:"tieupWithContact"`
	TieupWithStakeholderID string    `form:"tieupWithStakeholderID" json:"tieupWithStakeholderID"`
	TieupFile              string    `form:"-" json:"tieupfile"`
	StartDate              string    `form:"startDate" json:"startDate" binding:"required"`
	EndDate                string    `form:"endDate" json:"endDate"`
	EnablingFlag           string    `form:"enablingFlag" json:"enablingFlag"`
	CreationDate           time.Time `form:"-" json:"creationDate,omitempty"`
	LastUpdatedDate        time.Time `form:"-" json:"lastUpdatedDate,omitempty"`
}

// UnvSpecialOfferingsDBModel ...
type UnvSpecialOfferingsDBModel struct {
	ID                            int       `form:"-" json:"id"`
	StakeholderID                 string    `form:"-" json:"stakeholder,omitempty"`
	SpecialOfferingType           string    `form:"specialOfferingType" json:"specialOfferingType" binding:"required"`
	SpecialOfferingName           string    `form:"specialOfferingName" json:"specialOfferingName" binding:"required"`
	SpecialOfferingDescription    string    `form:"specialOfferingDescription" json:"specialOfferingDescription" binding:"required"`
	InternallyManagedFlag         string    `form:"internallyManagedFlag" json:"internallyManagedFlag"`
	OutsourcedVendorName          string    `form:"outsourcedVendorName" json:"outsourcedVendorName"`
	OutsourcedVendorContact       string    `form:"outsourcedVendorContact" json:"outsourcedVendorContact"`
	OutsourcedVendorStakeholderID string    `form:"outsourcedVendorStakeholderID" json:"outsourcedVendorStakeholderID"`
	SpecialOffersCol              string    `form:"specialOffersCol" json:"specialOffersCol"`
	SpecialOfferingFile           string    `form:"specialOfferingFile" json:"specialOfferingFile"`
	StartDate                     string    `form:"startDate" json:"startDate"`
	EndDate                       string    `form:"endDate" json:"endDate"`
	EnablingFlag                  string    `form:"enablingFlag" json:"enablingFlag"`
	CreationDate                  time.Time `form:"-" json:"creationDate,omitempty"`
	LastUpdatedDate               time.Time `form:"-" json:"lastUpdatedDate,omitempty"`
}

// UnvCEOsDBModel ...
type UnvCEOsDBModel struct {
	ID                            int       `form:"-" json:"id"`
	StakeholderID                 string    `form:"-" json:"stakeholder,omitempty"`
	CoeID                         string    `form:"coeID" json:"coeID,omitempty" binding:"required"`
	CoeType                       string    `form:"coeType" json:"coeType" binding:"required"`
	CoeName                       string    `form:"coeName" json:"coeName" binding:"required"`
	CoeDescription                string    `form:"coeDescription" json:"coeDescription" binding:"required"`
	InternallyManagedFlag         string    `form:"internallyManagedFlag" json:"internallyManagedFlag" binding:"required"`
	OutsourcedVendorName          string    `form:"outsourcedVendorName" json:"outsourcedVendorName"`
	OutsourcedVendorContact       string    `form:"outsourcedVendorContact" json:"outsourcedVendorContact"`
	OutsourcedVendorStakeholderID string    `form:"outsourcedVendorStakeholderID" json:"outsourcedVendorStakeholderID"`
	CoeFile                       string    `form:"coeFile" json:"coeFile"`
	StartDate                     string    `form:"startDate" json:"startDate" binding:"required"`
	EndDate                       string    `form:"endDate" json:"endDate" binding:"required"`
	EnablingFlag                  string    `form:"enablingFlag" json:"enablingFlag" binding:"required"`
	CreationDate                  time.Time `form:"-" json:"creationDate,omitempty"`
	LastUpdatedDate               time.Time `form:"-" json:"lastUpdatedDate,omitempty"`
}

// UnvProgramsDBModel ...
type UnvProgramsDBModel struct {
	ID              int       `form:"id" json:"id"`
	StakeholderID   string    `form:"-" json:"stakeholder,omitempty"`
	ProgramID       string    `form:"programID" json:"programID" binding:"required"`
	ProgramType     string    `form:"programType" json:"programType" binding:"required"`
	ProgramName     string    `form:"programName" json:"programName" binding:"required"`
	StartDate       string    `form:"startDate" json:"startDate"`
	EndDate         string    `form:"endDate" json:"endDate"`
	EnablingFlag    string    `form:"enablingFlag" json:"enablingFlag"`
	CreationDate    time.Time `form:"-" json:"creationDate,omitempty"`
	LastUpdatedDate time.Time `form:"-" json:"lastUpdatedDate,omitempty"`
}

// UnvPublishDBModel ...
type UnvPublishDBModel struct {
	StakeholderID            string    `form:"-" json:"stakeholder,omitempty"`
	PublishID                string    `form:"-" json:"publishID"`
	UniversityName           string    `form:"universityName" json:"universityName"`
	DateOfPublish            time.Time `form:"-" json:"dateOfPublish"`
	ProgramsPublished        bool      `form:"programsPublished" json:"programsPublished"`
	BranchesPublished        bool      `form:"branchesPublished" json:"branchesPublished"`
	StudentStrengthPublished bool      `form:"studentStrengthPublished" json:"studentStrengthPublished"`
	AcredPublished           bool      `form:"acredPublished" json:"acredPublished"`
	COEsPublished            bool      `form:"coesPublished" json:"coesPublished"`
	RankingPublished         bool      `form:"rankingPublished" json:"rankingPublished"`
	OtherPublished           bool      `form:"otherPublished" json:"otherPublished"`
	ProfilePublished         bool      `form:"profilePublished" json:"profilePublished"`
	InfoPublished            bool      `form:"infoPublished" json:"infoPublished"`
	GeneralNote              string    `form:"-" json:"generalNote"`
	PublishedData            string    `form:"publishData" json:"publishData" binding:"required"`
	CreationDate             time.Time `form:"-" json:"creationDate,omitEmpty"`
	LastUpdatedDate          time.Time `form:"-" json:"lastUpdatedDate,omitempty"`
}

// UniversityProposal ...
type UniversityProposal struct {
	Profile          UniversityProposalMasterDb    `form:"profile" json:"profile,omitempty"`
	Programs         []UnvProgramsDBModel          `form:"programs,omitempty" json:"programs,omitempty" binding:"dive"`
	Branches         []UnvProgramWiseBranchDBModel `form:"branches,omitempty" json:"branches,omitempty" binding:"dive"`
	Accredations     []UnvAccredationsDBModel      `form:"accredations,omitempty" json:"accredations,omitempty" binding:"dive"`
	Rankings         []UnvYearWiseRanking          `form:"rankings,omitempty" json:"rankings,omitempty" binding:"dive"`
	SpecialOfferings []UnvSpecialOfferingsDBModel  `form:"specialOfferings,omitempty" json:"specialOfferings,omitempty" binding:"dive"`
	Tieups           []UnvTieupsDBModel            `form:"tieups,omitempty" json:"tieups,omitempty" binding:"dive"`
	Coes             []UnvCEOsDBModel              `form:"coes,omitempty" json:"coes,omitempty" binding:"dive"`
}

// UnvOtherInformationModel ...
type UnvOtherInformationModel struct {
	StakeholderID   string    `form:"-" json:"-"`
	Title           string    `form:"title" json:"title" binding:"required"`
	Information     string    `form:"information" json:"information" binding:"required"`
	Attachment      []byte    `form:"-" json:"attachment,omitEmpty"`
	ID              int       `form:"-" json:"id,omitEmpty"`
	PublishID       string    `form:"-" json:"publishID"`
	PublishedFlag   string    `form:"-" json:"publishedFlag"`
	CreationDate    time.Time `form:"-" json:"creationDate,omitEmpty"`
	LastUpdatedDate time.Time `form:"-" json:"lastUpdatedDate,omitEmpty"`
}

// UniversityMasterDb ...
type UniversityMasterDb struct {
	StakeholderID                        string    `json:"stakeholderID"`
	UniversityName                       string    `json:"universityName"`
	UniversityCollegeID                  string    `json:"universityCollegeID,omitempty"`
	UniversityHQAddressLine1             string    `json:"universityHQAddressLine1,omitempty"`
	UniversityHQAddressLine2             string    `json:"universityHQAddressLine2,omitempty"`
	UniversityHQAddressLine3             string    `json:"universityHQAddressLine3,omitempty"`
	UniversityHQAddressCountry           string    `json:"universityHQAddressCountry,omitempty"`
	UniversityHQAddressState             string    `json:"universityHQAddressState,omitempty"`
	UniversityHQAddressCity              string    `json:"universityHQAddressCity,omitempty"`
	UniversityHQAddressDistrict          string    `json:"universityHQAddressDistrict,omitempty"`
	UniversityHQAddressZipcode           string    `json:"universityHQAddressZipcode,omitempty"`
	UniversityHQAddressPhone             string    `json:"universityHQAddressPhone,omitempty"`
	UniversityHQAddressemail             string    `json:"universityHQAddressemail,omitempty"`
	UniversityLocalBranchAddressLine1    string    `json:"universityLocalBranchAddressLine1,omitempty"`
	UniversityLocalBranchAddressLine2    string    `json:"universityLocalBranchAddressLine2,omitempty"`
	UniversityLocalBranchAddressLine3    string    `json:"universityLocalBranchAddressLine3,omitempty"`
	UniversityLocalBranchAddressCountry  string    `json:"universityLocalBranchAddressCountry,omitempty"`
	UniversityLocalBranchAddressState    string    `json:"universityLocalBranchAddressState,omitempty"`
	UniversityLocalBranchAddressCity     string    `json:"universityLocalBranchAddressCity,omitempty"`
	UniversityLocalBranchAddressDistrict string    `json:"universityLocalBranchAddressDistrict,omitempty"`
	UniversityLocalBranchAddressZipcode  string    `json:"universityLocalBranchAddressZipcode,omitempty"`
	UniversityLocalBranchAddressPhone    string    `json:"universityLocalBranchAddressPhone,omitempty"`
	UniversityLocalBranchAddressemail    string    `json:"universityLocalBranchAddressemail,omitempty"`
	PrimaryContactFirstName              string    `json:"primaryContactFirstName,omitempty"`
	PrimaryContactMiddleName             string    `json:"primaryContactMiddleName,omitempty"`
	PrimaryContactLastName               string    `json:"primaryContactLastName,omitempty"`
	PrimaryContactDesignation            string    `json:"primaryContactDesignation,omitempty"`
	PrimaryContactPhone                  string    `json:"primaryContactPhone" binding:"required"`
	PrimaryContactEmail                  string    `json:"primaryContactEmail" binding:"required,email"`
	UniversitySector                     string    `json:"universitySector" binding:"required"`
	YearOfEstablishment                  int64     `json:"yearOfEstablishment" binding:"required"`
	UniversityProfile                    string    `json:"universityProfile,omitempty"`
	Attachment                           []byte    `json:"attachment,omitempty"`
	DateOfJoining                        time.Time `json:"dateOfJoining,omitempty"`
}

// UniversityProposalMasterDb ...
type UniversityProposalMasterDb struct {
	StakeholderID                        string    `json:"stakeholderID"`
	UniversityName                       string    `json:"universityName"`
	UniversityCollegeID                  string    `json:"universityCollegeID,omitempty"`
	UniversityHQAddressLine1             string    `json:"universityHQAddressLine1,omitempty"`
	UniversityHQAddressLine2             string    `json:"universityHQAddressLine2,omitempty"`
	UniversityHQAddressLine3             string    `json:"universityHQAddressLine3,omitempty"`
	UniversityHQAddressCountry           string    `json:"universityHQAddressCountry,omitempty"`
	UniversityHQAddressState             string    `json:"universityHQAddressState,omitempty"`
	UniversityHQAddressCity              string    `json:"universityHQAddressCity,omitempty"`
	UniversityHQAddressDistrict          string    `json:"universityHQAddressDistrict,omitempty"`
	UniversityHQAddressZipcode           string    `json:"universityHQAddressZipcode,omitempty"`
	UniversityHQAddressPhone             string    `json:"universityHQAddressPhone,omitempty"`
	UniversityHQAddressemail             string    `json:"universityHQAddressemail,omitempty"`
	UniversityLocalBranchAddressLine1    string    `json:"universityLocalBranchAddressLine1,omitempty"`
	UniversityLocalBranchAddressLine2    string    `json:"universityLocalBranchAddressLine2,omitempty"`
	UniversityLocalBranchAddressLine3    string    `json:"universityLocalBranchAddressLine3,omitempty"`
	UniversityLocalBranchAddressCountry  string    `json:"universityLocalBranchAddressCountry,omitempty"`
	UniversityLocalBranchAddressState    string    `json:"universityLocalBranchAddressState,omitempty"`
	UniversityLocalBranchAddressCity     string    `json:"universityLocalBranchAddressCity,omitempty"`
	UniversityLocalBranchAddressDistrict string    `json:"universityLocalBranchAddressDistrict,omitempty"`
	UniversityLocalBranchAddressZipcode  string    `json:"universityLocalBranchAddressZipcode,omitempty"`
	UniversityLocalBranchAddressPhone    string    `json:"universityLocalBranchAddressPhone,omitempty"`
	UniversityLocalBranchAddressemail    string    `json:"universityLocalBranchAddressemail,omitempty"`
	PrimaryContactFirstName              string    `json:"primaryContactFirstName,omitempty"`
	PrimaryContactMiddleName             string    `json:"primaryContactMiddleName,omitempty"`
	PrimaryContactLastName               string    `json:"primaryContactLastName,omitempty"`
	PrimaryContactDesignation            string    `json:"primaryContactDesignation,omitempty"`
	PrimaryContactPhone                  string    `json:"primaryContactPhone"`
	PrimaryContactEmail                  string    `json:"primaryContactEmail"`
	UniversitySector                     string    `json:"universitySector"`
	YearOfEstablishment                  int64     `json:"yearOfEstablishment"`
	UniversityProfile                    string    `json:"universityProfile,omitempty"`
	Attachment                           []byte    `json:"attachment,omitempty"`
	DateOfJoining                        time.Time `json:"dateOfJoining,omitempty"`
}
