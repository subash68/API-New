// Package models ...
package models

import "fmt"

// DbModelKeys ...
var DbModelKeys map[string]string

// GetDbKey ...
func GetDbKey(apiKey string) (string, bool) {
	DbModelKeys = map[string]string{
		"corporateName":                        "Corporate_Name",
		"CIN":                                  "Corporate_CIN",
		"corporateHQAddressLine1":              "CorporateHQAddress_Line1",
		"corporateHQAddressLine2":              "CorporateHQAddress_Line2",
		"corporateHQAddressLine3":              "CorporateHQAddress_Line3",
		"corporateHQAddressCountry":            "CorporateHQAddress_Country",
		"corporateHQAddressState":              "CorporateHQAddress_State",
		"corporateHQAddressCity":               "CorporateHQAddress_City",
		"corporateHQAddressDistrict":           "CorporateHQAddress_District",
		"corporateHQAddressZipCode":            "CorporateHQAddress_ZipCode",
		"corporateHQAddressPhone":              "CorporateHQAddress_Phone",
		"corporateHQAddressEmail":              "CorporateHQAddress_Email",
		"corporateLocalBranchAddressLine1":     "CorporateLocal_BranchAddress_Line1",
		"corporateLocalBranchAddressLine2":     "CorporateLocal_BranchAddress_Line2",
		"corporateLocalBranchAddressLine3":     "CorporateLocal_BranchAddress_Line3",
		"corporateLocalBranchAddressCountry":   "CorporateLocal_BranchAddress_Country",
		"corporateLocalBranchAddressState":     "CorporateLocal_BranchAddress_State",
		"corporateLocalBranchAddressCity":      "CorporateLocal_BranchAddress_City",
		"corporateLocalBranchAddressDistrict":  "CorporateLocal_BranchAddress_District",
		"corporateLocalBranchAddressZipCode":   "CorporateLocal_BranchAddress_ZipCode",
		"corporateLocalBranchAddressPhone":     "CorporateLocal_BranchAddress_Phone",
		"corporateLocalBranchAddressEmail":     "CorporateLocal_BranchAddress_Email",
		"primaryContactFirstName":              "PrimaryContact_FirstName",
		"primaryContactMiddleName":             "PrimaryContact_MiddleName",
		"primaryContactLastName":               "PrimaryContact_LastName",
		"primaryContactDesignation":            "PrimaryContact_Designation",
		"secondaryContactFirstName":            "SecondaryContact_FirstName",
		"secondaryContactMiddleName":           "SecondaryContact_MiddleName",
		"secondaryContactLastName":             "SecondaryContact_LastName",
		"secondaryContactDesignation":          "SecondaryContact_Designation",
		"secondaryContactPhone":                "SecondaryContact_Phone",
		"secondaryContactEmail":                "SecondaryContact_Email",
		"corporateIndustry":                    "CorporateIndustry",
		"companyProfile":                       "CompanyProfile",
		"universityName":                       "University_Name",
		"universityID":                         "UniversityID",
		"universityCollegeID":                  "University_College_ID",
		"universityHQAddressLine1":             "UniversityHQAddress_Line1",
		"universityHQAddressLine2":             "UniversityHQAddress_Line2",
		"universityHQAddressLine3":             "UniversityHQAddress_Line3",
		"universityHQAddressCountry":           "UniversityHQAddress_Country",
		"universityHQAddressState":             "UniversityHQAddress_State",
		"universityHQAddressCity":              "UniversityHQAddress_City",
		"universityHQAddressDistrict":          "UniversityHQAddress_District",
		"universityHQAddressZipcode":           "UniversityHQAddress_Zipcode",
		"universityHQAddressPhone":             "UniversityHQAddress_Phone",
		"universityHQAddressemail":             "UniversityHQAddress_Email",
		"universityLocalBranchAddressLine1":    "UniversityLocal_BranchAddress_Line1",
		"universityLocalBranchAddressLine2":    "UniversityLocal_BranchAddress_Line2",
		"universityLocalBranchAddressLine3":    "UniversityLocal_BranchAddress_Line3",
		"universityLocalBranchAddressCountry":  "UniversityLocal_BranchAddress_Country",
		"universityLocalBranchAddressState":    "UniversityLocal_BranchAddress_State",
		"universityLocalBranchAddressCity":     "UniversityLocal_BranchAddress_City",
		"universityLocalBranchAddressDistrict": "UniversityLocal_BranchAddress_District",
		"universityLocalBranchAddressZipcode":  "UniversityLocal_BranchAddress_Zipcode",
		"universityLocalBranchAddressPhone":    "UniversityLocal_BranchAddress_Phone",
		"universityLocalBranchAddressemail":    "UniversityLocal_BranchAddress_Email",
		"universityProfile":                    "UniversityProfile",
		"noOfPrograms":                         "NoOfPrograms",
		"noOfBranches":                         "NoOfBranches",
		"TotalNoOfStudents":                    "TotalNoOfStudents",
		"noOfCOEs":                             "NoOfCOEs",
		"specialOffersFlag":                    "SpecialOffersFlag",
		"firstName":                            "Student_FirstName",
		"middleName":                           "Student_MiddleName",
		"lastName":                             "Student_LastName",
		"personalEmail":                        "Student_PersonalEmailID",
		"alternatePhoneNumber":                 "Student_AlternateContactNumber",
		"gender":                               "Student_Gender",
		"dateOfBirth":                          "Student_DateOfBirth",
		"aadharNumber":                         "Student_AadharNumber",
		"permanentAddressLine1":                "StudentPermanantAddress_Line1",
		"permanentAddressLine2":                "StudentPermanantAddress_Line2",
		"permanentAddressLine3":                "StudentPermanantAddress_Line3",
		"permanentAddressCountry":              "StudentPermanantAddress_Country",
		"permanentAddressState":                "StudentPermanantAddress_State",
		"permanentAddressCity":                 "StudentPermanantAddress_City",
		"permanentAddressDistrict":             "StudentPermanantAddress_District",
		"permanentAddressZipcode":              "StudentPermanantAddress_Zipcode",
		"permanentAddressPhone":                "StudentPermanantAddress_Phone",
		"presentAddressLine1":                  "StudentPresentAddress_Line1",
		"presentAddressLine2":                  "StudentPresentAddress_Line2",
		"presentAddressLine3":                  "StudentPresentAddress_Line3",
		"presentAddressCountry":                "StudentPresentAddress_Country",
		"presentAddressState":                  "StudentPresentAddress_State",
		"presentAddressCity":                   "StudentPresentAddress_City",
		"presentAddressDistrict":               "StudentPresentAddress_District",
		"presentAddressZipcode":                "StudentPresentAddress_Zipcode",
		"presentAddressPhone":                  "StudentPresentAddress_Phone",
		"fathersGuardianFullName":              "Father_Guardian_FullName",
		"fathersGuardianOccupation":            "Father_Guardian_Occupation",
		"fathersGuardianCompany":               "Father_Guardian_Company",
		"fathersGuardianPhoneNumber":           "Father_Guardian_PhoneNumber",
		"fathersGuardianEmailID":               "Father_Guardian_EmailID",
		"mothersGuardianFullName":              "Mother_Guardian_FullName",
		"mothersGuardianOccupation":            "Mother_Guardian_Occupation",
		"mothersGuardianCompany":               "Mother_Guardian_Comany",
		"mothersGuardianDesignation":           "Mother_Guardian_Designation",
		"mothersGuardianPhoneNumber":           "Mother_Guardian_PhoneNumber",
		"mothersGuardianEmailID":               "Mother_Guardian_EmailID",
		"aboutMe":                              "Student_AboutMe",
		"profilePicture":                       "ProfilePicture",
		"StudentUniversityID":                  "University_Stakeholder_ID",
		"programName":                          "ProgramName",
		"programID":                            "Program_ID",
		"branchName":                           "BranchName",
		"branchID":                             "Branch_ID",
		"collegeRollNumber":                    "Student_CollegeRollNo",
		"collegeEmailID":                       "Student_CollegeEmailID",
	}
	key := DbModelKeys[apiKey]
	fmt.Println("-----> ", key, key == "")
	if key == "" {
		return "", false
	}
	return key, true
}
