// Package models ...
package models

import (
	"github.com/jaswanth-gorripati/PGK/s0_Lookups/configuration"
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
		"Token_Allocation_ins": "INSERT INTO " + dbConfig.DbDatabaseName + "." + dbConfig.TokenAllocatinoDbName + "(Stakeholder_ID,NoOfTokensAllocated,TokenAllocationDate,Payment_ID,ModeOfIssue) VALUES(?,?,?,?,?)",
	}
}
