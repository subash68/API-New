// Package controllers ...
package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/jaswanth-gorripati/PGK/s1_Onboarding/configuration"
)

// TokenResp ...
type TokenResp struct {
	Token string `json:"token"`
}

// Logout resp

// CreateAuthRequest ...
func CreateAuthRequest(urlData map[string]string, endpointURL string) (body []byte, err error) {
	authService := configuration.AuthConfig()
	endpoint := "http://" + authService.Host + ":" + strconv.Itoa(authService.Port) + endpointURL
	data := url.Values{}

	for key, value := range urlData {
		data.Set(key, value)
	}
	fmt.Printf("\n x-Form-data : %+v , url : %v\n", data, endpoint)
	client := &http.Client{}
	r, err := http.NewRequest("POST", endpoint, strings.NewReader(data.Encode())) // URL-encoded payload
	if err != nil {
		return body, err
	}
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	res, err := client.Do(r)
	if err != nil {
		return body, err
	}
	log.Println("---------------->>", res.Status)
	if res.StatusCode > 300 || res.StatusCode < 200 {
		return body, fmt.Errorf("Request Existed with status code " + res.Status)
	}
	defer res.Body.Close()
	body, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return body, err
	}
	log.Println(string(body))
	return body, err
}

// CreateToken ...
func CreateToken(ctx context.Context, usrType string, usrUUID string, tokTyp string) (string, error) {
	urlData := map[string]string{"tokenType": tokTyp, "userType": usrType, "userId": usrUUID}
	body, err := CreateAuthRequest(urlData, "/a/createToken")
	if err != nil {
		fmt.Printf("\n Error while calling : %+v\n", err)
		return "", err
	}
	var tokenResp TokenResp
	err = json.Unmarshal(body, &tokenResp)
	return tokenResp.Token, err
}

// DeleteToken ...
func DeleteToken(token string) (string, error) {
	urlData := map[string]string{"token": token}
	body, err := CreateAuthRequest(urlData, "/a/delToken")
	fmt.Printf("\n ---> errr : %+v , %v\n", err, string(body))
	return string(body), err
}
