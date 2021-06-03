package models

import (
	"database/sql"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"time"
)

// InsertAcademics ...
func (sa *StudentAcademicsModelReq) InsertAcademics(form *multipart.Form) error {
	ques := ""
	val := []interface{}{}
	val = append(val, sa.StakeholderID)
	insFields := ""
	updSP := ""
	semIns := ""
	semVals := []interface{}{}
	currentTime := time.Now().Format(time.RFC3339)
	if sa.Tenth.Name != "" && sa.Twelfth.Name != "" {
		return fmt.Errorf("Failed to Parse Data, Required Only one Academic Details per request")
	}
	if sa.Tenth.Name != "" {
		tenthSP, _ := RetriveSP("STU_TENTH_INS")
		tenthUPD, _ := RetriveSP("STU_TENTH_UPD")
		vals := constructTTReq(sa.Tenth, form, "tenthFile")
		val = append(val, vals...)
		val = append(val, vals...)

		ques += "?,?,?,?,?,?"
		insFields += tenthSP
		updSP += tenthUPD
	}
	if sa.Twelfth.Name != "" {
		twelfthSP, _ := RetriveSP("STU_Twelfth_INS")
		tenthUPD, _ := RetriveSP("STU_Twelfth_UPD")
		vals := constructTTReq(sa.Twelfth, form, "twelfthFile")
		val = append(val, vals...)
		val = append(val, vals...)
		ques += "?,?,?,?,?,?"
		insFields += twelfthSP
		updSP += tenthUPD
	}
	if sa.Graduation.UniversityStakeholderIDUniv != "" {
		gradSP, _ := RetriveSP("STU_GRAD_INS")
		gradUPD, _ := RetriveSP("STU_GRAD_UPD")
		semIns, _ = RetriveSP("STU_SEM_INS")
		for _, sem := range sa.Graduation.Semesters {
			err := verifyAttachment(sem.AttachFile)
			if err != nil {
				return fmt.Errorf("Failed at semister Attachfile due to %v", err.Error())
			}
			semIns += "(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?),"
			semVals = append(semVals, sa.StakeholderID, sa.Graduation.UniversityStakeholderIDUniv, true, false, sem.Semester, sa.Graduation.CollegeRollNumber, sa.Graduation.ProgramName, sa.Graduation.ProgramID, sa.Graduation.BranchName, sa.Graduation.BranchID, sem.CGPA, sem.Percentage, sem.AttachFile, true, currentTime, currentTime, sem.SemesterCompletionDate)
		}
		semIns = semIns[0 : len(semIns)-1]
		vals := []interface{}{sa.Graduation.UniversityStakeholderIDUniv, sa.Graduation.CollegeRollNumber, sa.Graduation.ExpectedYearOfPassing, sa.Graduation.ProgramID, sa.Graduation.ProgramName, sa.Graduation.BranchID, sa.Graduation.BranchName, sa.Graduation.FinalCGPA, sa.Graduation.FinalPercentage, sa.Graduation.ActiveBacklogsNumber, sa.Graduation.TotalNumberOfBacklogs}
		val = append(val, vals...)
		val = append(val, vals...)
		ques += "?,?,?,?,?,?,?,?,?,?,?"
		insFields += gradSP
		updSP += gradUPD
	}
	if sa.PostGraduation.UniversityStakeholderIDUniv != "" {
		gradSP, _ := RetriveSP("STU_PG_INS")
		gradUPD, _ := RetriveSP("STU_PG_UPD")
		semIns, _ = RetriveSP("STU_SEM_INS")
		for _, sem := range sa.PostGraduation.Semesters {
			err := verifyAttachment(sem.AttachFile)
			if err != nil {
				return fmt.Errorf("Failed at semister Attachfile due to %v", err.Error())
			}
			semIns += "(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?),"
			semVals = append(semVals, sa.StakeholderID, sa.PostGraduation.UniversityStakeholderIDUniv, false, true, sem.Semester, sa.PostGraduation.CollegeRollNumber, sa.PostGraduation.ProgramName, sa.PostGraduation.ProgramID, sa.PostGraduation.BranchName, sa.PostGraduation.BranchID, sem.CGPA, sem.Percentage, sem.AttachFile, true, currentTime, currentTime, sem.SemesterCompletionDate)
		}
		semIns = semIns[0 : len(semIns)-1]
		vals := []interface{}{sa.PostGraduation.UniversityStakeholderIDUniv, sa.PostGraduation.CollegeRollNumber, sa.PostGraduation.ExpectedYearOfPassing, sa.PostGraduation.ProgramID, sa.PostGraduation.ProgramName, sa.PostGraduation.BranchID, sa.PostGraduation.BranchName, sa.PostGraduation.FinalCGPA, sa.PostGraduation.FinalPercentage}
		val = append(val, vals...)
		val = append(val, vals...)
		ques += "?,?,?,?,?,?,?,?,?"
		insFields += gradSP
		updSP += gradUPD
	}
	// TODO -- Graduation and PG
	saSP, _ := RetriveSP("STU_INS_ACADEMICS")
	insFields = insFields[0 : len(insFields)-1]
	saSP += "" + insFields + ") VALUES(?," + ques + ") ON DUPLICATE KEY UPDATE " + updSP
	saSP = saSP[0 : len(saSP)-1]
	fmt.Println("======================== " + saSP + " -----++++++++++++++++++++++++")
	fmt.Printf("\n%+v\n", val)
	insStmt, err := Db.Prepare(saSP)
	if err != nil {
		return fmt.Errorf("Failed Prepare Insert/Update Student Academics. Error:%v", err.Error())
	}

	_, err = insStmt.Exec(val...)
	if err != nil {

		return fmt.Errorf("Failed Insert/Update Student Academics. Error:%v", err.Error())
	}

	if semIns != "" {
		insStmt, err := Db.Prepare(semIns)
		if err != nil {
			return fmt.Errorf("Failed Prepare Insert/Update Student Semesters. Error:%v", err.Error())
		}
		_, err = insStmt.Exec(semVals...)
		if err != nil {

			return fmt.Errorf("Failed Insert/Update Student Semesters. Error:%v", err.Error())
		}
	}

	fmt.Println(val, ques)
	return nil
}

// GetAcademics ...
func GetAcademics(ID string) (StudentAcademicsModelReq, error) {
	var sa StudentAcademicsModelReq
	var tenN StudentNullableTTModel
	var twelfthN StudentNullableTTModel
	var grad StudentGradModel
	var pg StudentPGModel
	getByIDSP, _ := RetriveSP("STU_GET_ACADEMICS")
	err := Db.QueryRow(getByIDSP, ID).Scan(&tenN.Name, &tenN.Location, &tenN.MonthAndYearOfPassing, &tenN.Board, &tenN.Percentage, &tenN.AttachmentFile, &twelfthN.Name, &twelfthN.Location, &twelfthN.MonthAndYearOfPassing, &twelfthN.Board, &twelfthN.Percentage, &twelfthN.AttachmentFile, &grad.UniversityStakeholderIDUniv, &grad.CollegeRollNumber, &grad.ExpectedYearOfPassing, &grad.ProgramID, &grad.ProgramName, &grad.BranchID, &grad.BranchName, &grad.FinalCGPA, &grad.FinalPercentage, &grad.ActiveBacklogsNumber, &grad.TotalNumberOfBacklogs, &pg.UniversityStakeholderIDUniv, &pg.CollegeRollNumber, &pg.ExpectedYearOfPassing, &pg.ProgramID, &pg.ProgramName, &pg.BranchID, &pg.BranchName, &pg.FinalCGPA, &pg.FinalPercentage)
	if err != nil && err != sql.ErrNoRows {
		return sa, fmt.Errorf("Failed to retrieving Academics : %v", err.Error())
	}
	if err != nil && err == sql.ErrNoRows {
		return sa, nil
	}
	sa.Tenth = tenN.ConvertToJSONStruct()
	sa.Twelfth = twelfthN.ConvertToJSONStruct()
	sa.Graduation = grad
	sa.PostGraduation = pg
	allSems := getAllSemDetails(ID)
	if sa.Graduation.UniversityStakeholderIDUniv != "" {
		sa.Graduation.ParseSem(allSems)
	}
	if sa.PostGraduation.UniversityStakeholderIDUniv != "" {
		sa.PostGraduation.ParseSem(allSems)
	}
	return sa, nil
}

func getAllSemDetails(ID string) []StudentSemesterModel {
	var allSems []StudentSemesterModel
	getAllSemSP, _ := RetriveSP("STU_SEM_GET_ALL")
	semRows, err := Db.Query(getAllSemSP, ID) //.Scan()
	if err != nil && err != sql.ErrNoRows {
		fmt.Printf("Cannot get the Semester Rows %v", err.Error())
		return allSems
	} else if err == sql.ErrNoRows {
		return allSems
	}
	defer semRows.Close()
	for semRows.Next() {
		var newSem StudentSemesterModel
		err = semRows.Scan(&newSem.ID, &newSem.StudentStakeholderID, &newSem.UniversityStakeholderID, &newSem.IsGrad, &newSem.IsPG, &newSem.Semester, &newSem.StudentCollegeRollNo, &newSem.ProgramName, &newSem.ProgramID, &newSem.BranchName, &newSem.BranchID, &newSem.CGPA, &newSem.Percentage, &newSem.AttachFile, &newSem.EnabledFlag, &newSem.CreationDate, &newSem.LastUpdatedDate, &newSem.SemesterCompletionDate)
		if err != nil {
			fmt.Printf("\nInvalid semester details %v\n", err)
		}
		allSems = append(allSems, newSem)
	}
	return allSems
}

// ParseSem  ...
func (sg *StudentGradModel) ParseSem(allSems []StudentSemesterModel) {
	for index := range allSems {
		if allSems[index].IsGrad {
			sg.Semesters = append(sg.Semesters, allSems[index])
		}
	}
}

// ParseSem  ...
func (sp *StudentPGModel) ParseSem(allSems []StudentSemesterModel) {
	for index := range allSems {
		if allSems[index].IsPG {
			sp.Semesters = append(sp.Semesters, allSems[index])
		}
	}
}

// ConvertToJSONStruct ...
func (st *StudentNullableTTModel) ConvertToJSONStruct() StudentTTModel {
	var stt StudentTTModel
	emptyStt := StudentTTModel{}
	if st.Name.Valid {
		stt.Name = st.Name.String
	} else {
		return emptyStt
	}
	if st.Location.Valid {
		stt.Location = st.Location.String
	} else {
		return emptyStt
	}
	if st.MonthAndYearOfPassing.Valid {
		stt.MonthAndYearOfPassing = st.MonthAndYearOfPassing.String
	} else {
		return emptyStt
	}
	if st.Board.Valid {
		stt.Board = st.Board.String
	} else {
		return emptyStt
	}
	if st.Percentage.Valid {
		stt.Percentage = st.Percentage.String
	}
	stt.AttachmentFile = st.AttachmentFile
	return stt
}

func constructTTReq(st StudentTTModel, form *multipart.Form, fileName string) []interface{} {
	var err error
	vals := []interface{}{}
	files := form.File[fileName]
	for _, file := range files {
		fileContent, _ := file.Open()
		st.AttachmentFile, err = ioutil.ReadAll(fileContent)
		if err != nil {
			fmt.Println(fileName+" file ", err.Error)
		}
	}
	vals = append(vals, st.Name, st.Location, st.MonthAndYearOfPassing, st.Board, st.Percentage, st.AttachmentFile)
	return vals
}

func verifyAttachment(binaryfile string) (err error) {
	if binaryfile != "" {
		_, err := base64.StdEncoding.DecodeString(binaryfile)
		if err != nil {
			return fmt.Errorf("Attached file is not a base64 Encoded string")
		}
	} else {
		return fmt.Errorf("Attached file is reqquired")
	}
	return nil
}

// GetContactFromProfile ....
func (si *StudentCompleteProfileDataModel) GetContactFromProfile() {
	si.ContactInfo = StudentContactInfoModel{StakeholderID: si.Profile.StakeholderID, FirstName: si.Profile.FirstName, MiddleName: si.Profile.MiddleName, LastName: si.Profile.LastName, PersonalEmail: si.Profile.PersonalEmail, CollegeEmail: si.Profile.CollegeEmail, PhoneNumber: si.Profile.PhoneNumber, AlternatePhoneNumber: si.Profile.AlternatePhoneNumber, CollegeID: si.Profile.CollegeID, Gender: si.Profile.Gender, DateOfBirth: si.Profile.DateOfBirth, AadharNumber: si.Profile.AadharNumber, PermanentAddressLine1: si.Profile.PermanentAddressLine1, PermanentAddressLine2: si.Profile.PermanentAddressLine2, PermanentAddressLine3: si.Profile.PermanentAddressLine3, PermanentAddressCountry: si.Profile.PermanentAddressCountry, PermanentAddressState: si.Profile.PermanentAddressState, PermanentAddressCity: si.Profile.PermanentAddressCity, PermanentAddressDistrict: si.Profile.PermanentAddressDistrict, PermanentAddressZipcode: si.Profile.PermanentAddressZipcode, PermanentAddressPhone: si.Profile.PermanentAddressPhone, PresentAddressLine1: si.Profile.PresentAddressLine1, PresentAddressLine2: si.Profile.PresentAddressLine2, PresentAddressLine3: si.Profile.PresentAddressLine3, PresentAddressCountry: si.Profile.PresentAddressCountry, PresentAddressState: si.Profile.PresentAddressState, PresentAddressCity: si.Profile.PresentAddressCity, PresentAddressDistrict: si.Profile.PresentAddressDistrict, PresentAddressZipcode: si.Profile.PresentAddressZipcode, PresentAddressPhone: si.Profile.PresentAddressPhone, UniversityName: si.Profile.UniversityName, UniversityID: si.Profile.UniversityID, ProfilePicture: si.Profile.ProfilePicture, DateOfJoining: si.Profile.DateOfJoining}
}

// GetContactFromProfile ....
func (si *StudentProfileVerificationDataModel) GetContactFromProfile() {
	si.ContactInfo = StudentContactInfoModel{StakeholderID: si.Profile.StakeholderID, FirstName: si.Profile.FirstName, MiddleName: si.Profile.MiddleName, LastName: si.Profile.LastName, PersonalEmail: si.Profile.PersonalEmail, CollegeEmail: si.Profile.CollegeEmail, PhoneNumber: si.Profile.PhoneNumber, AlternatePhoneNumber: si.Profile.AlternatePhoneNumber, CollegeID: si.Profile.CollegeID, Gender: si.Profile.Gender, DateOfBirth: si.Profile.DateOfBirth, AadharNumber: si.Profile.AadharNumber, PermanentAddressLine1: si.Profile.PermanentAddressLine1, PermanentAddressLine2: si.Profile.PermanentAddressLine2, PermanentAddressLine3: si.Profile.PermanentAddressLine3, PermanentAddressCountry: si.Profile.PermanentAddressCountry, PermanentAddressState: si.Profile.PermanentAddressState, PermanentAddressCity: si.Profile.PermanentAddressCity, PermanentAddressDistrict: si.Profile.PermanentAddressDistrict, PermanentAddressZipcode: si.Profile.PermanentAddressZipcode, PermanentAddressPhone: si.Profile.PermanentAddressPhone, PresentAddressLine1: si.Profile.PresentAddressLine1, PresentAddressLine2: si.Profile.PresentAddressLine2, PresentAddressLine3: si.Profile.PresentAddressLine3, PresentAddressCountry: si.Profile.PresentAddressCountry, PresentAddressState: si.Profile.PresentAddressState, PresentAddressCity: si.Profile.PresentAddressCity, PresentAddressDistrict: si.Profile.PresentAddressDistrict, PresentAddressZipcode: si.Profile.PresentAddressZipcode, PresentAddressPhone: si.Profile.PresentAddressPhone, UniversityName: si.Profile.UniversityName, UniversityID: si.Profile.UniversityID, DateOfJoining: si.Profile.DateOfJoining, ProfilePicture: si.Profile.ProfilePicture}
}
