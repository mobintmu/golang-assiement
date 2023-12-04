package test

import (
	"aqary/entity"
	"aqary/repository/postgres"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var PostgresInstance *postgres.PostgresRepository

func getDB() *postgres.PostgresRepository {
	if PostgresInstance != nil {
		return PostgresInstance
	}

	//read from .env
	postgresURL := "postgresql://postgres:postgres@localhost:5432/user_db?sslmode=disable"

	ctx := context.Background()
	PostgresInstance, err := postgres.NewClient(ctx, postgresURL)
	if err != nil {
		log.Fatal(err)
	}

	return PostgresInstance

}

func TestMain(m *testing.M) {
	ctx := context.Background()
	_, err := getDB().DB.Exec(ctx, "TRUNCATE TABLE users CASCADE")
	if err != nil {
		log.Fatal(err)
	}
	exitCode := m.Run()
	getDB().DB.Close()
	os.Exit(exitCode)
}

func TestStoreUser(t *testing.T) {

	name := "John"
	phone := "+971212123453"

	data := url.Values{
		"name":         {name},
		"phone_number": {phone},
	}

	urlAddress := "http://localhost:8080/api"
	resp, err := http.PostForm(urlAddress+"/users", data)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	var res entity.User
	json.NewDecoder(resp.Body).Decode(&res)

	assert.Equal(t, resp.StatusCode, 200)
	assert.Equal(t, res.Name, name)

	//--------------------- send again

	resp, err = http.PostForm(urlAddress+"/users", data)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	json.NewDecoder(resp.Body).Decode(&res)

	assert.Equal(t, resp.StatusCode, 400)

	// -------------------------- Generate OTP

	data2 := url.Values{
		"phone_number": {phone},
	}

	resp, err = http.PostForm(urlAddress+"/users/generateotp", data2)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	var res2 entity.User
	json.NewDecoder(resp.Body).Decode(&res2)

	assert.Equal(t, resp.StatusCode, http.StatusOK)

	result, _ := getDB().Queries.GetUser(context.Background(), phone)

	assert.Equal(t, result.Name, name)
	assert.Equal(t, result.PhoneNumber, phone)
	assert.Equal(t, len(result.Otp.String), 4)

	// --------------------------- Verify OTP
	data3 := url.Values{
		"phone_number": {phone},
		"otp":          {result.Otp.String},
	}

	resp, err = http.PostForm(urlAddress+"/users/verifyotp", data3)
	if err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, resp.StatusCode, 200)

}

func TestOPTNotFind(t *testing.T) {
	phone := "+971212123453"
	urlAddress := "http://localhost:8080/api"

	data2 := url.Values{
		"phone_number": {phone},
	}

	resp, err := http.PostForm(urlAddress+"/users/generateotp", data2)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	var res2 entity.User
	json.NewDecoder(resp.Body).Decode(&res2)

	assert.Equal(t, resp.StatusCode, 404)
}
