package models

// CheckPing Checks if connection exists with Database
func CheckPing() error {
	err := Db.Ping()
	if err != nil {
		InitDataModel()
		err1 := Db.Ping()
		if err1 != nil {
			log.Fatalf("Failed to connect with Database, Error : %v", err1)
			return err1
		}
	}
	return nil
}
