package main

import (
	"aqary/handler"
	"aqary/repository/postgres"
	"aqary/router"
	"aqary/service"
	"context"
	"fmt"
	"log"
)

func main() {

	//read from .env
	postgresURL := "postgresql://postgres:postgres@localhost:5432/user_db?sslmode=disable"

	ctx := context.Background()
	postgresInstance, err := postgres.NewClient(ctx, postgresURL)

	userService := service.NewUser(postgresInstance)
	userHandler := handler.NewUser(userService)

	routerInstance, err := router.NewRouter(userHandler)
	if err != nil {
		log.Fatal(err)
	}

	//read from .env file
	fmt.Println("Starting ...")
	routerInstance.RunRouter("localhost:8080")

}
