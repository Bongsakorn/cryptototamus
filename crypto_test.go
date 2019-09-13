package cryptototamus

import (
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kms"
)

func TestEncrypt(t *testing.T) {
	keyName := "LogginKey" // Please insert your key

	inputMessage := "Hello World"
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		// Specify profile to load for the session's config
		Profile: "plugin",

		// Provide SDK Config options, such as Region.
		Config: aws.Config{
			Region: aws.String("ap-southeast-1"),
		},
	}))
	kmsClient := kms.New(sess, aws.NewConfig())

	encrypted, err := Encrypt(kmsClient, keyName, []byte(inputMessage))
	if err != nil {
		panic(err)
	}

	kmsClient = kms.New(sess, aws.NewConfig())
	decrypted, err := Decrypt(kmsClient, encrypted)
	if err != nil {
		panic(err)
	}

	if string(decrypted) != inputMessage {
		t.Errorf("Failed, expected %s but got %s", inputMessage, decrypted)
	} else {
		t.Logf("Success, expected %s but got %s", inputMessage, decrypted)
	}
}
