package models

import (
	"database/sql"
	"fmt"
)

// ScanTNB ...
func (ald *AllLutData) ScanTNB(rows *sql.Rows, err error) error {
	var l1b Lut10BoardsModel
	if err != nil {
		return err
	}
	for rows.Next() {
		err := rows.Scan(&l1b.BoardName, &l1b.CertificateName, &l1b.BoardID)
		if err != nil {
			log.Fatalf("%s, Error: %v ", ScanError, err)
			return fmt.Errorf("Failed to get 'tenthBoards'")
		}
		ald.TenthBoards = append(ald.TenthBoards, l1b)
	}
	cacheInMemoryStore.Cache.Set("tenthBoards", ald.TenthBoards, cacheDuration)
	return nil
}

// ScanTWB ...
func (ald *AllLutData) ScanTWB(rows *sql.Rows, err error) error {
	var l1b Lut10BoardsModel
	if err != nil {
		return err
	}
	for rows.Next() {
		err := rows.Scan(&l1b.BoardName, &l1b.CertificateName, &l1b.BoardID)
		if err != nil {
			log.Fatalf("%s, Error: %v ", ScanError, err)
			return fmt.Errorf("Failed to get 'twelfthBoards'")
		}
		ald.TwelfthBoards = append(ald.TwelfthBoards, l1b)
	}
	cacheInMemoryStore.Cache.Set("twelfthBoards", ald.TwelfthBoards, cacheDuration)
	return nil
}

// ScanAS ...
func (ald *AllLutData) ScanAS(rows *sql.Rows, err error) error {
	var ld LutAccountStatusModel
	if err != nil {
		return err
	}
	for rows.Next() {
		err := rows.Scan(&ld.AccountStatus, &ld.AccountStatusCode)
		if err != nil {
			log.Fatalf("%s, Error: %v ", ScanError, err)
			return fmt.Errorf("Failed to get 'accountStatus'")
		}
		ald.AccountStatus = append(ald.AccountStatus, ld)
	}
	cacheInMemoryStore.Cache.Set("accountStatus", ald.AccountStatus, cacheDuration)
	return nil
}

// ScanBranchCatalog ...
func (ald *AllLutData) ScanBranchCatalog(rows *sql.Rows, err error) error {
	var ld LutBranchCatalogModel
	if err != nil {
		return err
	}
	for rows.Next() {
		err := rows.Scan(&ld.BranchID, &ld.BranchName, &ld.Duration, &ld.ProgramID, &ld.ProgramType)
		if err != nil {
			log.Fatalf("%s, Error: %v ", ScanError, err)
			return fmt.Errorf("Failed to get 'branchCatalog'")
		}
		ald.BranchCatalog = append(ald.BranchCatalog, ld)
	}
	cacheInMemoryStore.Cache.Set("branchCatalog", ald.BranchCatalog, cacheDuration)
	return nil
}

// ScanCoporateCategory ...
func (ald *AllLutData) ScanCoporateCategory(rows *sql.Rows, err error) error {
	var ld LutCorporateCategoryModel
	if err != nil {
		return err
	}
	for rows.Next() {
		err := rows.Scan(&ld.Name, &ld.Code, &ld.OneLtrCode)
		if err != nil {
			log.Fatalf("%s, Error: %v ", ScanError, err)
			return fmt.Errorf("Failed to get 'corporateCategory'")
		}
		ald.CoporateCategory = append(ald.CoporateCategory, ld)
	}
	cacheInMemoryStore.Cache.Set("corporateCategory", ald.CoporateCategory, cacheDuration)
	return nil
}

// ScanCoporateIndustry ...
func (ald *AllLutData) ScanCoporateIndustry(rows *sql.Rows, err error) error {
	var ld LutCorporateIndustryModel
	if err != nil {
		return err
	}
	for rows.Next() {
		err := rows.Scan(&ld.Name, &ld.Code)
		if err != nil {
			log.Fatalf("%s, Error: %v ", ScanError, err)
			return fmt.Errorf("Failed to get 'corporateIndustry'")
		}
		ald.CoporateIndustry = append(ald.CoporateIndustry, ld)
	}
	cacheInMemoryStore.Cache.Set("corporateIndustry", ald.CoporateIndustry, cacheDuration)
	return nil
}

// ScanCoporateType ...
func (ald *AllLutData) ScanCoporateType(rows *sql.Rows, err error) error {
	var ld LutCorporateTypeModel
	if err != nil {
		return err
	}
	for rows.Next() {
		err := rows.Scan(&ld.Code, &ld.Name, &ld.OneLtrCode)
		if err != nil {
			log.Fatalf("%s, Error: %v ", ScanError, err)
			return fmt.Errorf("Failed to get 'corporateType'")
		}
		ald.CorporateType = append(ald.CorporateType, ld)
	}
	cacheInMemoryStore.Cache.Set("corporateType", ald.CorporateType, cacheDuration)
	return nil
}

// ScanJobType ...
func (ald *AllLutData) ScanJobType(rows *sql.Rows, err error) error {
	var ld LutJobTypeModel
	if err != nil {
		return err
	}
	for rows.Next() {
		err := rows.Scan(&ld.Name, &ld.Code)
		if err != nil {
			log.Fatalf("%s, Error: %v ", ScanError, err)
			return fmt.Errorf("Failed to get 'jobType'")
		}
		ald.JobType = append(ald.JobType, ld)
	}
	cacheInMemoryStore.Cache.Set("jobType", ald.JobType, cacheDuration)
	return nil
}

// ScanLangProf ...
func (ald *AllLutData) ScanLangProf(rows *sql.Rows, err error) error {
	var ld LutLangProficiencyModel
	if err != nil {
		return err
	}
	for rows.Next() {
		err := rows.Scan(&ld.Name, &ld.Code)
		if err != nil {
			log.Fatalf("%s, Error: %v ", ScanError, err)
			return fmt.Errorf("Failed to get 'languageProficiency'")
		}
		ald.LanguageProficiency = append(ald.LanguageProficiency, ld)
	}
	cacheInMemoryStore.Cache.Set("languageProficiency", ald.LanguageProficiency, cacheDuration)
	return nil
}

// ScanModeOfIssue ...
func (ald *AllLutData) ScanModeOfIssue(rows *sql.Rows, err error) error {
	var ld LutModeOfTokenIssueModel
	if err != nil {
		return err
	}
	for rows.Next() {
		err := rows.Scan(&ld.Name, &ld.Code)
		if err != nil {
			log.Fatalf("%s, Error: %v ", ScanError, err)
			return fmt.Errorf("Failed to get 'modeOfTokenIssue'")
		}
		ald.ModeOfTokenIssue = append(ald.ModeOfTokenIssue, ld)
	}
	cacheInMemoryStore.Cache.Set("modeOfTokenIssue", ald.ModeOfTokenIssue, cacheDuration)
	return nil
}

// ScanNftType ...
func (ald *AllLutData) ScanNftType(rows *sql.Rows, err error) error {
	var ld LutNotificationTypeModel
	if err != nil {
		return err
	}
	for rows.Next() {
		err := rows.Scan(&ld.Name, &ld.Code)
		if err != nil {
			log.Fatalf("%s, Error: %v ", ScanError, err)
			return fmt.Errorf("Failed to get 'notificationType'")
		}
		ald.NotificationType = append(ald.NotificationType, ld)
	}
	cacheInMemoryStore.Cache.Set("notificationType", ald.NotificationType, cacheDuration)
	return nil
}

// ScanPaymentMode ...
func (ald *AllLutData) ScanPaymentMode(rows *sql.Rows, err error) error {
	var ld LutPaymentModeModel
	if err != nil {
		return err
	}
	for rows.Next() {
		err := rows.Scan(&ld.Name, &ld.Code)
		if err != nil {
			log.Fatalf("%s, Error: %v ", ScanError, err)
			return fmt.Errorf("Failed to get 'paymentMode'")
		}
		ald.PaymentMode = append(ald.PaymentMode, ld)
	}
	cacheInMemoryStore.Cache.Set("paymentMode", ald.PaymentMode, cacheDuration)
	return nil
}

// ScanProgramCatalog ...
func (ald *AllLutData) ScanProgramCatalog(rows *sql.Rows, err error) error {
	var ld LutProgramCatalogModel
	if err != nil {
		return err
	}
	for rows.Next() {
		err := rows.Scan(&ld.Code, &ld.Name, &ld.Type)
		if err != nil {
			log.Fatalf("%s, Error: %v ", ScanError, err)
			return fmt.Errorf("Failed to get 'programCatalog'")
		}
		ald.ProgramCatalog = append(ald.ProgramCatalog, ld)
	}
	cacheInMemoryStore.Cache.Set("programCatalog", ald.ProgramCatalog, cacheDuration)
	return nil
}

// ScanProgramType ...
func (ald *AllLutData) ScanProgramType(rows *sql.Rows, err error) error {
	var ld LutProgramTypeModel
	if err != nil {
		return err
	}
	for rows.Next() {
		err := rows.Scan(&ld.Name, &ld.Code)
		if err != nil {
			log.Fatalf("%s, Error: %v ", ScanError, err)
			return fmt.Errorf("Failed to get 'programType'")
		}
		ald.ProgramType = append(ald.ProgramType, ld)
	}
	cacheInMemoryStore.Cache.Set("programType", ald.ProgramType, cacheDuration)
	return nil
}

// ScanSkillProficiency ...
func (ald *AllLutData) ScanSkillProficiency(rows *sql.Rows, err error) error {
	var ld LutSkillProficiencyModel
	if err != nil {
		return err
	}
	for rows.Next() {
		err := rows.Scan(&ld.Name, &ld.Code)
		if err != nil {
			log.Fatalf("%s, Error: %v ", ScanError, err)
			return fmt.Errorf("Failed to get 'skillProficiency'")
		}
		ald.SkillProficiency = append(ald.SkillProficiency, ld)
	}
	cacheInMemoryStore.Cache.Set("skillProficiency", ald.SkillProficiency, cacheDuration)
	return nil
}

// ScanSkills ...
func (ald *AllLutData) ScanSkills(rows *sql.Rows, err error) error {
	var ld LutSkillsModel
	if err != nil {
		return err
	}
	for rows.Next() {
		err := rows.Scan(&ld.Code, &ld.Name, &ld.Disabled)
		if err != nil {
			log.Fatalf("%s, Error: %v ", ScanError, err)
			return fmt.Errorf("Failed to get 'skills'")
		}
		ald.Skills = append(ald.Skills, ld)
	}
	cacheInMemoryStore.Cache.Set("skills", ald.Skills, cacheDuration)
	return nil
}

// ScanSortBy ...
func (ald *AllLutData) ScanSortBy(rows *sql.Rows, err error) error {
	var ld LutSortByModel
	if err != nil {
		return err
	}
	for rows.Next() {
		err := rows.Scan(&ld.Name, &ld.Code)
		if err != nil {
			log.Fatalf("%s, Error: %v ", ScanError, err)
			return fmt.Errorf("Failed to get 'sortBy'")
		}
		ald.SortBy = append(ald.SortBy, ld)
	}
	cacheInMemoryStore.Cache.Set("sortBy", ald.SortBy, cacheDuration)
	return nil
}

// ScanStakeholderTypes ...
func (ald *AllLutData) ScanStakeholderTypes(rows *sql.Rows, err error) error {
	var ld LutStakeholdersModel
	if err != nil {
		return err
	}
	for rows.Next() {
		err := rows.Scan(&ld.Name, &ld.Code, &ld.OneLtrCode)
		if err != nil {
			log.Fatalf("%s, Error: %v ", ScanError, err)
			return fmt.Errorf("Failed to get 'stakeholderType'")
		}
		ald.StakeholderType = append(ald.StakeholderType, ld)
	}
	cacheInMemoryStore.Cache.Set("stakeholderType", ald.StakeholderType, cacheDuration)
	return nil
}

// ScanStudentEventResult ...
func (ald *AllLutData) ScanStudentEventResult(rows *sql.Rows, err error) error {
	var ld LutStudentEventResultModel
	if err != nil {
		return err
	}
	for rows.Next() {
		err := rows.Scan(&ld.Name, &ld.Code)
		if err != nil {
			log.Fatalf("%s, Error: %v ", ScanError, err)
			return fmt.Errorf("Failed to get 'studentEventResult'")
		}
		ald.StudentEventResult = append(ald.StudentEventResult, ld)
	}
	cacheInMemoryStore.Cache.Set("studentEventResult", ald.StudentEventResult, cacheDuration)
	return nil
}

// ScanStudentEvent ...
func (ald *AllLutData) ScanStudentEvent(rows *sql.Rows, err error) error {
	var ld LutStudentEventModel
	if err != nil {
		return err
	}
	for rows.Next() {
		err := rows.Scan(&ld.Name, &ld.Code)
		if err != nil {
			log.Fatalf("%s, Error: %v ", ScanError, err)
			return fmt.Errorf("Failed to get 'studentEvent'")
		}
		ald.StudentEvent = append(ald.StudentEvent, ld)
	}
	cacheInMemoryStore.Cache.Set("studentEvent", ald.StudentEvent, cacheDuration)
	return nil
}

// ScanStudentVrfStatus ...
func (ald *AllLutData) ScanStudentVrfStatus(rows *sql.Rows, err error) error {
	var ld LutStudentVerfStatusModel
	if err != nil {
		return err
	}
	for rows.Next() {
		err := rows.Scan(&ld.Name, &ld.Code)
		if err != nil {
			log.Fatalf("%s, Error: %v ", ScanError, err)
			return fmt.Errorf("Failed to get 'studentVerificationStatus'")
		}
		ald.StudentVerificationStatus = append(ald.StudentVerificationStatus, ld)
	}
	cacheInMemoryStore.Cache.Set("studentVerificationStatus", ald.StudentVerificationStatus, cacheDuration)
	return nil
}

// ScanStudentVrfType ...
func (ald *AllLutData) ScanStudentVrfType(rows *sql.Rows, err error) error {
	var ld LutStudentVerfTypeModel
	if err != nil {
		return err
	}
	for rows.Next() {
		err := rows.Scan(&ld.Name, &ld.Code)
		if err != nil {
			log.Fatalf("%s, Error: %v ", ScanError, err)
			return fmt.Errorf("Failed to get 'studentVerificationType'")
		}
		ald.StudentVerificationType = append(ald.StudentVerificationType, ld)
	}
	cacheInMemoryStore.Cache.Set("studentVerificationType", ald.StudentVerificationType, cacheDuration)
	return nil
}

// ScanSubscriptionType ...
func (ald *AllLutData) ScanSubscriptionType(rows *sql.Rows, err error) error {
	var ld LutSubscriptionType
	if err != nil {
		return err
	}
	for rows.Next() {
		err := rows.Scan(&ld.Name, &ld.Code)
		if err != nil {
			log.Fatalf("%s, Error: %v ", ScanError, err)
			return fmt.Errorf("Failed to get 'subscriptionType'")
		}
		ald.SubscriptionType = append(ald.SubscriptionType, ld)
	}
	cacheInMemoryStore.Cache.Set("subscriptionType", ald.SubscriptionType, cacheDuration)
	return nil
}

// ScanTokenEvent ...
func (ald *AllLutData) ScanTokenEvent(rows *sql.Rows, err error) error {
	var ld LutTokenEvent
	if err != nil {
		return err
	}
	for rows.Next() {
		err := rows.Scan(&ld.Name, &ld.Code, &ld.StakeholderType)
		if err != nil {
			log.Fatalf("%s, Error: %v ", ScanError, err)
			return fmt.Errorf("Failed to get 'tokenEvent'")
		}
		ald.TokenEvent = append(ald.TokenEvent, ld)
	}
	cacheInMemoryStore.Cache.Set("tokenEvent", ald.TokenEvent, cacheDuration)
	return nil
}

// ScanUniversityAccreditation ...
func (ald *AllLutData) ScanUniversityAccreditation(rows *sql.Rows, err error) error {
	var ld LutUniversityAccreditationModel
	if err != nil {
		return err
	}
	for rows.Next() {
		err := rows.Scan(&ld.Name, &ld.Code)
		if err != nil {
			log.Fatalf("%s, Error: %v ", ScanError, err)
			return fmt.Errorf("Failed to get 'universityAccreditation'")
		}
		ald.UniversityAccreditation = append(ald.UniversityAccreditation, ld)
	}
	cacheInMemoryStore.Cache.Set("universityAccreditation", ald.UniversityAccreditation, cacheDuration)
	return nil
}

// ScanUniversityCatalog ...
func (ald *AllLutData) ScanUniversityCatalog(rows *sql.Rows, err error) error {
	var ld LutUniversityCatalogModel
	if err != nil {
		return err
	}
	for rows.Next() {
		err := rows.Scan(&ld.Code, &ld.Name, &ld.Address, &ld.Type, &ld.State, &ld.DateOfEst)
		if err != nil {
			log.Fatalf("%s, Error: %v ", ScanError, err)
			return fmt.Errorf("Failed to get 'universityCatalog'")
		}
		ald.UniversityCatalog = append(ald.UniversityCatalog, ld)
	}
	cacheInMemoryStore.Cache.Set("universityCatalog", ald.UniversityCatalog, cacheDuration)
	return nil
}

// ScanUniversityCoe ...
func (ald *AllLutData) ScanUniversityCoe(rows *sql.Rows, err error) error {
	var ld LutUniversityCoeModel
	if err != nil {
		return err
	}
	for rows.Next() {
		err := rows.Scan(&ld.Name, &ld.Code)
		if err != nil {
			log.Fatalf("%s, Error: %v ", ScanError, err)
			return fmt.Errorf("Failed to get 'universityCoe'")
		}
		ald.UniversityCoe = append(ald.UniversityCoe, ld)
	}
	cacheInMemoryStore.Cache.Set("universityCoe", ald.UniversityCoe, cacheDuration)
	return nil
}

// ScanUniversitySplOffType ...
func (ald *AllLutData) ScanUniversitySplOffType(rows *sql.Rows, err error) error {
	var ld LutUniversitySplOffTypeModel
	if err != nil {
		return err
	}
	for rows.Next() {
		err := rows.Scan(&ld.Name, &ld.Code)
		if err != nil {
			log.Fatalf("%s, Error: %v ", ScanError, err)
			return fmt.Errorf("Failed to get 'universitySpecialOfferingType'")
		}
		ald.UniversitySpecialOfferingType = append(ald.UniversitySpecialOfferingType, ld)
	}
	cacheInMemoryStore.Cache.Set("universitySpecialOfferingType", ald.UniversitySpecialOfferingType, cacheDuration)
	return nil
}

// ScanUniversityTieUpType ...
func (ald *AllLutData) ScanUniversityTieUpType(rows *sql.Rows, err error) error {
	var ld LutUniversityTieUpTypeModel
	if err != nil {
		return err
	}
	for rows.Next() {
		err := rows.Scan(&ld.Name, &ld.Code)
		if err != nil {
			log.Fatalf("%s, Error: %v ", ScanError, err)
			return fmt.Errorf("Failed to get 'universityTieUpType'")
		}
		ald.UniversityTieUpType = append(ald.UniversityTieUpType, ld)
	}
	cacheInMemoryStore.Cache.Set("universityTieUpType", ald.UniversityTieUpType, cacheDuration)
	return nil
}

// ScanUniversityType ...
func (ald *AllLutData) ScanUniversityType(rows *sql.Rows, err error) error {
	var ld LutUniversityTypeModel
	if err != nil {
		return err
	}
	for rows.Next() {
		err := rows.Scan(&ld.Name, &ld.Code, &ld.OneLtrCode)
		if err != nil {
			log.Fatalf("%s, Error: %v ", ScanError, err)
			return fmt.Errorf("Failed to get 'universityType'")
		}
		ald.UniversityType = append(ald.UniversityType, ld)
	}
	cacheInMemoryStore.Cache.Set("universityType", ald.UniversityType, cacheDuration)
	return nil
}
