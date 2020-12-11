package core

import (
	"github.com/Zzocker/bl-otp/core/ports"
	"github.com/Zzocker/bl-utils/pkg/errors"
)

type OTPBusiness struct {
	DS   ports.OTPDatastoreInterface
	SMTP ports.SMTPServiceInterface
	UP   ports.UserprofileInterface
}

func (o *OTPBusiness) Send(email string) *errors.Er {
	return errors.NewMsgln(errors.UNIMPLEMENTED, "Send Core business not implemented")
}
func (o *OTPBusiness) Verify(email string, otp string) *errors.Er {
	return errors.NewMsgln(errors.UNIMPLEMENTED, "Verify Core business not implemented")
}
