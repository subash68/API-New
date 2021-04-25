package services

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/jaswanth-gorripati/PGK/s2_Auth/configuration"
	"github.com/jaswanth-gorripati/PGK/s2_Auth/models"
)

// CreateToken ...
func CreateToken(tokenDb *models.TokenMasterDB) (token string, err error) {

	jwtConfig := configuration.JwtConfig()
	jobdb := make(chan models.DbModelError, 1)

	//Creating Access Token
	namespace, _ := uuid.NewRandom()
	tokenDb.AccessUUID = uuid.NewSHA1(namespace, []byte(tokenDb.UserID)).String()

	// Adding Claims
	atClaims := jwt.MapClaims{}
	atClaims["tokTyp"] = tokenDb.TokenType
	atClaims["accessUUID"] = tokenDb.AccessUUID
	atClaims["exp"] = tokenDb.AtExpires
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	tokenDb.AccessToken, err = at.SignedString([]byte(jwtConfig.JwtAccessSecret))
	if err != nil {
		return "", err
	}

	go func() {
		fmt.Println("sending job for insert")
		select {
		case insertJobChan := <-tokenDb.InsertToken():
			jobdb <- insertJobChan
		}
	}()

	insertJob := <-jobdb
	fmt.Printf("\n insert job : %+v\n", insertJob)
	if insertJob.ErrTyp != "000" {
		return "", insertJob.Err
	}
	return tokenDb.AccessToken, nil
}

// FetchAccessUUID ...
func FetchAccessUUID(vrfToken string) (string, error) {
	jwtConfig := configuration.JwtConfig()
	token, err := jwt.Parse(vrfToken, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(jwtConfig.JwtAccessSecret), nil
	})
	if err != nil {
		return "", err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return "", fmt.Errorf("Invalid Token, token expired")
	}
	accessUUID, ok := claims["accessUUID"].(string)
	if !ok {
		return "", fmt.Errorf("Invalid token , cannot find the access key")
	}
	return accessUUID, nil
}

// VerifyToken ...
func VerifyToken(vrfToken string) (tokenClaims models.TokenClaims, err error) {

	accessUUID, err := FetchAccessUUID(vrfToken)
	if err != nil {
		return tokenClaims, err
	}
	tokenDb := models.TokenMasterDB{AccessToken: accessUUID}
	tokenClaims, err = tokenDb.GetClaimsForToken()
	if err != nil {
		return tokenClaims, err
	}
	fmt.Printf("\n Token Claims : %+v\n", tokenClaims)
	return tokenClaims, err
}

// DelToken ...
func DelToken(vrfToken string) error {
	fmt.Printf("--->>>> in services ")
	accessUUID, err := FetchAccessUUID(vrfToken)
	if err != nil {
		return err
	}
	tokenDb := models.TokenMasterDB{AccessToken: accessUUID}
	isDeleted := tokenDb.DeleteToken()
	if isDeleted {
		return nil
	} else {
		return fmt.Errorf("Cannot Delete token")
	}
}
