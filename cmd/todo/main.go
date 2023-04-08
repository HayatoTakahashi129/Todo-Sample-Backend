//go:build release

package main

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	fiberadapter "github.com/awslabs/aws-lambda-go-api-proxy/fiber"
)

var fiberLambda *fiberadapter.FiberLambda

func init() {

	app := CreateRoute()
	fiberLambda = fiberadapter.New(app)
}

/**
* AWS-Gateway/AWS-Lambda 構成で来たイベント情報を元に、Fiberでハンドリングできるようプロキシライブラリで変換
 */
func FiberHandler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// If no name is provided in the HTTP request body, throw an error
	return fiberLambda.ProxyWithContext(ctx, req)
}

func main() {
	lambda.Start(FiberHandler)
}
