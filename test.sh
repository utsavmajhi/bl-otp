#/bin/bash

ADDRESS=localhost:8081

case $1 in 
    "send")
        EMAIL_ID=$2
        grpcurl --plaintext -d '{"EmailID":"'"$EMAIL_ID"'"}' $ADDRESS otp.OTPServices.SendOTP
    ;;
    "verify")
        EMAIL_ID=$2
        OTP=$3
        grpcurl --plaintext -d '{"EmailID":"'"$EMAIL_ID"'","OTP":"'"$OTP"'"}' $ADDRESS otp.OTPServices.VerifyOTP
    ;;
esac 

# grpcurl --plaintext -msg-template localhost:8081 describe .otp.VerifyOTPRequest