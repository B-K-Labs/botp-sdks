package main

import (
	"fmt"

	"./sdk"
)

func main() {
	// Replace 'YOUR_API_KEY' with your actual API key
	botpSDK := sdk.NewBOTPSDK("YOUR_API_KEY")

	// Example usage for sending messages
	userAddresses := []string{"0xDB026e60C1083375167094ae3531352f47f05b0F", "0xC0c0b84907b5b93aAF37936eC5d9D1fDF7A60aD5"}
	messages := []string{"keythinh1", "keythinh2"}
	notifyMessages := []string{"[khiem-2] Test analyser1", "[khiem-2] Test analyser2"}

	if err := botpSDK.sendMessage(userAddresses, messages, notifyMessages); err != nil {
		fmt.Println("Error sending message:", err)
	} else {
		fmt.Println("Messages sent successfully")
	}

	// Example usage for OTP validation
	userAddr := "0xf0465189F703fAb578e2A040C6906460463115d9"
	OTP := "3982995"
	message := "agennenwgwj"

	if err := botpSDK.agentValidateOTP(userAddr, OTP, message); err != nil {
		fmt.Println("Error validating OTP:", err)
	} else {
		fmt.Println("OTP validated successfully")
	}
}
