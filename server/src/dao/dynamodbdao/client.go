package dynamodbdao

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/drabinowitz/ChartHopInterview/server/src/dao"
)

const (
	userIDFieldName       = "UserID"
	recordIDFieldName     = "RecordID"
	attrNotExistsTemplate = "attribute_not_exists(%v)"
	attrExistsTemplate    = "attribute_exists(%v)"
	userIDAttributeKey    = ":user_id"
)

var (
	tableName         = "ChartHopInterviewScratch"
	userIDNotExists   = fmt.Sprintf(attrNotExistsTemplate, userIDFieldName)
	recordIDNotExists = fmt.Sprintf(attrNotExistsTemplate, recordIDFieldName)
	recordIDExists    = fmt.Sprintf(attrExistsTemplate, recordIDFieldName)

	nullAttr = dynamodb.AttributeValue{
		NULL: aws.Bool(true),
	}

	userIDEqualsExpression = userIDFieldName + " = " + userIDAttributeKey
)

func toLookupKey(userID, recordID string) map[string]*dynamodb.AttributeValue {
	return map[string]*dynamodb.AttributeValue{
		userIDFieldName: {
			S: &userID,
		},
		recordIDFieldName: {
			S: &recordID,
		},
	}
}

func unmarshal(item map[string]*dynamodb.AttributeValue) (map[string]interface{}, error) {
	var record map[string]interface{}
	err := dynamodbattribute.UnmarshalMap(item, &record)
	if err != nil {
		return record, fmt.Errorf("unexpected error unmarshaling,  item=[%v], err=[%w]", item, err)
	}

	return record, nil
}

// Dependencies requirements for instantiating dynamodb implementation of dao
type Dependencies struct {
	DynamoDB *dynamodb.DynamoDB
}

// New instantiates dynamo db implementation of dao.
func New(deps Dependencies) (dao.DAO, error) {
	if deps.DynamoDB == nil {
		return nil, fmt.Errorf("missing dynamo db dependency")
	}

	return &dynamodbdao{
		ddb: deps.DynamoDB,
	}, nil
}

type dynamodbdao struct {
	ddb *dynamodb.DynamoDB
}

func (dynamodbdao *dynamodbdao) Remove(userID, recordID string) error {
	_, err := dynamodbdao.ddb.DeleteItem(&dynamodb.DeleteItemInput{
		Key:       toLookupKey(userID, recordID),
		TableName: &tableName,
	})

	if err != nil {
		return fmt.Errorf("unexpected error removing record, userID=[%v], recordID=[%v], err=[%w]", userID, recordID, err)
	}

	return nil
}

func (dynamodbdao *dynamodbdao) Read(userID, recordID string) (map[string]interface{}, error) {
	consistentRead := true
	resp, err := dynamodbdao.ddb.GetItem(&dynamodb.GetItemInput{
		ConsistentRead: &consistentRead,
		Key:            toLookupKey(userID, recordID),
		TableName:      &tableName,
	})

	if err != nil {
		return nil, fmt.Errorf("unexpected error reading record, userID=[%v], recordID=[%v], err=[%w]", userID, recordID, err)
	}

	if resp.Item == nil {
		return nil, fmt.Errorf("failed to find record, userID=[%v], recordID=[%v], err=[%w]", userID, recordID, err)
	}

	return unmarshal(resp.Item)
}

func (dynamodbdao *dynamodbdao) List(userID string) ([]map[string]interface{}, error) {
	queryInput := &dynamodb.QueryInput{
		ConsistentRead: aws.Bool(true),
		// return results in reverse order
		ScanIndexForward:       aws.Bool(false),
		KeyConditionExpression: aws.String("UserID = :partitionkeyval"),
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":partitionkeyval": {
				S: &userID,
			},
		},
		TableName: &tableName,
	}
	queryOutput, err := dynamodbdao.ddb.Query(queryInput)

	if err != nil {
		return nil, fmt.Errorf("unexpected error obtaining records for user, userID=[%v], err=[%w]", userID, err)
	}

	records := make([]map[string]interface{}, 0, len(queryOutput.Items))
	err = dynamodbattribute.UnmarshalListOfMaps(queryOutput.Items, &records)

	if err != nil {
		return nil, fmt.Errorf("unexpected error unmarshaling records, userID=[%v], err=[%w]", userID, err)
	}

	return records, nil
}

func (dynamodbdao *dynamodbdao) Create(record interface{}) error {
	ddbItem, err := dynamodbattribute.MarshalMap(record)

	if err != nil {
		return fmt.Errorf(
			"unexpected error marshaling record for dynamodb, record=[%v], err=[%w]", record, err,
		)
	}

	_, err = dynamodbdao.ddb.PutItem(&dynamodb.PutItemInput{
		Item:                ddbItem,
		ConditionExpression: &recordIDNotExists,
		TableName:           &tableName,
	})

	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			if awsErr.Code() == dynamodb.ErrCodeConditionalCheckFailedException {
				return fmt.Errorf("record with this id already exists record=[%v], err=[%w]", record, awsErr)
			}
		}
		return fmt.Errorf("unexpected error writing record to dynamo, record=[%v], err=[%w]", record, err)
	}

	return nil
}

func (dynamodbdao *dynamodbdao) Update(record interface{}) error {
	ddbItem, err := dynamodbattribute.MarshalMap(record)

	if err != nil {
		return fmt.Errorf(
			"unexpected error marshaling record for dynamodb, record=[%v], err=[%w]", record, err,
		)
	}

	_, err = dynamodbdao.ddb.PutItem(&dynamodb.PutItemInput{
		Item:                ddbItem,
		ConditionExpression: &recordIDExists,
		TableName:           &tableName,
	})

	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			if awsErr.Code() == dynamodb.ErrCodeConditionalCheckFailedException {
				return fmt.Errorf("record with this id does not already exist record=[%v], err=[%w]", record, awsErr)
			}
		}
		return fmt.Errorf("unexpected error writing record to dynamo, record=[%v], err=[%w]", record, err)
	}

	return nil
}
