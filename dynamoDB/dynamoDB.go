package dynamoDB

import (
	"errors"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws/session"
	db "github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/byuoitav/pi-credentials-microservice/kms"
	"github.com/fatih/color"
)

var dbClient *db.DynamoDB

func Init(session *session.Session) {

	dbClient = db.New(session)
}

func AddEntry(entry *Entry) error {

	//encrypt entry
	username, password, err := kms.EncryptEntry(entry)
	if err != nil {
		return err
	}

	//add entry to DB
	input := &db.PutItemInput{
		Item: map[string]*db.AttributeValue{
			string(username): {
				B: password,
			},
		},
	}

	result, err := dbClient.PutItem(input)
	if err != nil {
		msg := fmt.Sprintf("item not added: %s", err.Error())
		log.Printf("%s", color.HiRedString("[dynamoDB] %s", msg))
		return errors.New(msg)
	}

	return nil
}
