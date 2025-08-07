package main

import (
	"context"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-gonic/gin"

	"content-service/internal/server"
)

var ginLambda *ginadapter.GinLambda

func init() {
	log.Printf("Gin cold start")
	ginServer := server.NewServer()
	ginLambda = ginadapter.New(ginServer.Handler.(*gin.Engine))
}

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return ginLambda.ProxyWithContext(ctx, req)
}

func main() {
	lambda.Start(Handler)
}
