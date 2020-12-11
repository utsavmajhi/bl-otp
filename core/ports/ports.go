package ports

import (
	"github.com/Zzocker/bl-otp/model"
	"github.com/Zzocker/bl-utils/pkg/errors"
)

// OTPDatastoreInterface :
type OTPDatastoreInterface interface {
	Create(model.EmailOTP) *errors.Er
	Read(email string) (*model.EmailOTP, *errors.Er)
	Update(model.EmailOTP) *errors.Er
	Delete(email string) *errors.Er
}

// SMTPServiceInterface :
type SMTPServiceInterface interface {
	SendMail(email, otp string) *errors.Er
}

// UserprofileInterface :
type UserprofileInterface interface {
	IsEmailAvailable(email string) *errors.Er
}
