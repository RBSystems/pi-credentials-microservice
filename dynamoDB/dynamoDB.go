package dynamoDB

import (
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

	log.Printf("%s", color.HiGreenString("[dynamodb] adding entry for host %s", entry.Hostname))

	log.Printf("[dynamodb] encrypting entry...")
	_, password, err := kms.EncryptDbEntry(entry)
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
				B: password.CiphertextBlob,
			},
		},
		TableName: aws.String(os.Getenv("AWS_DYNAMO_TABLE")),
	}

	log.Printf("[dynamodb] adding struct to DB...")
	return dbClient.PutItem(input)
}

func GetEntry(hostname string) (*db.GetItemOutput, error) {

	log.Printf("%s", color.HiGreenString("[dynamodb] collecting indices for host %s", hostname))

	input := &db.GetItemInput{
		Key: map[string]*db.AttributeValue{
			"hostname": {
				S: &hostname,
			},
		},
		TableName: aws.String(os.Getenv("AWS_DYNAMO_TABLE")),
	}

	return dbClient.GetItem(input)
}
