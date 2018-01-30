package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

// Create structs to hold info about new item
type Item struct {
	ID    string `json:"id"`
	Value string `json:"value"`
}

type Repository interface {
	GetItem(id string) string
}

type DynamoRepo struct {
	Region       string
	TableName    string
	IDColumnName string
}

func (repo DynamoRepo) GetItem(id string) {

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(repo.Region)},
	)

	// Create DynamoDB client
	svc := dynamodb.New(sess)

	// Create table Movies
	input := &dynamodb.CreateTableInput{
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String("id"),
				AttributeType: aws.String("S"),
			},
		},
		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String("id"),
				KeyType:       aws.String("HASH"),
			},
		},
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(10),
			WriteCapacityUnits: aws.Int64(10),
		},
		TableName: aws.String(repo.TableName),
	}

	_, err = svc.CreateTable(input)

	if err != nil {
		fmt.Println("Got error calling CreateTable:")
		fmt.Println(err.Error())
		//os.Exit(1)
	}

	fmt.Println("Created the table")

	sess, err = session.NewSession(&aws.Config{
		Region: aws.String(repo.Region)},
	)

	// Create DynamoDB client
	svc = dynamodb.New(sess)

	item := Item{
		ID:    "222",
		Value: "The Big New Movie",
	}

	av, err := dynamodbattribute.MarshalMap(item)

	if err != nil {
		fmt.Println("Got error marshalling map:")
		fmt.Println(err.Error())
		os.Exit(1)
	}

	// Create item in table Movies
	pii := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(repo.TableName),
	}

	_, err = svc.PutItem(pii)

	if err != nil {
		fmt.Println("Got error calling PutItem:")
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Println("Successfully added")

	// Initialize a session in us-west-2 that the SDK will use to load
	// credentials from the shared credentials file ~/.aws/credentials.
	sess, err = session.NewSession(&aws.Config{
		Region: aws.String(repo.Region)},
	)

	// Create DynamoDB client
	svc = dynamodb.New(sess)

	result, err := svc.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(repo.TableName),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String("222"),
			},
		},
	})

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	item = Item{}

	err = dynamodbattribute.UnmarshalMap(result.Item, &item)

	if err != nil {
		panic(fmt.Sprintf("Failed to unmarshal Record, %v", err))
	}

	fmt.Println("Found item:")
	fmt.Println("Year:  ", item.Value)
}
