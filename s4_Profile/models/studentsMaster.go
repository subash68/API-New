// Package models ...
package models

import "time"

// StudentMasterDb ...
type StudentMasterDb struct {
	StakeholderID              string     `form:"stakeholderID" json:"stakeholderID,omitempty"`
	FirstName                  string     `form:"firstName" json:"firstName" binding:"required"`
	MiddleName                 string     `form:"middleName" json:"middleName"`
	LastName                   string     `form:"lastName" json:"lastName" binding:"required"`
	PersonalEmail              string     `form:"personalEmail" json:"personalEmail"`
	CollegeEmail               string     `form:"collegeEmail" json:"collegeEmail" binding:"required,email"`
	PhoneNumber                string     `form:"phoneNuber" json:"phoneNuber" binding:"required"`
	AlternatePhoneNumber       string     `form:"alternatePhoneNuber" json:"alternatePhoneNuber"`
	CollegeID                  string     `form:"collegeID" json:"collegeID" binding:"required"`
	Gender                     string     `form:"gender" json:"gender" binding:"required"`
	DateOfBirth                string     `form:"dateOfBirth" json:"dateOfBirth" binding:"required"`
	AadharNumber               string     `form:"aadharNumber" json:"aadharNumber" binding:"required"`
	PermanentAddressLine1      string     `form:"permanentAddressLine1" json:"permanentAddressLine1" binding:"required"`
	PermanentAddressLine2      string     `form:"permanentAddressLine2" json:"permanentAddressLine2" binding:"required"`
	PermanentAddressLine3      string     `form:"permanentAddressLine3" json:"permanentAddressLine3"`
	PermanentAddressCountry    string     `form:"permanentAddressCountry" json:"permanentAddressCountry" binding:"required"`
	PermanentAddressState      string     `form:"permanentAddressState" json:"permanentAddressState" binding:"required"`
	PermanentAddressCity       string     `form:"permanentAddressCity" json:"permanentAddressCity" binding:"required"`
	PermanentAddressDistrict   string     `form:"permanentAddressDistrict" json:"permanentAddressDistrict" binding:"required"`
	PermanentAddressZipcode    string     `form:"permanentAddressZipcode" json:"permanentAddressZipcode" binding:"required"`
	PermanentAddressPhone      string     `form:"permanentAddressPhone" json:"permanentAddressPhone" binding:"required"`
	PresentAddressLine1        string     `form:"presentAddressLine1" json:"presentAddressLine1" binding:"required"`
	PresentAddressLine2        string     `form:"presentAddressLine2" json:"presentAddressLine2" binding:"required"`
	PresentAddressLine3        string     `form:"presentAddressLine3" json:"presentAddressLine3"`
	PresentAddressCountry      string     `form:"presentAddressCountry" json:"presentAddressCountry" binding:"required"`
	PresentAddressState        string     `form:"presentAddressState" json:"presentAddressState" binding:"required"`
	PresentAddressCity         string     `form:"presentAddressCity" json:"presentAddressCity" binding:"required"`
	PresentAddressDistrict     string     `form:"presentAddressDistrict" json:"presentAddressDistrict" binding:"required"`
	PresentAddressZipcode      string     `form:"presentAddressZipcode" json:"presentAddressZipcode" binding:"required"`
	PresentAddressPhone        string     `form:"presentAddressPhone" json:"presentAddressPhone" binding:"required"`
	FathersGuardianFullName    string     `form:"fathersGuardianFullName" json:"fathersGuardianFullName" binding:"required"`
	FathersGuardianOccupation  string     `form:"fathersGuardianOccupation" json:"fathersGuardianOccupation" binding:"required"`
	FathersGuardianCompany     string     `form:"fathersGuardianCompany" json:"fathersGuardianCompany" binding:"required"`
	FathersGuardianPhoneNumber string     `form:"fathersGuardianPhoneNumber" json:"fathersGuardianPhoneNumber" binding:"required"`
	FathersGuardianEmailID     string     `form:"fathersGuardianEmailID" json:"fathersGuardianEmailID" binding:"required,email"`
	MothersGuardianFullName    string     `form:"mothersGuardianFullName" json:"mothersGuardianFullName"`
	MothersGuardianOccupation  string     `form:"mothersGuardianOccupation" json:"mothersGuardianOccupation"`
	MothersGuardianCompany     string     `form:"mothersGuardianCompany" json:"mothersGuardianCompany"`
	MothersGuardianDesignation string     `form:"mothersGuardianDesignation" json:"mothersGuardianDesignation"`
	MothersGuardianPhoneNumber string     `form:"mothersGuardianPhoneNumber" json:"mothersGuardianPhoneNumber" `
	MothersGuardianEmailID     string     `form:"mothersGuardianEmailID" json:"mothersGuardianEmailID" `
	AccountStatus              string     `form:"accountStatus" json:"accountStatus"`
	Password                   string     `form:"password" json:"password,omitempty" binding:"required,min=8,max=15" binding:"required"`
	PrimaryPhoneVerified       bool       `form:"primaryPhoneVerified" json:"primaryPhoneVerified"`
	PrimaryEmailVerified       bool       `form:"primaryEmailVerified" json:"primaryEmailVerified"`
	DateOfJoining              time.Time  `json:"dateOfJoining,omitempty" time_format="2006-12-01T21:23:34.409Z"`
	ProfilePicture             []byte     `form:"-" json:"profilePicture"`
	AccountExpiryDate          time.Time  `form:"-" json:"accountExpiryDate" time_format="2006-12-01T21:23:34.409Z"`
	AboutMeNullable            NullString `json:"-"`
	AboutMe                    string     `form:"aboutMe" json:"aboutMe"`
}

// StudentTTModel ...
type StudentTTModel struct {
	Name                  string `form:"name" json:"name" binding:"required"`
	Location              string `form:"location" json:"location" binding:"required"`
	MonthAndYearOfPassing string `form:"monthAndYearOfPassing" json:"monthAndYearOfPassing" binding:"required"`
	Board                 string `form:"schoolBoard" json:"schoolBoard" binding:"required"`
	Percentage            string `form:"percentage" json:"percentage" binding:"required"`
	AttachmentFile        []byte `form:"attachment" json:"attachment" binding:"required"`
	EnablingFlag          bool   `form:"-" json:"enablingFlag"`
}

// StudentNullableTTModel ...
type StudentNullableTTModel struct {
	Name                  NullString `form:"name" json:"name"`
	Location              NullString `form:"location" json:"location"`
	MonthAndYearOfPassing NullString `form:"monthAndYearOfPassing" json:"monthAndYearOfPassing"`
	Board                 NullString `form:"schoolBoard" json:"schoolBoard"`
	Percentage            NullString `form:"percentage" json:"percentage"`
	AttachmentFile        []byte     `form:"attachment" json:"attachment"`
	EnablingFlag          NullBool   `form:"-" json:"enablingFlag"`
}

// StudentAcademicsModelReq ...
type StudentAcademicsModelReq struct {
	StakeholderID string         `form:"-" json:"-"`
	Tenth         StudentTTModel `form:"tenth,omitempty" json:"tenth,omitempty" binding:"dive"`
	Twelfth       StudentTTModel `form:"twelfth,omitempty" json:"twelfth,omitempty" binding:"dive"`
}

// StudentLangModel ...
type StudentLangModel struct {
	StakeholderID    string    `form:"-" json:"-"`
	ID               int       `form:"-" json:"id"`
	LanguageName     string    `form:"languageName" json:"languageName" binding:"required"`
	SpeakProficiency string    `form:"speakProficiency" json:"speakProficiency" binding:"required"`
	ReadProficiency  string    `form:"readProficiency" json:"readProficiency" binding:"required"`
	WriteProficiency string    `form:"writeProficiency" json:"writeProficiency" binding:"required"`
	IsMotherTongue   bool      `form:"isMotherTongue" json:"isMotherTongue" binding:"required"`
	EnabledFlag      bool      `form:"-" json:"enabledFlag"`
	CreationDate     time.Time `form:"-" json:"creationDate"`
	LastUpdatedDate  time.Time `form:"-" json:"lastUpdatedDate"`
}

// StudentAllLanguagesModel ...
type StudentAllLanguagesModel struct {
	StakeholderID string             `form:"-" json:"-"`
	Languages     []StudentLangModel `form:"languages" json:"languages" binding:"dive"`
}

// StudentCertsModel ...
type StudentCertsModel struct {
	StakeholderID    string    `form:"-" json:"-"`
	ID               int       `form:"-" json:"id"`
	Name             string    `form:"name" json:"name" binding:"required"`
	IssuingAuthority string    `form:"issuingAuthority" json:"issuingAuthority" binding:"required"`
	StartDate        time.Time `form:"startDate" json:"startDate" binding:"required"`
	EndDate          string    `form:"endDate" json:"endDate"`
	Result           string    `form:"result" json:"result"`
	Description      string    `form:"description" json:"description"`
	Attachment       []byte    `form:"attachment" json:"attachment"`
	EnabledFlag      bool      `form:"-" json:"enabledFlag"`
	CreationDate     time.Time `form:"-" json:"creationDate"`
	LastUpdatedDate  time.Time `form:"-" json:"lastUpdatedDate"`
}

// StudentAllCertsModel ...
type StudentAllCertsModel struct {
	StakeholderID  string              `form:"-" json:"-"`
	Certifications []StudentCertsModel `form:"certifications" json:"certifications" binding:"dive"`
}

// StudentAssessmentModel ...
type StudentAssessmentModel struct {
	StakeholderID    string    `form:"-" json:"-"`
	ID               int       `form:"-" json:"id"`
	Name             string    `form:"name" json:"name" binding:"required"`
	Score            string    `form:"score" json:"score" binding:"required"`
	IssuingAuthority string    `form:"issuingAuthority" json:"issuingAuthority" binding:"required"`
	AssessmentDate   time.Time `form:"assessmentDate" json:"assessmentDate" binding:"required"`
	Description      string    `form:"description" json:"description"`
	Attachment       []byte    `form:"attachment" json:"attachment"`
	EnabledFlag      bool      `form:"-" json:"enabledFlag"`
	CreationDate     time.Time `form:"-" json:"creationDate"`
	LastUpdatedDate  time.Time `form:"-" json:"lastUpdatedDate"`
}

// StudentAllAssessmentModel ...
type StudentAllAssessmentModel struct {
	StakeholderID string                   `form:"-" json:"-"`
	Assessments   []StudentAssessmentModel `form:"assessments" json:"assessments" binding:"dive"`
}

// StudentInternshipModel ...
type StudentInternshipModel struct {
	StakeholderID    string    `form:"-" json:"-"`
	ID               int       `form:"-" json:"id"`
	Name             string    `form:"name" json:"name" binding:"required"`
	OrganizationName string    `form:"organizationName" json:"organizationName" binding:"required"`
	OrganizationCity string    `form:"organizationCity" json:"organizationCity" binding:"required"`
	FieldOfWork      string    `form:"fieldOfWork" json:"fieldOfWork" binding:"required"`
	StartDate        time.Time `form:"startDate" json:"startDate" binding:"required" time_format="2006-12-01T21:23:34.409Z"`
	EndDate          time.Time `form:"endDate" json:"endDate"  binding:"required" time_format="2006-12-01T21:23:34.409Z"`
	Description      string    `form:"description" json:"description"`
	Attachment       []byte    `form:"attachment" json:"attachment"`
	EnabledFlag      bool      `form:"-" json:"enabledFlag"`
	CreationDate     time.Time `form:"-" json:"creationDate"`
	LastUpdatedDate  time.Time `form:"-" json:"lastUpdatedDate"`
}

// StudentAllInternshipModel ...
type StudentAllInternshipModel struct {
	StakeholderID string                   `form:"-" json:"-"`
	Internships   []StudentInternshipModel `form:"internships" json:"internships" binding:"dive"`
}

// StudentAwardsModel ...
type StudentAwardsModel struct {
	StakeholderID    string    `form:"-" json:"-"`
	ID               int       `form:"-" json:"id"`
	RecognitionName  string    `form:"recognitionName" json:"recognitionName" binding:"required"`
	RecognitionDate  time.Time `form:"recognitionDate" json:"recognitionDate" binding:"required" time_format="2006-12-01T21:23:34.409Z"`
	IssuingAuthority string    `form:"issuingAuthority" json:"issuingAuthority" binding:"required"`
	Attachment       []byte    `form:"attachment" json:"attachment"`
	EnabledFlag      bool      `form:"-" json:"enabledFlag"`
	CreationDate     time.Time `form:"-" json:"creationDate"`
	LastUpdatedDate  time.Time `form:"-" json:"lastUpdatedDate"`
}

// StudentAllAwardsModel ...
type StudentAllAwardsModel struct {
	StakeholderID string               `form:"-" json:"-"`
	Awards        []StudentAwardsModel `form:"awards" json:"awards" binding:"dive"`
}

// StudentCompetitionModel ...
type StudentCompetitionModel struct {
	StakeholderID   string    `form:"-" json:"-"`
	ID              int       `form:"-" json:"id"`
	Name            string    `form:"name" json:"name" binding:"required"`
	Date            time.Time `form:"date" json:"date" binding:"required" time_format="2006-12-01T21:23:34.409Z"`
	Rank            string    `form:"rank" json:"rank" binding:"required"`
	Attachment      []byte    `form:"attachment" json:"attachment"`
	EnabledFlag     bool      `form:"-" json:"enabledFlag"`
	CreationDate    time.Time `form:"-" json:"creationDate"`
	LastUpdatedDate time.Time `form:"-" json:"lastUpdatedDate"`
}

// StudentAllCompetitionModel ...
type StudentAllCompetitionModel struct {
	StakeholderID string                    `form:"-" json:"-"`
	Competitions  []StudentCompetitionModel `form:"competitions" json:"competitions" binding:"dive"`
}

// StudentConfWorkshopModel ...
type StudentConfWorkshopModel struct {
	StakeholderID   string    `form:"-" json:"-"`
	ID              int       `form:"-" json:"id"`
	Name            string    `form:"name" json:"name" binding:"required"`
	Date            time.Time `form:"date" json:"date" binding:"required" time_format="2006-12-01T21:23:34.409Z"`
	Attachment      []byte    `form:"attachment" json:"attachment"`
	EnabledFlag     bool      `form:"-" json:"enabledFlag"`
	CreationDate    time.Time `form:"-" json:"creationDate"`
	LastUpdatedDate time.Time `form:"-" json:"lastUpdatedDate"`
}

// StudentAllConfWorkshopModel ...
type StudentAllConfWorkshopModel struct {
	StakeholderID string                     `form:"-" json:"-"`
	ConfWorkshop  []StudentConfWorkshopModel `form:"confWorkshop" json:"confWorkshop" binding:"dive"`
}

// StudentExtraCurricularModel ...
type StudentExtraCurricularModel struct {
	StakeholderID   string    `form:"-" json:"-"`
	ID              int       `form:"-" json:"id"`
	Name            string    `form:"name" json:"name" binding:"required"`
	Attachment      []byte    `form:"attachment" json:"attachment"`
	EnabledFlag     bool      `form:"-" json:"enabledFlag"`
	CreationDate    time.Time `form:"-" json:"creationDate"`
	LastUpdatedDate time.Time `form:"-" json:"lastUpdatedDate"`
}

// StudentAllExtraCurricularModel ...
type StudentAllExtraCurricularModel struct {
	StakeholderID   string                        `form:"-" json:"-"`
	ExtraCurricular []StudentExtraCurricularModel `form:"extraCurricular" json:"extraCurricular" binding:"dive"`
}

// StudentPatentsModel ...
type StudentPatentsModel struct {
	StakeholderID   string    `form:"-" json:"-"`
	ID              int       `form:"-" json:"id"`
	Name            string    `form:"name" json:"name" binding:"required"`
	PatentType      string    `form:"patentType" json:"patentType" binding:"required"`
	PatentNumber    string    `form:"patentNumber" json:"patentNumber" binding:"required"`
	PatentStatus    string    `form:"patentStatus" json:"patentStatus" binding:"required"`
	Attachment      []byte    `form:"attachment" json:"attachment"`
	EnabledFlag     bool      `form:"-" json:"enabledFlag"`
	CreationDate    time.Time `form:"-" json:"creationDate"`
	LastUpdatedDate time.Time `form:"-" json:"lastUpdatedDate"`
}

// StudentAllPatentsModel ...
type StudentAllPatentsModel struct {
	StakeholderID string                `form:"-" json:"-"`
	Patents       []StudentPatentsModel `form:"patents" json:"patents" binding:"dive"`
}

// StudentProjectsModel ...
type StudentProjectsModel struct {
	StakeholderID   string    `form:"-" json:"-"`
	ID              int       `form:"-" json:"id"`
	Name            string    `form:"name" json:"name" binding:"required"`
	ProjectAbstract string    `form:"projectAbstract" json:"projectAbstract" binding:"required"`
	GuideName       string    `form:"guideName" json:"guideName"`
	GuideEmail      string    `form:"guideEmail" json:"guideEmail"`
	StartDate       time.Time `form:"startDate" json:"startDate" binding:"required" time_format="2006-12-01T21:23:34.409Z"`
	EndDate         time.Time `form:"endDate" json:"endDate,omitempty
	" time_format="2006-12-01T21:23:34.409Z"`
	Attachment      []byte    `form:"attachment" json:"attachment"`
	EnabledFlag     bool      `form:"-" json:"enabledFlag"`
	CreationDate    time.Time `form:"-" json:"creationDate"`
	LastUpdatedDate time.Time `form:"-" json:"lastUpdatedDate"`
}

// StudentAllProjectsModel ...
type StudentAllProjectsModel struct {
	StakeholderID string                 `form:"-" json:"-"`
	Projects      []StudentProjectsModel `form:"projects" json:"projects" binding:"dive"`
}

// StudentPublicationsModel ...
type StudentPublicationsModel struct {
	StakeholderID       string    `form:"-" json:"-"`
	ID                  int       `form:"-" json:"id"`
	Name                string    `form:"name" json:"name" binding:"required"`
	PublishingAuthority string    `form:"publishingAuthority" json:"publishingAuthority" binding:"required"`
	GuideName           string    `form:"guideName" json:"guideName" binding:"required"`
	GuideEmail          string    `form:"guideEmail" json:"guideEmail" binding:"required"`
	StartDate           time.Time `form:"startDate" json:"startDate" binding:"required" time_format="2006-12-01T21:23:34.409Z"`
	EndDate             time.Time `form:"endDate" json:"endDate" binding:"required" time_format="2006-12-01T21:23:34.409Z"`
	Attachment          []byte    `form:"attachment" json:"attachment"`
	EnabledFlag         bool      `form:"-" json:"enabledFlag"`
	CreationDate        time.Time `form:"-" json:"creationDate"`
	LastUpdatedDate     time.Time `form:"-" json:"lastUpdatedDate"`
}

// StudentAllPublicationsModel ...
type StudentAllPublicationsModel struct {
	StakeholderID string                     `form:"-" json:"-"`
	Publications  []StudentPublicationsModel `form:"publications" json:"publications" binding:"dive"`
}

// StudentScholarshipsModel ...
type StudentScholarshipsModel struct {
	StakeholderID       string    `form:"-" json:"-"`
	ID                  int       `form:"-" json:"id"`
	Name                string    `form:"name" json:"name" binding:"required"`
	ScholarshipIssuedBy string    `form:"scholarshipIssuedBy" json:"publishingAuthority" binding:"required"`
	ScholarshipDate     time.Time `form:"scholarshipDate" json:"scholarshipDate" binding:"required" time_format="2006-12-01T21:23:34.409Z"`
	Attachment          []byte    `form:"attachment" json:"attachment"`
	EnabledFlag         bool      `form:"-" json:"enabledFlag"`
	CreationDate        time.Time `form:"-" json:"creationDate"`
	LastUpdatedDate     time.Time `form:"-" json:"lastUpdatedDate"`
}

// StudentAllScholarshipsModel ...
type StudentAllScholarshipsModel struct {
	StakeholderID string                     `form:"-" json:"-"`
	Scholarships  []StudentScholarshipsModel `form:"scholarships" json:"scholarships" binding:"dive"`
}

// StudentSocialAccountModel ...
type StudentSocialAccountModel struct {
	StakeholderID   string    `form:"-" json:"-"`
	ID              int       `form:"-" json:"id"`
	UserID          string    `form:"userID" json:"userID" binding:"required"`
	EnabledFlag     bool      `form:"-" json:"enabledFlag"`
	CreationDate    time.Time `form:"-" json:"creationDate"`
	LastUpdatedDate time.Time `form:"-" json:"lastUpdatedDate"`
}

// StudentAllSocialAccountModel ...
type StudentAllSocialAccountModel struct {
	StakeholderID  string                      `form:"-" json:"-"`
	SocialAccounts []StudentSocialAccountModel `form:"socialAccounts" json:"socialAccounts" binding:"dive"`
}

// StudentTechSkillsModel ...
type StudentTechSkillsModel struct {
	StakeholderID   string    `form:"-" json:"-"`
	ID              int       `form:"-" json:"id"`
	SkillID         string    `form:"skillID" json:"skillID" binding:"required"`
	SkillName       string    `form:"skillName" json:"skillName" binding:"required"`
	EnabledFlag     bool      `form:"-" json:"enabledFlag"`
	CreationDate    time.Time `form:"-" json:"creationDate"`
	LastUpdatedDate time.Time `form:"-" json:"lastUpdatedDate"`
}

// StudentAllTechSkillsModel ...
type StudentAllTechSkillsModel struct {
	StakeholderID string                   `form:"-" json:"-"`
	TechSkills    []StudentTechSkillsModel `form:"techSkills" json:"techSkills" binding:"dive"`
}

// StudentTestScoresModel ...
type StudentTestScoresModel struct {
	StakeholderID   string    `form:"-" json:"-"`
	ID              int       `form:"-" json:"id"`
	Name            string    `form:"name" json:"name" binding:"required"`
	TestScoreDate   time.Time `form:"testScoreDate" json:"testScoreDate" binding:"required" time_format="2006-12-01T21:23:34.409Z"`
	TestScore       string    `form:"testScore" json:"testScore" binding:"required"`
	TestScoreTotal  string    `form:"testScoreTotal" json:"testScoreTotal" binding:"required"`
	Attachment      []byte    `form:"attachment" json:"attachment"`
	EnabledFlag     bool      `form:"-" json:"enabledFlag"`
	CreationDate    time.Time `form:"-" json:"creationDate"`
	LastUpdatedDate time.Time `form:"-" json:"lastUpdatedDate"`
}

// StudentAllTestScoresModel ...
type StudentAllTestScoresModel struct {
	StakeholderID string                   `form:"-" json:"-"`
	TestScores    []StudentTestScoresModel `form:"testScores" json:"testScores" binding:"dive"`
}

// StudentVolunteerExperienceModel ...
type StudentVolunteerExperienceModel struct {
	StakeholderID   string    `form:"-" json:"-"`
	ID              int       `form:"-" json:"id"`
	Name            string    `form:"name" json:"name" binding:"required"`
	Organisation    string    `form:"organisation" json:"organisation" binding:"required"`
	Location        string    `form:"location" json:"location" binding:"required"`
	StartDate       time.Time `form:"startDate" json:"startDate" binding:"required" time_format="2006-12-01T21:23:34.409Z"`
	EndDate         time.Time `form:"endDate" json:"endDate" binding:"required" time_format="2006-12-01T21:23:34.409Z"`
	Attachment      []byte    `form:"attachment" json:"attachment"`
	EnabledFlag     bool      `form:"-" json:"enabledFlag"`
	CreationDate    time.Time `form:"-" json:"creationDate"`
	LastUpdatedDate time.Time `form:"-" json:"lastUpdatedDate"`
}

// StudentAllVolunteerExperienceModel ...
type StudentAllVolunteerExperienceModel struct {
	StakeholderID       string                            `form:"-" json:"-"`
	VolunteerExperience []StudentVolunteerExperienceModel `form:"volunteerExperience" json:"VolunteerExperience" binding:"dive"`
}

// StudentContactInfoModel ...
type StudentContactInfoModel struct {
	StakeholderID            string `form:"stakeholderID" json:"stakeholderID,omitempty"`
	FirstName                string `form:"firstName" json:"firstName" binding:"required"`
	MiddleName               string `form:"middleName" json:"middleName"`
	LastName                 string `form:"lastName" json:"lastName" binding:"required"`
	PersonalEmail            string `form:"personalEmail" json:"personalEmail"`
	CollegeEmail             string `form:"collegeEmail" json:"collegeEmail" binding:"required,email"`
	PhoneNumber              string `form:"phoneNuber" json:"phoneNuber" binding:"required"`
	AlternatePhoneNumber     string `form:"alternatePhoneNuber" json:"alternatePhoneNuber"`
	CollegeID                string `form:"collegeID" json:"collegeID" binding:"required"`
	Gender                   string `form:"gender" json:"gender" binding:"required"`
	DateOfBirth              string `form:"dateOfBirth" json:"dateOfBirth" binding:"required"`
	AadharNumber             string `form:"aadharNumber" json:"aadharNumber" binding:"required"`
	PermanentAddressLine1    string `form:"permanentAddressLine1" json:"permanentAddressLine1" binding:"required"`
	PermanentAddressLine2    string `form:"permanentAddressLine2" json:"permanentAddressLine2" binding:"required"`
	PermanentAddressLine3    string `form:"permanentAddressLine3" json:"permanentAddressLine3"`
	PermanentAddressCountry  string `form:"permanentAddressCountry" json:"permanentAddressCountry" binding:"required"`
	PermanentAddressState    string `form:"permanentAddressState" json:"permanentAddressState" binding:"required"`
	PermanentAddressCity     string `form:"permanentAddressCity" json:"permanentAddressCity" binding:"required"`
	PermanentAddressDistrict string `form:"permanentAddressDistrict" json:"permanentAddressDistrict" binding:"required"`
	PermanentAddressZipcode  string `form:"permanentAddressZipcode" json:"permanentAddressZipcode" binding:"required"`
	PermanentAddressPhone    string `form:"permanentAddressPhone" json:"permanentAddressPhone" binding:"required"`
	PresentAddressLine1      string `form:"presentAddressLine1" json:"presentAddressLine1" binding:"required"`
	PresentAddressLine2      string `form:"presentAddressLine2" json:"presentAddressLine2" binding:"required"`
	PresentAddressLine3      string `form:"presentAddressLine3" json:"presentAddressLine3"`
	PresentAddressCountry    string `form:"presentAddressCountry" json:"presentAddressCountry" binding:"required"`
	PresentAddressState      string `form:"presentAddressState" json:"presentAddressState" binding:"required"`
	PresentAddressCity       string `form:"presentAddressCity" json:"presentAddressCity" binding:"required"`
	PresentAddressDistrict   string `form:"presentAddressDistrict" json:"presentAddressDistrict" binding:"required"`
	PresentAddressZipcode    string `form:"presentAddressZipcode" json:"presentAddressZipcode" binding:"required"`
	PresentAddressPhone      string `form:"presentAddressPhone" json:"presentAddressPhone" binding:"required"`
	AboutMe                  string `form:"aboutMe" json:"aboutMe"`
}

// StudentCompleteProfileDataModel ...
type StudentCompleteProfileDataModel struct {
	Profile             StudentMasterDb           `form:"-" json:"-"`
	ContactInfo         StudentContactInfoModel   `form:"contactInfo" json:"contactInfo"`
	Academics           StudentAcademicsModelReq  `form:"academics" json:"academics"`
	Languages           StudentAllLanguagesModel  `form:"-" json:"-"`
	Certifications      StudentAllCertsModel      `form:"-" json:"-"`
	Assessments         StudentAllAssessmentModel `form:"-" json:"-"`
	Internships         StudentAllInternshipModel `form:"-" json:"-"`
	LanguagesArray      []StudentLangModel        `form:"languages" json:"languages"`
	CertificationsArray []StudentCertsModel       `form:"certifications" json:"certifications"`
	AssessmentsArray    []StudentAssessmentModel  `form:"assessments" json:"assessments"`
	InternshipsArray    []StudentInternshipModel  `form:"internships" json:"internships"`
}