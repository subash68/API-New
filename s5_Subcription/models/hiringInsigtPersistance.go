package models

import (
	"database/sql"
	"fmt"
	"strings"
	"time"
)

// Insert ...
func (chi *CrpHiringInsightsModel) Insert() (*CrpHiringInsightsModel, error) {
	chi.Top5LocationsRecruited = []string{"Banglore", "Hyderabad", "Pune", "NCR", "Mumbai"}
	chi.Top5SkillsRecruited = []string{"Java", "Golang", "Nodejs", "Blockchain", "Backend"}
	chi.Top5ProgramsRecruited = []string{"Mtech", "Btech", "MCA", "Bcom", "MBA"}
	chi.AverageCutoffCGPA = 6.5
	chi.AverageCutoffPercentage = 69
	chi.AverageSalary = "700000 LPA"
	chi.TentativeMonthOfHiring = "MAY"

	currentTime := time.Now().Format(time.RFC3339)

	ct, _ := time.Parse(time.RFC3339, currentTime)

	chi.SubscribedDate, chi.CreationDate, chi.LastUpdatedDate = ct, ct, ct
	var err error
	chi.SubscriptionID, err = createSudID(chi.SubscribedStakeholderID, "CORP_HCI_Get_Last_ID", "SUBCHI")
	if err != nil {
		return chi, err
	}

	newCHISubIns, _ := RetriveSP("CORP_HCI_INS")
	subInsStmt, err := Db.Prepare(newCHISubIns)
	if err != nil {
		return chi, fmt.Errorf("Cannot prepare Hiring Insights Subscription insert due to %v %v", newCHISubIns, err.Error())
	}

	_, err = subInsStmt.Exec(chi.SubscriptionID, chi.SubscriberStakeholderID, chi.SubscribedStakeholderID, true, strings.Join(chi.Top5LocationsRecruited, ","), strings.Join(chi.Top5SkillsRecruited, ","), strings.Join(chi.Top5ProgramsRecruited, ","), chi.TentativeMonthOfHiring, chi.AverageCutoffCGPA, chi.AverageCutoffPercentage, chi.AverageSalary, currentTime, currentTime, currentTime)
	if err != nil {
		return chi, fmt.Errorf("Cannot Insert Hiring Insights Subscription due to %v %v", newCHISubIns, err.Error())
	}
	return chi, nil
}

// GetHiringInsightBySubID ...
func (chi *CrpHiringInsightsModel) GetHiringInsightBySubID() (*CrpHiringInsightsModel, error) {
	newUISubGet, _ := RetriveSP("CORP_HCI_GET_SUB")
	var tp, ts, tl string
	err := Db.QueryRow(newUISubGet, chi.SubscribedStakeholderID, chi.SubscriptionID).Scan(&chi.SubscriptionID, &chi.SubscriberStakeholderID, &chi.SubscribedStakeholderID, &chi.SubscriptionValidityFlag, &tl, &ts, &tp, &chi.TentativeMonthOfHiring, &chi.AverageCutoffCGPA, &chi.AverageCutoffPercentage, &chi.AverageSalary, &chi.SubscribedDate, &chi.CreationDate, &chi.LastUpdatedDate)
	if err != nil && err != sql.ErrNoRows {
		return chi, fmt.Errorf("Cannot prepare Hiring Insights Subscription insert due to %v %v", newUISubGet, err.Error())
	}
	if err != nil && err == sql.ErrNoRows {
		return chi, fmt.Errorf("Inavalid / unauthorized subscriptionID")
	}

	chi.Top5LocationsRecruited = strings.Split(tl, ",")
	chi.Top5SkillsRecruited = strings.Split(ts, ",")
	chi.Top5ProgramsRecruited = strings.Split(tp, ",")
	return chi, nil
}

// GetHiringInsightAll ...
func (chi *CrpHiringInsightsModel) GetHiringInsightAll() ([]CrpHiringInsightsModel, error) {
	var chis []CrpHiringInsightsModel
	newUISubGet, _ := RetriveSP("CORP_HCI_GET_ALL")
	rows, err := Db.Query(newUISubGet, chi.SubscribedStakeholderID)
	if err != nil {
		return chis, fmt.Errorf("Cannot GET Hiring Insights Subscription due to %v %v", newUISubGet, err.Error())
	}
	defer rows.Close()
	for rows.Next() {
		var newChi CrpHiringInsightsModel
		var tp, ts, tl string
		err := rows.Scan(&newChi.SubscriptionID, &newChi.SubscriberStakeholderID, &newChi.SubscribedStakeholderID, &newChi.SubscriptionValidityFlag, &tl, &ts, &tp, &newChi.TentativeMonthOfHiring, &newChi.AverageCutoffCGPA, &newChi.AverageCutoffPercentage, &newChi.AverageSalary, &newChi.SubscribedDate, &newChi.CreationDate, &newChi.LastUpdatedDate)
		fmt.Println(newChi, "----", tp, ts, tl)
		if err != nil {
			fmt.Errorf("Cannot Scan fetched Hiring Insights Subscription insert due to %v %v", newUISubGet, err.Error())
		}
		newChi.Top5LocationsRecruited = strings.Split(tl, ",")
		newChi.Top5SkillsRecruited = strings.Split(ts, ",")
		newChi.Top5ProgramsRecruited = strings.Split(tp, ",")
		chis = append(chis, newChi)
	}

	return chis, nil
}
