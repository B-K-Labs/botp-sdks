// Import the BOTP SDK
const BOTPSDK = require('./sdk.js');

// Replace 'YOUR_API_KEY' with your actual API key
const botpSDK = new BOTPSDK('30fdb9d7-6cfd-4e8d-abe9-bdd7b00f6363');

// Example usage
async function exampleUsage() {
  try {
    // Example usage for sending messages
    const userAddresses = ['0x48216CDb36624fa6B3BE9Fc6a571Ea2f33c88251'];
    const messages = ['keyKhiem'];
    const notifyMessages = ['[Tiktok] Reset password di ne ban eii 8'];

    const sendMessageResponse = await botpSDK.sendMessage(userAddresses, messages, notifyMessages);
    console.log('Message sent:', sendMessageResponse);

    // Example usage for OTP validation
    // const userAddr = '0xf0465189F703fAb578e2A040C6906460463115d9';
    // const OTP = '3982995';
    // const message = 'agennenwgwj';

    // const validateOTPResponse = await botpSDK.agentValidateOTP(userAddr, OTP, message);
    // console.log('OTP validated:', validateOTPResponse);
  } catch (error) {
    console.error('Error:', error.message);
  }
}

// Call the example usage function
exampleUsage();
