package models

import "time"

// StuPublishDBModel ...
type StuPublishDBModel struct {
	StakeholderID             string    `form:"-" json:"stakeholder,omitempty"`
	PublishID                 string    `form:"-" json:"publishID"`
	StudentName               string    `form:"StudentName" json:"StudentName"`
	DateOfPublish             time.Time `form:"-" json:"dateOfPublish"`
	ContactInfoPublished      bool      `form:"contactInfoPublished" json:"contactInfoPublished"`
	EducationPublished        bool      `form:"educationPublished" json:"educationPublished"`
	LanguagesPublished        bool      `form:"languagesPublished" json:"languagesPublished"`
	CertificationsPublished   bool      `form:"certificationsPublished" json:"certificationsPublished"`
	AssessmentsPublished      bool      `form:"assessmentsPublished" json:"assessmentsPublished"`
	InternshipPublished       bool      `form:"internshipPublished" json:"internshipPublished"`
	OtherInformationPublished bool      `form:"otherInformationPublished" json:"otherInformationPublished"`
	GeneralNote               string    `form:"-" json:"generalNote"`
	PublishedData             string    `form:"publishData" json:"publishData" binding:"required"`
	CreationDate              time.Time `form:"-" json:"creationDate,omitEmpty"`
	LastUpdatedDate           time.Time `form:"-" json:"lastUpdatedDate,omitempty"`
}

// StuOtherInformationModel ...
type StuOtherInformationModel struct {
	StakeholderID   string    `form:"-" json:"-"`
	Title           string    `form:"title" json:"title" binding:"required"`
	Information     string    `form:"information" json:"information" binding:"required"`
	Attachment      []byte    `form:"-" json:"attachment,omitempty"`
	ID              int       `form:"-" json:"id"`
	PublishID       string    `form:"-" json:"publishID,omitempty"`
	PublishedFlag   string    `form:"-" json:"publishedFlag"`
	CreationDate    time.Time `form:"-" json:"creationDate"`
	LastUpdatedDate time.Time `form:"-" json:"lastUpdatedDate"`
}
