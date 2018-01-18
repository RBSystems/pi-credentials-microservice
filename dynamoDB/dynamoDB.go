package dynamoDB

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

var dbClient *dynamodb.DynamoDB

func Init(session *session.Session) {

	dbClient = dynamodb.New(session)
}
