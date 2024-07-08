package main

import (
	"log"
	"net"

	"github.com/kunal768/go-grpc-tc/db"
	pb "github.com/kunal768/go-grpc-tc/proto"
	"github.com/kunal768/go-grpc-tc/user"
	"google.golang.org/grpc"
)

func main() {

	repo := user.NewRepository(db.InitDb())

	service := user.NewService(repo)

	server := grpc.NewServer()
	pb.RegisterUserServiceServer(server, user.NewUserServiceServer(service))

	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("server listening at %v", lis.Addr())
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
