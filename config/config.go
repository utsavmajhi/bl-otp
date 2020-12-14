package config

import "os"

const (
	OTP_EXPIRY_DURATION = 3 //m
	SERVER_RUNNING_PORT = "8081"
)

var (
	EMAIL_HEADER   = "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	SMTP_HOST      = os.Getenv("SMTP_HOST") //"smtp.gmail.com"
	SMTP_PORT      = os.Getenv("SMTP_PORT") //"587"
	EMAIL_FROM     = os.Getenv("EMAIL_FROM")
	PASSWORD       = os.Getenv("PASSWORD")
	MYSQL_URL      = os.Getenv("MYSQL_URL")
	MYSQL_USERNAME = os.Getenv("MYSQL_USERNAME")
	MYSQL_PASSWORD = os.Getenv("MYSQL_PASSWORD")
	MYSQL_DATABASE = os.Getenv("MYSQL_DATABASE")
)

type MailChannelMsg struct {
	From string
	OTP  string
}
