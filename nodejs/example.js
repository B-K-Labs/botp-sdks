// Import the BOTP SDK
const BOTPSDK = require('./sdk.js');

// Replace 'YOUR_API_KEY' with your actual API key
const botpSDK = new BOTPSDK('YOUR_API_KEY');

// Example usage
async function exampleUsage() {
  try {
    // Example usage for sending messages
    const userAddresses = ['0xDB026e60C1083375167094ae3531352f47f05b0F', '0xC0c0b84907b5b93aAF37936eC5d9D1fDF7A60aD5'];
    const messages = ['keythinh1', 'keythinh2'];
    const notifyMessages = ['[khiem-2] Test analyser1', '[khiem-2] Test analyser2'];

    const sendMessageResponse = await botpSDK.sendMessage(userAddresses, messages, notifyMessages);
    console.log('Message sent:', sendMessageResponse);

    // Example usage for OTP validation
    const userAddr = '0xf0465189F703fAb578e2A040C6906460463115d9';
    const OTP = '3982995';
    const message = 'agennenwgwj';

    const validateOTPResponse = await botpSDK.agentValidateOTP(userAddr, OTP, message);
    console.log('OTP validated:', validateOTPResponse);
  } catch (error) {
    console.error('Error:', error.message);
  }
}

// Call the example usage function
exampleUsage();
