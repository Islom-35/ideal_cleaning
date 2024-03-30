package main

import (
	grpc_client "example.com/m/internal/api-gateway/delivery/grpc"
	"example.com/m/internal/api-gateway/delivery/rest"
	"fmt"
	"log"
	"net/http"
	"os"
)

// @title Clean-sweep-solution_App 
// @version 1.0
// @description API Server for Clean-sweep-solution Application

// @host localhost:8000
// @BasePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {
	auditClient, err := grpc_client.NewClient(os.Getenv("PRODUCT_PORT"))
	if err != nil {
		log.Println(err, "port")
	}
	defer auditClient.CloseConnection()

	handlers := rest.NewHandler(*auditClient)

	srv := &http.Server{
		Addr:    fmt.Sprintf("%v", os.Getenv("HTTP_PORT")),
		Handler: handlers.InitRouters(),
	}

	log.Println("SERVER STARTED")

	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
