// Package models ...
package models

import (
	"github.com/jaswanth-gorripati/PGK/s7_PaymentGateway/configuration"
)

// SP is a Mapping of procedure name to Procedure query
var SP map[string]string

// RetriveSP takes in the required Proceder name and returns the procedure
func RetriveSP(procedureName string) (string, bool) {
	if SP[procedureName] == "" {
		return "", false
	}
	return SP[procedureName], true
}

// CreateSP Creates default stored procedures for Database
func CreateSP() {
	dbConfig := configuration.DbConfig()
	SP = map[string]string{
		"Payment_Insert":     "INSERT INTO " + dbConfig.DbDatabaseName + ".RPLCE (Stakeholder_ID,AmountPaid,Payment_ID,Payment_Mode,PaymentDate) VALUES(?,?,?,?,?)",
		"GET_Payments_BY_ID": "SELECT AmountPaid,Payment_ID,Payment_Mode,PaymentDate FROM " + dbConfig.DbDatabaseName + ".RPLCE WHERE Stakeholder_ID=? GROUP BY AmountPaid,Payment_ID,Payment_Mode,PaymentDate ORDER BY PaymentDate DESC ",
	}
}
