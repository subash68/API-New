package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/jaswanth-gorripati/PGK/s4_Profile/configuration"
)

type twiloOtpResp struct {
	SID    string `json:"sid"`
	Status string `json:"status"`
	Valid  bool   `json:"valid"`
}

// SendSmsOtp ...
func SendSmsOtp(to string) (bool, error) {
	twilioConfig := configuration.TwilioConfig()
	apiURL := "https://verify.twilio.com/"
	resource := "/v2/Services/" + twilioConfig.VrfSID + "/Verifications"
	data := url.Values{}
	data.Set("To", to)
	data.Set("Channel", "sms")
	data.Set("Code", "123456")

	u, _ := url.ParseRequestURI(apiURL)
	u.Path = resource
	urlStr := u.String()

	client := &http.Client{}

	r, _ := http.NewRequest(http.MethodPost, urlStr, strings.NewReader(data.Encode()))
	r.SetBasicAuth(twilioConfig.AccSID, twilioConfig.AccSecret)
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	resp, err := client.Do(r)
	if err != nil {
		return false, err
	}
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		var data twiloOtpResp
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		err := json.Unmarshal(bodyBytes, &data)
		if err == nil {
			return true, nil
		}
		return false, fmt.Errorf("Failed to get the SID")
	} else {
		return false, fmt.Errorf("Failed to send otp , Http response :%v", resp.StatusCode)
	}
}

// ValidateOTP ...
func ValidateOTP(otp string, to string) (bool, error) {

	twilioConfig := configuration.TwilioConfig()

	apiURL := "https://verify.twilio.com/"
	resource := "/v2/Services/" + twilioConfig.VrfSID + "/VerificationCheck"
	data := url.Values{}
	data.Set("To", to)
	data.Set("Code", otp)

	u, _ := url.ParseRequestURI(apiURL)
	u.Path = resource
	urlStr := u.String()

	client := &http.Client{}

	r, _ := http.NewRequest(http.MethodPost, urlStr, strings.NewReader(data.Encode()))
	r.SetBasicAuth(twilioConfig.AccSID, twilioConfig.AccSecret)
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	resp, err := client.Do(r)
	if err != nil {
		return false, err
	}
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		var data twiloOtpResp
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		err := json.Unmarshal(bodyBytes, &data)
		if err == nil && data.Status == "approved" {
			return true, nil
		} else if !data.Valid {
			return false, fmt.Errorf("Invalid OTP")
		}
		return false, fmt.Errorf("Failed to get the SID")
	} else {
		return false, fmt.Errorf("Invalid / Expired OTP")
	}
}
