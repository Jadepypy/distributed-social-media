package sms

import "context"

type Service interface {
	SendSingle(ctx context.Context, mobile string, msg string) error
}
