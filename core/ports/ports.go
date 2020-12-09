package ports

import (
	"github.com/Zzocker/bl-otp/model"
	"github.com/Zzocker/bl-utils/pkg/errors"
)

// ClientPortsInterface port interface to used by client adapter
type ClientPortsInterface interface {
	GetOTP(email string) *errors.Er
	VerifyOTP(email string, otp string) *errors.Er
}

// OTPDatastoreInterface :
type OTPDatastoreInterface interface {
	InsertOTP(model.EmailOTP) *errors.Er
	UpdateExpiryTime(int64) *errors.Er
	DeleteOTP(email string) *errors.Er
}
// SMTPServiceInterface :
type SMTPServiceInterface interface{
	SendMail(email,otp string)
}

// UserprofileInterface :
type UserprofileInterface interface{
	IsEmailAvailable(email string) *errors.Er
}