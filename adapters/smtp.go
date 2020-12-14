package adapters

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/smtp"

	"github.com/Zzocker/bl-otp/config"
)

type SMTP struct {
	Auth smtp.Auth
	Tmt  *template.Template
	Ch   <-chan config.MailChannelMsg
}

func (s *SMTP) StartDaemon() {
	log.Println("smtp daemon started")
	for {
		msg := <-s.Ch
		go func(msg config.MailChannelMsg) {
			var body bytes.Buffer
			body.Write([]byte(fmt.Sprintf("Subject: Signup: E-mail verification \n%s\n\n", config.EMAIL_HEADER)))
			s.Tmt.Execute(&body, struct {
				OTP string
			}{
				OTP: msg.OTP,
			})
			err := smtp.SendMail(config.SMTP_HOST+":"+config.SMTP_PORT, s.Auth, config.EMAIL_FROM, []string{msg.From}, body.Bytes())
			if err != nil {
				log.Println(err)
				return
			}
		}(msg)
	}
}
