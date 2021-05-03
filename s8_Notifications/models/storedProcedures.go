// Package models ...
package models

import (
	"github.com/jaswanth-gorripati/PGK/s8_Notifications/configuration"
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
		"NFT_INS":             "INSERT INTO " + dbConfig.DbDatabaseName + "." + dbConfig.NotificationsDbName + "(Notification_ID,Initiator_Stakeholder_ID,SenderUserRole,Receiver_Stakeholder_ID,DateofNotification,NotificationType,Notification_Content,Notification_AttachFile,RedirectedURL,Publish_ID,PublishFlag,CreationDate,LastUpdatedDate,NotificationType_ID,GenericMessage) VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)",
		"NFT_Get_Last_ID":     "SELECT Notification_ID FROM " + dbConfig.DbDatabaseName + "." + dbConfig.NotificationsDbName + " where Initiator_Stakeholder_ID=? ORDER BY Notification_ID DESC LIMIT 1",
		"NFT_GET_ALL":         "SELECT Notification_ID,Initiator_Stakeholder_ID,SenderUserRole,ifnull(Receiver_Stakeholder_ID,''),DateofNotification,NotificationType,Notification_Content,ifnull(Notification_AttachFile,''),ifnull(RedirectedURL,''),ifnull(Publish_ID,''),ifnull(PublishFlag,false),CreationDate,LastUpdatedDate FROM " + dbConfig.DbDatabaseName + "." + dbConfig.NotificationsDbName + " WHERE Initiator_Stakeholder_ID != ? AND (Receiver_Stakeholder_ID='' OR Receiver_Stakeholder_ID=null OR Receiver_Stakeholder_ID=?) ",
		"NFT_GROUP_COND":      " GROUP BY Notification_ID,Initiator_Stakeholder_ID,SenderUserRole,Receiver_Stakeholder_ID,DateofNotification,NotificationType,Notification_Content,Notification_AttachFile,RedirectedURL,Publish_ID,PublishFlag,CreationDate,LastUpdatedDate ORDER BY CreationDate DESC LIMIT ?,?",
		"GET_CORPORATE_NAME":  "SELECT Corporate_Name FROM CollabToHire.Corporate_Master where Stakeholder_ID=?",
		"GET_UNIVERSITY_NAME": "SELECT University_Name FROM CollabToHire.University_Master where Stakeholder_ID=?",
		"GET_STUDENT_NAME":    "SELECT group_concat(Student_FirstName,' ',Student_MiddleName,' ',Student_LastName) as name FROM CollabToHire.Student_Master where Stakeholder_ID=?",
	}
}
