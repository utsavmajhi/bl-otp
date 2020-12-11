package adapters

import (
	"bytes"
	"fmt"
	"html/template"
	"net/smtp"

	"github.com/Zzocker/bl-otp/config"
	"github.com/Zzocker/bl-utils/pkg/errors"
)

type SMTP struct {
	Auth smtp.Auth
	Tmt  *template.Template
}

func (s *SMTP) SendMail(email, otp string) *errors.Er {
	var body bytes.Buffer
	body.Write([]byte(fmt.Sprintf("Subject: This is a test subject \n%s\n\n", config.EMAIL_HEADER)))
	s.Tmt.Execute(&body, struct {
		OTP string
	}{
		OTP: otp,
	})
	err := smtp.SendMail(config.SMTP_SERVER_ADDR, s.Auth, config.EMAIL_FROM, []string{email}, body.Bytes())
	if err != nil {
		errors.NewMsgln(errors.INTERNAL, err.Error())
	}
	return nil
}
