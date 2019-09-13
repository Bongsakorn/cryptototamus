package cryptototamus

import (
	"bytes"
	"crypto/rand"
	b64 "encoding/base64"
	"encoding/gob"
	"fmt"

	"github.com/aws/aws-sdk-go/service/kms"
	"golang.org/x/crypto/nacl/secretbox"
)

const (
	keyLength   = 32
	nonceLength = 24
)

type payload struct {
	Key     []byte
	Nonce   *[nonceLength]byte
	Message []byte
}

// Encrypt TODO
func Encrypt(kmsClient *kms.KMS, kmsKeyName string, plaintext []byte) (string, error) {
	// Generate data key

	//provide either the key's arn OR its alias, as shown below:
	//keyId := "arn:aws:kms:us-east-1:779993255822:key/bb1a147c-8600-4558-910d-8b841c8f7493"
	keyId := "alias/" + kmsKeyName
	keySpec := "AES_128"
	dataKeyInput := kms.GenerateDataKeyInput{KeyId: &keyId, KeySpec: &keySpec}

	dataKeyOutput, err := kmsClient.GenerateDataKey(&dataKeyInput)
	if err == nil { // dataKeyOutput is now filled
		fmt.Println(dataKeyOutput)
	} else {
		return "", err
	}

	// Initialize payload
	p := &payload{
		Key:   dataKeyOutput.CiphertextBlob,
		Nonce: &[nonceLength]byte{},
	}

	// Set nonce
	if _, err = rand.Read(p.Nonce[:]); err != nil {
		return "", err
	}

	// Create key
	key := &[keyLength]byte{}
	copy(key[:], dataKeyOutput.Plaintext)

	// Encrypt message
	p.Message = secretbox.Seal(p.Message, plaintext, p.Nonce, key)

	buf := &bytes.Buffer{}
	if err := gob.NewEncoder(buf).Encode(p); err != nil {
		return "", err
	}

	return b64.StdEncoding.EncodeToString(buf.Bytes()), nil
}

// Decrypt TODO
func Decrypt(kmsClient *kms.KMS, encrypted string) ([]byte, error) {
	// Decode ciphertext with gob
	var p payload
	decoded, _ := b64.StdEncoding.DecodeString(encrypted)
	gob.NewDecoder(bytes.NewReader(decoded)).Decode(&p)

	//Decrypt a ciphertext that was previously encrypted.
	//Note that we dont actually specify the key name.
	//I guess the ciphertext already encodes it?
	dataKeyOutput, err := kmsClient.Decrypt(&kms.DecryptInput{
		CiphertextBlob: p.Key,
	})
	if err == nil { // dataKeyOutput is now filled
		fmt.Println(dataKeyOutput)
	} else {
		return nil, err
	}

	key := &[keyLength]byte{}
	copy(key[:], dataKeyOutput.Plaintext)

	// Decrypt message
	var plaintext []byte
	plaintext, ok := secretbox.Open(plaintext, p.Message, p.Nonce, key)
	if !ok {
		return nil, fmt.Errorf("Failed to open secretbox")
	}
	return plaintext, nil
}
