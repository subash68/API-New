// Package models ...
package models

import "time"

// StudentMasterDb ...
type StudentMasterDb struct {
	StakeholderID               string     `form:"stakeholderID" json:"stakeholderID,omitempty"`
	FirstName                   string     `form:"firstName" json:"firstName" binding:"required"`
	MiddleName                  string     `form:"middleName" json:"middleName"`
	LastName                    string     `form:"lastName" json:"lastName" binding:"required"`
	PersonalEmail               string     `form:"personalEmail" json:"personalEmail"`
	CollegeEmail                string     `form:"collegeEmail" json:"collegeEmail" binding:"required,email"`
	PhoneNumber                 string     `form:"phoneNumber" json:"phoneNumber" binding:"required"`
	AlternatePhoneNumber        string     `form:"alternatePhoneNumber" json:"alternatePhoneNumber"`
	Gender                      string     `form:"gender" json:"gender" binding:"required"`
	DateOfBirth                 string     `form:"dateOfBirth" json:"dateOfBirth" binding:"required"`
	AadharNumber                string     `form:"aadharNumber" json:"aadharNumber" binding:"required"`
	PermanentAddressLine1       string     `form:"permanentAddressLine1" json:"permanentAddressLine1" binding:"required"`
	PermanentAddressLine2       string     `form:"permanentAddressLine2" json:"permanentAddressLine2" binding:"required"`
	PermanentAddressLine3       string     `form:"permanentAddressLine3" json:"permanentAddressLine3"`
	PermanentAddressCountry     string     `form:"permanentAddressCountry" json:"permanentAddressCountry" binding:"required"`
	PermanentAddressState       string     `form:"permanentAddressState" json:"permanentAddressState" binding:"required"`
	PermanentAddressCity        string     `form:"permanentAddressCity" json:"permanentAddressCity" binding:"required"`
	PermanentAddressDistrict    string     `form:"permanentAddressDistrict" json:"permanentAddressDistrict" binding:"required"`
	PermanentAddressZipcode     string     `form:"permanentAddressZipcode" json:"permanentAddressZipcode" binding:"required"`
	PermanentAddressPhone       string     `form:"permanentAddressPhone" json:"permanentAddressPhone" binding:"required"`
	PresentAddressLine1         string     `form:"presentAddressLine1" json:"presentAddressLine1" binding:"required"`
	PresentAddressLine2         string     `form:"presentAddressLine2" json:"presentAddressLine2" binding:"required"`
	PresentAddressLine3         string     `form:"presentAddressLine3" json:"presentAddressLine3"`
	PresentAddressCountry       string     `form:"presentAddressCountry" json:"presentAddressCountry" binding:"required"`
	PresentAddressState         string     `form:"presentAddressState" json:"presentAddressState" binding:"required"`
	PresentAddressCity          string     `form:"presentAddressCity" json:"presentAddressCity" binding:"required"`
	PresentAddressDistrict      string     `form:"presentAddressDistrict" json:"presentAddressDistrict" binding:"required"`
	PresentAddressZipcode       string     `form:"presentAddressZipcode" json:"presentAddressZipcode" binding:"required"`
	PresentAddressPhone         string     `form:"presentAddressPhone" json:"presentAddressPhone" binding:"required"`
	FathersGuardianFullName     string     `form:"fathersGuardianFullName" json:"fathersGuardianFullName" binding:"required"`
	FathersGuardianOccupation   string     `form:"fathersGuardianOccupation" json:"fathersGuardianOccupation" binding:"required"`
	FathersGuardianCompany      string     `form:"fathersGuardianCompany" json:"fathersGuardianCompany" binding:"required"`
	FathersGuardianPhoneNumber  string     `form:"fathersGuardianPhoneNumber" json:"fathersGuardianPhoneNumber" binding:"required"`
	FathersGuardianEmailID      string     `form:"fathersGuardianEmailID" json:"fathersGuardianEmailID" binding:"required,email"`
	MothersGuardianFullName     string     `form:"mothersGuardianFullName" json:"mothersGuardianFullName"`
	MothersGuardianOccupation   string     `form:"mothersGuardianOccupation" json:"mothersGuardianOccupation"`
	MothersGuardianCompany      string     `form:"mothersGuardianCompany" json:"mothersGuardianCompany"`
	MothersGuardianDesignation  string     `form:"mothersGuardianDesignation" json:"mothersGuardianDesignation"`
	MothersGuardianPhoneNumber  string     `form:"mothersGuardianPhoneNumber" json:"mothersGuardianPhoneNumber" `
	MothersGuardianEmailID      string     `form:"mothersGuardianEmailID" json:"mothersGuardianEmailID" `
	AccountStatus               string     `form:"accountStatus" json:"accountStatus"`
	Password                    string     `form:"password" json:"password,omitempty" binding:"required,min=8,max=15" binding:"required"`
	PrimaryPhoneVerified        bool       `form:"primaryPhoneVerified" json:"primaryPhoneVerified"`
	PrimaryEmailVerified        bool       `form:"primaryEmailVerified" json:"primaryEmailVerified"`
	DateOfJoining               time.Time  `json:"dateOfJoining" time_format="2006-12-01T21:23:34.409Z"`
	ProfilePicture              []byte     `form:"-" json:"profilePicture"`
	AccountExpiryDate           time.Time  `form:"-" json:"accountExpiryDate" time_format="2006-12-01T21:23:34.409Z"`
	AboutMeNullable             NullString `json:"-"`
	UniversityName              string     `form:"universityName" json:"universityName"`
	UniversityID                string     `form:"universityID" json:"universityID"`
	ProgramName                 string     `form:"programName" json:"programName"`
	ProgramID                   string     `form:"programID" json:"programID"`
	BranchName                  string     `form:"branchName" json:"branchName"`
	BranchID                    string     `form:"branchID" json:"branchID"`
	CollegeID                   string     `form:"collegeRollNumber" json:"collegeRollNumber" `
	CollegeEmailID              string     `form:"collegeEmailID" json:"collegeEmailID"`
	Attachment                  []byte     `form:"-" json:"attachment"`
	AttachmentName              string     `form:"-" json:"attachmentName,omitempty"`
	ReferralCode                string     `form:"-" json:"referralCode,omitempty"`
	SentforVerification         bool       `form:"-" json:"sentforVerification"`
	DateSentforVerification     string     `form:"-" json:"dateSentforVerification,omitempty"`
	Verified                    bool       `form:"-" json:"verified"`
	DateVerified                string     `form:"-" json:"dateVerified,omitempty"`
	SentbackforRevalidation     bool       `form:"-" json:"sentbackforRevalidation"`
	DateSentBackForRevalidation string     `form:"-" json:"dateSentBackForRevalidation,omitempty"`
	ValidatorRemarks            string     `form:"-" json:"validatorRemarks,omitempty"`
	VerificationType            string     `form:"-" json:"verificationType,omitempty"`
	VerifiedByStakeholderID     string     `form:"-" json:"verifiedByStakeholderID,omitempty"`
	VerifiedByEmailID           string     `form:"-" json:"verifiedByEmailID,omitempty"`
}

// StudentTTModel ...
type StudentTTModel struct {
	Name                        string `form:"name" json:"name" binding:"required"`
	Location                    string `form:"location" json:"location" binding:"required"`
	MonthAndYearOfPassing       string `form:"monthAndYearOfPassing" json:"monthAndYearOfPassing" binding:"required"`
	Board                       string `form:"schoolBoard" json:"schoolBoard" binding:"required"`
	Percentage                  string `form:"percentage" json:"percentage" binding:"required"`
	AttachmentFile              []byte `form:"attachment" json:"attachment"`
	AttachmentName              string `form:"attachmentName" json:"attachmentName"`
	SentforVerification         bool   `form:"-" json:"sentforVerification"`
	DateSentforVerification     string `form:"-" json:"dateSentforVerification,omitempty"`
	Verified                    bool   `form:"-" json:"verified"`
	DateVerified                string `form:"-" json:"dateVerified,omitempty"`
	SentbackforRevalidation     bool   `form:"-" json:"sentbackforRevalidation"`
	DateSentBackForRevalidation string `form:"-" json:"dateSentBackForRevalidation,omitempty"`
	ValidatorRemarks            string `form:"-" json:"validatorRemarks,omitempty"`
	VerificationType            string `form:"-" json:"verificationType,omitempty"`
	VerifiedByStakeholderID     string `form:"-" json:"verifiedByStakeholderID,omitempty"`
	VerifiedByEmailID           string `form:"-" json:"verifiedByEmailID,omitempty"`
}

// StudentNullableTTModel ...
type StudentNullableTTModel struct {
	Name                        NullString `form:"name" json:"name"`
	Location                    NullString `form:"location" json:"location"`
	MonthAndYearOfPassing       NullString `form:"monthAndYearOfPassing" json:"monthAndYearOfPassing"`
	Board                       NullString `form:"schoolBoard" json:"schoolBoard"`
	Percentage                  NullString `form:"percentage" json:"percentage"`
	AttachmentFile              []byte     `form:"attachment" json:"attachment"`
	AttachmentName              string     `form:"attachmentName" json:"attachmentName"`
	SentforVerification         bool       `form:"-" json:"sentforVerification"`
	DateSentforVerification     string     `form:"-" json:"dateSentforVerification,omitempty"`
	Verified                    bool       `form:"-" json:"verified"`
	DateVerified                string     `form:"-" json:"dateVerified,omitempty"`
	SentbackforRevalidation     bool       `form:"-" json:"sentbackforRevalidation"`
	DateSentBackForRevalidation string     `form:"-" json:"dateSentBackForRevalidation,omitempty"`
	ValidatorRemarks            string     `form:"-" json:"validatorRemarks,omitempty"`
	VerificationType            string     `form:"-" json:"verificationType,omitempty"`
	VerifiedByStakeholderID     string     `form:"-" json:"verifiedByStakeholderID,omitempty"`
	VerifiedByEmailID           string     `form:"-" json:"verifiedByEmailID,omitempty"`
}

// StudentGradModel ...
type StudentGradModel struct {
	UniversityStakeholderIDUniv string                 `form:"universityID" json:"universityID" binding:"required"`
	CollegeRollNumber           string                 `form:"collegeRollNumber" json:"collegeRollNumber"`
	ExpectedYearOfPassing       string                 `form:"expectedYearOfPassing" json:"expectedYearOfPassing" binding:"required"`
	ProgramID                   string                 `form:"programID" json:"programID" binding:"required"`
	ProgramName                 string                 `form:"programName" json:"programName" binding:"required"`
	BranchID                    string                 `form:"branchID" json:"branchID" binding:"required"`
	BranchName                  string                 `form:"branchName" json:"branchName" binding:"required"`
	FinalCGPA                   string                 `form:"finalCGPA" json:"finalCGPA"`
	FinalPercentage             string                 `form:"finalPercentage" json:"finalPercentage"`
	ActiveBacklogsNumber        int                    `form:"activeBacklogsNumber" json:"activeBacklogsNumber"`
	TotalNumberOfBacklogs       int                    `form:"totalNumberOfBacklogs" json:"totalNumberOfBacklogs"`
	Semesters                   []StudentSemesterModel `form:"semesters" json:"semesters" binding="dive"`
}

// StudentPGModel ...
type StudentPGModel struct {
	UniversityStakeholderIDUniv string                 `form:"universityID" json:"universityID" binding:"required"`
	CollegeRollNumber           string                 `form:"collegeRollNumber" json:"collegeRollNumber"`
	ExpectedYearOfPassing       string                 `form:"expectedYearOfPassing" json:"expectedYearOfPassing" binding:"required"`
	ProgramID                   string                 `form:"programID" json:"programID" binding:"required"`
	ProgramName                 string                 `form:"programName" json:"programName" binding:"required"`
	BranchID                    string                 `form:"branchID" json:"branchID" binding:"required"`
	BranchName                  string                 `form:"branchName" json:"branchName" binding:"required"`
	FinalCGPA                   string                 `form:"finalCGPA" json:"finalCGPA"`
	FinalPercentage             string                 `form:"finalPercentage" json:"finalPercentage"`
	Semesters                   []StudentSemesterModel `form:"semesters" json:"semesters" binding="dive"`
}

// StudentSemesterModel ...
type StudentSemesterModel struct {
	ID                          int    `form:"id" json:"id"`
	StudentStakeholderID        string `form:"studentStakeholderID" json:"studentStakeholderID"`
	UniversityStakeholderID     string `form:"universityStakeholderID" json:"universityStakeholderID"`
	IsGrad                      bool   `form:"isGrad" json:"isGrad"`
	IsPG                        bool   `form:"isPG" json:"isPG"`
	Semester                    string `form:"semester" json:"semester" binding:"required"`
	StudentCollegeRollNo        string `form:"studentCollegeRollNo" json:"studentCollegeRollNo"`
	ProgramName                 string `form:"programName" json:"programName"`
	ProgramID                   string `form:"programID" json:"programID" `
	BranchName                  string `form:"branchName" json:"branchName" `
	BranchID                    string `form:"branchID" json:"branchID"`
	CGPA                        string `form:"cgpa" json:"cgpa"`
	Percentage                  string `form:"percentage" json:"percentage"`
	AttachFile                  string `form:"attachFile" json:"attachFile" binding:"required"`
	AttachmentName              string `form:"attachFileName" json:"attachFileName" binding:"required"`
	EnabledFlag                 string `form:"enabledFlag" json:"enabledFlag"`
	CreationDate                string `form:"creationDate" json:"creationDate"`
	LastUpdatedDate             string `form:"lastUpdatedDate" json:"lastUpdatedDate"`
	SemesterCompletionDate      string `form:"semesterCompletionDate" json:"semesterCompletionDate" binding:"required"`
	SentforVerification         bool   `form:"-" json:"sentforVerification"`
	DateSentforVerification     string `form:"-" json:"dateSentforVerification,omitempty"`
	Verified                    bool   `form:"-" json:"verified"`
	DateVerified                string `form:"-" json:"dateVerified,omitempty"`
	SentbackforRevalidation     bool   `form:"-" json:"sentbackforRevalidation"`
	DateSentBackForRevalidation string `form:"-" json:"dateSentBackForRevalidation,omitempty"`
	ValidatorRemarks            string `form:"-" json:"validatorRemarks,omitempty"`
	VerificationType            string `form:"-" json:"verificationType,omitempty"`
	VerifiedByStakeholderID     string `form:"-" json:"verifiedByStakeholderID,omitempty"`
	VerifiedByEmailID           string `form:"-" json:"verifiedByEmailID,omitempty"`
}

// StudentAcademicsModelReq ...
type StudentAcademicsModelReq struct {
	StakeholderID  string           `form:"-" json:"-"`
	Tenth          StudentTTModel   `form:"tenth,omitempty" json:"tenth,omitempty"`
	Twelfth        StudentTTModel   `form:"twelfth,omitempty" json:"twelfth,omitempty"`
	Graduation     StudentGradModel `form:"graduation" json:"graduation,omitempty"`
	PostGraduation StudentPGModel   `form:"postGraduation" json:"postGraduation,omitempty"`
}

// StudentTenthAcademicsModelReq ...
type StudentTenthAcademicsModelReq struct {
	StakeholderID string         `form:"-" json:"-"`
	Tenth         StudentTTModel `form:"tenth" json:"tenth,omitempty" binding:"dive"`
}

// StudentTwelfthAcademicsModelReq ...
type StudentTwelfthAcademicsModelReq struct {
	StakeholderID string         `form:"-" json:"-"`
	Twelfth       StudentTTModel `form:"twelfth" json:"twelfth,omitempty"  binding:"dive"`
}

// StudentGradAcademicsModelReq ...
type StudentGradAcademicsModelReq struct {
	StakeholderID string           `form:"-" json:"-"`
	Graduation    StudentGradModel `form:"graduation" json:"graduation"  binding:"dive"`
}

// StudentPGAcademicsModelReq ...
type StudentPGAcademicsModelReq struct {
	StakeholderID  string         `form:"-" json:"-"`
	PostGraduation StudentPGModel `form:"postGraduation" json:"postGraduation"  binding:"dive"`
}

// StudentLangModel ...
type StudentLangModel struct {
	StakeholderID    string    `form:"-" json:"-"`
	ID               int       `form:"-" json:"id"`
	LanguageName     string    `form:"languageName" json:"languageName" binding:"required"`
	SpeakProficiency string    `form:"speakProficiency" json:"speakProficiency" binding:"required"`
	ReadProficiency  string    `form:"readProficiency" json:"readProficiency" binding:"required"`
	WriteProficiency string    `form:"writeProficiency" json:"writeProficiency" binding:"required"`
	IsMotherTongue   bool      `form:"isMotherTongue" json:"isMotherTongue"`
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
	// Result           string    `form:"result" json:"result"`
	// Description      string    `form:"description" json:"description"`
	Attachment                  []byte    `form:"attachment" json:"attachment"`
	AttachmentName              string    `form:"attachmentName" json:"attachmentName"`
	EnabledFlag                 bool      `form:"-" json:"enabledFlag"`
	CreationDate                time.Time `form:"-" json:"creationDate"`
	LastUpdatedDate             time.Time `form:"-" json:"lastUpdatedDate"`
	SentforVerification         bool      `form:"-" json:"sentforVerification"`
	DateSentforVerification     string    `form:"-" json:"dateSentforVerification,omitempty"`
	Verified                    bool      `form:"-" json:"verified"`
	DateVerified                string    `form:"-" json:"dateVerified,omitempty"`
	SentbackforRevalidation     bool      `form:"-" json:"sentbackforRevalidation"`
	DateSentBackForRevalidation string    `form:"-" json:"dateSentBackForRevalidation,omitempty"`
	ValidatorRemarks            string    `form:"-" json:"validatorRemarks,omitempty"`
	VerificationType            string    `form:"-" json:"verificationType,omitempty"`
	VerifiedByStakeholderID     string    `form:"-" json:"verifiedByStakeholderID,omitempty"`
	VerifiedByEmailID           string    `form:"-" json:"verifiedByEmailID,omitempty"`
}

// StudentAllCertsModel ...
type StudentAllCertsModel struct {
	StakeholderID  string              `form:"-" json:"-"`
	Certifications []StudentCertsModel `form:"certifications" json:"certifications" binding:"dive"`
}

// StudentAssessmentModel ...
type StudentAssessmentModel struct {
	StakeholderID               string    `form:"-" json:"-"`
	ID                          int       `form:"-" json:"id"`
	Name                        string    `form:"name" json:"name" binding:"required"`
	Score                       string    `form:"score" json:"score" binding:"required"`
	IssuingAuthority            string    `form:"issuingAuthority" json:"issuingAuthority" binding:"required"`
	AssessmentDate              time.Time `form:"assessmentDate" json:"assessmentDate" binding:"required"`
	Description                 string    `form:"description" json:"description"`
	Attachment                  []byte    `form:"attachment" json:"attachment"`
	AttachmentName              string    `form:"attachmentName" json:"attachmentName"`
	EnabledFlag                 bool      `form:"-" json:"enabledFlag"`
	CreationDate                time.Time `form:"-" json:"creationDate"`
	LastUpdatedDate             time.Time `form:"-" json:"lastUpdatedDate"`
	SentforVerification         bool      `form:"-" json:"sentforVerification"`
	DateSentforVerification     string    `form:"-" json:"dateSentforVerification,omitempty"`
	Verified                    bool      `form:"-" json:"verified"`
	DateVerified                string    `form:"-" json:"dateVerified,omitempty"`
	SentbackforRevalidation     bool      `form:"-" json:"sentbackforRevalidation"`
	DateSentBackForRevalidation string    `form:"-" json:"dateSentBackForRevalidation,omitempty"`
	ValidatorRemarks            string    `form:"-" json:"validatorRemarks,omitempty"`
	VerificationType            string    `form:"-" json:"verificationType,omitempty"`
	VerifiedByStakeholderID     string    `form:"-" json:"verifiedByStakeholderID,omitempty"`
	VerifiedByEmailID           string    `form:"-" json:"verifiedByEmailID,omitempty"`
}

// StudentAllAssessmentModel ...
type StudentAllAssessmentModel struct {
	StakeholderID string                   `form:"-" json:"-"`
	Assessments   []StudentAssessmentModel `form:"assessments" json:"assessments" binding:"dive"`
}

// StudentInternshipModel ...
type StudentInternshipModel struct {
	StakeholderID               string    `form:"-" json:"-"`
	ID                          int       `form:"-" json:"id"`
	Name                        string    `form:"name" json:"name" binding:"required"`
	OrganizationName            string    `form:"organizationName" json:"organizationName" binding:"required"`
	OrganizationCity            string    `form:"organizationCity" json:"organizationCity" binding:"required"`
	FieldOfWork                 string    `form:"fieldOfWork" json:"fieldOfWork" binding:"required"`
	StartDate                   time.Time `form:"startDate" json:"startDate" binding:"required" time_format="2006-12-01T21:23:34.409Z"`
	EndDate                     time.Time `form:"endDate" json:"endDate"  binding:"required" time_format="2006-12-01T21:23:34.409Z"`
	Description                 string    `form:"description" json:"description"`
	Attachment                  []byte    `form:"attachment" json:"attachment"`
	AttachmentName              string    `form:"attachmentName" json:"attachmentName"`
	EnabledFlag                 bool      `form:"-" json:"enabledFlag"`
	CreationDate                time.Time `form:"-" json:"creationDate"`
	LastUpdatedDate             time.Time `form:"-" json:"lastUpdatedDate"`
	SentforVerification         bool      `form:"-" json:"sentforVerification"`
	DateSentforVerification     string    `form:"-" json:"dateSentforVerification,omitempty"`
	Verified                    bool      `form:"-" json:"verified"`
	DateVerified                string    `form:"-" json:"dateVerified,omitempty"`
	SentbackforRevalidation     bool      `form:"-" json:"sentbackforRevalidation"`
	DateSentBackForRevalidation string    `form:"-" json:"dateSentBackForRevalidation,omitempty"`
	ValidatorRemarks            string    `form:"-" json:"validatorRemarks,omitempty"`
	VerificationType            string    `form:"-" json:"verificationType,omitempty"`
	VerifiedByStakeholderID     string    `form:"-" json:"verifiedByStakeholderID,omitempty"`
	VerifiedByEmailID           string    `form:"-" json:"verifiedByEmailID,omitempty"`
}

// StudentAllInternshipModel ...
type StudentAllInternshipModel struct {
	StakeholderID string                   `form:"-" json:"-"`
	Internships   []StudentInternshipModel `form:"internships" json:"internships" binding:"dive"`
}

// StudentAwardsModel ...
type StudentAwardsModel struct {
	StakeholderID               string    `form:"-" json:"-"`
	ID                          int       `form:"-" json:"id"`
	RecognitionName             string    `form:"recognitionName" json:"recognitionName" binding:"required"`
	RecognitionDate             time.Time `form:"recognitionDate" json:"recognitionDate" binding:"required" time_format="2006-12-01T21:23:34.409Z"`
	IssuingAuthority            string    `form:"issuingAuthority" json:"issuingAuthority" binding:"required"`
	Attachment                  []byte    `form:"attachment" json:"attachment"`
	AttachmentName              string    `form:"attachmentName" json:"attachmentName"`
	EnabledFlag                 bool      `form:"-" json:"enabledFlag"`
	CreationDate                time.Time `form:"-" json:"creationDate"`
	LastUpdatedDate             time.Time `form:"-" json:"lastUpdatedDate"`
	SentforVerification         bool      `form:"-" json:"sentforVerification"`
	DateSentforVerification     string    `form:"-" json:"dateSentforVerification,omitempty"`
	Verified                    bool      `form:"-" json:"verified"`
	DateVerified                string    `form:"-" json:"dateVerified,omitempty"`
	SentbackforRevalidation     bool      `form:"-" json:"sentbackforRevalidation"`
	DateSentBackForRevalidation string    `form:"-" json:"dateSentBackForRevalidation,omitempty"`
	ValidatorRemarks            string    `form:"-" json:"validatorRemarks,omitempty"`
	VerificationType            string    `form:"-" json:"verificationType,omitempty"`
	VerifiedByStakeholderID     string    `form:"-" json:"verifiedByStakeholderID,omitempty"`
	VerifiedByEmailID           string    `form:"-" json:"verifiedByEmailID,omitempty"`
}

// StudentAllAwardsModel ...
type StudentAllAwardsModel struct {
	StakeholderID string               `form:"-" json:"-"`
	Awards        []StudentAwardsModel `form:"awards" json:"awards" binding:"dive"`
}

// StudentCompetitionModel ...
type StudentCompetitionModel struct {
	StakeholderID               string    `form:"-" json:"-"`
	ID                          int       `form:"-" json:"id"`
	Name                        string    `form:"name" json:"name" binding:"required"`
	Date                        time.Time `form:"date" json:"date" binding:"required" time_format="2006-12-01T21:23:34.409Z"`
	Rank                        string    `form:"rank" json:"rank" binding:"required"`
	Attachment                  []byte    `form:"attachment" json:"attachment"`
	AttachmentName              string    `form:"attachmentName" json:"attachmentName"`
	EnabledFlag                 bool      `form:"-" json:"enabledFlag"`
	CreationDate                time.Time `form:"-" json:"creationDate"`
	LastUpdatedDate             time.Time `form:"-" json:"lastUpdatedDate"`
	SentforVerification         bool      `form:"-" json:"sentforVerification"`
	DateSentforVerification     string    `form:"-" json:"dateSentforVerification,omitempty"`
	Verified                    bool      `form:"-" json:"verified"`
	DateVerified                string    `form:"-" json:"dateVerified,omitempty"`
	SentbackforRevalidation     bool      `form:"-" json:"sentbackforRevalidation"`
	DateSentBackForRevalidation string    `form:"-" json:"dateSentBackForRevalidation,omitempty"`
	ValidatorRemarks            string    `form:"-" json:"validatorRemarks,omitempty"`
	VerificationType            string    `form:"-" json:"verificationType,omitempty"`
	VerifiedByStakeholderID     string    `form:"-" json:"verifiedByStakeholderID,omitempty"`
	VerifiedByEmailID           string    `form:"-" json:"verifiedByEmailID,omitempty"`
}

// StudentAllCompetitionModel ...
type StudentAllCompetitionModel struct {
	StakeholderID string                    `form:"-" json:"-"`
	Competitions  []StudentCompetitionModel `form:"competitions" json:"competitions" binding:"dive"`
}

// StudentEventsModel ...
type StudentEventsModel struct {
	StakeholderID               string    `form:"-" json:"-"`
	ID                          int       `form:"-" json:"id"`
	Name                        string    `form:"name" json:"name" binding:"required"`
	Date                        time.Time `form:"date" json:"date" binding:"required" time_format="2006-12-01T21:23:34.409Z"`
	Attachment                  []byte    `form:"attachment" json:"attachment"`
	AttachmentName              string    `form:"attachmentName" json:"attachmentName"`
	OrganizedBy                 string    `form:"organizedBy" json:"organizedBy" binding:"required"`
	OrganizedByEmail            string    `form:"organizedByEmail" json:"organizedByEmail" binding:"required"`
	OrganizedByPhone            string    `form:"organizedByPhone" json:"organizedByPhone"  binding:"required"`
	EventType                   string    `form:"eventType" json:"eventType"  binding:"required"`
	EventTypeOther              string    `form:"eventTypeOther" json:"eventTypeOther"`
	EventResult                 string    `form:"eventResult" json:"eventResult"`
	EventResultOther            string    `form:"eventResultOther" json:"eventResultOther"`
	EnabledFlag                 bool      `form:"-" json:"enabledFlag"`
	CreationDate                time.Time `form:"-" json:"creationDate"`
	LastUpdatedDate             time.Time `form:"-" json:"lastUpdatedDate"`
	SentforVerification         bool      `form:"-" json:"sentforVerification"`
	DateSentforVerification     string    `form:"-" json:"dateSentforVerification,omitempty"`
	Verified                    bool      `form:"-" json:"verified"`
	DateVerified                string    `form:"-" json:"dateVerified,omitempty"`
	SentbackforRevalidation     bool      `form:"-" json:"sentbackforRevalidation"`
	DateSentBackForRevalidation string    `form:"-" json:"dateSentBackForRevalidation,omitempty"`
	ValidatorRemarks            string    `form:"-" json:"validatorRemarks,omitempty"`
	VerificationType            string    `form:"-" json:"verificationType,omitempty"`
	VerifiedByStakeholderID     string    `form:"-" json:"verifiedByStakeholderID,omitempty"`
	VerifiedByEmailID           string    `form:"-" json:"verifiedByEmailID,omitempty"`
}

// StudentAllEventsModel ...
type StudentAllEventsModel struct {
	StakeholderID string               `form:"-" json:"-"`
	Events        []StudentEventsModel `form:"events" json:"events" binding:"dive"`
}

// StudentExtraCurricularModel ...
type StudentExtraCurricularModel struct {
	StakeholderID               string    `form:"-" json:"-"`
	ID                          int       `form:"-" json:"id"`
	Name                        string    `form:"name" json:"name" binding:"required"`
	Attachment                  []byte    `form:"attachment" json:"attachment"`
	AttachmentName              string    `form:"attachmentName" json:"attachmentName"`
	EnabledFlag                 bool      `form:"-" json:"enabledFlag"`
	CreationDate                time.Time `form:"-" json:"creationDate"`
	LastUpdatedDate             time.Time `form:"-" json:"lastUpdatedDate"`
	SentforVerification         bool      `form:"-" json:"sentforVerification"`
	DateSentforVerification     string    `form:"-" json:"dateSentforVerification,omitempty"`
	Verified                    bool      `form:"-" json:"verified"`
	DateVerified                string    `form:"-" json:"dateVerified,omitempty"`
	SentbackforRevalidation     bool      `form:"-" json:"sentbackforRevalidation"`
	DateSentBackForRevalidation string    `form:"-" json:"dateSentBackForRevalidation,omitempty"`
	ValidatorRemarks            string    `form:"-" json:"validatorRemarks,omitempty"`
	VerificationType            string    `form:"-" json:"verificationType,omitempty"`
	VerifiedByStakeholderID     string    `form:"-" json:"verifiedByStakeholderID,omitempty"`
	VerifiedByEmailID           string    `form:"-" json:"verifiedByEmailID,omitempty"`
}

// StudentAllExtraCurricularModel ...
type StudentAllExtraCurricularModel struct {
	StakeholderID   string                        `form:"-" json:"-"`
	ExtraCurricular []StudentExtraCurricularModel `form:"extraCurricular" json:"extraCurricular" binding:"dive"`
}

// StudentPatentsModel ...
type StudentPatentsModel struct {
	StakeholderID               string    `form:"-" json:"-"`
	ID                          int       `form:"-" json:"id"`
	Name                        string    `form:"name" json:"name" binding:"required"`
	PatentType                  string    `form:"patentType" json:"patentType" binding:"required"`
	PatentNumber                string    `form:"patentNumber" json:"patentNumber" binding:"required"`
	PatentStatus                string    `form:"patentStatus" json:"patentStatus" binding:"required"`
	Attachment                  []byte    `form:"attachment" json:"attachment"`
	AttachmentName              string    `form:"attachmentName" json:"attachmentName"`
	EnabledFlag                 bool      `form:"-" json:"enabledFlag"`
	CreationDate                time.Time `form:"-" json:"creationDate"`
	LastUpdatedDate             time.Time `form:"-" json:"lastUpdatedDate"`
	SentforVerification         bool      `form:"-" json:"sentforVerification"`
	DateSentforVerification     string    `form:"-" json:"dateSentforVerification,omitempty"`
	Verified                    bool      `form:"-" json:"verified"`
	DateVerified                string    `form:"-" json:"dateVerified,omitempty"`
	SentbackforRevalidation     bool      `form:"-" json:"sentbackforRevalidation"`
	DateSentBackForRevalidation string    `form:"-" json:"dateSentBackForRevalidation,omitempty"`
	ValidatorRemarks            string    `form:"-" json:"validatorRemarks,omitempty"`
	VerificationType            string    `form:"-" json:"verificationType,omitempty"`
	VerifiedByStakeholderID     string    `form:"-" json:"verifiedByStakeholderID,omitempty"`
	VerifiedByEmailID           string    `form:"-" json:"verifiedByEmailID,omitempty"`
}

// StudentAllPatentsModel ...
type StudentAllPatentsModel struct {
	StakeholderID string                `form:"-" json:"-"`
	Patents       []StudentPatentsModel `form:"patents" json:"patents" binding:"dive"`
}

// StudentProjectsModel ...
type StudentProjectsModel struct {
	StakeholderID               string    `form:"-" json:"-"`
	ID                          int       `form:"-" json:"id"`
	Name                        string    `form:"name" json:"name" binding:"required"`
	ProjectAbstract             string    `form:"projectAbstract" json:"projectAbstract" binding:"required"`
	GuideName                   string    `form:"guideName" json:"guideName"`
	GuideEmail                  string    `form:"guideEmail" json:"guideEmail"`
	StartDate                   time.Time `form:"startDate" json:"startDate" binding:"required" time_format="2006-12-01T21:23:34.409Z"`
	EndDate                     time.Time `form:"endDate" json:"endDate,omitempty" time_format="2006-12-01T21:23:34.409Z"`
	Attachment                  []byte    `form:"attachment" json:"attachment"`
	AttachmentName              string    `form:"attachmentName" json:"attachmentName"`
	EnabledFlag                 bool      `form:"-" json:"enabledFlag"`
	CreationDate                time.Time `form:"-" json:"creationDate"`
	LastUpdatedDate             time.Time `form:"-" json:"lastUpdatedDate"`
	SentforVerification         bool      `form:"-" json:"sentforVerification"`
	DateSentforVerification     string    `form:"-" json:"dateSentforVerification,omitempty"`
	Verified                    bool      `form:"-" json:"verified"`
	DateVerified                string    `form:"-" json:"dateVerified,omitempty"`
	SentbackforRevalidation     bool      `form:"-" json:"sentbackforRevalidation"`
	DateSentBackForRevalidation string    `form:"-" json:"dateSentBackForRevalidation,omitempty"`
	ValidatorRemarks            string    `form:"-" json:"validatorRemarks,omitempty"`
	VerificationType            string    `form:"-" json:"verificationType,omitempty"`
	VerifiedByStakeholderID     string    `form:"-" json:"verifiedByStakeholderID,omitempty"`
	VerifiedByEmailID           string    `form:"-" json:"verifiedByEmailID,omitempty"`
}

// StudentAllProjectsModel ...
type StudentAllProjectsModel struct {
	StakeholderID string                 `form:"-" json:"-"`
	Projects      []StudentProjectsModel `form:"projects" json:"projects" binding:"dive"`
}

// StudentPublicationsModel ...
type StudentPublicationsModel struct {
	StakeholderID               string    `form:"-" json:"-"`
	ID                          int       `form:"-" json:"id"`
	Name                        string    `form:"name" json:"name" binding:"required"`
	PublishingAuthority         string    `form:"publishingAuthority" json:"publishingAuthority" binding:"required"`
	GuideName                   string    `form:"guideName" json:"guideName" binding:"required"`
	GuideEmail                  string    `form:"guideEmail" json:"guideEmail" binding:"required"`
	StartDate                   time.Time `form:"startDate" json:"startDate" binding:"required" time_format="2006-12-01T21:23:34.409Z"`
	EndDate                     time.Time `form:"endDate" json:"endDate" binding:"required" time_format="2006-12-01T21:23:34.409Z"`
	Attachment                  []byte    `form:"attachment" json:"attachment"`
	AttachmentName              string    `form:"attachmentName" json:"attachmentName"`
	EnabledFlag                 bool      `form:"-" json:"enabledFlag"`
	CreationDate                time.Time `form:"-" json:"creationDate"`
	LastUpdatedDate             time.Time `form:"-" json:"lastUpdatedDate"`
	SentforVerification         bool      `form:"-" json:"sentforVerification"`
	DateSentforVerification     string    `form:"-" json:"dateSentforVerification,omitempty"`
	Verified                    bool      `form:"-" json:"verified"`
	DateVerified                string    `form:"-" json:"dateVerified,omitempty"`
	SentbackforRevalidation     bool      `form:"-" json:"sentbackforRevalidation"`
	DateSentBackForRevalidation string    `form:"-" json:"dateSentBackForRevalidation,omitempty"`
	ValidatorRemarks            string    `form:"-" json:"validatorRemarks,omitempty"`
	VerificationType            string    `form:"-" json:"verificationType,omitempty"`
	VerifiedByStakeholderID     string    `form:"-" json:"verifiedByStakeholderID,omitempty"`
	VerifiedByEmailID           string    `form:"-" json:"verifiedByEmailID,omitempty"`
}

// StudentAllPublicationsModel ...
type StudentAllPublicationsModel struct {
	StakeholderID string                     `form:"-" json:"-"`
	Publications  []StudentPublicationsModel `form:"publications" json:"publications" binding:"dive"`
}

// StudentScholarshipsModel ...
type StudentScholarshipsModel struct {
	StakeholderID               string    `form:"-" json:"-"`
	ID                          int       `form:"-" json:"id"`
	Name                        string    `form:"name" json:"name" binding:"required"`
	ScholarshipIssuedBy         string    `form:"scholarshipIssuedBy" json:"publishingAuthority" binding:"required"`
	ScholarshipDate             time.Time `form:"scholarshipDate" json:"scholarshipDate" binding:"required" time_format="2006-12-01T21:23:34.409Z"`
	Attachment                  []byte    `form:"attachment" json:"attachment"`
	AttachmentName              string    `form:"attachmentName" json:"attachmentName"`
	EnabledFlag                 bool      `form:"-" json:"enabledFlag"`
	CreationDate                time.Time `form:"-" json:"creationDate"`
	LastUpdatedDate             time.Time `form:"-" json:"lastUpdatedDate"`
	SentforVerification         bool      `form:"-" json:"sentforVerification"`
	DateSentforVerification     string    `form:"-" json:"dateSentforVerification,omitempty"`
	Verified                    bool      `form:"-" json:"verified"`
	DateVerified                string    `form:"-" json:"dateVerified,omitempty"`
	SentbackforRevalidation     bool      `form:"-" json:"sentbackforRevalidation"`
	DateSentBackForRevalidation string    `form:"-" json:"dateSentBackForRevalidation,omitempty"`
	ValidatorRemarks            string    `form:"-" json:"validatorRemarks,omitempty"`
	VerificationType            string    `form:"-" json:"verificationType,omitempty"`
	VerifiedByStakeholderID     string    `form:"-" json:"verifiedByStakeholderID,omitempty"`
	VerifiedByEmailID           string    `form:"-" json:"verifiedByEmailID,omitempty"`
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
	StakeholderID               string    `form:"-" json:"-"`
	ID                          int       `form:"-" json:"id"`
	Name                        string    `form:"name" json:"name" binding:"required"`
	TestScoreDate               time.Time `form:"testScoreDate" json:"testScoreDate" binding:"required" time_format="2006-12-01T21:23:34.409Z"`
	TestScore                   string    `form:"testScore" json:"testScore" binding:"required"`
	TestScoreTotal              string    `form:"testScoreTotal" json:"testScoreTotal" binding:"required"`
	Attachment                  []byte    `form:"attachment" json:"attachment"`
	AttachmentName              string    `form:"attachmentName" json:"attachmentName"`
	EnabledFlag                 bool      `form:"-" json:"enabledFlag"`
	CreationDate                time.Time `form:"-" json:"creationDate"`
	LastUpdatedDate             time.Time `form:"-" json:"lastUpdatedDate"`
	SentforVerification         bool      `form:"-" json:"sentforVerification"`
	DateSentforVerification     string    `form:"-" json:"dateSentforVerification,omitempty"`
	Verified                    bool      `form:"-" json:"verified"`
	DateVerified                string    `form:"-" json:"dateVerified,omitempty"`
	SentbackforRevalidation     bool      `form:"-" json:"sentbackforRevalidation"`
	DateSentBackForRevalidation string    `form:"-" json:"dateSentBackForRevalidation,omitempty"`
	ValidatorRemarks            string    `form:"-" json:"validatorRemarks,omitempty"`
	VerificationType            string    `form:"-" json:"verificationType,omitempty"`
	VerifiedByStakeholderID     string    `form:"-" json:"verifiedByStakeholderID,omitempty"`
	VerifiedByEmailID           string    `form:"-" json:"verifiedByEmailID,omitempty"`
}

// StudentAllTestScoresModel ...
type StudentAllTestScoresModel struct {
	StakeholderID string                   `form:"-" json:"-"`
	TestScores    []StudentTestScoresModel `form:"testScores" json:"testScores" binding:"dive"`
}

// StudentVolunteerExperienceModel ...
type StudentVolunteerExperienceModel struct {
	StakeholderID               string    `form:"-" json:"-"`
	ID                          int       `form:"-" json:"id"`
	Name                        string    `form:"name" json:"name" binding:"required"`
	Organisation                string    `form:"organisation" json:"organisation" binding:"required"`
	Location                    string    `form:"location" json:"location" binding:"required"`
	StartDate                   time.Time `form:"startDate" json:"startDate" binding:"required" time_format="2006-12-01T21:23:34.409Z"`
	EndDate                     time.Time `form:"endDate" json:"endDate" binding:"required" time_format="2006-12-01T21:23:34.409Z"`
	Attachment                  []byte    `form:"attachment" json:"attachment"`
	AttachmentName              string    `form:"attachmentName" json:"attachmentName"`
	EnabledFlag                 bool      `form:"-" json:"enabledFlag"`
	CreationDate                time.Time `form:"-" json:"creationDate"`
	LastUpdatedDate             time.Time `form:"-" json:"lastUpdatedDate"`
	SentforVerification         bool      `form:"-" json:"sentforVerification"`
	DateSentforVerification     string    `form:"-" json:"dateSentforVerification,omitempty"`
	Verified                    bool      `form:"-" json:"verified"`
	DateVerified                string    `form:"-" json:"dateVerified,omitempty"`
	SentbackforRevalidation     bool      `form:"-" json:"sentbackforRevalidation"`
	DateSentBackForRevalidation string    `form:"-" json:"dateSentBackForRevalidation,omitempty"`
	ValidatorRemarks            string    `form:"-" json:"validatorRemarks,omitempty"`
	VerificationType            string    `form:"-" json:"verificationType,omitempty"`
	VerifiedByStakeholderID     string    `form:"-" json:"verifiedByStakeholderID,omitempty"`
	VerifiedByEmailID           string    `form:"-" json:"verifiedByEmailID,omitempty"`
}

// StudentAllVolunteerExperienceModel ...
type StudentAllVolunteerExperienceModel struct {
	StakeholderID       string                            `form:"-" json:"-"`
	VolunteerExperience []StudentVolunteerExperienceModel `form:"volunteerExperience" json:"VolunteerExperience" binding:"dive"`
}

// StudentContactInfoModel ...
type StudentContactInfoModel struct {
	StakeholderID               string    `form:"stakeholderID" json:"stakeholderID,omitempty"`
	FirstName                   string    `form:"firstName" json:"firstName" binding:"required"`
	MiddleName                  string    `form:"middleName" json:"middleName"`
	LastName                    string    `form:"lastName" json:"lastName" binding:"required"`
	PersonalEmail               string    `form:"personalEmail" json:"personalEmail"`
	CollegeEmail                string    `form:"collegeEmail" json:"collegeEmail" binding:"required,email"`
	PhoneNumber                 string    `form:"phoneNumber" json:"phoneNumber" binding:"required"`
	AlternatePhoneNumber        string    `form:"alternatePhoneNumber" json:"alternatePhoneNumber"`
	CollegeID                   string    `form:"collegeID" json:"collegeID" binding:"required"`
	Gender                      string    `form:"gender" json:"gender" binding:"required"`
	DateOfBirth                 string    `form:"dateOfBirth" json:"dateOfBirth" binding:"required"`
	AadharNumber                string    `form:"aadharNumber" json:"aadharNumber" binding:"required"`
	PermanentAddressLine1       string    `form:"permanentAddressLine1" json:"permanentAddressLine1" binding:"required"`
	PermanentAddressLine2       string    `form:"permanentAddressLine2" json:"permanentAddressLine2" binding:"required"`
	PermanentAddressLine3       string    `form:"permanentAddressLine3" json:"permanentAddressLine3"`
	PermanentAddressCountry     string    `form:"permanentAddressCountry" json:"permanentAddressCountry" binding:"required"`
	PermanentAddressState       string    `form:"permanentAddressState" json:"permanentAddressState" binding:"required"`
	PermanentAddressCity        string    `form:"permanentAddressCity" json:"permanentAddressCity" binding:"required"`
	PermanentAddressDistrict    string    `form:"permanentAddressDistrict" json:"permanentAddressDistrict" binding:"required"`
	PermanentAddressZipcode     string    `form:"permanentAddressZipcode" json:"permanentAddressZipcode" binding:"required"`
	PermanentAddressPhone       string    `form:"permanentAddressPhone" json:"permanentAddressPhone" binding:"required"`
	PresentAddressLine1         string    `form:"presentAddressLine1" json:"presentAddressLine1" binding:"required"`
	PresentAddressLine2         string    `form:"presentAddressLine2" json:"presentAddressLine2" binding:"required"`
	PresentAddressLine3         string    `form:"presentAddressLine3" json:"presentAddressLine3"`
	PresentAddressCountry       string    `form:"presentAddressCountry" json:"presentAddressCountry" binding:"required"`
	PresentAddressState         string    `form:"presentAddressState" json:"presentAddressState" binding:"required"`
	PresentAddressCity          string    `form:"presentAddressCity" json:"presentAddressCity" binding:"required"`
	PresentAddressDistrict      string    `form:"presentAddressDistrict" json:"presentAddressDistrict" binding:"required"`
	PresentAddressZipcode       string    `form:"presentAddressZipcode" json:"presentAddressZipcode" binding:"required"`
	PresentAddressPhone         string    `form:"presentAddressPhone" json:"presentAddressPhone" binding:"required"`
	AboutMe                     string    `form:"aboutMe" json:"aboutMe"`
	UniversityName              string    `form:"universityName" json:"universityName"`
	UniversityID                string    `form:"universityID" json:"-"`
	DateOfJoining               time.Time `json:"dateOfJoining,omitempty" time_format="2006-12-01T21:23:34.409Z"`
	ProfilePicture              []byte    `form:"-" json:"profilePicture"`
	SentforVerification         bool      `form:"-" json:"sentforVerification"`
	DateSentforVerification     string    `form:"-" json:"dateSentforVerification,omitempty"`
	Verified                    bool      `form:"-" json:"verified"`
	DateVerified                string    `form:"-" json:"dateVerified,omitempty"`
	SentbackforRevalidation     bool      `form:"-" json:"sentbackforRevalidation"`
	DateSentBackForRevalidation string    `form:"-" json:"dateSentBackForRevalidation,omitempty"`
	ValidatorRemarks            string    `form:"-" json:"validatorRemarks,omitempty"`
	VerificationType            string    `form:"-" json:"verificationType,omitempty"`
	VerifiedByStakeholderID     string    `form:"-" json:"verifiedByStakeholderID,omitempty"`
	VerifiedByEmailID           string    `form:"-" json:"verifiedByEmailID,omitempty"`
}

// StudentCompleteProfileDataModel ...
type StudentCompleteProfileDataModel struct {
	Profile                  StudentMasterDb                   `form:"-" json:"profile"`
	ContactInfo              StudentContactInfoModel           `form:"contactInfo" json:"contactInfo"`
	Academics                StudentAcademicsModelReq          `form:"academics" json:"academics"`
	Languages                StudentAllLanguagesModel          `form:"-" json:"-"`
	Certifications           StudentAllCertsModel              `form:"-" json:"-"`
	Assessments              StudentAllAssessmentModel         `form:"-" json:"-"`
	Internships              StudentAllInternshipModel         `form:"-" json:"-"`
	LanguagesArray           []StudentLangModel                `form:"languages" json:"languages"`
	CertificationsArray      []StudentCertsModel               `form:"certifications" json:"certifications"`
	AssessmentsArray         []StudentAssessmentModel          `form:"assessments" json:"assessments"`
	InternshipsArray         []StudentInternshipModel          `form:"internships" json:"internships"`
	AwardsArray              []StudentAwardsModel              `form:"awards" json:"awards"`
	EventsArray              []StudentEventsModel              `form:"events" json:"events"`
	CompetitionsArray        []StudentCompetitionModel         `form:"competitions" json:"competitions,omitempty"`
	ExtraCurricularArray     []StudentExtraCurricularModel     `form:"extraCurricular" json:"extraCurricular"`
	PatentsArray             []StudentPatentsModel             `form:"patents" json:"patents"`
	ProjectsArray            []StudentProjectsModel            `form:"projects" json:"projects"`
	PublicationsArray        []StudentPublicationsModel        `form:"publications" json:"publications"`
	ScholarshipsArray        []StudentScholarshipsModel        `form:"scholarships" json:"scholarships"`
	SocialAccountArray       []StudentSocialAccountModel       `form:"socialAccount" json:"socialAccount"`
	TechSkillsArray          []StudentTechSkillsModel          `form:"techSkills" json:"techSkills"`
	TestScoresArray          []StudentTestScoresModel          `form:"testScores" json:"testScores"`
	VolunteerExperienceArray []StudentVolunteerExperienceModel `form:"volunteerExperience" json:"volunteerExperience"`
}

// StudentProfileVerificationDataModel ...
type StudentProfileVerificationDataModel struct {
	Profile     StudentMasterDb          `form:"-" json:"-"`
	ContactInfo StudentContactInfoModel  `form:"contactInfo" json:"contactInfo"`
	Academics   StudentAcademicsModelReq `form:"academics" json:"academics"`
}

// StudentAllProfiles ...
type StudentAllProfiles struct {
	StudentPlatformID string `json:"studentPlatformID"`
	StudentFirstName  string `json:"studentFirstName"`
	StudentMiddleName string `json:"studentMiddleName"`
	StudentLastName   string `json:"studentLastName"`
	UniversityID      string `json:"UniversityID`
	Program           string `json:"program"`
	BranchName        string `json:"branch"`
	Year              string `json:"year"`
}
