package dynamodbdao

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

var ddb *dynamodb.DynamoDB

// InitializeTest creates a local test version of the db
func InitializeTest() *dynamodb.DynamoDB {
	if ddb == nil {
		// will need to run dynamo db local with docker for this
		// https://hub.docker.com/r/amazon/dynamodb-local
		sess, _ := session.NewSession(&aws.Config{
			Region:   aws.String("us-east-1"),
			Endpoint: aws.String("http://localhost:8000"),
		})
		ddb = dynamodb.New(sess)
		createRecordsTableIfItDoesNotExist()
	}

	return ddb
}

func createRecordsTableIfItDoesNotExist() {
	existingTables, err := ddb.ListTables(nil)
	if err != nil {
		panic("unable to determine if records table exists")
	}

	for _, tName := range existingTables.TableNames {
		if *tName == tableName {
			return
		}
	}

	createTableInput := &dynamodb.CreateTableInput{
		BillingMode: aws.String("PAY_PER_REQUEST"),
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String("UserID"),
				AttributeType: aws.String("S"),
			},
			{
				AttributeName: aws.String("RecordID"),
				AttributeType: aws.String("S"),
			},
		},
		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String("UserID"),
				KeyType:       aws.String("HASH"),
			},
			{
				AttributeName: aws.String("RecordID"),
				KeyType:       aws.String("RANGE"),
			},
		},
		TableName: &tableName,
	}

	_, err = ddb.CreateTable(createTableInput)
	if err != nil {
		panic(fmt.Sprintf("failed to create records table, err=%v", err))
	}
}
