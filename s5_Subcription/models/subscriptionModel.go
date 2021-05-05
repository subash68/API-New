package models

import "time"

// SubscriptionModel ...
type SubscriptionModel struct {
	Subscriber         string    `form:"-" json:"subscriber,omitempty"`
	Publisher          string    `form:"-" json:"publisher"`
	DateOfSubscription time.Time `form:"-" json:"dateOfSubscription"`
	PublishID          string    `form:"publishId" json:"publishId" binding:"required"`
	TokensUsed         float64   `form:"tokensUsed" json:"tokensUsed,omitempty" binding:"required"`
	TransactionID      string    `form:"transactionID" json:"transactionID,omitempty"`
	CorporateName      string    `form:"-" json:"publisherName,omitempty"`
	GeneralNote        string    `form:"-" json:"generalNote,omitempty"`
	SubscriptionID     string    `form:"-" json:"subscriptionID,omitempty"`
	PublisherLocation  string    `form:"-" json:"location"`
}

// SubscriptionReq ...
type SubscriptionReq struct {
	Subscriber         string    `form:"-" json:"subscriber,omitempty"`
	Publisher          string    `form:"-" json:"publisher"`
	DateOfSubscription time.Time `form:"-" json:"dateOfSubscription"`
	PublishID          string    `form:"publishId" json:"publishId" binding:"required"`
	BonusTokensUsed    float64   `form:"bonusTokensUsed" json:"bonusTokensUsed,omitempty"`
	PaidTokensUsed     float64   `form:"paidTokensUsed" json:"paidTokensUsed,omitempty" `
	TransactionID      string    `form:"transactionID" json:"transactionID,omitempty"`
}

// SubSuccessResp ...
type SubSuccessResp struct {
	Message string `json:"message"`
}

// AllSubscriptionsModel ...
type AllSubscriptionsModel struct {
	Subscriptions []SubscriptionModel `json:"subscriptions"`
}

// SubscriptionPaymentModel ...
type SubscriptionPaymentModel struct {
	Message                   string  `json:"string"`
	TokensRequired            float64 `json:"tokensrequired"`
	BonusTokenUsagePercentage float64 `json:"bonusTokenUsagePercentage"`
}
