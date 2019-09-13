# Crytototamus
Encrypt your message with passphrase

### Install
```
go get gopkg.in/Bongsakorn/cryptototamus.v1
```

### Example
``` GO
package main

import (
	"encoding/hex"
	"fmt"

	crypt "gopkg.in/Bongsakorn/cryptototamus.v1"
)

func main() {
	ciphertext, err := crypt.Encrypt([]byte("Hello World"), "password")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("Encrypted: %s\n", hex.EncodeToString(ciphertext))

	plaintext, err := crypt.Decrypt(ciphertext, "password")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("Decrypted: %s\n", plaintext)
}
```

### Credits
[Nic Raboy](https://www.thepolyglotdeveloper.com/2018/02/encrypt-decrypt-data-golang-application-crypto-packages/)
