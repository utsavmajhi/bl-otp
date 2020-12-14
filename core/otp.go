package core

import (
	"time"

	"github.com/Zzocker/bl-otp/config"
	"github.com/Zzocker/bl-otp/core/ports"
	"github.com/Zzocker/bl-otp/model"
	"github.com/Zzocker/bl-utils/pkg/errors"
)

type OTPBusiness struct {
	DS ports.OTPDatastoreInterface
	Ch chan<- config.MailChannelMsg
}

func (o *OTPBusiness) Send(email string) *errors.Er {
	emailOTP := model.New(email)
	err := o.DS.Update(emailOTP)
	if err != nil {
		if err.Status == errors.NOT_FOUND {
			if newErr := o.DS.Create(emailOTP); newErr != nil {
				return newErr
			}
		} else {
			return err
		}
	}
	o.Ch <- config.MailChannelMsg{
		From: email,
		OTP:  emailOTP.OTP,
	}
	return nil
}
func (o *OTPBusiness) Verify(email string, otp string) *errors.Er {
	emailOTP, err := o.DS.Read(email)
	if err != nil {
		return errors.NewMsgln(errors.INVALID_ARGUMENT, "invalid OTP")
	}
	if emailOTP.OTP != otp || emailOTP.ExpiryTime < time.Now().Unix() {
		return errors.NewMsgln(errors.INVALID_ARGUMENT, "invalid OTP")
	}
	return o.DS.Delete(email)
}
