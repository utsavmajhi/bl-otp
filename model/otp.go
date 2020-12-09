package model

// EmailOTP :
type EmailOTP struct {
	Email      string
	OTP        string
	ExpiryTime int64
}
