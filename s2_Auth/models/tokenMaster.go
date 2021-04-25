package models

// TokenMasterDB ...
type TokenMasterDB struct {
	TokenType   string `form:"tokenType" binding:"required"`
	UserType    string `form:"userType" binding:"required"`
	UserID      string `form:"userId" binding:"required"`
	AccessToken string `form:"accessToken"`
	AccessUUID  string `form:"accessUUID"`
	AtExpires   int64  `form:"atExpires"`
}

// TokenClaims ...
type TokenClaims struct {
	TokenType string `json:"tokenType"`
	UserType  string `json:"userType"`
	UserID    string `json:"userId"`
}
