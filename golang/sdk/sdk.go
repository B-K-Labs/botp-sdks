package sdk

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type BOTPSDK struct {
	APIKey  string
	BaseURL string
	Client  *http.Client
}

func NewBOTPSDK(apiKey string) *BOTPSDK {
	return &BOTPSDK{
		APIKey:  apiKey,
		BaseURL: "https://api.blockey.co/api/v1",
		Client:  &http.Client{},
	}
}

type sendMessageParams struct {
	APIKey           string             `json:"APIKey"`
	ObjectListParams []messageObjParams `json:"ObjectListParams"`
}

type messageObjParams struct {
	UserAddress   string `json:"userAddress"`
	Message       string `json:"message"`
	NotifyMessage string `json:"notifyMessage"`
}

type agentValidateOTPParams struct {
	APIKey           string             `json:"APIKey"`
	ObjectListParams []validateOTPParam `json:"ObjectListParams"`
	Period           int                `json:"period"`
	Digits           int                `json:"digits"`
	Algorithm        string             `json:"algorithm"`
}

type validateOTPParam struct {
	UserAddr string `json:"userAddr"`
	OTP      string `json:"OTP"`
	Message  string `json:"message"`
}

func (b *BOTPSDK) SendMessage(userAddresses []string, messages []string, notifyMessages []string) ([]byte, error) {
	params := sendMessageParams{
		APIKey:           b.APIKey,
		ObjectListParams: make([]messageObjParams, len(userAddresses)),
	}

	for i, userAddress := range userAddresses {
		params.ObjectListParams[i] = messageObjParams{
			UserAddress:   userAddress,
			Message:       messages[i],
			NotifyMessage: notifyMessages[i],
		}
	}

	reqBody, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	resp, err := b.Client.Post(fmt.Sprintf("%s/message/sendMessage", b.BaseURL), "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(string(respBody))
	}

	return respBody, nil
}

func (b *BOTPSDK) AgentValidateOTP(userAddr, OTP, message string, period, digits int, algorithm string) ([]byte, error) {
	params := agentValidateOTPParams{
		APIKey:    b.APIKey,
		Period:    period,
		Digits:    digits,
		Algorithm: algorithm,
		ObjectListParams: []validateOTPParam{
			{
				UserAddr: userAddr,
				OTP:      OTP,
				Message:  message,
			},
		},
	}

	reqBody, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	resp, err := b.Client.Post(fmt.Sprintf("%s/otp/agentValidateOTP", b.BaseURL), "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(string(respBody))
	}

	return respBody, nil
}
