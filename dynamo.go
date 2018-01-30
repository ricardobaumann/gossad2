package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

//Item represents a rown in the database
type Item struct {
	ID    string `json:"id"`
	Value string `json:"value"`
}

//Repository describes the method to interface with database
type Repository interface {
	GetItem(id string) string
}

//DynamoRepo is a dynamo db Repository implementation
type DynamoRepo struct {
	Region       string
	TableName    string
	IDColumnName string
}

//GetItem fetches a record from database
func (repo DynamoRepo) GetItem(id string) string {

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(repo.Region)},
	)

	// Create DynamoDB client
	svc := dynamodb.New(sess)

	result, err := svc.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(repo.TableName),
		Key: map[string]*dynamodb.AttributeValue{
			repo.IDColumnName: {
				S: aws.String(id),
			},
		},
	})

	if err != nil {
		fmt.Println(err.Error())
		return ""
	}

	item := Item{}

	err = dynamodbattribute.UnmarshalMap(result.Item, &item)

	if err != nil {
		panic(fmt.Sprintf("Failed to unmarshal Record, %v", err))
	}

	fmt.Println("Found item:")
	fmt.Println("Value:  ", item.Value)
	return item.Value
}
