package models

import "time"

// TokenBalanceModel ...
type TokenBalanceModel struct {
	StakeholderID     string    `form:"stakeholderID" json:"stakeholderID,omitempty"`
	BonusTokenBalance float64   `form:"-" json:"bonusTokenBalance"`
	PaidTokenBalance  float64   `form:"-" json:"paidTokenBalance"`
	BalanceDate       time.Time `form:"-" json:"balanceDate"`
	LastUpdatedDate   time.Time `form:"-" json:"lastUpdatedDate"`
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

// TokenTxAllocationModel ...
type TokenTxAllocationModel struct {
	StakeholderID    string    `form:"stakeholderID" json:"stakeholderID,omitempty" binding:"required"`
	PaymentID        string    `form:"paymentID" json:"paymentID" binding:"required"`
	AmountPaid       float64   `json:"amountPaid"`
	ModeOfTokenissue string    `form:"modeOfTokenissue" json:"modeOfTokenissue" binding:"required"`
	AllocatedTokens  float64   `form:"allocatedTokens" json:"allocatedTokens"  binding:"required"`
	AllocatedDate    time.Time `form:"allocatedDate" json:"allocatedDate"`
	LastUpdatedDate  time.Time `form:"-" json:"lastUpdatedDate"`
}

// TokenTransactionsModel ...
type TokenTransactionsModel struct {
	StakeholderID         string    `form:"stakeholderID" json:"stakeholderID,omitempty" binding:"required"`
	TransactionID         string    `form:"transactionID" json:"transactionID" binding:"required"`
	BonusTokensTransacted float64   `form:"bonusTokensTransacted" json:"bonusTokensTransacted"`
	PaidTokensTransacted  float64   `form:"paidTokensTransacted" json:"paidTokensTransacted"`
	TransactionDate       time.Time `form:"transactionDate" json:"transactionDate"`
	LastUpdatedDate       time.Time `form:"-" json:"lastUpdatedDate"`
}

// AllocatedTokens ...
type AllocatedTokens struct {
	AllocatedTokens []TokenAllocationModel `json:"allocatedToken"`
}

// TxTokens ...
type TxTokens struct {
	AllocatedTokens []TokenTxAllocationModel `json:"paymentTransactions"`
}

// TokenTransactions ...
type TokenTransactions struct {
	Transactions []TokenTransactionsModel `json:"transactions"`
}

// TokenDbResp ...
type TokenDbResp struct {
	Message string `json:"message"`
}
