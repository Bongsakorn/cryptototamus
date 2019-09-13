# Crytototamus
Sample encrypt and decrypt your message with AWS KMS

### Install
```
go get gopkg.in/Bongsakorn/cryptototamus.v2
```

### Example
``` GO
package main

import (
	"encoding/hex"
	"fmt"

	crypt "gopkg.in/Bongsakorn/cryptototamus.v2"
)

func main() {
    keyName := "LogginKey" // insert your key here
    inputMessage := "Hello World" // message would like to encrypt
    
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		// Specify profile to load for the session's config
		Profile: "your_profile",

		// Provide SDK Config options, such as Region.
		Config: aws.Config{
			Region: aws.String("ap-southeast-1"),
		},
	}))
	kmsClient := kms.New(sess, aws.NewConfig())

	// Encrypt plaintext
	encrypted, err := Encrypt(kmsClient, keyName, []byte(inputMessage))
	if err != nil {
		panic(err)
	}
	fmt.Println("Encrypted: ", encoded)

	// Decrypt ciphertext
	kmsClient = kms.New(sess, aws.NewConfig())
	decrypted, err := Decrypt(kmsClient, encrypted)
	if err != nil {
		panic(err)
	}
	fmt.Println("Decrypted: ", string(decrypted))
}
```
