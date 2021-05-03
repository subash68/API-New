package models

import (
	"fmt"
	"io/ioutil"
	"mime/multipart"
)

// InsertAcademics ...
func (sa *StudentAcademicsModelReq) InsertAcademics(form *multipart.Form) error {
	ques := ""
	val := []interface{}{}
	val = append(val, sa.StakeholderID)
	insFields := ""
	updSP := ""
	if sa.Tenth.Name != "" && sa.Twelfth.Name != "" {
		return fmt.Errorf("Failed to Parse Data, Required Only one Academic Details per request")
	}
	if sa.Tenth.Name != "" {
		tenthSP, _ := RetriveSP("STU_TENTH_INS")
		tenthUPD, _ := RetriveSP("STU_TENTH_UPD")
		vals := constructTTReq(sa.Tenth, form, "tenth")
		val = append(val, vals...)
		val = append(val, vals...)

		ques += "?,?,?,?,?,?,?"
		insFields += tenthSP
		updSP += tenthUPD
	}
	if sa.Twelfth.Name != "" {
		twelfthSP, _ := RetriveSP("STU_Twelfth_INS")
		tenthUPD, _ := RetriveSP("STU_Twelfth_UPD")
		vals := constructTTReq(sa.Twelfth, form, "twelfth")
		val = append(val, vals...)
		val = append(val, vals...)
		ques += "?,?,?,?,?,?,?"
		insFields += twelfthSP
		updSP += tenthUPD
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

	fmt.Println(val, ques)
	return nil
}

// GetAcademics ...
func GetAcademics(ID string) (StudentAcademicsModelReq, error) {
	var sa StudentAcademicsModelReq
	var tenN StudentNullableTTModel
	var twelfthN StudentNullableTTModel
	getByIDSP, _ := RetriveSP("STU_GET_ACADEMICS")
	err := Db.QueryRow(getByIDSP, ID).Scan(&tenN.Name, &tenN.Location, &tenN.MonthAndYearOfPassing, &tenN.Board, &tenN.Percentage, &tenN.AttachmentFile, &twelfthN.Name, &twelfthN.Location, &twelfthN.MonthAndYearOfPassing, &twelfthN.Board, &twelfthN.Percentage, &twelfthN.AttachmentFile)
	if err != nil {
		return sa, fmt.Errorf("Failed to retrieving Academics : %v", err.Error())
	}
	sa.Tenth = tenN.ConvertToJSONStruct()
	sa.Twelfth = twelfthN.ConvertToJSONStruct()
	return sa, nil
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
	if st.EnablingFlag.Valid {
		stt.EnablingFlag = st.EnablingFlag.Bool
	} else {
		return emptyStt
	}
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
	vals = append(vals, st.Name, st.Location, st.MonthAndYearOfPassing, st.Board, st.Percentage, st.AttachmentFile, true)
	return vals
}

// GetContactFromProfile ....
func (si *StudentCompleteProfileDataModel) GetContactFromProfile() {
	si.ContactInfo = StudentContactInfoModel{StakeholderID: si.Profile.StakeholderID, FirstName: si.Profile.FirstName, MiddleName: si.Profile.MiddleName, LastName: si.Profile.LastName, PersonalEmail: si.Profile.PersonalEmail, CollegeEmail: si.Profile.CollegeEmail, PhoneNumber: si.Profile.PhoneNumber, AlternatePhoneNumber: si.Profile.AlternatePhoneNumber, CollegeID: si.Profile.CollegeID, Gender: si.Profile.Gender, DateOfBirth: si.Profile.DateOfBirth, AadharNumber: si.Profile.AadharNumber, PermanentAddressLine1: si.Profile.PermanentAddressLine1, PermanentAddressLine2: si.Profile.PermanentAddressLine2, PermanentAddressLine3: si.Profile.PermanentAddressLine3, PermanentAddressCountry: si.Profile.PermanentAddressCountry, PermanentAddressState: si.Profile.PermanentAddressState, PermanentAddressCity: si.Profile.PermanentAddressCity, PermanentAddressDistrict: si.Profile.PermanentAddressDistrict, PermanentAddressZipcode: si.Profile.PermanentAddressZipcode, PermanentAddressPhone: si.Profile.PermanentAddressPhone, PresentAddressLine1: si.Profile.PresentAddressLine1, PresentAddressLine2: si.Profile.PresentAddressLine2, PresentAddressLine3: si.Profile.PresentAddressLine3, PresentAddressCountry: si.Profile.PresentAddressCountry, PresentAddressState: si.Profile.PresentAddressState, PresentAddressCity: si.Profile.PresentAddressCity, PresentAddressDistrict: si.Profile.PresentAddressDistrict, PresentAddressZipcode: si.Profile.PresentAddressZipcode, PresentAddressPhone: si.Profile.PresentAddressPhone, AboutMe: si.Profile.AboutMe, UniversityName: si.Profile.UniversityName, UniversityID: si.Profile.UniversityID}
}

// GetContactFromProfile ....
func (si *StudentProfileVerificationDataModel) GetContactFromProfile() {
	si.ContactInfo = StudentContactInfoModel{StakeholderID: si.Profile.StakeholderID, FirstName: si.Profile.FirstName, MiddleName: si.Profile.MiddleName, LastName: si.Profile.LastName, PersonalEmail: si.Profile.PersonalEmail, CollegeEmail: si.Profile.CollegeEmail, PhoneNumber: si.Profile.PhoneNumber, AlternatePhoneNumber: si.Profile.AlternatePhoneNumber, CollegeID: si.Profile.CollegeID, Gender: si.Profile.Gender, DateOfBirth: si.Profile.DateOfBirth, AadharNumber: si.Profile.AadharNumber, PermanentAddressLine1: si.Profile.PermanentAddressLine1, PermanentAddressLine2: si.Profile.PermanentAddressLine2, PermanentAddressLine3: si.Profile.PermanentAddressLine3, PermanentAddressCountry: si.Profile.PermanentAddressCountry, PermanentAddressState: si.Profile.PermanentAddressState, PermanentAddressCity: si.Profile.PermanentAddressCity, PermanentAddressDistrict: si.Profile.PermanentAddressDistrict, PermanentAddressZipcode: si.Profile.PermanentAddressZipcode, PermanentAddressPhone: si.Profile.PermanentAddressPhone, PresentAddressLine1: si.Profile.PresentAddressLine1, PresentAddressLine2: si.Profile.PresentAddressLine2, PresentAddressLine3: si.Profile.PresentAddressLine3, PresentAddressCountry: si.Profile.PresentAddressCountry, PresentAddressState: si.Profile.PresentAddressState, PresentAddressCity: si.Profile.PresentAddressCity, PresentAddressDistrict: si.Profile.PresentAddressDistrict, PresentAddressZipcode: si.Profile.PresentAddressZipcode, PresentAddressPhone: si.Profile.PresentAddressPhone, AboutMe: si.Profile.AboutMe, UniversityName: si.Profile.UniversityName, UniversityID: si.Profile.UniversityID}
}
