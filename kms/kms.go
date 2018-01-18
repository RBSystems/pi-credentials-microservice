package kms

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	k "github.com/aws/aws-sdk-go/service/kms"
	"github.com/byuoitav/pi-credentials-microservice/structs"
	"github.com/fatih/color"
)

var kmsClient *k.KMS

func Init(session *session.Session) {

	kmsClient = k.New(session)
}

func EncryptDbEntry(entry *structs.Entry) (*k.EncryptOutput, *k.EncryptOutput, error) {

	hostInput := &k.EncryptInput{
		KeyId:     aws.String(os.Getenv("AWS_KMS_KEY_ID")),
		Plaintext: []byte(entry.Hostname),
	}

	passwordInput := &k.EncryptInput{
		KeyId:     aws.String(os.Getenv("AWS_KMS_KEY_ID")),
		Plaintext: []byte(entry.Password),
	}

	encryptedHost, err := kmsClient.Encrypt(hostInput)
	if err != nil {
		msg := fmt.Sprintf("failed to encrypt hostname: %s", err.Error())
		log.Printf("%s", color.HiRedString("[kms] %s", msg))
		return &k.EncryptOutput{}, &k.EncryptOutput{}, errors.New(msg)
	}

	log.Printf("result: %s", encryptedHost)

	encryptedPassword, err := kmsClient.Encrypt(passwordInput)
	if err != nil {
		msg := fmt.Sprintf("failed to encrypt hostname: %s", err.Error())
		log.Printf("%s", color.HiRedString("[kms] %s", msg))
		return &k.EncryptOutput{}, &k.EncryptOutput{}, errors.New(msg)
	}

	log.Printf("result: %s", encryptedPassword)

	return encryptedHost, encryptedPassword, nil
}
