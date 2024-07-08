package user

import (
	"context"

	pb "github.com/kunal768/go-grpc-tc/proto"
	"github.com/kunal768/go-grpc-tc/utility"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Service interface {
	AddUser(ctx context.Context, req *pb.User) (*pb.UserResponse, error)
	GetUserByID(ctx context.Context, req *pb.UserIDRequest) (*pb.UserResponse, error)
	GetUsersByIDs(ctx context.Context, req *pb.UserIDsRequest) (*pb.UsersResponse, error)
	SearchUsers(ctx context.Context, req *pb.SearchRequest) (*pb.UsersResponse, error)
	ListUsers(ctx context.Context, req *pb.ListUsersRequest) (*pb.UsersResponse, error)
}

type svc struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &svc{
		repo: repo,
	}
}

func (s svc) AddUser(ctx context.Context, req *pb.User) (*pb.UserResponse, error) {
	user, err := s.repo.AddUser(ctx, User{
		ID:      UserId(req.Id),
		FName:   req.Fname,
		City:    req.City,
		Phone:   req.Phone,
		Height:  req.Height,
		Married: req.Married,
	})

	if err != nil {
		if err == utility.ErrUserIdAlreadyExists {
			return nil, status.Errorf(codes.AlreadyExists, err.Error())
		}
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	return &pb.UserResponse{User: &pb.User{
		Id:      int32(user.ID),
		Fname:   user.FName,
		City:    user.City,
		Phone:   user.Phone,
		Height:  user.Height,
		Married: user.Married,
	}}, nil
}

func (s svc) GetUserByID(ctx context.Context, req *pb.UserIDRequest) (*pb.UserResponse, error) {
	user, err := s.repo.GetUserById(ctx, int(req.Id))
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}
	return &pb.UserResponse{User: &pb.User{
		Id:      int32(user.ID),
		Fname:   user.FName,
		City:    user.City,
		Phone:   user.Phone,
		Height:  user.Height,
		Married: user.Married,
	}}, nil
}

func (s svc) GetUsersByIDs(ctx context.Context, req *pb.UserIDsRequest) (*pb.UsersResponse, error) {
	users := s.repo.GetUsersById(ctx, convertToIntSlice(req.Ids))
	var pbUsers []*pb.User
	for _, user := range users {
		pbUsers = append(pbUsers, &pb.User{
			Id:      int32(user.ID),
			Fname:   user.FName,
			City:    user.City,
			Phone:   user.Phone,
			Height:  user.Height,
			Married: user.Married,
		})
	}
	return &pb.UsersResponse{Users: pbUsers}, nil
}

func (s svc) SearchUsers(ctx context.Context, req *pb.SearchRequest) (*pb.UsersResponse, error) {
	users, err := s.repo.SearchUsers(ctx, UsersSearchRequest{
		ID:          int(req.Id),
		FName:       req.Fname,
		City:        req.City,
		Phone:       req.Phone,
		Height:      req.Height,
		Married:     req.Married,
		FindMarried: req.Searchmarried,
	})

	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	var pbUsers []*pb.User
	for _, user := range users {
		pbUsers = append(pbUsers, &pb.User{
			Id:      int32(user.ID),
			Fname:   user.FName,
			City:    user.City,
			Phone:   user.Phone,
			Height:  user.Height,
			Married: user.Married,
		})
	}
	return &pb.UsersResponse{Users: pbUsers}, nil
}

func (s svc) ListUsers(ctx context.Context, req *pb.ListUsersRequest) (*pb.UsersResponse, error) {
	users := s.repo.ListUsers(ctx, int(req.PageSize), int(req.Page))

	var pbUsers []*pb.User
	for _, user := range users {
		pbUsers = append(pbUsers, &pb.User{
			Id:      int32(user.ID),
			Fname:   user.FName,
			City:    user.City,
			Phone:   user.Phone,
			Height:  user.Height,
			Married: user.Married,
		})
	}

	return &pb.UsersResponse{Users: pbUsers}, nil
}

func convertToIntSlice(ids []int32) []int {
	var intIds []int
	for _, id := range ids {
		intIds = append(intIds, int(id))
	}
	return intIds
}
