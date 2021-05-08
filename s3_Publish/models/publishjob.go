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
	HiringCriteriaID        string     `form:"-" json:"hiringCriteriaID"`
	HiringCriteriaName      string     `form:"hiringCriteriaName" json:"hiringCriteriaName" binding:"required" validate:"required"`
	StakeholderID           string     `form:"stakeholderID" json:"stakeholderID,omitempty"`
	ProgramID               string     `form:"programID" json:"programID" binding:"required"`
	DepartmentID            string     `form:"courseID" json:"courseID" binding:"required"`
	CutOffCategory          string     `form:"cutOffCategory" json:"cutOffCategory" binding:"required"`
	CutOff                  float64    `form:"cutOff" json:"cutOff" binding:"required"`
	EduGapsSchoolAllowed    bool       `form:"eduGapsSchoolAllowed" json:"eduGapsSchoolAllowed"`
	EduGaps11N12Allowed     bool       `form:"eduGaps11N12Allowed" json:"eduGaps11N12Allowed"`
	EduGapsGradAllowed      bool       `form:"eduGapsGradAllowed" json:"eduGapsGradAllowed"`
	EduGapsPGAllowed        bool       `form:"eduGapsPGAllowed" json:"eduGapsPGAllowed"`
	AllowActiveBacklogs     bool       `form:"allowActiveBacklogs" json:"allowActiveBacklogs"`
	NumberOfAllowedBacklogs int        `form:"numberOfAllowedBacklogs" json:"numberOfAllowedBacklogs"`
	YearOfPassing           int        `form:"yearOfPassing" json:"yearOfPassing" binding:"required"`
	Remarks                 string     `form:"remarks" json:"remarks"`
	CreationDate            time.Time  `form:"-" json:"creationDate"`
	PublishedFlagNull       NullBool   `form:"-" json:"-"`
	PublishIDNull           NullString `form:"-" json:"-"`
	PublishedFlag           bool       `form:"-" json:"publishedFlag"`
	PublishID               string     `form:"-" json:"publishID,omitempty"`
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
