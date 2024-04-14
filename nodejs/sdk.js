const axios = require('axios');

class BOTPSDK {
  constructor(apiKey) {
    this.apiKey = apiKey;
    this.baseURL = 'https://api.blockey.co/api/v1';
  }

  async sendMessage(userAddresses, messages, notifyMessages) {
    try {
      const response = await axios.post(`${this.baseURL}/message/sendMessage`, {
        APIKey: this.apiKey,
        ObjectListParams: userAddresses.map((userAddress, index) => ({
          userAddress,
          message: messages[index],
          notifyMessage: notifyMessages[index]
        }))
      }, {
        headers: {
          'Content-Type': 'application/json'
        }
      });
      return response.data;
    } catch (error) {
      throw new Error(error.response.data);
    }
  }

  async agentValidateOTP(userAddr, OTP, message, period = 120, digits = 7, algorithm = 'SHA-512') {
    try {
      const response = await axios.post(`${this.baseURL}/otp/agentValidateOTP`, {
        APIKey: this.apiKey,
        ObjectListParams: [{
          userAddr,
          OTP,
          message
        }],
        period,
        digits,
        algorithm
      }, {
        headers: {
          'Content-Type': 'application/json'
        }
      });
      return response.data;
    } catch (error) {
      throw new Error(error.response.data);
    }
  }
}

module.exports = BOTPSDK;
