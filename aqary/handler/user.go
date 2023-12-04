package handler

import (
	"aqary/entity"
	"aqary/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService *service.User
}

func NewUser(userService *service.User) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

const (
	ErrDuplicate           = "ERROR: duplicate key value violates unique constraint \"users_phone_number_key\" (SQLSTATE 23505)"
	ErrInternalServerError = "Internal Server Error"
)

func (u *UserHandler) CreateUser(ct *gin.Context) {

	if len(ct.PostForm("name")) == 0 || len(ct.PostForm("phone_number")) == 0 {
		ct.JSON(http.StatusBadRequest, gin.H{"error": "name and phone are required"})
		return
	}

	user := entity.User{
		Name:        ct.PostForm("name"),
		PhoneNumber: ct.PostForm("phone_number"),
	}

	err := u.userService.StoreUser(ct, user)
	if err != nil {
		if err.Error() == ErrDuplicate {
			ct.JSON(http.StatusBadRequest, gin.H{"error": "Phone number already exists"})
			return
		}
		ct.JSON(http.StatusInternalServerError, gin.H{"error": ErrInternalServerError})
		return
	}

	ct.IndentedJSON(http.StatusOK, user)
}

func (u *UserHandler) GenerateOTP(ct *gin.Context) {
	phoneNumber := ct.PostForm("phone_number")
	err := u.userService.GenerateOTP(ct, phoneNumber)
	if err != nil {
		if err.Error() == "no rows in result set" {
			ct.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}
		ct.JSON(http.StatusInternalServerError, gin.H{"error": ErrInternalServerError})
		return
	}
	ct.IndentedJSON(http.StatusOK, gin.H{"message": "OTP generated successfully"})
}

func (u *UserHandler) VerifyOTP(ct *gin.Context) {
	phoneNumber := ct.PostForm("phone_number")
	otp := ct.PostForm("otp")
	err := u.userService.VerifyOTP(ct, phoneNumber, otp)
	if err != nil {
		if err.Error() == "no rows in result set" {
			ct.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}
		if err.Error() == "otp expired" {
			ct.JSON(http.StatusBadRequest, gin.H{"error": "OTP expired"})
			return
		}
		if err.Error() == "otp not match" {
			ct.JSON(http.StatusBadRequest, gin.H{"error": "OTP not match"})
			return
		}
		ct.JSON(http.StatusInternalServerError, gin.H{"error": ErrInternalServerError})
		return
	}
	ct.IndentedJSON(http.StatusOK, gin.H{"message": "OTP verified successfully"})
}
