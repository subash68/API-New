package middleware

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/jaswanth-gorripati/PGK/s3_Publish/configuration"
)

// TokenResp ...
type TokenResp struct {
	Token string `json:"token"`
}

// SubmitHTTP ...
func SubmitHTTP(serviceHost string, servicePort int, method string, endpointURL string, reqData map[string]string, autToken string) ([]byte, error) {
	endpoint := "http://" + serviceHost + ":" + strconv.Itoa(servicePort) + endpointURL
	data := url.Values{}

	for key, value := range reqData {
		data.Set(key, value)
	}
	client := &http.Client{}
	r, err := http.NewRequest(method, endpoint, strings.NewReader(data.Encode())) // URL-encoded payload
	if err != nil {
		return nil, err
	}
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))
	if autToken != "" {
		r.Header.Add("Authorization", "Bearer "+autToken)
	}

	res, err := client.Do(r)
	if err != nil {
		return nil, err
	}
	log.Println("---------------->>", res.Status)

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	if res.StatusCode > 300 || res.StatusCode < 200 {
		return nil, fmt.Errorf("Request Existed with status code %v due to %s", res.Status, string(body))
	}
	return body, nil

}

// MakeInternalServiceCall ...
func MakeInternalServiceCall(serviceHost string, servicePort int, method string, endpointURL string, reqData map[string]string) ([]byte, error) {

	authService := configuration.AuthConfig()
	internalTokenData := map[string]string{"tokenType": "INTERNAL", "userType": "S3", "userId": "Publish"}
	tokenBody, err := SubmitHTTP(authService.Host, authService.Port, "POST", "/a/createInternalToken", internalTokenData, "")
	if err != nil {
		return nil, err
	}
	var tokenResp TokenResp
	err = json.Unmarshal(tokenBody, &tokenResp)
	if err != nil {
		return nil, err
	}
	httpBody, err := SubmitHTTP(serviceHost, servicePort, method, endpointURL, reqData, tokenResp.Token)
	if err != nil {
		return nil, err
	}
	return httpBody, nil

}
