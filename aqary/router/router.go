package router

import (
	"aqary/handler"

	"github.com/gin-gonic/gin"
)

type Router struct {
	router *gin.Engine
}

func NewRouter(
	userHandler *handler.UserHandler,
) (*Router, error) {

	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	apiGroup := router.Group("/api")
	{
		apiGroup.POST("/users", userHandler.CreateUser)
		apiGroup.POST("/users/generateotp", userHandler.GenerateOTP)
		apiGroup.POST("/users/verifyotp", userHandler.VerifyOTP)
	}

	return &Router{
		router: router,
	}, nil
}

func (r *Router) RunRouter(address string) {
	r.router.Run(address)
}
