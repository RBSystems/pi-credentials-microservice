# Basic Server Password Management
## Installation:
```go get github.com/byuoitav/pi-credential-store```

## Envrionment:
The client must have the proper AWS Credentials
* ```AWS_ACCESS_KEY_ID```
* ```AWS_SECRET_ACCESS_KEY```
* ```AWS_DEFAULT_REGION```
* ```AWS_KMS_KEY_ID``` - the KMS ID set in the console
* ```AWS_DYNAMO_TABLE``` - the name of the table to store the data in. The table must have `hostname` set 
as a primary key and `password` as a secondary global index

# How it's configured in AWS

##KMS
*Alias
	`credstash`

*Tag Key
	`raspi`

*Tag Value
	`creds`

*IAM Administrators
	`PowerUser`

*IAM Users
	* `PowerUser`
	* `aws-elasticbeanstalk-service-role`

##DynamoDB
