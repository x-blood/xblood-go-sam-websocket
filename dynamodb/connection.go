package dynamodb

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/pkg/errors"
)

const tableName = "xblood-go-sam-websocket-table"

type Connection struct {
	ConnectionID string `dynamo"connectionId,hash"`
}

func Put(connectionID string) error {
	db, err := connect()
	if err != nil {
		return errors.WithStack(err)
	}
	table := getTable(db, tableName)
	putModel := Connection{ConnectionID: connectionID}
	table.Put(putModel).Run()

	return nil
}

func GetAll() ([]Connection, error) {
	db, err := connect()
	if err != nil {
		return nil, errors.WithStack(err)
	}
	table := getTable(db, tableName)
	scan := table.Scan()

	var results []Connection
	err = scan.All(&results)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return results, nil
}

func Delete(connectionID string) error {
	db, err := connect()
	if err != nil {
		return errors.WithStack(err)
	}
	table := getTable(db, tableName)
	table.Delete("connectionId", connectionID).Run()

	return nil
}


func getTable(db *dynamo.DB, tableName string) dynamo.Table {
	return db.Table(tableName)
}

func connect() (*dynamo.DB, error) {
	dynamoSession, err := session.NewSession(&aws.Config{Region: aws.String("ap-northeast-1")})
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return dynamo.New(dynamoSession), nil
}


