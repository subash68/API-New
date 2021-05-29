// Package models ...
package models

import (
	"time"
)

// MyTime ...
type MyTime struct {
	time.Time
}

// UnmarshalJSON ....
func (m *MyTime) UnmarshalJSON(data []byte) error {
	// Ignore null, like in the main JSON package.
	if string(data) == "null" || string(data) == `""` {
		return nil
	}
	// Fractional seconds are handled implicitly by Parse.
	tt, err := time.Parse(`"`+time.RFC3339+`"`, string(data))
	*m = MyTime{tt}
	return err
}

// HiringCriteriaDB ...
type HiringCriteriaDB struct {
	StakeholderID               string            `form:"-" json:"-"`
	HiringCriteriaID            string            `form:"-" json:"hiringCriteriaID"`
	HiringCriteriaName          string            `form:"hiringCriteriaName" json:"hiringCriteriaName" binding:"required" validate:"required"`
	MinimumCutoffPercentage10th float64           `form:"minimumCutoffPercentage10th" json:"minimumCutoffPercentage10th"`
	MinimumCutoffPercentage12th float64           `form:"minimumCutoffPercentage12th" json:"minimumCutoffPercentage12th"`
	MinimumCutoffCGPAGrad       float64           `form:"minimumCutoffCGPAGrad" json:"minimumCutoffCGPAGrad"`
	MinimumCutoffPercentageGrad float64           `form:"minimumCutoffPercentageGrad" json:"minimumCutoffPercentageGrad"`
	EduGapsSchoolAllowed        bool              `form:"eduGapsSchoolAllowed" json:"eduGapsSchoolAllowed"`
	EduGaps11N12Allowed         bool              `form:"eduGaps11N12Allowed" json:"eduGaps11N12Allowed"`
	EduGaps12NGradAllowed       bool              `form:"eduGaps12NGradAllowed" json:"eduGaps12NGradAllowed"`
	EduGapsGradAllowed          bool              `form:"eduGapsGradAllowed" json:"eduGapsGradAllowed"`
	EduGapsGradNPGAllowed       bool              `form:"eduGapsGradNPGAllowed" json:"eduGapsGradNPGAllowed"`
	EduGapsSchool               int               `form:"eduGapsSchool" json:"eduGapsSchool"`
	EduGaps11N12                int               `form:"eduGaps11N12" json:"eduGaps11N12"`
	EduGaps12NGrad              int               `form:"eduGaps12NGrad" json:"eduGaps12NGrad"`
	EduGapsGrad                 int               `form:"eduGapsGrad" json:"eduGapsGrad"`
	EduGapsGradNPG              int               `form:"eduGapsGradNPG" json:"eduGapsGradNPG"`
	AllowActiveBacklogs         bool              `form:"allowActiveBacklogs" json:"allowActiveBacklogs"`
	NumberOfAllowedBacklogs     int               `form:"numberOfAllowedBacklogs" json:"numberOfAllowedBacklogs"`
	YearOfPassing               int               `form:"yearOfPassing" json:"yearOfPassing" binding:"required"`
	Remarks                     string            `form:"remarks" json:"remarks"`
	Programs                    []HcProgramsModel `form:"hcPrograms" json:"hcPrograms,omitempty" binding:"dive"`
	CreationDate                time.Time         `form:"-" json:"creationDate" time_format="2006-12-01T21:23:34.409Z"`
	LastUpdatedDate             time.Time         `form:"-" json:"lastUpdatedDate" time_format="2006-12-01T21:23:34.409Z"`
	PublishedFlag               bool              `form:"-" json:"publishedFlag"`
	PublishID                   string            `form:"-" json:"publishID"`
	ProgramsInString            string            `json:"hcProgramsInString"`
}

// HcProgramsModel ...
type HcProgramsModel struct {
	HiringCriteriaID   string `json:"hiringCriteriaID,omitempty"`
	HiringCriteriaName string `json:"hiringCriteriaName,omitempty"`
	StakeholderID      string `json:"stakeholderID,omitempty"`
	ProgramName        string `json:"programName" binding:"required"`
	ProgramID          string `json:"programID" binding:"required"`
	BranchName         string `json:"branchName" binding:"required"`
	BranchID           string `json:"branchID" binding:"required"`
	CreationDate       string `json:"creationDate"`
	LastUpdatedDate    string `json:"lastUpdatedDate,omitempty"`
	PublishFlag        string `json:"publishFlag,omitempty"`
}

// MultipleHC ...
type MultipleHC struct {
	HiringCriterias []HiringCriteriaDB `form:"hiringCriterias" json:"hiringCriterias" binding:"dive"`
}

// JobHcMappingDB ...
type JobHcMappingDB struct {
	JobID              string     `form:"jobID" json:"jobID"`
	StakeholderID      string     `form:"-" json:"stakeholderID,omitempty"`
	HiringCriteriaID   NullString `form:"-" json:"-"`
	HiringCriteriaName NullString `form:"-" json:"-"`
	HcID               string     `form:"hiringCriteriaID" json:"hiringCriteriaID"`
	HcName             string     `form:"hiringCriteriaName" json:"hiringCriteriaName"`
	JobName            string     `form:"jobName" json:"jobName" binding:"required"`
	CreationDate       time.Time  `form:"-" json:"creationDate"`
	PublishedFlag      bool       `form:"-" json:"publishedFlag"`
	PublishID          string     `form:"-" json:"publishID"`
}

// JobSkillsMapping ...
type JobSkillsMapping struct {
	ID            int       `form:"id" json:"id"`
	JobID         string    `form:"jobID" json:"jobID"`
	JobName       string    `form:"jobName" json:"jobName"`
	StakeholderID string    `form:"-" json:"stakeholder,omitempty"`
	SkillID       string    `form:"skillID" json:"skillID"`
	Skill         string    `form:"skill" json:"skill"`
	NoOfPositions int       `form:"noOfPositions" json:"noOfPositions"`
	Location      string    `form:"location" json:"location"`
	SalaryRange   string    `form:"salaryRange" json:"salaryRange"`
	DateOfHiring  time.Time `form:"dateOfHiring" json:"dateOfHiring" binding:"required"`
	Status        string    `form:"status" json:"status"`
	Remarks       string    `form:"remarks" json:"remarks"`
	Attachment    []byte    `form:"attachment" json:"attachment"`
	CreationDate  time.Time `form:"creationDate" json:"creationDate"`
}

// FullJobDb ...
type FullJobDb struct {
	JobHcMappingDB
	Jobs []JobSkillsMapping `form:"skills" json:"skills" binding:"dive"`
}

// SkillsUpdateJobDb ...
type SkillsUpdateJobDb struct {
	JobID         string             `form:"-" json:"jobID"`
	JobName       string             `form:"jobName" json:"jobName" binding:"required"`
	StakeholderID string             `form:"-"`
	Jobs          []JobSkillsMapping `form:"skills" json:"skills" binding:"dive"`
}

// PublishedJobsDB ...
type PublishedJobsDB struct {
	PublishID     string    `form:"-" json:"publishID"`
	JobID         string    `form:"jobID" json:"jobID" binding:"required"`
	JobName       string    `form:"-" json:"jobName"`
	StakeholderID string    `form:"-" json:"stakeholderID"`
	CreationDate  time.Time `form:"-" json:"creationDate"`
}

// PublishJobs ...
type PublishJobs struct {
	PublishedJobs []PublishedJobsDB `form:"publishJobs" json:"publishJobs" binding:"dive"`
}

// PublishHiringCriteriasModel ...
type PublishHiringCriteriasModel struct {
	StakeholderID     string   `form:"-" json:"-"`
	HiringCriteriaIDs []string `form:"hiringCriteriaIDs" json:"hiringCriteriaIDs" binding:"required"`
}

// PublishDataModel ...
type PublishDataModel struct {
	StakeholderID string `form:"-" json:"-"`
	PublishData   string `form:"publishData" json:"publishData" binding:"required"`
}

// OtherInformationModel ...
type OtherInformationModel struct {
	StakeholderID   string    `form:"-" json:"-"`
	Title           string    `form:"title" json:"title" binding:"required"`
	Information     string    `form:"information" json:"information" binding:"required"`
	Attachment      []byte    `form:"-" json:"-"`
	ID              int       `form:"-" json:"id"`
	PublishID       string    `form:"-" json:"publishID,omitempty"`
	PublishedFlag   bool      `form:"-" json:"publishedFlag"`
	CreationDate    time.Time `form:"-" json:"creationDate"`
	LastUpdatedDate time.Time `form:"-" json:"lastUpdatedDate"`
}
