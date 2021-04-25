// Package models ...
package models

// LutCorporateType ...
type LutCorporateType struct {
	CodeDescription string `json:"codeDescription"`
	Code            string `json:"code"`
	CharCode        string `json:"charCode,omitempty"`
}

// LutSkillsMaster ...
type LutSkillsMaster struct {
	SkillID string `json:"skillID`
	Skill   string `json:"skill"`
}

// LutProgramMaster ...
type LutProgramMaster struct {
	ProgramID string `json:"programID`
	Program   string `json:"program"`
}

// LutDepartmentMaster ...
type LutDepartmentMaster struct {
	DepartmentID string `json:"departmentID"`
	ProgramID    string `json:"programID`
	Department   string `json:"department"`
}

// LutResponse ...
type LutResponse struct {
	CorporateTypes     []LutCorporateType    `json:"corporateTypes,omitempty"`
	CorporateCategory  []LutCorporateType    `json:"corporateCategory,omitempty"`
	CorporateIndustry  []LutCorporateType    `json:"corporateIndustry,omitempty"`
	UniversityCategory []LutCorporateType    `json:"universityCategory,omitempty"`
	Skills             []LutSkillsMaster     `json:"skills,omitempty"`
	Programs           []LutProgramMaster    `json:"programs,omitempty"`
	Departments        []LutDepartmentMaster `json:"departments,omitempty"`
}
