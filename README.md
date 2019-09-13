# Crytototamus
Encrypt your message with passphrase

### Install
```
go get gopkg.in/Bongsakorn/cryptototamus.v1
```

### Example
``` GO
package cryptototamus

import (
	"encoding/hex"
	"fmt"
)

func main() {
	ciphertext, err := encrypt([]byte("Hello World"), "password")
	if err != nil {
		fmt.Println(err.Error())
	}
    fmt.Printf("Encrypted: %s\n", hex.EncodeToString(ciphertext))
    
	ciphertext, _ := hex.DecodeString(c)
	plaintext, err := decrypt(ciphertext, "password")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("Decrypted: %s\n", plaintext)
}

```

### Credits
[Nic Raboy](https://www.thepolyglotdeveloper.com/2018/02/encrypt-decrypt-data-golang-application-crypto-packages/)
