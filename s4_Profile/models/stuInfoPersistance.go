package models

import (
	"database/sql"
	"fmt"
)

// StudentInfoService ...
var (
	StudentInfoService studentInfoService = studentInfoService{}
)

type studentInfoService struct {
}

func (si *studentInfoService) AddToStudentInfo(query string, vals []interface{}) error {
	// Preparing Database insert
	fmt.Printf("\n========= Add info query : %s ====\n ====== Vals : %v\n", query, vals)
	siInfoCmd, _ := RetriveSP(query)

	stmt, err := Db.Prepare(siInfoCmd)
	if err != nil {
		return fmt.Errorf("Cannot prepare query -- %v  -- due to %v", siInfoCmd, err.Error())
	}
	_, err = stmt.Exec(vals...)
	if err != nil {
		return fmt.Errorf("Failed to insert into database -- %v -- insert due to %v", siInfoCmd, err.Error())
	}
	return nil
}

// GetAllInternships ....
func (si *studentInfoService) GetAllStudentInfo(query string, ID string) (*sql.Rows, error) {
	// Preparing Database insert
	siGetAllCmd, _ := RetriveSP(query)
	fmt.Println("Query : ", siGetAllCmd)
	slRows, err := Db.Query(siGetAllCmd, ID)
	if err != nil && err != sql.ErrNoRows {
		return slRows, fmt.Errorf("Cannot get the Rows %v", err.Error())
	} else if err == sql.ErrNoRows {
		return slRows, nil
	}

	return slRows, nil
}

// UpdateInternship ....
func (si *studentInfoService) UpdateStudentInfo(query string, vals []interface{}) error {
	// Preparing Database insert
	siUpdateCmd, _ := RetriveSP(query)

	stmt, err := Db.Prepare(siUpdateCmd)
	if err != nil {
		return fmt.Errorf("Cannot prepare -- %v -- insert due to %v", siUpdateCmd, err.Error())
	}
	_, err = stmt.Exec(vals...)
	if err != nil {
		return fmt.Errorf("Failed to update in database -- %v  -- insert due to %v", siUpdateCmd, err.Error())
	}
	return nil
}

// DeleteInternship ....
func (si *studentInfoService) DeleteStudentInfo(query string, vals []interface{}) error {
	// Preparing Database insert
	siDelCmd, _ := RetriveSP(query)

	stmt, err := Db.Prepare(siDelCmd)
	if err != nil {
		return fmt.Errorf("Cannot prepare -- %v -- insert due to %v", siDelCmd, err.Error())
	}
	_, err = stmt.Exec(vals)
	if err != nil {
		return fmt.Errorf("Failed to update in database -- %v  -- insert due to %v", siDelCmd, err.Error())
	}
	return nil
}
