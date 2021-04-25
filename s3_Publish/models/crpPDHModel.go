package models

import "time"

// JobPdhModel ...
type JobPdhModel struct {
	JobID                        string `json:"jobID"`
	JobName                      string `json:"jobName"`
	CorporateName                string `json:"corporateName"`
	ProgramID                    string `json:"programID"`
	BranchID                     string `json:"branchID"`
	MinimumCutoffCategory        string `json:"minimumCutoffCategory"`
	MinimumCutoff                string `json:"minimumCutoff"`
	ActiveBacklogsAllowed        string `json:"activeBacklogsAllowed"`
	TotalNumberOfBacklogsAllowed string `json:"totalNumberOfBacklogsAllowed"`
	EduGaps11N12Allowed          string `json:"eduGaps11N12Allowed"`
	EduGapsGradAllowed           string `json:"eduGapsGradAllowed"`
	EduGapsSchoolAllowed         string `json:"eduGapsSchoolAllowed"`
	EduGapsPGAllowed             string `json:"eduGapsPGAllowed"`
	YearOfPassing                string `json:"yearOfPassing"`
	Remarks                      string `json:"remarks"`
	Skills                       string `json:"skills"`
}

// HcPdhModel ...
type HcPdhModel struct {
	HcID                         string `json:"hiringCriteriaID"`
	HcName                       string `json:"hiringCriteriaName"`
	CorporateName                string `json:"corporateName"`
	ProgramID                    string `json:"programID"`
	BranchID                     string `json:"branchID"`
	MinimumCutoffCategory        string `json:"minimumCutoffCategory"`
	MinimumCutoff                string `json:"minimumCutoff"`
	ActiveBacklogsAllowed        string `json:"activeBacklogsAllowed"`
	TotalNumberOfBacklogsAllowed string `json:"totalNumberOfBacklogsAllowed"`
	EduGaps11N12Allowed          string `json:"eduGaps11N12Allowed"`
	EduGapsGradAllowed           string `json:"eduGapsGradAllowed"`
	EduGapsSchoolAllowed         string `json:"eduGapsSchoolAllowed"`
	EduGapsPGAllowed             string `json:"eduGapsPGAllowed"`
	YearOfPassing                string `json:"yearOfPassing"`
	Remarks                      string `json:"remarks"`
}

// OtherInformationSubModel ...
type OtherInformationSubModel struct {
	Title       string `json:"title"`
	Information string `json:"information"`
	tempAttach  NullString
	Attachment  []byte `json:"attachment,omitempty"`
}

// CorpPushedDataModel ...
type CorpPushedDataModel struct {
	StakeholderID           string    `form:"-" json:"stakeholder,omitempty"`
	PublishID               string    `form:"-" json:"publishID"`
	StudentName             string    `form:"StudentName" json:"StudentName"`
	DateOfPublish           string    `form:"-" json:"dateOfPublish"`
	HiringCriteriaPublished bool      `form:"hiringCriteriaPublished" json:"hiringCriteriaPublished"`
	JobsPublished           bool      `form:"jobsPublished" json:"jobsPublished"`
	ProfilePublished        bool      `form:"profilePublished" json:"profilePublished"`
	OtherPublished          bool      `form:"otherPublished" json:"otherPublished"`
	GeneralNote             string    `form:"-" json:"generalNote"`
	PublishedData           string    `form:"publishData" json:"publishData" binding:"required"`
	CreationDate            time.Time `form:"-" json:"creationDate,omitEmpty"`
	LastUpdatedDate         time.Time `form:"-" json:"lastUpdatedDate,omitempty"`
}
