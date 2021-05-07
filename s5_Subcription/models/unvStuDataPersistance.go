package models

import (
	"database/sql"
	"fmt"
	"time"
)

// Subscribe ...
func (usd *UnvStuDataModel) Subscribe() (*UnvStuDataModel, error) {
	var err error
	currentTime := time.Now().Format(time.RFC3339)

	ct, _ := time.Parse(time.RFC3339, currentTime)

	usd.SubscribedDate, usd.CreationDate, usd.LastUpdatedDate = ct, ct, ct
	usd.SubscriptionID, err = createSudID(usd.SubscribedStakeholderID, "UNV_STU_DB_Get_Last_ID", "SUBUSD")
	if err != nil {
		return usd, err
	}
	newUISubIns, _ := RetriveSP("UNV_STU_DB_SUB_INIT")
	newUISubIns += "(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)"
	subInsStmt, err := Db.Prepare(newUISubIns)
	if err != nil {
		return usd, fmt.Errorf("Cannot prepare University Student Database Subscription insert due to %v %v", newUISubIns, err.Error())
	}

	_, err = subInsStmt.Exec(usd.SubscriptionID, usd.SubscriberStakeholderID, usd.SubscribedStakeholderID, "", "", "", "", "", "", "", "", currentTime, false, currentTime, currentTime, "")
	if err != nil {
		return usd, fmt.Errorf("Cannot Insert University Student Database Subscription due to %v %v", newUISubIns, err.Error())
	}
	usd.StudentDataExists = false
	return usd, nil
}

// StoreStudentData ...
func (usd *UnvStuDataModel) StoreStudentData(query string, search string) (*UnvStuDataModel, error) {
	newUSDSubGet, _ := RetriveSP("UNV_STU_DB_VAL_INS")
	var collegeID string
	var count int
	err := Db.QueryRow(newUSDSubGet, usd.SubscriptionID, usd.SubscribedStakeholderID, usd.SubscriberStakeholderID).Scan(&usd.SubscriptionID, &usd.SubscriberStakeholderID, &usd.SubscribedStakeholderID, &collegeID, &count)
	if err != nil && err != sql.ErrNoRows {
		return usd, fmt.Errorf("Cannot prepare University Student database Subscription insert due to %v %v", newUSDSubGet, err.Error())
	} else if err == sql.ErrNoRows {
		return usd, fmt.Errorf("Invalid Subscription ID")
	}
	if count > 1 || collegeID != "" {
		return usd, fmt.Errorf("Data already Subscribed, Unauthorized database search")
	}
	usd.StudentsData, err = GetStudentsList(query, usd.SubscriberStakeholderID)
	if err != nil {
		return usd, err
	}
	newUSDDel, _ := RetriveSP("UNV_STU_DB_DEL_BFR_UPDATE")
	stmt, err := Db.Prepare(newUSDDel)
	if err != nil {
		return usd, fmt.Errorf("Failed to Prepare delete and update the student db Due to %v", err)
	}
	_, err = stmt.Exec(usd.SubscriptionID, usd.SubscribedStakeholderID, usd.SubscriberStakeholderID)
	if err != nil {
		return usd, fmt.Errorf("Failed to delete and update the student db Due to %v", err)
	}

	currentTime := time.Now().Format(time.RFC3339)

	ct, _ := time.Parse(time.RFC3339, currentTime)

	usd.SubscribedDate, usd.CreationDate, usd.LastUpdatedDate = ct, ct, ct
	newUISubIns, _ := RetriveSP("UNV_STU_DB_SUB_INIT")
	vals := []interface{}{}
	for i := range usd.StudentsData {
		newUISubIns += "(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?),"
		vals = append(vals, usd.SubscriptionID, usd.SubscriberStakeholderID, usd.SubscribedStakeholderID, usd.StudentsData[i].CollegeID, usd.StudentsData[i].ProgramName, usd.StudentsData[i].ProgramID, usd.StudentsData[i].BranchName, usd.StudentsData[i].BranchID, usd.StudentsData[i].AvgPercentage, usd.StudentsData[i].AvgPercentage, usd.StudentsData[i].StakeholderID, currentTime, true, currentTime, currentTime, search)
	}
	newUISubIns = newUISubIns[0 : len(newUISubIns)-1]
	subInsStmt, err := Db.Prepare(newUISubIns)
	if err != nil {
		return usd, fmt.Errorf("Cannot prepare University Student Database insert due to %v %v", newUISubIns, err.Error())
	}

	_, err = subInsStmt.Exec(vals...)
	if err != nil {
		return usd, fmt.Errorf("Cannot Insert University Student Database due to %v %v", newUISubIns, err.Error())
	}
	usd.StudentDataExists = true
	return usd, nil
}

// GetStudentsList ...
func GetStudentsList(query string, unvID string) ([]StuInfoFromUnvDatabaseModel, error) {
	var stuData []StuInfoFromUnvDatabaseModel
	newStudentList, _ := RetriveSP("UNV_GET_STU_DATA_FOR_QUERY")
	newStudentList += query
	rows, err := Db.Query(newStudentList, unvID)
	if err != nil && err != sql.ErrNoRows {
		return stuData, fmt.Errorf("Cannot prepare University Student database Subscription insert due to %v %v", newStudentList, err.Error())
	} else if err == sql.ErrNoRows {
		return stuData, fmt.Errorf("No data found for the Query, Try With another Search criteria")
	}
	defer rows.Close()
	for rows.Next() {
		var nsd StuInfoFromUnvDatabaseModel
		err = rows.Scan(&nsd.StakeholderID, &nsd.AvgPercentage, &nsd.ProgramName, &nsd.ProgramID, &nsd.BranchName, &nsd.BranchID, &nsd.CollegeID, &nsd.Name)
		if err != nil {
			return stuData, err
		}
		nsd.AvgCgpa = nsd.AvgPercentage
		stuData = append(stuData, nsd)
		stuData = append(stuData, nsd)
		stuData = append(stuData, nsd)
	}

	return stuData, nil
}

// RetrieveStudentData ... UNV_STU_DB_SUB_GET_ALL
func (usd *UnvStuDataModel) RetrieveStudentData() (*UnvStuDataModel, error) {
	newStudentList, _ := RetriveSP("UNV_STU_DB_SUB_GET_ALL")
	rows, err := Db.Query(newStudentList, usd.SubscriptionID, usd.SubscribedStakeholderID)
	if err != nil && err != sql.ErrNoRows {
		return usd, fmt.Errorf("Cannot prepare University Student database Subscription insert due to %v %v", newStudentList, err.Error())
	} else if err == sql.ErrNoRows {
		return usd, fmt.Errorf("Invalid Subscription ID")
	}
	defer rows.Close()
	for rows.Next() {
		var nsd StuInfoFromUnvDatabaseModel
		err = rows.Scan(&usd.SubscriptionID, &usd.SubscriberStakeholderID, &usd.SubscribedStakeholderID, &nsd.CollegeID, &nsd.ProgramName, &nsd.ProgramID, &nsd.BranchName, &nsd.BranchID, &nsd.AvgCgpa, &nsd.AvgPercentage, &nsd.StakeholderID, &usd.SubscribedDate, &usd.SubscriptionValidityTag, &usd.CreationDate, &usd.LastUpdatedDate)
		if err != nil {
			return usd, err
		}
		if nsd.StakeholderID == "" {
			usd.StudentDataExists = false
			return usd, nil
		}
		nsd.AvgCgpa = nsd.AvgPercentage
		usd.StudentsData = append(usd.StudentsData, nsd)
	}
	usd.StudentDataExists = true
	return usd, nil
}
