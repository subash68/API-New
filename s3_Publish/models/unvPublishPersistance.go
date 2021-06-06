// Package models ...
package models

import (
	"database/sql"
	"encoding/base64"
	"fmt"
	"mime/multipart"
	"strconv"
	"strings"
	"time"
)

type dbStatements struct {
	InsStmt string
	Values  []interface{}
}

// Insert ...
func (up *UniversityProposal) Insert(ID string, files *multipart.Form) <-chan DbModelError {
	Job := make(chan DbModelError, 1)
	successResp := map[string]string{}
	fmt.Printf("\n Branches: %v, Programs: %v, Ac:%v, Tie: %v, coes:%v\n", len(up.Branches), len(up.Programs), len(up.Accredations), len(up.Tieups), len(up.Coes))
	var customError DbModelError
	var allStmts []dbStatements
	if CheckPing(&customError); customError.Err != nil {
		Job <- customError
		return Job
	}
	lastUpdatedDate := time.Now().String()
	if len(up.Programs) > 0 {
		dbStmt, err := addPrograms(up.Programs, ID, lastUpdatedDate)
		if err != nil {
			customError.Err = fmt.Errorf("Cannot prepare Programs insert due to %v", err.Error())
			Job <- customError
			return Job
		}
		allStmts = append(allStmts, dbStmt)
	}
	if len(up.Branches) > 0 {
		dbStmt, err := addBranches(up.Branches, ID)
		if err != nil {
			customError.Err = fmt.Errorf("Cannot prepare Branches insert due to %v", err.Error())
			Job <- customError
			return Job
		}
		allStmts = append(allStmts, dbStmt)
	}
	if len(up.Coes) > 0 {
		dbStmt, err := addCoes(up.Coes, ID, files)
		if err != nil {
			customError.Err = fmt.Errorf("Cannot prepare Coes insert due to %v", err.Error())
			Job <- customError
			return Job
		}
		allStmts = append(allStmts, dbStmt)
	}
	if len(up.Accredations) > 0 {
		dbStmt, err := addAccredations(up.Accredations, ID, files)
		if err != nil {
			customError.Err = fmt.Errorf("Cannot prepare Accredations insert due to %v", err.Error())
			Job <- customError
			return Job
		}
		allStmts = append(allStmts, dbStmt)
	}
	if len(up.Rankings) > 0 {
		dbStmt, err := addRankings(up.Rankings, ID, files)
		if err != nil {
			customError.Err = fmt.Errorf("Cannot prepare Rankings insert due to %v", err.Error())
			Job <- customError
			return Job
		}
		allStmts = append(allStmts, dbStmt)
	}
	if len(up.SpecialOfferings) > 0 {
		dbStmt, err := addSplOfferings(up.SpecialOfferings, ID, files)
		if err != nil {
			customError.Err = fmt.Errorf("Cannot prepare Special offerings insert due to %v", err.Error())
			Job <- customError
			return Job
		}
		allStmts = append(allStmts, dbStmt)
	}
	if len(up.Tieups) > 0 {
		dbStmt, err := addTieUps(up.Tieups, ID, files)
		if err != nil {
			customError.Err = fmt.Errorf("Cannot prepare Tieups insert due to %v", err.Error())
			Job <- customError
			return Job
		}
		allStmts = append(allStmts, dbStmt)
	}

	if len(allStmts) == 0 {
		customError.ErrCode = "UNVPROPOSAL012"
		customError.ErrTyp = "500"
		customError.Err = fmt.Errorf("Need to send atleast one Catagory in programs,branches,accredations,rankings,specialOfferings,tieups,coes")
		customError.SuccessResp = successResp
		Job <- customError
		return Job
	}
	customError = execUnvProposal(allStmts, ID)

	customError.SuccessResp = successResp

	Job <- customError
	fmt.Printf("\n --> ins : %+v\n", customError)
	return Job

}

// GetUnvProposal ...
func (up *UniversityProposal) GetUnvProposal(ID string) <-chan DbModelError {
	Job := make(chan DbModelError, 1)
	successResp := map[string]string{}
	var customError DbModelError
	if CheckPing(&customError); customError.Err != nil {
		Job <- customError
		return Job
	}

	customError = getProfileByID(ID, up)
	if customError.ErrTyp != "000" {
		Job <- customError
		return Job
	}
	customError = getPrograms(ID, up)
	if customError.ErrTyp != "000" {
		Job <- customError
		return Job
	}

	customError = getAccredations(ID, up)
	if customError.ErrTyp != "000" {
		Job <- customError
		return Job
	}

	customError = getBranches(ID, up)
	if customError.ErrTyp != "000" {
		Job <- customError
		return Job
	}

	customError = getRankings(ID, up)
	if customError.ErrTyp != "000" {
		Job <- customError
		return Job
	}

	customError = getSpclOffering(ID, up)
	if customError.ErrTyp != "000" {
		Job <- customError
		return Job
	}

	customError = getCoes(ID, up)
	if customError.ErrTyp != "000" {
		Job <- customError
		return Job
	}

	customError = getTieups(ID, up)
	if customError.ErrTyp != "000" {
		Job <- customError
		return Job
	}

	customError.ErrTyp = "000"
	customError.SuccessResp = successResp
	fmt.Printf("\n coes: %+v --<\n", up.Coes)
	Job <- customError
	fmt.Printf("\n --> ins : %+v\n", customError)
	return Job

}

// Publish ...
func (up *UnvPublishDBModel) Publish(ID string) <-chan DbModelError {
	Job := make(chan DbModelError, 1)
	successResp := map[string]string{}
	var customError DbModelError
	customError.SuccessResp = successResp
	if CheckPing(&customError); customError.Err != nil {
		Job <- customError
		return Job
	}

	// Creating HCID
	up.PublishID, customError = CreateUnvPublishID(ID, "UPDH", "UNV_PDH_Get_Last_ID")
	if customError.ErrTyp != "000" || up.PublishID == "" {
		fmt.Printf("\nFailed to Generate PublishID :%+v\n", customError)
		Job <- customError
		return Job
	}
	if up.UniversityName == "" {
		unvNameCmd, _ := RetriveSP("UNV_GET_Name")
		_ = Db.QueryRow(unvNameCmd, up.StakeholderID).Scan(&up.UniversityName)
	}
	var puberr error
	fmt.Println(up, "data is updated", up.ProfilePublished)
	//preapraing publish list
	if up.AcredPublished == true {
		puberr = fnUpdatepublishflag(up.StakeholderID, "UNV_UPDATE_Accredations")
	}
	if up.COEsPublished == true {
		puberr = fnUpdatepublishflag(up.StakeholderID, "UNV_UPDATE_Coes")
	}
	if up.BranchesPublished == true {
		puberr = fnUpdatepublishflag(up.StakeholderID, "UNV_UPDATE_Branch")
	}
	if up.OtherPublished == true {
		puberr = fnUpdatepublishflag(up.StakeholderID, "UNV_UPDATE_OI")
	}
	if up.ProfilePublished == true {
		fmt.Println("entered in to the mster")
		puberr = fnUpdatepublishflag(up.StakeholderID, "UNV_UPDATE_UnvMaster")
	}
	if up.ProgramsPublished == true {
		puberr = fnUpdatepublishflag(up.StakeholderID, "UNV_UPDATE_Program")
	}
	if up.StudentStrengthPublished == true {
		//		err := fnUpdatepublishflag(up.StakeholderID, "")
	}
	if up.RankingPublished == true {
		puberr = fnUpdatepublishflag(up.StakeholderID, "UNV_UPDATE_Rank")
	}
	if up.InfoPublished == true {
		// err := fnUpdatepublishflag(up.StakeholderID,"")
	}

	if puberr != nil {
		customError.ErrTyp = "500"
		customError.Err = fmt.Errorf("Cannot prepare -- %v , -- update due to %v", puberr.Error())
		customError.ErrCode = "S3PJ002"
		Job <- customError
		return Job
	}

	// Preparing Database insert
	unvPublishCmd, _ := RetriveSP("UNV_PDH_INS_NEW")
	unvPublishCmd += "(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)"

	stmt, err := Db.Prepare(unvPublishCmd)
	if err != nil {
		customError.ErrTyp = "500"
		customError.Err = fmt.Errorf("Cannot prepare -- %v , -- insert due to %v", unvPublishCmd, err.Error())
		customError.ErrCode = "S3PJ002"
		Job <- customError
		return Job
	}
	currentTime := time.Now()
	fmt.Printf("\n==============>  %+v <=====================\n", up)
	_, err = stmt.Exec(up.StakeholderID, up.PublishID, up.UniversityName, up.DateOfPublish, up.ProgramsPublished, up.BranchesPublished, up.StudentStrengthPublished, up.AcredPublished, up.COEsPublished, up.RankingPublished, up.OtherPublished, up.ProfilePublished, up.InfoPublished, "Profile has been published", currentTime, currentTime, up.PublishedData)
	if err != nil {
		customError.ErrTyp = "500"
		customError.Err = fmt.Errorf("Failed to insert in database -- %v , -- insert due to %v", unvPublishCmd, err.Error())
		customError.ErrCode = "S3PJ002"
		Job <- customError
		return Job
	}

	customError.ErrTyp = "000"
	successResp["publishID"] = fmt.Sprintf("%v", up.PublishID)
	customError.SuccessResp = successResp

	Job <- customError
	fmt.Printf("\n --> ins : %+v\n", customError)
	return Job
}
func fnUpdatepublishflag(stakeHolderID string, retrivesp string) error {

	updateSP, _ := RetriveSP(retrivesp)
	stmtWhere, _ := RetriveSP("UNV_UPDATE_GENWHERE")

	updateSP = updateSP + " PublishFlag=1 " + stmtWhere
	updateStm, err := Db.Prepare(updateSP)
	if err != nil {
		fmt.Println(updateSP)
		return fmt.Errorf("Cannot prepare database update due to %v", err.Error())

	}
	fmt.Println("fnupdatecall ", retrivesp, stakeHolderID)
	_, err = updateStm.Exec(stakeHolderID)
	if err != nil {
		return fmt.Errorf("Failed to update the database due to : %v", err.Error())
	}
	return nil
}

// func fnUpdatecoes(stakeHolderID string) error {

// 	updateSP, _ := RetriveSP("UNV_UPDATE_Coes")
// 	stmtWhere, _ := RetriveSP("UNV_UPDATE_GENWHERE")

// 	updateSP = updateSP + " PublishFlag=1 " + stmtWhere
// 	updateStm, err := Db.Prepare(updateSP)
// 	if err != nil {
// 		fmt.Println(updateSP)
// 		return fmt.Errorf("Cannot prepare database update due to %v", err.Error())

// 	}
// 	_, err = updateStm.Exec(stakeHolderID)
// 	if err != nil {
// 		return fmt.Errorf("Failed to update the database due to : %v", err.Error())
// 	}
// 	return nil
// }

// func fnUpdateAcc(stakeHolderID string) error {

// 	updateSP, _ := RetriveSP("UNV_UPDATE_Accredations")
// 	stmtWhere, _ := RetriveSP("UNV_UPDATE_GENWHERE")

// 	updateSP = updateSP + " PublishFlag=1 " + stmtWhere
// 	updateStm, err := Db.Prepare(updateSP)
// 	if err != nil {
// 		fmt.Println(updateSP)
// 		return fmt.Errorf("Cannot prepare database update due to %v", err.Error())

// 	}
// 	_, err = updateStm.Exec(stakeHolderID)
// 	if err != nil {
// 		return fmt.Errorf("Failed to update the database due to : %v", err.Error())
// 	}
// 	return nil
// }

// GetAllPublishedData ...
func (up *UnvPublishDBModel) GetAllPublishedData() ([]UnvPublishDBModel, error) {
	var ups []UnvPublishDBModel
	programRows, customError := getQueryRows(up.StakeholderID, "UNV_PDH_GET_BY_ID")
	if customError.ErrTyp != "000" {
		return ups, customError.Err
	}
	defer programRows.Close()
	for programRows.Next() {
		var newPD UnvPublishDBModel
		err := programRows.Scan(&newPD.StakeholderID, &newPD.PublishID, &newPD.DateOfPublish, &newPD.ProgramsPublished, &newPD.BranchesPublished, &newPD.StudentStrengthPublished, &newPD.AcredPublished, &newPD.COEsPublished, &newPD.RankingPublished, &newPD.OtherPublished, &newPD.ProfilePublished, &newPD.InfoPublished, &newPD.GeneralNote, &newPD.CreationDate, &newPD.LastUpdatedDate, &newPD.PublishedData)
		if err != nil {
			customError.ErrTyp = "500"
			customError.Err = fmt.Errorf("Cannot Scan Published data rows %v", err)
			customError.ErrCode = "S3UNVP001"
			return ups, customError.Err
		}
		ups = append(ups, newPD)
	}
	return ups, nil
}

// CreateUnvPublishID ...
func CreateUnvPublishID(crpID string, code string, queryStr string) (string, DbModelError) {
	rowSP, _ := RetriveSP(queryStr)
	lastID := ""
	err := Db.QueryRow(rowSP, crpID).Scan(&lastID)
	fmt.Println("---------------> ", lastID)
	var idCreationError DbModelError
	if err != nil && err != sql.ErrNoRows {
		idCreationError.ErrTyp = "500"
		idCreationError.Err = fmt.Errorf("Failed to create Published Job ID ", err)
		idCreationError.ErrCode = "S3PJ001"
		return "", idCreationError
	}
	if err == sql.ErrNoRows || lastID == "" {
		lastID = "0000000000000"
	}
	corporateNum, _ := strconv.Atoi(crpID[8:])
	fmt.Println("---------------> ", corporateNum)
	countNum, _ := strconv.Atoi(lastID[len(lastID)-8:])
	fmt.Println("---------------> ", countNum)
	idCreationError.ErrTyp = "000"
	return code + strconv.Itoa(corporateNum) + (fmt.Sprintf("%08d", (countNum + (1)))), idCreationError
}

func addPrograms(programs []UnvProgramsDBModel, ID string, lastUpdatedDate string) (dbStatements, error) {
	insSP, stmtExists := RetriveSP("UNV_Add_Program")
	if !stmtExists {
		return dbStatements{}, fmt.Errorf("Failed to retrieve the database queries")
	}
	vals := []interface{}{}
	for _, value := range programs {
		insSP += "(?,?,?,?,?,?,?,?),"
		vals = append(vals, ID, value.ProgramID, value.ProgramName, value.ProgramType, value.StartDate, value.EndDate, value.EnablingFlag, lastUpdatedDate)
	}
	insSP = insSP[0 : len(insSP)-1]

	return dbStatements{insSP, vals}, nil
}

func addAccredations(accrs []UnvAccredationsDBModel, ID string, form *multipart.Form) (dbStatements, error) {
	insSP, stmtExists := RetriveSP("UNV_Add_Accredations")
	if !stmtExists {
		return dbStatements{}, fmt.Errorf("Failed to retrieve the database queries")
	}
	vals := []interface{}{}
	for _, value := range accrs {
		if value.AccredationFile != "" {
			_, err := base64.StdEncoding.DecodeString(string(value.AccredationFile))
			if err != nil {
				return dbStatements{}, fmt.Errorf("AccredationFile file is not a base64 Encoded string")
			}
			if len(strings.TrimSpace(value.AccredationFileName)) == 0 {
				return dbStatements{}, fmt.Errorf("AccredationFileName is required ")
			}
		}

		insSP += "(?,?,?,?,?,?,?,?,?,?,?),"
		vals = append(vals, ID, value.AccredationName, value.AccredationType, value.AccredationDescription, value.IssuingAuthority, value.AccredationFile, value.AccredationFileName, value.StartDate, value.EndDate, value.EnablingFlag, value.LastUpdatedDate)
	}
	insSP = insSP[0 : len(insSP)-1]

	return dbStatements{insSP, vals}, nil
}

func addRankings(rankings []UnvYearWiseRanking, ID string, form *multipart.Form) (dbStatements, error) {
	insSP, stmtExists := RetriveSP("UNV_Add_Ranking")
	if !stmtExists {
		return dbStatements{}, fmt.Errorf("Failed to retrieve the database queries")
	}
	vals := []interface{}{}
	for _, value := range rankings {
		if value.RankingFile != "" {
			_, err := base64.StdEncoding.DecodeString(string(value.RankingFile))
			if err != nil {
				return dbStatements{}, fmt.Errorf("RankingFile file is not a base64 Encoded string")
			}
			if len(strings.TrimSpace(value.RankingFileName)) == 0 {
				return dbStatements{}, fmt.Errorf("RankingfileName is required ")
			}
		}
		insSP += "(?,?,?,?,?),"
		vals = append(vals, ID, value.Rank, value.IssuingAuthority, value.RankingFile, value.RankingFileName)
	}
	insSP = insSP[0 : len(insSP)-1]

	return dbStatements{insSP, vals}, nil
}

func addSplOfferings(splOff []UnvSpecialOfferingsDBModel, ID string, form *multipart.Form) (dbStatements, error) {
	insSP, stmtExists := RetriveSP("UNV_Add_SpecialOfferings")
	if !stmtExists {
		return dbStatements{}, fmt.Errorf("Failed to retrieve the database queries")
	}
	vals := []interface{}{}
	for _, value := range splOff {
		if value.SpecialOfferingFile != "" {
			_, err := base64.StdEncoding.DecodeString(string(value.SpecialOfferingFile))
			if err != nil {
				return dbStatements{}, fmt.Errorf("SpecialOfferingFile file is not a base64 Encoded string")
			}
			if len(strings.TrimSpace(value.SpecialOfferingFileName)) == 0 {
				return dbStatements{}, fmt.Errorf("SpecialOfferingFileName is required ")
			}
		}

		insSP += "(?,?,?,?,?,?,?,?,?,?,?,?,?,?),"
		vals = append(vals, ID, value.SpecialOfferingName, value.SpecialOfferingType, value.SpecialOfferingDescription, value.InternallyManagedFlag, value.OutsourcedVendorName, value.OutsourcedVendorPhoneNumber, value.OutsourcedVendorEmailID, value.OutsourcedVendorStakeholderID, value.SpecialOfferingFile, value.SpecialOfferingFileName, value.StartDate, value.EndDate, value.EnablingFlag)
	}
	insSP = insSP[0 : len(insSP)-1]

	return dbStatements{insSP, vals}, nil
}

func addTieUps(tieups []UnvTieupsDBModel, ID string, form *multipart.Form) (dbStatements, error) {
	insSP, stmtExists := RetriveSP("UNV_Add_Tieups")
	if !stmtExists {
		return dbStatements{}, fmt.Errorf("Failed to retrieve the database queries")
	}
	vals := []interface{}{}
	for _, value := range tieups {
		if value.TieupFile != "" {
			_, err := base64.StdEncoding.DecodeString(string(value.TieupFile))
			if err != nil {
				return dbStatements{}, fmt.Errorf("Tieup file is not a base64 Encoded string")
			}
			if len(strings.TrimSpace(value.TieupFileName)) == 0 {
				return dbStatements{}, fmt.Errorf("TieupfileName is required ")
			}

		}
		insSP += "(?,?,?,?,?,?,?,?,?,?,?,?,?),"
		vals = append(vals, ID, value.TieupName, value.TieupType, value.TieupDescription, value.TieupWithName, value.TieupWithPhoneNumber, value.TieupWithEmail, value.TieupWithStakeholderID, value.TieupFile, value.TieupFileName, value.StartDate, value.EndDate, value.EnablingFlag)
	}
	insSP = insSP[0 : len(insSP)-1]

	return dbStatements{insSP, vals}, nil
}

func addCoes(tieups []UnvCEOsDBModel, ID string, form *multipart.Form) (dbStatements, error) {
	insSP, stmtExists := RetriveSP("UNV_Add_Coes")
	//var err error
	if !stmtExists {
		return dbStatements{}, fmt.Errorf("Failed to retrieve the database queries")
	}
	vals := []interface{}{}
	for _, value := range tieups {
		if value.CoeFile != "" {
			_, err := base64.StdEncoding.DecodeString(value.CoeFile)
			if err != nil {
				return dbStatements{}, fmt.Errorf("Coes file is not a base64 Encoded string")
			}
			if len(strings.TrimSpace(value.CoeFileName)) == 0 {
				return dbStatements{}, fmt.Errorf("CoefileName is required ")
			}
		}

		insSP += "(?,?,?,?,?,?,?,?,?,?,?,?,?,?),"
		vals = append(vals, ID, value.CoeName, value.CoeType, value.CoeDescription, value.InternallyManagedFlag, value.OutsourcedVendorName, value.OutsourcedVendorStakeholderID, value.CoeFile, value.CoeFileName, value.StartDate, value.EndDate, value.EnablingFlag, value.OutsourcedVendorEmailID, value.OutsourcedVendorPhoneNumber)
	}
	insSP = insSP[0 : len(insSP)-1]

	return dbStatements{insSP, vals}, nil
}

func addBranches(tieups []UnvProgramWiseBranchDBModel, ID string) (dbStatements, error) {
	insSP, stmtExists := RetriveSP("UNV_Add_Branches")
	if !stmtExists {
		return dbStatements{}, fmt.Errorf("Failed to retrieve the database queries")
	}
	vals := []interface{}{}
	for _, value := range tieups {
		insSP += "(?,?,?,?,?,?,?,?,?,?,?),"
		vals = append(vals, ID, value.ProgramID, value.ProgramName, value.BranchID, value.BranchName, value.StartDate, value.EndDate, value.EnablingFlag, value.NoOfPassingStudents, value.MonthYearOfPassing, value.LastUpdatedDate)
	}
	insSP = insSP[0 : len(insSP)-1]

	return dbStatements{insSP, vals}, nil
}

func execUnvProposal(execStmts []dbStatements, ID string) DbModelError {
	var customError DbModelError
	for _, stmt := range execStmts {
		insStmt, err := Db.Prepare(stmt.InsStmt)
		if err != nil {
			customError.ErrTyp = "500"
			customError.ErrCode = "S3PJ011"
			customError.Err = fmt.Errorf("Failed Prepare Insert/Update University Proposal. Error:%v", err.Error())
			return customError
		}

		_, err = insStmt.Exec(stmt.Values...)
		if err != nil {
			customError.ErrTyp = "500"
			customError.Err = fmt.Errorf("Failed Insert/Update University Proposal, Error : %v", err.Error())
			customError.ErrCode = "S3PJ012"
			return customError
		}

	}
	customError.ErrTyp = "000"
	return customError
}

func getProfileByID(ID string, up *UniversityProposal) DbModelError {
	//unvProfile, customError := getQueryRows(ID, "UNV_GET_Program")
	var dbCustomError DbModelError
	qrySP, stmtExists := RetriveSP("UNV_GET_PROFILE")
	if !stmtExists {
		dbCustomError.ErrTyp = "500"
		dbCustomError.Err = fmt.Errorf("Cannot retrive Profile query SP")
		dbCustomError.ErrCode = "S3UNVP001"
		return dbCustomError
	}
	row := Db.QueryRow(qrySP, ID)
	err := row.Scan(&up.Profile.StakeholderID, &up.Profile.UniversityName, &up.Profile.UniversityCollegeID, &up.Profile.UniversityHQAddressLine1, &up.Profile.UniversityHQAddressLine2, &up.Profile.UniversityHQAddressLine3, &up.Profile.UniversityHQAddressCountry, &up.Profile.UniversityHQAddressState, &up.Profile.UniversityHQAddressCity, &up.Profile.UniversityHQAddressDistrict, &up.Profile.UniversityHQAddressZipcode, &up.Profile.UniversityHQAddressPhone, &up.Profile.UniversityHQAddressemail, &up.Profile.UniversityLocalBranchAddressLine1, &up.Profile.UniversityLocalBranchAddressLine2, &up.Profile.UniversityLocalBranchAddressLine3, &up.Profile.UniversityLocalBranchAddressCountry, &up.Profile.UniversityLocalBranchAddressState, &up.Profile.UniversityLocalBranchAddressCity, &up.Profile.UniversityLocalBranchAddressDistrict, &up.Profile.UniversityLocalBranchAddressZipcode, &up.Profile.UniversityLocalBranchAddressPhone, &up.Profile.UniversityLocalBranchAddressemail, &up.Profile.PrimaryContactFirstName, &up.Profile.PrimaryContactMiddleName, &up.Profile.PrimaryContactLastName, &up.Profile.PrimaryContactDesignation, &up.Profile.PrimaryContactPhone, &up.Profile.PrimaryContactEmail, &up.Profile.UniversitySector, &up.Profile.UniversityProfile, &up.Profile.YearOfEstablishment, &up.Profile.Attachment, &up.Profile.AttachmentName, &up.Profile.PublishedFlag, &up.Profile.DateOfJoining)
	if err != nil {
		dbCustomError.ErrTyp = "500"
		dbCustomError.Err = fmt.Errorf("Cannot Scan Profile %s", err.Error())
		dbCustomError.ErrCode = "S3UNVP001"
		return dbCustomError
	}
	dbCustomError.ErrTyp = "000"
	return dbCustomError

}

func getPrograms(ID string, up *UniversityProposal) DbModelError {
	programRows, customError := getQueryRows(ID, "UNV_GET_Program")
	if customError.ErrTyp != "000" {
		return customError
	}
	defer programRows.Close()
	for programRows.Next() {
		var newProgram UnvProgramsDBModel
		err := programRows.Scan(&newProgram.ID, &newProgram.ProgramID, &newProgram.ProgramName, &newProgram.ProgramType, &newProgram.StartDate, &newProgram.EndDate, &newProgram.EnablingFlag, &newProgram.PublishedFlag, &newProgram.CreationDate, &newProgram.LastUpdatedDate)
		if err != nil {
			customError.ErrTyp = "500"
			customError.Err = fmt.Errorf("Cannot Scan programs rows")
			customError.ErrCode = "S3UNVP001"
			return customError
		}
		up.Programs = append(up.Programs, newProgram)
	}
	customError.ErrTyp = "000"
	return customError
}

func getAccredations(ID string, up *UniversityProposal) DbModelError {
	accrRows, customError := getQueryRows(ID, "UNV_GET_Accredations")
	if customError.ErrTyp != "000" {
		return customError
	}
	defer accrRows.Close()
	for accrRows.Next() {
		var newAccr UnvAccredationsDBModel
		err := accrRows.Scan(&newAccr.ID, &newAccr.AccredationName, &newAccr.AccredationType, &newAccr.AccredationDescription, &newAccr.IssuingAuthority, &newAccr.AccredationFile, &newAccr.AccredationFileName, &newAccr.StartDate, &newAccr.EndDate, &newAccr.EnablingFlag, &newAccr.PublishedFlag, &newAccr.CreationDate, &newAccr.LastUpdatedDate)
		if err != nil {
			customError.ErrTyp = "500"
			customError.Err = fmt.Errorf("Cannot Scan Accredations rows")
			customError.ErrCode = "S3UNVP001"
			return customError
		}
		up.Accredations = append(up.Accredations, newAccr)
	}
	customError.ErrTyp = "000"
	return customError
}

func getRankings(ID string, up *UniversityProposal) DbModelError {
	RankingRows, customError := getQueryRows(ID, "UNV_GET_Ranking")
	if customError.ErrTyp != "000" {
		return customError
	}
	defer RankingRows.Close()
	for RankingRows.Next() {
		var newRanking UnvYearWiseRanking
		err := RankingRows.Scan(&newRanking.ID, &newRanking.Rank, &newRanking.IssuingAuthority, &newRanking.RankingFile, &newRanking.RankingFileName, &newRanking.PublishedFlag, &newRanking.CreationDate, &newRanking.LastUpdatedDate)
		if err != nil {
			customError.ErrTyp = "500"
			customError.Err = fmt.Errorf("Cannot Scan Accredations rows")
			customError.ErrCode = "S3UNVP001"
			return customError
		}
		up.Rankings = append(up.Rankings, newRanking)
	}
	customError.ErrTyp = "000"
	return customError
}

func getSpclOffering(ID string, up *UniversityProposal) DbModelError {
	sofRows, customError := getQueryRows(ID, "UNV_GET_SpecialOfferings")
	if customError.ErrTyp != "000" {
		return customError
	}
	defer sofRows.Close()
	for sofRows.Next() {
		var newOffering UnvSpecialOfferingsDBModel
		err := sofRows.Scan(&newOffering.ID, &newOffering.SpecialOfferingName, &newOffering.SpecialOfferingType, &newOffering.SpecialOfferingDescription, &newOffering.InternallyManagedFlag, &newOffering.OutsourcedVendorName, &newOffering.OutsourcedVendorPhoneNumber, &newOffering.OutsourcedVendorEmailID, &newOffering.OutsourcedVendorStakeholderID, &newOffering.SpecialOfferingFile, &newOffering.SpecialOfferingFileName, &newOffering.StartDate, &newOffering.EndDate, &newOffering.EnablingFlag, &newOffering.PublishedFlag, &newOffering.CreationDate, &newOffering.LastUpdatedDate)
		if err != nil {
			customError.ErrTyp = "500"
			customError.Err = fmt.Errorf("Cannot Scan Specialoffering rows")
			customError.ErrCode = "S3UNVP001"
			return customError
		}
		up.SpecialOfferings = append(up.SpecialOfferings, newOffering)
	}
	customError.ErrTyp = "000"
	return customError
}

func getTieups(ID string, up *UniversityProposal) DbModelError {
	tuRows, customError := getQueryRows(ID, "UNV_GET_Tieups")
	if customError.ErrTyp != "000" {
		return customError
	}
	defer tuRows.Close()
	for tuRows.Next() {
		var newTU UnvTieupsDBModel
		err := tuRows.Scan(&newTU.ID, &newTU.TieupName, &newTU.TieupType, &newTU.TieupDescription, &newTU.TieupWithName, &newTU.TieupWithPhoneNumber, &newTU.TieupWithEmail, &newTU.TieupWithStakeholderID, &newTU.TieupFile, &newTU.TieupFileName, &newTU.StartDate, &newTU.EndDate, &newTU.EnablingFlag, &newTU.PublishedFlag, &newTU.CreationDate, &newTU.LastUpdatedDate)
		if err != nil {
			customError.ErrTyp = "500"
			customError.Err = fmt.Errorf("Cannot Scan Tieup rows")
			customError.ErrCode = "S3UNVP001"
			return customError
		}
		up.Tieups = append(up.Tieups, newTU)
	}
	customError.ErrTyp = "000"
	return customError
}

func getCoes(ID string, up *UniversityProposal) DbModelError {
	coeRows, customError := getQueryRows(ID, "UNV_GET_Coes")
	if customError.ErrTyp != "000" {
		return customError
	}
	defer coeRows.Close()
	for coeRows.Next() {
		var newCoe UnvCEOsDBModel
		err := coeRows.Scan(&newCoe.ID, &newCoe.CoeName, &newCoe.CoeType, &newCoe.CoeDescription, &newCoe.InternallyManagedFlag, &newCoe.OutsourcedVendorName, &newCoe.OutsourcedVendorStakeholderID, &newCoe.CoeFile, &newCoe.CoeFileName, &newCoe.StartDate, &newCoe.EndDate, &newCoe.EnablingFlag, &newCoe.PublishedFlag, &newCoe.OutsourcedVendorEmailID, &newCoe.OutsourcedVendorPhoneNumber, &newCoe.CreationDate, &newCoe.LastUpdatedDate)
		if err != nil {
			customError.ErrTyp = "500"
			customError.Err = fmt.Errorf("Cannot Scan COE rows %v", err.Error())
			customError.ErrCode = "S3UNVP001"
			return customError
		}
		up.Coes = append(up.Coes, newCoe)
	}
	customError.ErrTyp = "000"
	return customError
}

func getBranches(ID string, up *UniversityProposal) DbModelError {
	branchRows, customError := getQueryRows(ID, "UNV_GET_Branches")
	if customError.ErrTyp != "000" {
		return customError
	}
	defer branchRows.Close()
	for branchRows.Next() {
		var newBranch UnvProgramWiseBranchDBModel
		err := branchRows.Scan(&newBranch.ID, &newBranch.ProgramID, &newBranch.ProgramName, &newBranch.BranchID, &newBranch.BranchName, &newBranch.StartDate, &newBranch.EndDate, &newBranch.EnablingFlag, &newBranch.PublishedFlag, &newBranch.NoOfPassingStudents, &newBranch.MonthYearOfPassing, &newBranch.CreationDate, &newBranch.LastUpdatedDate)
		if err != nil {
			customError.ErrTyp = "500"
			customError.Err = fmt.Errorf("Cannot Scan Branch rows")
			customError.ErrCode = "S3UNVP001"
			return customError
		}
		up.Branches = append(up.Branches, newBranch)
	}
	customError.ErrTyp = "000"
	return customError
}

func getQueryRows(ID string, query string) (rows *sql.Rows, dbCustomError DbModelError) {
	qrySP, stmtExists := RetriveSP(query)
	if !stmtExists {
		dbCustomError.ErrTyp = "500"
		dbCustomError.Err = fmt.Errorf("Cannot retrive query SP %s", query)
		dbCustomError.ErrCode = "S3UNVP001"
		return rows, dbCustomError
	}
	rows, err := Db.Query(qrySP, ID) //.Scan()
	fmt.Println(qrySP)
	if err != nil && err != sql.ErrNoRows {
		dbCustomError.ErrTyp = "500"
		dbCustomError.Err = fmt.Errorf("Cannot get the Rows %v", err.Error())
		dbCustomError.ErrCode = "S3UNVP001"
		return rows, dbCustomError
	}

	dbCustomError.ErrTyp = "000"
	return rows, dbCustomError
}

// DeleteUnvProposalByID ...
func DeleteUnvProposalByID(ID string, deleteID string, query string) error {
	qrySP, stmtExists := RetriveSP(query)
	if !stmtExists {
		return fmt.Errorf("Cannot retrive query SP %s", query)
	}
	_, err := Db.Exec(qrySP, deleteID, ID)
	if err != nil {

		return fmt.Errorf("Deleting University Proposal failed with %v", err.Error())
	}

	return nil
}
