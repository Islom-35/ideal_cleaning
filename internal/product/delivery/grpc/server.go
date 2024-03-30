package grpc

import (
	"example.com/m/internal/genproto/product/pb"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

type Server struct {
	grpcSrv       *grpc.Server
	productServer pb.ProductServiceServer
}

func NewServer(productServer pb.ProductServiceServer) Server{
	return Server{
		grpcSrv:       grpc.NewServer(),
		productServer: productServer,
	}
}

func(s *Server) ListenAndServe(port string)error{
	addr :=fmt.Sprintf("%v",port)

	lis, err :=net.Listen("tcp",addr)
	if err!=nil{
		return err
	}
	pb.RegisterProductServiceServer(s.grpcSrv,s.productServer)
	log.Println("listen on port: ",addr)

	err = s.grpcSrv.Serve(lis) 
	if err!=nil{
		return err
	}
	return nil
}
