package models

import "time"

// NotificationsModel ...
type NotificationsModel struct {
	NotificationID     string    `form:"notificationID" json:"notificationID"`
	SenderID           string    `form:"senderID" json:"senderID" binding:"required"`
	SenderUserRole     string    `form:"senderUserRole" json:"senderUserRole" binding:"required"`
	SenderName         string    `form:"-" json:"senderName"`
	ReceiverID         string    `form:"receiverID" json:"receiverID"`
	ReceiverName       string    `form:"-" json:"receiverName,omitempty"`
	DateofNotification time.Time `form:"dateofNotification" json:"dateofNotification"  time_format="2006-12-01T21:23:34.409Z"`
	NotificationType   string    `form:"notificationType" json:"notificationType" binding:"required"`
	Content            string    `form:"content" json:"content" binding:"required"`
	AttachFile         []byte    `form:"attachFile" json:"attachFile"`
	RedirectedURL      string    `form:"redirectedURL" json:"redirectedURL"`
	PublishID          string    `form:"publishID" json:"publishID"`
	PublishFlag        bool      `form:"publishFlag" json:"publishFlag"`
	CreationDate       time.Time `form:"creationDate" json:"creationDate" time_format="2006-12-01T21:23:34.409Z"`
	LastUpdatedDate    time.Time `form:"lastUpdatedDate" json:"lastUpdatedDate" time_format="2006-12-01T21:23:34.409Z"`
}
