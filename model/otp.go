package model

import (
	"crypto/rand"
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/Zzocker/bl-otp/config"
)

// EmailOTP :
type EmailOTP struct {
	Email      string
	OTP        string
	ExpiryTime int64
}

func New(email string) EmailOTP {
	return EmailOTP{
		Email:      email,
		OTP:        generateOTP(),
		ExpiryTime: time.Now().Add(config.OTP_EXPIRY_DURATION * time.Minute).Unix(),
	}
}

func generateOTP() string {
	max := big.NewInt(9999)
	n, err := rand.Int(rand.Reader, max)
	if err != nil {
		log.Fatal(err)
	}
	return fmt.Sprintf("%04d", n.Int64())
}
