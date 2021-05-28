package models

var (
	queryList []string = []string{
		"tenthBoards",
		"twelfthBoards",
		"accountStatus",
		"branchCatalog",
		"corporateCategory",
		"corporateIndustry",
		"corporateType",
		"jobType",
		"languageProficiency",
		"modeOfTokenIssue",
		"notificationType",
		"paymentMode",
		"programCatalog",
		"programType",
		"skillProficiency",
		"skills",
		"sortBy",
		"stakeholderType",
	}
)

// Lut10BoardsModel ...
type Lut10BoardsModel struct {
	BoardName       string `json:"boardName"`
	CertificateName string `json:"certificateName"`
	BoardID         string `json:"boardID"`
}

// LutAccountStatusModel ...
type LutAccountStatusModel struct {
	AccountStatus     string `json:"accountStatus"`
	AccountStatusCode string `json:"accountStatusCode"`
}

// LutBranchCatalogModel ...
type LutBranchCatalogModel struct {
	BranchID    string `json:"branchID"`
	BranchName  string `json:"branchName"`
	Duration    string `json:"duration"`
	ProgramID   string `json:"programID"`
	ProgramType string `json:"programType"`
}

// LutCorporateCategoryModel ...
type LutCorporateCategoryModel struct {
	Name       string `json:"categoryName"`
	Code       string `json:"categoryCode"`
	OneLtrCode string `json:"oneLtrCode"`
}

// LutCorporateIndustryModel ...
type LutCorporateIndustryModel struct {
	Name string `json:"industryName"`
	Code string `json:"industryCode"`
}

// LutCorporateTypeModel ...
type LutCorporateTypeModel struct {
	Name       string `json:"corporateTypeName"`
	Code       string `json:"corporateTypeCode"`
	OneLtrCode string `json:"oneLtrCode"`
}

// LutJobTypeModel ...
type LutJobTypeModel struct {
	Name string `json:"jobType"`
	Code string `json:"jobTypeCode"`
}

// LutLangProficiencyModel ...
type LutLangProficiencyModel struct {
	Name string `json:"proficiency"`
	Code string `json:"proficiencyID"`
}

// LutModeOfTokenIssueModel ...
type LutModeOfTokenIssueModel struct {
	Name string `json:"modeOfTokenIssue"`
	Code string `json:"modeOfTokenIssueID"`
}

// LutNotificationTypeModel ...
type LutNotificationTypeModel struct {
	Name string `json:"notificationType"`
	Code string `json:"notificationTypeID"`
}

// LutPaymentModeModel ...
type LutPaymentModeModel struct {
	Name string `json:"paymentMode"`
	Code string `json:"paymentModeID"`
}

// LutProgramCatalogModel ...
type LutProgramCatalogModel struct {
	Name string `json:"programName"`
	Code string `json:"programCode"`
	Type string `json:"programType"`
}

// LutProgramTypeModel ...
type LutProgramTypeModel struct {
	Name string `json:"programType"`
	Code string `json:"programTypeID"`
}

// LutSkillProficiencyModel ...
type LutSkillProficiencyModel struct {
	Name string `json:"proficiency"`
	Code string `json:"proficiencyID"`
}

// LutSkillsModel ...
type LutSkillsModel struct {
	Name     string `json:"skillName"`
	Code     string `json:"skillCode"`
	Disabled bool   `json:"disabled"`
}

// LutSortByModel ...
type LutSortByModel struct {
	Name string `json:"sortBy"`
	Code string `json:"sortByCode"`
}

// LutStakeholdersModel ...
type LutStakeholdersModel struct {
	Name       string `json:"stakeholdertypeName"`
	Code       string `json:"stakeholderTypeCode"`
	OneLtrCode string `json:"oneLtrCode"`
}

// AllLutData ...
type AllLutData struct {
	TenthBoards         []Lut10BoardsModel          `json:"tenthBoards,omitempty"`
	TwelfthBoards       []Lut10BoardsModel          `json:"twelfthBoards,omitempty"`
	AccountStatus       []LutAccountStatusModel     `json:"accountStatus,omitempty"`
	BranchCatalog       []LutBranchCatalogModel     `json:"branchCatalog,omitempty"`
	CoporateCategory    []LutCorporateCategoryModel `json:"corporateCategory,omitempty"`
	CoporateIndustry    []LutCorporateIndustryModel `json:"coporateIndustry,omitempty"`
	CorporateType       []LutCorporateTypeModel     `json:"corporateType,omitempty"`
	JobType             []LutJobTypeModel           `json:"jobType,omitempty"`
	LanguageProficiency []LutLangProficiencyModel   `json:"languageProficiency,omitempty"`
	ModeOfTokenIssue    []LutModeOfTokenIssueModel  `json:"modeOfTokenIssue,omitempty"`
	NotificationType    []LutNotificationTypeModel  `json:"notificationType,omitempty"`
	PaymentMode         []LutPaymentModeModel       `json:"paymentMode,omitempty"`
	ProgramCatalog      []LutProgramCatalogModel    `json:"programCatalog,omitempty"`
	ProgramType         []LutProgramTypeModel       `json:"programType,omitempty"`
	SkillProficiency    []LutSkillProficiencyModel  `json:"skillProficiency,omitempty"`
	Skills              []LutSkillsModel            `json:"skills,omitempty"`
	SortBy              []LutSortByModel            `json:"sortBy,omitempty"`
	StakeholderType     []LutStakeholdersModel      `json:"stakeholderType,omitempty"`
}
