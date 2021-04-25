// Package models ...
package models

import (
	"fmt"
)

// GetData ...
func (ld *LutResponse) GetData(reqLuts []string) <-chan DbModelError {
	job := make(chan DbModelError, 1)
	successResp := map[string]string{}
	var err error
	var customError DbModelError
	if CheckPing(&customError); customError.Err != nil {
		job <- customError
		return job
	}
	for _, v := range reqLuts {
		fmt.Printf("\n checking : %v\n", v)
		switch v {
		case "corporateType":
			ld.CorporateTypes, err = getCrpType("GET_LUT_CRP_TPY")
			break
		case "corporateCategory":
			ld.CorporateCategory, err = getCrpType("GET_LUT_CRP_CAT")
			break
		case "corporateIndustry":
			ld.CorporateIndustry, err = getCrpInd("GET_LUT_CRP_IND")
			break
		case "universityCategory":
			ld.UniversityCategory, err = getCrpType("GET_LUT_UNV_CAT")
			break
		case "skills":
			ld.Skills, err = getSkills()
			break
		case "programs":
			ld.Programs, err = getPrograms()
			break
		case "branches":
			ld.Departments, err = getDepartments()
			break
		default:
			customError.ErrTyp = "500"
			customError.Err = fmt.Errorf("Invalid search query " + v)
			customError.ErrCode = "S4PRFLUT"
			customError.SuccessResp = successResp
			job <- customError
			return job
		}
		if err != nil {
			customError.ErrTyp = "500"
			customError.Err = fmt.Errorf("Failed to get details  %v", err.Error())
			customError.ErrCode = "S4PRFLUT"
			customError.SuccessResp = successResp
			job <- customError
			return job
		}
	}
	customError.ErrTyp = "000"
	job <- customError
	return job
}

func getCrpType(spName string) (lct []LutCorporateType, err error) {
	sp, _ := RetriveSP(spName)
	rows, err := Db.Query(sp)
	if err != nil {
		return lct, fmt.Errorf("Cannot get the Rows %v", err.Error())
	}
	defer rows.Close()
	for rows.Next() {
		var row LutCorporateType
		err = rows.Scan(&row.CodeDescription, &row.Code, &row.CharCode)
		if err != nil {
			return lct, fmt.Errorf("Cannot read the Rows %v", err.Error())
		}
		lct = append(lct, row)
	}
	return lct, nil

}
func getCrpInd(spName string) (lct []LutCorporateType, err error) {
	sp, _ := RetriveSP(spName)
	rows, err := Db.Query(sp)
	if err != nil {
		return lct, fmt.Errorf("Cannot get the Rows %v", err.Error())
	}
	defer rows.Close()
	for rows.Next() {
		var row LutCorporateType
		err = rows.Scan(&row.CodeDescription, &row.Code)
		if err != nil {
			return lct, fmt.Errorf("Cannot read the Rows %v", err.Error())
		}
		lct = append(lct, row)
	}
	return lct, nil

}

func getSkills() (lct []LutSkillsMaster, err error) {
	sp, _ := RetriveSP("GET_LUT_SKILLS")
	rows, err := Db.Query(sp)
	if err != nil {
		return lct, fmt.Errorf("Cannot get the Rows %v", err.Error())
	}
	defer rows.Close()
	for rows.Next() {
		var row LutSkillsMaster
		err = rows.Scan(&row.SkillID, &row.Skill)
		if err != nil {
			return lct, fmt.Errorf("Cannot read the Rows %v", err.Error())
		}
		lct = append(lct, row)
	}
	return lct, nil

}

func getPrograms() (lct []LutProgramMaster, err error) {
	sp, _ := RetriveSP("GET_LUT_PROGRAMS")
	rows, err := Db.Query(sp)
	if err != nil {
		return lct, fmt.Errorf("Cannot get the Rows %v", err.Error())
	}
	defer rows.Close()
	for rows.Next() {
		var row LutProgramMaster
		err = rows.Scan(&row.ProgramID, &row.Program)
		if err != nil {
			return lct, fmt.Errorf("Cannot read the Rows %v", err.Error())
		}
		lct = append(lct, row)
	}
	return lct, nil

}

func getDepartments() (lct []LutDepartmentMaster, err error) {
	sp, _ := RetriveSP("GET_LUT_DEPART")
	rows, err := Db.Query(sp)
	if err != nil {
		return lct, fmt.Errorf("Cannot get the Rows %v", err.Error())
	}
	defer rows.Close()
	for rows.Next() {
		var row LutDepartmentMaster
		err = rows.Scan(&row.DepartmentID, &row.ProgramID, &row.Department)
		if err != nil {
			return lct, fmt.Errorf("Cannot read the Rows %v", err.Error())
		}
		lct = append(lct, row)
	}
	return lct, nil

}
