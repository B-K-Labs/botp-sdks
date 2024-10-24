# SDK Usage Guide

This guide provides examples of how to use the SDK in various programming languages.

## Node.js

### Installation

```sh
npm install your-sdk-package
```

### Usage

```javascript
// Import the BOTP SDK
const BOTPSDK = require("./sdk.js");

// Replace 'YOUR_API_KEY' with your actual API key
const botpSDK = new BOTPSDK("YOUR_API_KEY");

// Example usage
async function exampleUsage() {
  try {
    const userAddresses = ["0x48216CDb36624fa6B3BE9Fc6a571Ea2f33c88251"];
    const messages = ["keyKhiem"];
    const notifyMessages = ["[Tiktok] Reset password di ne ban eii 8"];

    const sendMessageResponse = await botpSDK.sendMessage(
      userAddresses,
      messages,
      notifyMessages
    );
    console.log("Message sent:", sendMessageResponse);
  } catch (error) {
    console.error("Error:", error.message);
  }
}

// Call the example usage function
exampleUsage();
```

## Golang

### Installation

```sh
go get github.com/your-sdk-package
```

### Usage

```go
package main

import (
    "fmt"
    "log"
    "github.com/your-sdk-package/sdk"
)

func main() {
    apiKey := "YOUR_API_KEY"
    client := sdk.NewClient(apiKey)

    userAddresses := []string{"0x48216CDb36624fa6B3BE9Fc6a571Ea2f33c88251"}
    messages := []string{"keyKhiem"}
    notifyMessages := []string{"[Tiktok] Reset password di ne ban eii 8"}

    response, err := client.SendMessage(userAddresses, messages, notifyMessages)
    if err != nil {
        log.Fatalf("Error sending message: %v", err)
    }

    fmt.Printf("Message sent: %v\n", response)
}
```

## Python

### Installation

```sh
pip install your-sdk-package
```

### Usage

```python
from your_sdk_package import BOTPSDK

# Replace 'YOUR_API_KEY' with your actual API key
botp_sdk = BOTPSDK('YOUR_API_KEY')

try:
    user_addresses = ['0x48216CDb36624fa6B3BE9Fc6a571Ea2f33c88251']
    messages = ['keyKhiem']
    notify_messages = ['[Tiktok] Reset password di ne ban eii 8']

    send_message_response = botp_sdk.send_message(user_addresses, messages, notify_messages)
    print('Message sent:', send_message_response)
except Exception as e:
    print('Error:', str(e))
```

## Java

### Installation

```xml
<dependency>
    <groupId>com.your-sdk-package</groupId>
    <artifactId>sdk</artifactId>
    <version>1.0.0</version>
</dependency>
```

### Usage

```java
import com.your_sdk_package.BOTPSDK;

public class ExampleUsage {
    public static void main(String[] args) {
        String apiKey = "YOUR_API_KEY";
        BOTPSDK botpSDK = new BOTPSDK(apiKey);

        String[] userAddresses = {"0x48216CDb36624fa6B3BE9Fc6a571Ea2f33c88251"};
        String[] messages = {"keyKhiem"};
        String[] notifyMessages = {"[Tiktok] Reset password di ne ban eii 8"};

        try {
            String response = botpSDK.sendMessage(userAddresses, messages, notifyMessages);
            System.out.println("Message sent: " + response);
        } catch (Exception e) {
            System.err.println("Error: " + e.getMessage());
        }
    }
}
```
