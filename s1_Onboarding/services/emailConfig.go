// Package services ...
package services

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/jaswanth-gorripati/PGK/s1_Onboarding/configuration"
	"github.com/jaswanth-gorripati/PGK/s1_Onboarding/models"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/gmail/v1"
	"google.golang.org/api/option"
)

// EncodeToString ...
func EncodeToString(max int) string {
	b := make([]byte, max)
	n, err := io.ReadAtLeast(rand.Reader, b, max)
	if n != max {
		panic(err)
	}
	for i := 0; i < len(b); i++ {
		b[i] = table[int(b[i])%len(table)]
	}
	return string(b)
}

var table = [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}

// GmailService : Gmail client for sending email
var GmailService *gmail.Service

// ConfigureOAuthMailService to configure the mail using Gmail OAuth credentials
func ConfigureOAuthMailService() {
	// Fetch oauth details from config
	emailConfig := configuration.EmailConfig()
	config := oauth2.Config{
		ClientID:     emailConfig.EmailClientID,
		ClientSecret: emailConfig.EmailClientSecret,
		Endpoint:     google.Endpoint,
		RedirectURL:  emailConfig.EmailRedirectURL,
	}

	// Access token configuration
	token := oauth2.Token{
		AccessToken:  emailConfig.EmailAccessToken,
		RefreshToken: emailConfig.EmailRefreshToken,
		TokenType:    "Bearer",
		Expiry:       time.Now(),
	}

	var tokenSource = config.TokenSource(context.Background(), &token)

	srv, err := gmail.NewService(context.Background(), option.WithTokenSource(tokenSource))
	if err != nil {
		log.Printf("Unable to retrieve Gmail client: %v", err)
	}

	GmailService = srv
	if GmailService != nil {
		fmt.Println("Email service is configured")
	}
}

// SendEmail to send an email to the recipeint email
func SendEmail(to string, sub string, body string) (bool, error) {

	//emailBody := "hi this is a test mail"

	var message gmail.Message

	emailTo := "To: " + to + "\r\n"
	subject := "Subject: " + sub + "\n"
	mime := "MIME-version: 1.0;\nContent-Type: text/plain; charset=\"UTF-8\";\n\n"
	msg := []byte(emailTo + subject + mime + "\n" + body)

	message.Raw = base64.URLEncoding.EncodeToString(msg)

	// Send the message
	_, err := GmailService.Users.Messages.Send("me", &message).Do()
	if err != nil {
		return false, err
	}
	return true, nil
}

// SendOTPEmail to send an email to the recipeint email
func SendOTPEmail(to string, stakeholderID string) (bool, error) {
	fmt.Println(" ---> otp for ", stakeholderID)
	if to == "" {
		return false, fmt.Errorf("Invalid Email Address ")
	}
	var message gmail.Message
	OTP := EncodeToString(6)
	otpExp := time.Unix(time.Now().Add(time.Minute*10).Unix(), 0)
	emailTo := "To: " + to + "\r\n"
	subject := "Subject: Verificaton OTP for Application\n"
	body := "Your verification OTP is :<b>" + OTP + " </b> <br> ** It will expire in 10 minutes."
	mime := "MIME-version: 1.0;\nContent-Type: text/plain; charset=\"UTF-8\";\n\n"
	msg := []byte(emailTo + subject + mime + "\n" + body)

	message.Raw = base64.URLEncoding.EncodeToString(msg)
	// check if redis online
	var customError models.DbModelError
	if models.CheckRedisPing(&customError); customError.Err != nil {
		return false, fmt.Errorf("Redis not available")
	}
	// Send the message
	_, err := GmailService.Users.Messages.Send("me", &message).Do()
	if err != nil {
		fmt.Printf("\nFailed to Send OTP Email to %v, due to : %+v\n", to, err)
		return false, err
	}
	// Store otp in REdis for verification
	key := fmt.Sprint(stakeholderID, "emailOTP")

	errAccess := models.RedisClient.Set(key, OTP, otpExp.Sub(time.Now()))
	if errAccess.Err() != nil {
		fmt.Printf("\nFailed to store OTP : %+v\n", errAccess)
		return false, errAccess.Err()
	}

	return true, nil
}

// VerifyEmailOtp ...
func VerifyEmailOtp(pid string, otp string) (bool, error) {
	storedOtp, err := models.RedisClient.Get(pid + "emailOTP").Result()
	if err != nil || storedOtp != otp {
		fmt.Printf("\n email otp redis err : %+v\n", err)
		return false, fmt.Errorf("Invalid OTP, Try again")
	}
	if storedOtp == otp {
		_ = models.RedisClient.Del(pid + "emailOTP")
		return true, nil
	}
	return false, fmt.Errorf("Invalid OTP, Try again")

}
