package service

import (
	"aqary/entity"
	"aqary/repository/postgres"
	"context"
	"database/sql"
	"errors"
	"log"
	"math/rand"
	"strconv"
	"time"
)

type User struct {
	postgres *postgres.PostgresRepository
}

func NewUser(postgres *postgres.PostgresRepository) *User {
	return &User{
		postgres: postgres,
	}
}

func (u *User) StoreUser(ctx context.Context, user entity.User) error {

	err := u.postgres.Queries.CreateUser(ctx, postgres.CreateUserParams{Name: user.Name, PhoneNumber: user.PhoneNumber})
	if err != nil {
		log.Print(err)
		return err
	}

	return nil
}

func (u *User) GenerateOTP(ctx context.Context, phoneNumber string) error {

	_, err := u.postgres.Queries.GetUser(ctx, phoneNumber)
	if err != nil {
		log.Print(err)
		return errors.New("no rows in result set")
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	otp := strconv.Itoa(r.Intn(9000) + 1000)
	sqlTime := sql.NullTime{
		Time:  time.Now().Add(time.Minute * 1),
		Valid: true,
	}
	otpString := sql.NullString{
		String: otp,
		Valid:  true,
	}

	err = u.postgres.Queries.UpdateUser(ctx, postgres.UpdateUserParams{
		PhoneNumber:       phoneNumber,
		OtpExpirationTime: sqlTime,
		Otp:               otpString,
	})
	if err != nil {
		log.Print(err)
		return err
	}

	return nil
}

func (u *User) VerifyOTP(ctx context.Context, phoneNumber, otp string) error {

	user, err := u.postgres.Queries.GetUser(ctx, phoneNumber)
	if err != nil {
		return errors.New("no rows in result set")
	}
	if user.Otp.String != otp {
		return errors.New("otp not match")
	}

	if user.OtpExpirationTime.Time.Before(time.Now()) {
		return errors.New("otp expired")
	}

	err = u.postgres.Queries.UpdateUser(ctx, postgres.UpdateUserParams{
		PhoneNumber:       phoneNumber,
		OtpExpirationTime: sql.NullTime{},
		Otp:               sql.NullString{},
	})
	if err != nil {
		log.Print(err)
		return err
	}

	return nil
}
