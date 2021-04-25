package models

import "time"

// TokenBalanceModel ...
type TokenBalanceModel struct {
	StakeholderID   string    `form:"stakeholderID" json:"stakeholderID,omitempty"`
	Balance         float64   `form:"-" json:"balance"`
	BalanceDate     time.Time `form:"-" json:"balanceDate"`
	LastUpdatedDate time.Time `form:"-" json:"lastUpdatedDate"`
}

// TokenAllocationModel ...
type TokenAllocationModel struct {
	StakeholderID    string    `form:"stakeholderID" json:"stakeholderID,omitempty" binding:"required"`
	PaymentID        string    `form:"paymentID" json:"paymentID" binding:"required"`
	ModeOfTokenissue string    `form:"modeOfTokenissue" json:"modeOfTokenissue" binding:"required"`
	AllocatedTokens  float64   `form:"allocatedTokens" json:"allocatedTokens"  binding:"required"`
	AllocatedDate    time.Time `form:"allocatedDate" json:"allocatedDate"`
	LastUpdatedDate  time.Time `form:"-" json:"lastUpdatedDate"`
}

// TokenTransactionsModel ...
type TokenTransactionsModel struct {
	StakeholderID    string    `form:"stakeholderID" json:"stakeholderID,omitempty" binding:"required"`
	TransactionID    string    `form:"transactionID" json:"transactionID" binding:"required"`
	TokensTransacted float64   `form:"tokensTransacted" json:"tokensTransacted" binding:"required"`
	TransactionDate  time.Time `form:"transactionDate" json:"transactionDate"`
	LastUpdatedDate  time.Time `form:"-" json:"lastUpdatedDate"`
}

// AllocatedTokens ...
type AllocatedTokens struct {
	AllocatedTokens []TokenAllocationModel `json:"allocatedToken"`
}

// TokenTransactions ...
type TokenTransactions struct {
	Transactions []TokenTransactionsModel `json:"transactions"`
}

// TokenDbResp ...
type TokenDbResp struct {
	Message string `json:"message"`
}
