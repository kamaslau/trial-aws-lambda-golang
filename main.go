package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

type MyEvent struct {
	AppId     string `json:"appId"`
	AppSecret string `json:"appSecret"`
	Code      string `json:"code"`
}

func HandleRequest(ctx context.Context, event *MyEvent) (*string, error) {
	if event == nil {
		return nil, fmt.Errorf("received nil event")
	}
	message := fmt.Sprintf("AppId: %s\r\n", event.AppId)
	message += fmt.Sprintf("AppSecret: %s\r\n", event.AppSecret)
	message += fmt.Sprintf("Code: %s", event.Code)
	
	return &message, nil
}

func main() {
	lambda.Start(HandleRequest)
}
