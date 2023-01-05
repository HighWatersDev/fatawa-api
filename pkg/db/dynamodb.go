package db

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"log"
)

type DynamoInstance struct {
	Client *dynamodb.Client
	Table  string
}

const TableName = "Fatawa"

var DI DynamoInstance

func ConnectDynamoDB() {
	sdkConfig, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	client := dynamodb.NewFromConfig(sdkConfig)

	fmt.Println("DynamoDB Connected!")

	DI = DynamoInstance{
		Client: client,
		Table:  TableName,
	}

}

//func waitForTable(ctx context.Context, db *dynamodb.Client, tn string) error {
//	w := dynamodb.NewTableExistsWaiter(db)
//	err := w.Wait(ctx,
//		&dynamodb.DescribeTableInput{
//			TableName: aws.String(tn),
//		},
//		2*time.Minute,
//		func(o *dynamodb.TableExistsWaiterOptions) {
//			o.MaxDelay = 5 * time.Second
//			o.MinDelay = 5 * time.Second
//		})
//	if err != nil {
//		return errors.Wrap(err, "timed out while waiting for table to become active")
//	}
//
//	return err
//}
