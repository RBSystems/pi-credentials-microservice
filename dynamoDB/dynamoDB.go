package dynamoDB

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	db "github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/byuoitav/pi-credentials-microservice/kms"
	"github.com/byuoitav/pi-credentials-microservice/structs"
	"github.com/fatih/color"
)

var dbClient *db.DynamoDB

func Init(session *session.Session) {

	dbClient = db.New(session)
}

//encrypts the password field in an entry
//writes the hostname and encrypted password to DB
//duplicate primary keys overwrite indices (?)
func AddEntry(entry *structs.Entry) (*db.PutItemOutput, error) {

	log.Printf("%s", color.HiGreenString("[dynamodb] adding entry for host %s...", entry.Hostname))

	//encrypt password
	log.Printf("[dynamodb] encrypting entry...")
	cipherText, err := kms.Encrypt(entry.Password)
	if err != nil {
		return &db.PutItemOutput{}, err
	}

	//add entry to DB
	log.Printf("[dynamodb] building DB input struct...")
	input := &db.PutItemInput{
		Item: map[string]*db.AttributeValue{
			"hostname": {
				S: &entry.Hostname,
			},
			"password": {
				B: cipherText,
			},
		},
		TableName: aws.String(os.Getenv("AWS_DYNAMO_TABLE")),
	}

	log.Printf("[dynamodb] adding struct to DB...")
	return dbClient.PutItem(input)
}

func GetEntry(hostname string) (*structs.Entry, error) {

	log.Printf("%s", color.HiGreenString("[dynamodb] collecting indices for host %s...", hostname))

	//build request struct
	input := &db.GetItemInput{
		Key: map[string]*db.AttributeValue{
			"hostname": {
				S: &hostname,
			},
		},
		TableName: aws.String(os.Getenv("AWS_DYNAMO_TABLE")),
	}

	//get entry
	log.Printf("[dynamodb] searching for item...")
	result, err := dbClient.GetItem(input)
	if err != nil {
		msg := fmt.Sprintf("entry not found: %s", err.Error())
		log.Printf("%s", color.HiRedString("[dynamodb] %s", msg))
		return &structs.Entry{}, errors.New(msg)
	}

	//decrypt entry
	log.Printf("[dynamodb] decrypting password...")
	plaintext, err := kms.Decrypt(result.Item["password"].B)
	if err != nil {
		msg := fmt.Sprintf("failed to decrypt password: %s", err.Error())
		log.Printf("%s", color.HiRedString("[dynamodb] %s", msg))
		return &structs.Entry{}, errors.New(msg)
	}

	return &structs.Entry{
		Hostname: hostname,
		Password: plaintext,
	}, nil

}

//deletes item
//multiple deletes does not result in error
func DeleteEntry(hostname string) error {

	log.Printf("%s", color.HiGreenString("[dynamodb] deleting indices for host: %s", hostname))

	//build request struct
	log.Printf("[dynamodb] building request struct...")
	input := &db.DeleteItemInput{
		Key: map[string]*db.AttributeValue{
			"hostname": {
				S: &hostname,
			},
		},
		TableName: aws.String(os.Getenv("AWS_DYNAMO_TABLE")),
	}

	//remove entry
	log.Printf("[dynamodb] removing host...")
	_, err := dbClient.DeleteItem(input)
	if err != nil {
		msg := fmt.Sprintf("item not deleted: %s", err.Error())
		log.Printf("%s", color.HiRedString("[dynamodb] %s", msg))
		return errors.New(msg)
	}

	return nil
}
