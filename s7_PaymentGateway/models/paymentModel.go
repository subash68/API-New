package models

// PaymentDbModel ...
type PaymentDbModel struct {
	StakeholderID string  `form:"stakeholderID" json:"stakeholderID,omitempty" binding:"required"`
	PaymentID     string  `form:"paymentID" json:"paymentID" binding:"required"`
	PaymentMode   string  `form:"paymentMode" json:"paymentMode" binding:"required"`
	PayedAmount   float64 `form:"payedAmount" json:"payedAmount" binding:"required"`
	PayedDate     string  `form:"payedDate" json:"payedDate" binding:"required"`
}

// AllPaymentsModel ...
type AllPaymentsModel struct {
	Payments []PaymentDbModel `json:"payments"`
}

// CreatePaymentModel ...
type CreatePaymentModel struct {
	StakeholderID   string                 `form:"_" json:"stakeholderID"`
	StakeholderType string                 `form:"_" json:"stakeholderType"`
	PayAmount       float64                `form:"-" json:"amountToPay"`
	TokensUsed      float64                `form:"tokenUsed"`
	TokensToAdd     float64                `form:"tokensToAdd"`
	PayType         string                 `form:"payType" json:"payType" binding:"required"`
	Notes           map[string]interface{} `form:"-" json:"notes"`
}

// CreatePayRespModel ...
type CreatePayRespModel struct {
	OrderID string  `json:"orderID"`
	Amount  float64 `json:"amount"`
	//Notes   map[string]interface{} `json:"notes,omitempty"`
}

// PaySuccessReqModel ...
type PaySuccessReqModel struct {
	OrderID string `form:"orderID" json:"orderID" binding:"required"`
}

// PaySuccessRespModel ...
type PaySuccessRespModel struct {
	Notes map[string]interface{} `json:"notes,omitempty"`
}
