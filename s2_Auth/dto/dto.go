package dto

// LoginTokenDTO ...
type LoginTokenDTO struct {
	UserType string `form:"userType" bindig:"required"`
	UserID   string `form:"userID" binding:"required"`
}

// LoginTokenResp ...
type LoginTokenResp struct {
	Token string `json:"token"`
}

// VrfToken ...
type VrfToken struct {
	Token string `form:"token" binding:"required"`
}
