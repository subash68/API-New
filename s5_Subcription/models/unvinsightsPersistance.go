package models

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"
	"time"
)

// Insert ...
func (uim *UnvInsightsModel) Insert() (*UnvInsightsModel, error) {
	uim.AverageCGPA = 7.5
	uim.AveragePercentage = 66.8
	uim.HighestCGPA = 9.0
	uim.HighestPercentage = 89.5
	uim.HighestPackageReceived = "3400000 LPA"
	uim.AveragePackageReceived = "700000 LPA"
	uim.UniversityConversionRatio = 89
	uim.TentativeMonthOfPassing = "MAY"
	uim.Top5Recruiters = []string{"Oracle", "PGK", "Infosys", "Capgemini", "TCS"}
	uim.Top5Skills = []string{"Java", "Golang", "Nodejs", "Blockchain", "Backend"}

	currentTime := time.Now().Format(time.RFC3339)

	ct, _ := time.Parse(time.RFC3339, currentTime)

	uim.SubscribedDate, uim.CreationDate, uim.LastUpdatedDate = ct, ct, ct

	err := uim.createSudID()
	if err != nil {
		return uim, err
	}

	newUISubIns, _ := RetriveSP("UNV_INSIGHTS_INS")
	subInsStmt, err := Db.Prepare(newUISubIns)
	if err != nil {
		return uim, fmt.Errorf("Cannot prepare University Insights Subscription insert due to %v %v", newUISubIns, err.Error())
	}

	_, err = subInsStmt.Exec(uim.SubscriptionID, uim.SubscriberStakeholderID, uim.SubscribedStakeholderID, uim.AverageCGPA, uim.AveragePercentage, uim.HighestCGPA, uim.HighestPercentage, uim.HighestPackageReceived, uim.AveragePackageReceived, uim.UniversityConversionRatio, uim.TentativeMonthOfPassing, strings.Join(uim.Top5Recruiters, ","), strings.Join(uim.Top5Skills, ","), currentTime, currentTime, currentTime)
	if err != nil {
		return uim, fmt.Errorf("Cannot Insert University Insights Subscription due to %v %v", newUISubIns, err.Error())
	}
	return uim, nil
}

// GetUnvInsightBySubID ...
func (uim *UnvInsightsModel) GetUnvInsightBySubID() (*UnvInsightsModel, error) {
	newUISubGet, _ := RetriveSP("UNV_INSIGHTS_GET_SUB")
	var tr, ts string
	err := Db.QueryRow(newUISubGet, uim.SubscribedStakeholderID, uim.SubscriptionID).Scan(&uim.SubscriptionID, &uim.SubscriberStakeholderID, &uim.SubscribedStakeholderID, &uim.AverageCGPA, &uim.AveragePercentage, &uim.HighestCGPA, &uim.HighestPercentage, &uim.HighestPackageReceived, &uim.AveragePackageReceived, &uim.UniversityConversionRatio, &uim.TentativeMonthOfPassing, &tr, &ts, &uim.SubscribedDate, &uim.CreationDate, &uim.LastUpdatedDate)
	if err != nil {
		return uim, fmt.Errorf("Cannot prepare University Insights Subscription insert due to %v %v", newUISubGet, err.Error())
	}
	uim.Top5Recruiters = strings.Split(tr, ",")
	uim.Top5Skills = strings.Split(ts, ",")
	return uim, nil
}

// GetUnvInsightAll ...
func (uim *UnvInsightsModel) GetUnvInsightAll() ([]UnvInsightsModel, error) {
	var uims []UnvInsightsModel
	newUISubGet, _ := RetriveSP("UNV_INSIGHTS_GET_ALL")

	rows, err := Db.Query(newUISubGet, uim.SubscribedStakeholderID)
	if err != nil {
		return uims, fmt.Errorf("Cannot GET University Insights Subscription due to %v %v", newUISubGet, err.Error())
	}
	defer rows.Close()
	for rows.Next() {
		var newUIM UnvInsightsModel
		var tr, ts string
		err = rows.Scan(&newUIM.SubscriptionID, &newUIM.SubscriberStakeholderID, &newUIM.SubscribedStakeholderID, &newUIM.AverageCGPA, &newUIM.AveragePercentage, &newUIM.HighestCGPA, &newUIM.HighestPercentage, &newUIM.HighestPackageReceived, &newUIM.AveragePackageReceived, &newUIM.UniversityConversionRatio, &newUIM.TentativeMonthOfPassing, &tr, &ts, &newUIM.SubscribedDate, &newUIM.CreationDate, &newUIM.LastUpdatedDate)
		if err != nil {
			fmt.Errorf("Cannot Scan fetched University Insights Subscription insert due to %v %v", newUISubGet, err.Error())
		}
		newUIM.Top5Recruiters = strings.Split(tr, ",")
		newUIM.Top5Skills = strings.Split(ts, ",")
		uims = append(uims, newUIM)
	}

	return uims, nil
}

// createSudID ...
func (uim *UnvInsightsModel) createSudID() error {
	rowSP, _ := RetriveSP("UNV_INSIGHTS_Get_Last_ID")
	lastID := ""
	err := Db.QueryRow(rowSP, uim.SubscribedStakeholderID).Scan(&lastID)

	if err != nil && err != sql.ErrNoRows {
		return err
	}
	if err == sql.ErrNoRows {
		lastID = "0000000000000"
	}
	corporateNum, _ := strconv.Atoi(uim.SubscribedStakeholderID[7:])
	countNum, _ := strconv.Atoi(lastID[len(lastID)-7:])
	fmt.Println("--------------------> ", lastID, countNum)
	uim.SubscriptionID = "SUBUI" + strconv.Itoa(corporateNum) + (fmt.Sprintf("%07d", (countNum + 1)))

	return nil
}
