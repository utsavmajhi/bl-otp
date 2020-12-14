package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net"
	"net/smtp"

	"github.com/Zzocker/bl-utils/pkg/datastore"

	"github.com/Zzocker/bl-otp/adapters"
	"github.com/Zzocker/bl-otp/config"
	"github.com/Zzocker/bl-otp/core"
	"github.com/Zzocker/bl-otp/core/ports"
	"github.com/Zzocker/bl-otp/userside/grpcside"
	pb "github.com/Zzocker/bl-proto-go/otp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	lis, err := net.Listen("tcp", ":"+config.SERVER_RUNNING_PORT)
	if err != nil {
		log.Fatal(err)
	}
	srv := grpc.NewServer()
	otpServ, err := createOTPServerSide()
	if err != nil {
		log.Fatal(err)
	}
	reflection.Register(srv)
	fmt.Printf("Serevr started at port : %s", config.SERVER_RUNNING_PORT)
	pb.RegisterOTPServicesServer(srv, otpServ)
	if err := srv.Serve(lis); err != nil {
		log.Fatal(err)
	}

}

func createSMTP() (ports.SMTPServiceInterface, error) {
	auth := smtp.PlainAuth("", config.EMAIL_FROM, config.PASSWORD, config.SMTP_HOST)
	tmt, err := template.ParseFiles("template/otp_mail.html")
	if err != nil {
		return nil, err
	}
	return &adapters.SMTP{
		Auth: auth,
		Tmt:  tmt,
	}, nil
}

func createOTPDatastore() (ports.OTPDatastoreInterface, error) {
	dbcfg := datastore.DatastoreConfig{
		Code:     "mysql",
		URL:      config.MYSQL_URL,
		Username: config.MYSQL_USERNAME,
		Password: config.MYSQL_PASSWORD,
		DBName:   config.MYSQL_DATABASE,
	}
	db, err := datastore.FromFactory("mysql").Build(dbcfg)
	if err != nil {
		return nil, err
	}
	return &adapters.OTPDatastore{
		DB: db.(*sql.DB),
	}, nil
}

func createOTPCore() (core.OTP, error) {
	ds, err := createOTPDatastore()
	if err != nil {
		return nil, err
	}
	smt, err := createSMTP()
	if err != nil {
		return nil, err
	}
	msgchannel := make(chan config.MailChannelMsg, 30)
	smt.(*adapters.SMTP).Ch = msgchannel
	go smt.StartDaemon()
	return &core.OTPBusiness{
		DS: ds,
		Ch: msgchannel,
	}, nil
}

func createOTPServerSide() (pb.OTPServicesServer, error) {
	cre, err := createOTPCore()
	if err != nil {
		return nil, err
	}
	return &grpcside.OTPSide{
		Core: cre,
	}, nil
}
