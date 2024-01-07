package sns

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sns"
)

type service struct {
	client *sns.Client
}

func NewService() *service {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		panic("configuration error, " + err.Error())
	}
	client := sns.NewFromConfig(cfg)
	return &service{
		client: client,
	}
}

type SNSPublishAPI interface {
	Publish(ctx context.Context,
		params *sns.PublishInput,
		optFns ...func(*sns.Options)) (*sns.PublishOutput, error)
}

func PublishMessage(c context.Context, api SNSPublishAPI, input *sns.PublishInput) (*sns.PublishOutput, error) {
	return api.Publish(c, input)
}

func (s *service) SendSingle(ctx context.Context, mobile string, msg string) error {
	input := &sns.PublishInput{
		Message:     &msg,
		PhoneNumber: &mobile,
	}

	result, err := PublishMessage(context.TODO(), s.client, input)
	if err != nil {
		return err
	}

	fmt.Println("Message ID: " + *result.MessageId)
	return nil
}
