package server

import (
	"example.com/m/internal/genproto/user_pb/pb"
	"example.com/m/internal/user/adapter"
	"example.com/m/internal/user/app"
	gr "example.com/m/internal/user/delivery/grpc"
	"example.com/m/pkg/common"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

func RunGrpcUserServer() {
	db, err := common.ConnectToDb(
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_DATABASE"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
	)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	repo := adapter.NewUserRepo(db)
	usecase := app.NewUserUseCase(repo)
	userGrpc := gr.NewUserGrpcServer(usecase)

	listener, err := net.Listen("tcp", os.Getenv("USER_PORT"))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	fmt.Println("listening op port 8080")
	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, userGrpc)
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to server: %v", err)
	}
}
