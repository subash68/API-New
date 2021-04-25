// Package models ...
package models

// DbModelKeys ...
var DbModelKeys map[string]string

// GetDbKey ...
func GetDbKey(apiKey string) (string, bool) {
	DbModelKeys = map[string]string{
		"hiringCriteriaName":      "HiringCriteria_Name",
		"programID":               "Program_ID",
		"courseID":                "Branch_ID",
		"cutOffCategory":          "MinimumCutoffCategory",
		"cutOff":                  "MinimumCutoff",
		"eduGapsSchoolAllowed":    "EduGaps_School_Allowed",
		"eduGaps11N12Allowed":     "EduGaps_11N12_Allowed",
		"eduGapsGradAllowed":      "EduGaps_Grad_Allowed",
		"eduGapsPGAllowed":        "EduGaps_PG_Allowed",
		"allowActiveBacklogs":     "ActiveBacklogsAllowed",
		"numberOfAllowedBacklogs": "TotalNumberOfBacklogsAllowed",
		"yearOfPassing":           "YearOfPassing",
		"remarks":                 "Remarks",
		"status":                  "Job_Status",
		"jobID":                   "Job_ID",
		"jobName":                 "JobName",
		"skillID":                 "Skill_ID",
		"skill":                   "SkillName",
		"noOfPositions":           "NoOfPositions",
		"location":                "Location",
		"salaryType":              "SalaryRange",
		"dateOfHiring":            "DateOfHiring",
	}
	key := DbModelKeys[apiKey]
	if key == "" {
		return "", false
	}
	return key, true
}
