package service

import "testing"

func Test_generateSixDigits(t *testing.T) {
	o := &otpService{}
	got := o.generateSixDigits()
	if len(got) != 6 {
		t.Errorf("generateSixDigits() = %v, want %v", got, 6)
	}
}
