package sdk

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

// BOTPSDK represents the SDK for interacting with BOTP APIs
type BOTPSDK struct {
	APIKey  string
	BaseURL string
}

// NewBOTPSDK creates a new instance of the BOTP SDK
func NewBOTPSDK(apiKey string) *BOTPSDK {
	return &BOTPSDK{
		APIKey:  apiKey,
		BaseURL: "https://botp-backend-logic-api.herokuapp.com/api/v1/",
	}
}

// sendMessage sends transaction messages to users
func (sdk *BOTPSDK) sendMessage(userAddresses []string, messages []string, notifyMessages []string) error {
	url := sdk.BaseURL + "message/sendMessage"
	data := map[string]interface{}{
		"APIKey": sdk.APIKey,
		"ObjectListParams": []map[string]string{
			{"userAddress": userAddresses[0], "message": messages[0], "notifyMessage": notifyMessages[0]},
			{"userAddress": userAddresses[1], "message": messages[1], "notifyMessage": notifyMessages[1]},
		},
	}

	if err := sdk.postRequest(url, data); err != nil {
		return err
	}
	return nil
}

// agentValidateOTP validates the OTP received from the BOTP app
func (sdk *BOTPSDK) agentValidateOTP(userAddr, OTP, message string) error {
	url := sdk.BaseURL + "otp/agentValidateOTP"
	data := map[string]interface{}{
		"APIKey": sdk.APIKey,
		"ObjectListParams": []map[string]string{
			{"userAddr": userAddr, "OTP": OTP, "message": message},
		},
		"period":    120,
		"digits":    7,
		"algorithm": "SHA-512",
	}

	if err := sdk.postRequest(url, data); err != nil {
		return err
	}
	return nil
}

// postRequest makes a POST request to the specified URL with the provided data
func (sdk *BOTPSDK) postRequest(url string, data interface{}) error {
	payload, err := json.Marshal(data)
	if err != nil {
		return err
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(payload))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		return errors.New(fmt.Sprintf("API request failed with status code %d: %s", resp.StatusCode, string(body)))
	}

	return nil
}
