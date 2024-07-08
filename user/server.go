package user

import (
	"context"

	pb "github.com/kunal768/go-grpc-tc/proto"
)

type userServiceServer struct {
	pb.UnimplementedUserServiceServer
	service Service
}

func NewUserServiceServer(service Service) pb.UserServiceServer {
	return &userServiceServer{
		service: service,
	}
}

func (s *userServiceServer) GetUserByID(ctx context.Context, req *pb.UserIDRequest) (*pb.UserResponse, error) {
	return s.service.GetUserByID(ctx, req)
}

func (s *userServiceServer) GetUsersByIDs(ctx context.Context, req *pb.UserIDsRequest) (*pb.UsersResponse, error) {
	return s.service.GetUsersByIDs(ctx, req)
}

func (s *userServiceServer) SearchUsers(ctx context.Context, req *pb.SearchRequest) (*pb.UsersResponse, error) {
	return s.service.SearchUsers(ctx, req)
}

func (s *userServiceServer) AddUser(ctx context.Context, req *pb.User) (*pb.UserResponse, error) {
	return s.service.AddUser(ctx, req)
}

func (s *userServiceServer) ListUsers(ctx context.Context, req *pb.ListUsersRequest) (*pb.UsersResponse, error) {
	return s.service.ListUsers(ctx, req)
}
