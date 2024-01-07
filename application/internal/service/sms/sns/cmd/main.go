package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sns"
)

func main() {
	//msg := flag.String("m", "", "The message to send to the subscribed users of the topic")
	//topicARN := flag.String("t", "", "The ARN of the topic to which the user subscribes")

	//flag.Parse()
	//
	//if *msg == "" {
	//	fmt.Println("You must supply a message and topic ARN")
	//	fmt.Println("-m MESSAGE -t TOPIC-ARN")
	//	return
	//}

	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		panic("configuration error, " + err.Error())
	}

	client := sns.NewFromConfig(cfg)

	mobile := ""
	msg := "Hello World!"
	input := &sns.PublishInput{
		Message:     &msg,
		PhoneNumber: &mobile,
	}

	result, err := PublishMessage(context.TODO(), client, input)
	if err != nil {
		fmt.Println("Got an error publishing the message:")
		fmt.Println(err)
		return
	}

	fmt.Println("Message ID: " + *result.MessageId)
}

type SNSPublishAPI interface {
	Publish(ctx context.Context,
		params *sns.PublishInput,
		optFns ...func(*sns.Options)) (*sns.PublishOutput, error)
}

func PublishMessage(c context.Context, api SNSPublishAPI, input *sns.PublishInput) (*sns.PublishOutput, error) {
	return api.Publish(c, input)
}
