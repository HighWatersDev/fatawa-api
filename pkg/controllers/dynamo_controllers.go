package controllers

import (
	"context"
	"errors"
	"fatawa-api/pkg/models"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
)

type TableFatawa struct {
	DynamoDbClient *dynamodb.Client
	TableName      string
}

func (fatawa TableFatawa) TableExists() (bool, error) {
	exists := true
	_, err := fatawa.DynamoDbClient.DescribeTable(
		context.TODO(), &dynamodb.DescribeTableInput{TableName: aws.String(fatawa.TableName)},
	)
	if err != nil {
		var notFoundEx *types.ResourceNotFoundException
		if errors.As(err, &notFoundEx) {
			log.Printf("Table %v does not exist.\n", fatawa.TableName)
			err = nil
		} else {
			log.Printf("Couldn't determine existence of table %v. Here's why: %v\n", fatawa.TableName, err)
		}
		exists = false
	}
	return exists, err
}

func (fatawa TableFatawa) addFatwa(fatwa *models.Fatwa, ctx *gin.Context) {
	item, err := attributevalue.MarshalMap(fatwa)
	if err != nil {
		panic(err)
	}
	out, err := fatawa.DynamoDbClient.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName: aws.String(fatawa.TableName), Item: item,
	})
	if err := ctx.ShouldBindJSON(&fatwa); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	if err != nil {
		if strings.Contains(err.Error(), "title already exists") {
			ctx.JSON(http.StatusConflict, gin.H{"status": "fail", "message": err.Error()})
			return
		}

		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": out.Attributes})
}
