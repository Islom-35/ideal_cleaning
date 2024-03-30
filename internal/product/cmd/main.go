package main

import (
	"example.com/m/internal/product/adapter"
	"example.com/m/internal/product/app"
	"example.com/m/pkg/common"
	server "example.com/m/internal/product/delivery/grpc"
	"log"
	"os"
)

func main() {
	db, err := common.ConnectToDb(
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_DATABASE"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
	)
	if err != nil {
		log.Printf("Failed to connect to database: %v", err)
	}

	defer db.Close()

	productRepo := adapter.NewProductRepo(db)
	productUseCase := app.NewProductUseCase(productRepo)
	productServer := server.NewProductServer(productUseCase)

	srv :=server.NewServer(&productServer)
	if err := srv.ListenAndServe(os.Getenv("PRODUCT_PORT")); err != nil {
		log.Println(err)
	}
	log.Println(os.Getenv("PRODUCT_PORT"))

}
