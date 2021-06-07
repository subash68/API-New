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
		"Token_Allocation_ins":         "INSERT INTO " + dbConfig.DbDatabaseName + "." + dbConfig.TokenAllocationDbName + "(Stakeholder_ID,NoOfTokensAllocated,TokenAllocationDate,Payment_ID,ModeOfIssue) VALUES(?,?,?,?,?)",
		"GET_Token_Allocation_BY_ID":   "SELECT NoOfTokensAllocated,TokenAllocationDate,Payment_ID,ModeOfIssue FROM " + dbConfig.DbDatabaseName + "." + dbConfig.TokenAllocationDbName + " WHERE Stakeholder_ID=? GROUP BY NoOfTokensAllocated,TokenAllocationDate,Payment_ID,ModeOfIssue ORDER BY TokenAllocationDate DESC ",
		"GET_Token_Transactions_BY_ID": "SELECT NoOfTokensAllocated,TokenAllocationDate,Payment_ID,ifnull((select AmountPaid FROM " + dbConfig.DbDatabaseName + ".REPLACE as t where t.Payment_ID=tx.Payment_ID),0),ModeOfIssue FROM " + dbConfig.DbDatabaseName + "." + dbConfig.TokenAllocationDbName + " as tx WHERE Stakeholder_ID=? GROUP BY NoOfTokensAllocated,TokenAllocationDate,Payment_ID,ModeOfIssue ORDER BY TokenAllocationDate DESC ",
		"TKN_BLNCE_SH_EXISTS":          "SELECT IF(COUNT(*),'true','false') FROM " + dbConfig.DbDatabaseName + "." + dbConfig.TokenBalanceDbName + " WHERE Stakeholder_ID=? ",
		"Token_Balance_ins":            "INSERT INTO " + dbConfig.DbDatabaseName + "." + dbConfig.TokenBalanceDbName + "(Stakeholder_ID,BonusTokensBalance,PurchasedTokensBalance,TokenBalanceDate) VALUES(?,?,?,?) ON DUPLICATE KEY UPDATE BonusTokensBalance=BonusTokensBalance+?,PurchasedTokensBalance=PurchasedTokensBalance+?,TokenBalanceDate=?",
		"GET_Balance_BY_ID":            "SELECT BonusTokensBalance,PurchasedTokensBalance,TokenBalanceDate,LastUpdatedDate FROM " + dbConfig.DbDatabaseName + "." + dbConfig.TokenBalanceDbName + " WHERE Stakeholder_ID=? GROUP BY BonusTokensBalance,PurchasedTokensBalance,TokenBalanceDate,LastUpdatedDate ",
		"Token_TX_ins":                 "INSERT INTO " + dbConfig.DbDatabaseName + "." + dbConfig.TokenTransactionsDbName + "(Stakeholder_ID,BonusTokensTransacted,PurchasedTokensTransacted,Transaction_ID,TokenTransactionDate,Publisher_Stakeholder_ID,Subscription_ID,SubscriptionType,Publisher_Stakeholder_Type) VALUES(?,?,?,?,?,?,?,?,?)",
		"GET_Token_TX_BY_ID":           "SELECT t.BonusTokensTransacted,t.PurchasedTokensTransacted,t.Transaction_ID,t.TokenTransactionDate,t.LastUpdatedDate,ifnull(t.Publisher_Stakeholder_ID,''),ifnull(t.Subscription_ID,''),ifnull(t.SubscriptionType,''),ifnull(t.Publisher_Stakeholder_Type,''),(SELECT json_object('exists',if(count(*),true,false),'name',ifnull(Corporate_Name,''),'location',ifnull(CorporateHQAddress_City,'')) as corporateDetails FROM CollabToHire.Corporate_Master WHERE Stakeholder_ID=t.Publisher_Stakeholder_ID),(SELECT json_object('exists',if(count(*),true,false),'name',ifnull(University_Name,''),'location',ifnull(UniversityHQAddress_City,'')) as corporateDetails FROM CollabToHire.University_Master WHERE Stakeholder_ID=t.Publisher_Stakeholder_ID) FROM " + dbConfig.DbDatabaseName + "." + dbConfig.TokenTransactionsDbName + " as t WHERE t.Stakeholder_ID=? ORDER BY TokenTransactionDate DESC ",
	}
}
