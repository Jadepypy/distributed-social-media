package service

import (
	"context"
	"fmt"
	"github.com/Jadepypy/distributed-social-media/application/internal/domain"
	"github.com/Jadepypy/distributed-social-media/application/internal/service/sms"
	"math/rand"
)

type OtpService interface {
	SendOtp(ctx context.Context, biz string, phoneNumber string) error
	VerifyOtp(ctx context.Context, user domain.User, otp string) error
}

type otpService struct {
	smsService sms.Service
}

func NewOtpService(smsService sms.Service) OtpService {
	return &otpService{
		smsService: smsService,
	}
}

func (o *otpService) SendOtp(ctx context.Context, biz string, phoneNumber string) error {
	otp := o.generateSixDigits()
	msg := "Your otp is " + otp
	if err := o.smsService.SendSingle(ctx, phoneNumber, msg); err != nil {
		return err
	}

	return nil
}

func (o *otpService) VerifyOtp(ctx context.Context, user domain.User, otp string) error {
	return nil
}

func (o *otpService) generateSixDigits() string {
	num := rand.Intn(1000000)
	return fmt.Sprintf("%06d", num)
}
