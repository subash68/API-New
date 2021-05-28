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
	return nil
}

// ScanCoporateType ...
func (ald *AllLutData) ScanCoporateType(rows *sql.Rows, err error) error {
	var ld LutCorporateTypeModel
	if err != nil {
		return err
	}
	for rows.Next() {
		err := rows.Scan(&ld.Name, &ld.Code, &ld.OneLtrCode)
		if err != nil {
			log.Fatalf("%s, Error: %v ", ScanError, err)
			return fmt.Errorf("Failed to get 'corporateType'")
		}
		ald.CorporateType = append(ald.CorporateType, ld)
	}
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
			return fmt.Errorf("Failed to get 'paymentMode'")
		}
		ald.ProgramCatalog = append(ald.ProgramCatalog, ld)
	}
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
			return fmt.Errorf("Failed to get 'skillProficiency'")
		}
		ald.Skills = append(ald.Skills, ld)
	}
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
	return nil
}
