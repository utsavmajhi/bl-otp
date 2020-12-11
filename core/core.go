package core

import "github.com/Zzocker/bl-utils/pkg/errors"

// OTP interface for otp core business
type OTP interface {
	Send(email string) *errors.Er
	Verify(email string, otp string) *errors.Er
}
