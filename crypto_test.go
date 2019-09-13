package cryptototamus

import (
	"testing"
)

func TestEncrypt(t *testing.T) {
	inputMessage := "Hello World"
	ciphertext, _ := encrypt([]byte(inputMessage), "mypassword")
	plaintext, _ := decrypt(ciphertext, "mypassword")
	if string(plaintext) != inputMessage {
		t.Errorf("Failed, expected %s but got %s", inputMessage, plaintext)
	} else {
		t.Logf("Success, expected %s but got %s", inputMessage, plaintext)
	}
}
