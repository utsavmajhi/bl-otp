package main

import (
	"log"
	"net"

	"github.com/Zzocker/bl-otp/adapters"
	"github.com/Zzocker/bl-otp/core"
	"github.com/Zzocker/bl-otp/core/ports"
	"github.com/Zzocker/bl-otp/userside/grpcside"
	pb "github.com/Zzocker/bl-proto-go/otp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	lis, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatal(err)
	}
	srv := grpc.NewServer()
	otpServ, err := createOTPServerSide()
	if err != nil {
		log.Fatal(err)
	}
	reflection.Register(srv)
	pb.RegisterOTPServicesServer(srv, otpServ)
	if err := srv.Serve(lis); err != nil {
		log.Fatal(err)
	}

}

func createOTPDatastore() (ports.OTPDatastoreInterface, error) {
	return &adapters.OTPDatastore{}, nil
}

func createOTPCore() (core.OTP, error) {
	ds, err := createOTPDatastore()
	if err != nil {
		return nil, err
	}
	return &core.OTPBusiness{
		DS: ds,
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
