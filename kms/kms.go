package kms

import (
	"github.com/aws/aws-sdk-go/aws/session"
	k "github.com/aws/aws-sdk-go/service/kms"
)

var kmsClient k.KMS

func Init(session *session.Session) {

	kmsClient = k.New(session)
}
