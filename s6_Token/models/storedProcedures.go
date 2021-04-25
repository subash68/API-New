// Package models ...
package models

import (
	"github.com/jaswanth-gorripati/PGK/s6_Token/configuration"
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
		"Token_Allocation_ins":       "INSERT INTO " + dbConfig.DbDatabaseName + "." + dbConfig.TokenAllocatinoDbName + "(Stakeholder_ID,NoOfTokensAllocated,TokenAllocationDate,Payment_ID,ModeOfIssue) VALUES(?,?,?,?,?)",
		"GET_Token_Allocation_BY_ID": "SELECT NoOfTokensAllocated,TokenAllocationDate,Payment_ID,ModeOfIssue FROM " + dbConfig.DbDatabaseName + "." + dbConfig.TokenAllocatinoDbName + " WHERE Stakeholder_ID=? GROUP BY NoOfTokensAllocated,TokenAllocationDate,Payment_ID,ModeOfIssue ORDER BY TokenAllocationDate DESC ",
		"TKN_BLNCE_SH_EXISTS":        "SELECT IF(COUNT(*),'true','false') FROM " + dbConfig.DbDatabaseName + "." + dbConfig.TokenBalanceDbName + " WHERE Stakeholder_ID=? ",
		"Token_Balance_ins":          "INSERT INTO " + dbConfig.DbDatabaseName + "." + dbConfig.TokenBalanceDbName + "(Stakeholder_ID,BonusTokensBalance,PurchasedTokensBalance,TokenBalanceDate) VALUES(?,?,?,?) ON DUPLICATE KEY UPDATE BonusTokensBalance=BonusTokensBalance+?,PurchasedTokensBalance=PurchasedTokensBalance+?,TokenBalanceDate=?",
		"GET_Balance_BY_ID":          "SELECT BonusTokensBalance,PurchasedTokensBalance,TokenBalanceDate,LastUpdatedDate FROM " + dbConfig.DbDatabaseName + "." + dbConfig.TokenBalanceDbName + " WHERE Stakeholder_ID=? GROUP BY BonusTokensBalance,PurchasedTokensBalance,TokenBalanceDate,LastUpdatedDate ",
		"Token_TX_ins":               "INSERT INTO " + dbConfig.DbDatabaseName + "." + dbConfig.TokenTransactionsDbName + "(Stakeholder_ID,BonusTokensTransacted,PurchasedTokensTransacted,Transaction_ID,TokenTransactionDate) VALUES(?,?,?,?,?)",
		"GET_Token_TX_BY_ID":         "SELECT BonusTokensTransacted,PurchasedTokensTransacted,Transaction_ID,TokenTransactionDate,LastUpdatedDate FROM " + dbConfig.DbDatabaseName + "." + dbConfig.TokenTransactionsDbName + " WHERE Stakeholder_ID=? GROUP BY BonusTokensTransacted,PurchasedTokensTransacted,Transaction_ID,TokenTransactionDate,LastUpdatedDate ORDER BY TokenTransactionDate DESC ",
	}
}
