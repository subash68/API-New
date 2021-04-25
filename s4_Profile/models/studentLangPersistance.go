package models

import (
	"database/sql"
	"fmt"
	"time"
)

// InsertLanguages ....
func (sl *StudentAllLanguagesModel) InsertLanguages() error {
	// Preparing Database insert
	slInsertCmd, _ := RetriveSP("STU_LANG_INS")

	vals := []interface{}{}

	for _, lng := range sl.Languages {
		slInsertCmd += "(?,?,?,?,?,?,?),"
		vals = append(vals, sl.StakeholderID, lng.LanguageName, lng.SpeakProficiency, lng.IsMotherTongue, lng.ReadProficiency, lng.WriteProficiency, true)
	}
	slInsertCmd = slInsertCmd[0 : len(slInsertCmd)-1]
	stmt, err := Db.Prepare(slInsertCmd)
	if err != nil {
		return fmt.Errorf("Cannot prepare -- %v , %v -- insert due to %v", slInsertCmd, vals, err.Error())
	}
	_, err = stmt.Exec(vals...)
	if err != nil {
		return fmt.Errorf("Failed to insert in database -- %v , %v -- insert due to %v", slInsertCmd, vals, err.Error())
	}
	return nil
}

// GetAllLanguages ....
func (sl *StudentAllLanguagesModel) GetAllLanguages() error {
	// Preparing Database insert
	slInsertCmd, _ := RetriveSP("STU_LANG_GETALL")

	slRows, err := Db.Query(slInsertCmd, sl.StakeholderID)
	if err != nil && err != sql.ErrNoRows {
		return fmt.Errorf("Cannot get the Rows %v", err.Error())
	} else if err == sql.ErrNoRows {
		return nil
	}
	defer slRows.Close()
	for slRows.Next() {
		var newSl StudentLangModel
		err = slRows.Scan(&newSl.ID, &newSl.LanguageName, &newSl.SpeakProficiency, &newSl.IsMotherTongue, &newSl.ReadProficiency, &newSl.WriteProficiency, &newSl.EnabledFlag, &newSl.CreationDate, &newSl.LastUpdatedDate)
		if err != nil {
			return fmt.Errorf("Cannot read the Rows %v", err.Error())
		}
		sl.Languages = append(sl.Languages, newSl)
	}

	return nil
}

// UpdateLanguage ....
func (sl *StudentLangModel) UpdateLanguage() error {
	// Preparing Database insert
	slInsertCmd, _ := RetriveSP("STU_LANG_UPD")

	stmt, err := Db.Prepare(slInsertCmd)
	if err != nil {
		return fmt.Errorf("Cannot prepare -- %v -- insert due to %v", slInsertCmd, err.Error())
	}

	_, err = stmt.Exec(sl.LanguageName, sl.SpeakProficiency, sl.IsMotherTongue, sl.ReadProficiency, sl.WriteProficiency, time.Now(), sl.ID, sl.StakeholderID)
	if err != nil {
		return fmt.Errorf("Failed to update in database -- %v  -- insert due to %v", slInsertCmd, err.Error())
	}
	return nil
}

// DeleteLanguage ....
func (sl *StudentLangModel) DeleteLanguage() error {
	// Preparing Database insert
	slInsertCmd, _ := RetriveSP("STU_LANG_DLT")

	stmt, err := Db.Prepare(slInsertCmd)
	if err != nil {
		return fmt.Errorf("Cannot prepare -- %v -- insert due to %v", slInsertCmd, err.Error())
	}
	_, err = stmt.Exec(sl.ID, sl.StakeholderID)
	if err != nil {
		return fmt.Errorf("Failed to update in database -- %v  -- insert due to %v", slInsertCmd, err.Error())
	}
	return nil
}
