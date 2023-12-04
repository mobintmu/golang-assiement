package entity

import "time"

type User struct {
	ID                string    `json:"id"`
	Name              string    `json:"name"`
	PhoneNumber       string    `json:"phone_number"`
	OTP               string    `json:"otp,omitempty"`
	OTPExpirationTime time.Time `json:"otp_expiration_time,omitempty"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}
