package kms

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	k "github.com/aws/aws-sdk-go/service/kms"
	"github.com/fatih/color"
)

var kmsClient *k.KMS

//creates the KMS client given an AWS session
func Init(session *session.Session) {

	kmsClient = k.New(session)
}

//checks out the KMS key stored in AWS_KMS_KEY_ID and
//encrypts a string
//clients must call Init() before using
func Encrypt(secret string) ([]byte, error) {

	input := &k.EncryptInput{
		KeyId:     aws.String(os.Getenv("AWS_KMS_KEY_ID")),
		Plaintext: []byte(secret),
	}

	result, err := kmsClient.Encrypt(input)
	if err != nil {
		msg := fmt.Sprintf("failed to encrypt data: %s", err.Error())
		log.Printf("%s", color.HiRedString("[kms] %s", msg))
		return []byte{}, errors.New(msg)
	}

	return result.CiphertextBlob, nil
}

func Decrypt(secret []byte) (string, error) {

	input := &k.DecryptInput{
		CiphertextBlob: secret,
	}

	result, err := kmsClient.Decrypt(input)
	if err != nil {
		msg := fmt.Sprintf("failed to encrypt data: %s", err.Error())
		log.Printf("%s", color.HiRedString("[kms] %s", msg))
		return "", errors.New(msg)
	}

	return string(result.Plaintext), nil
}
