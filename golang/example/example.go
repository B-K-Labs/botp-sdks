package example

import (
	"fmt"

	"golang/sdk"
)

func Example() {
	apiKey := "30fdb9d7-6cfd-4e8d-abe9-bdd7b00f6363"
	botpsdk := sdk.NewBOTPSDK(apiKey)

	userAddresses := []string{"0x48216CDb36624fa6B3BE9Fc6a571Ea2f33c88251"}
	messages := []string{"message1"}
	notifyMessages := []string{"notify1"}

	response, err := botpsdk.SendMessage(userAddresses, messages, notifyMessages)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Response:", string(response))

	// Example usage for OTP validation
	// userAddr := "0xf0465189F703fAb578e2A040C6906460463115d9"
	// OTP := "3982995"
	// message := "agennenwgwj"

	// if err := botpSDK.AgentValidateOTP(userAddr, OTP, message); err != nil {
	// 	fmt.Println("Error validating OTP:", err)
	// } else {
	// 	fmt.Println("OTP validated successfully")
	// }
}
