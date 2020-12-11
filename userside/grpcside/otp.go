package grpcside

import (
	"context"
	"log"

	"github.com/Zzocker/bl-otp/core"
	"github.com/Zzocker/bl-proto-go/common"
	pb "github.com/Zzocker/bl-proto-go/otp"
)

type OTPSide struct {
	Core core.OTP
	pb.UnsafeOTPServicesServer
}

// SendOTP :
func (o *OTPSide) SendOTP(ctx context.Context, in *pb.GetOTPRequest) (*pb.GetOTPResponse, error) {
	log.Println("inside SendOTP")
	header := common.Header{
		Status: common.StatusCode_OK,
	}
	err := o.Core.Send(in.EmailID)
	if err != nil {
		header.Status = common.StatusCode(err.Status)
		header.Description = err.Error()
	}
	return &pb.GetOTPResponse{
		Header: &header,
	}, nil
}

// VerifyOTP :
func (o *OTPSide) VerifyOTP(ctx context.Context, in *pb.VerifyOTPRequest) (*pb.VerifyOTPResponse, error) {
	log.Println("inside VerifyOTP")
	header := common.Header{
		Status: common.StatusCode_OK,
	}
	err := o.Core.Verify(in.EmailID, in.OTP)
	if err != nil {
		header.Status = common.StatusCode(err.Status)
		header.Description = err.Error()
	}
	return &pb.VerifyOTPResponse{
		Header: &header,
	}, nil
}
