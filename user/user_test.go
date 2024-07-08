package user

import (
	"context"
	"testing"

	pb "github.com/kunal768/go-grpc-tc/proto"
	"github.com/kunal768/go-grpc-tc/utility"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestUserRepository_GetUserById(t *testing.T) {
	t.Run("User found", func(t *testing.T) {
		repo := NewRepository(UserDB{
			1: {ID: 1, FName: "John", City: "New York", Phone: 1234567890, Height: 180.5, Married: true},
		})
		user, err := repo.GetUserById(context.Background(), 1)
		assert.NoError(t, err)
		assert.Equal(t, User{ID: 1, FName: "John", City: "New York", Phone: 1234567890, Height: 180.5, Married: true}, user)
	})

	t.Run("User not found", func(t *testing.T) {
		repo := NewRepository(UserDB{})
		_, err := repo.GetUserById(context.Background(), 1)
		assert.ErrorIs(t, err, utility.ErrUserNotFound)
	})
}

func TestUserRepository_GetUsersById(t *testing.T) {
	t.Run("All users found", func(t *testing.T) {
		repo := NewRepository(UserDB{
			1: {ID: 1, FName: "John", City: "New York", Phone: 1234567890, Height: 180.5, Married: true},
			2: {ID: 2, FName: "Jane", City: "Los Angeles", Phone: 9876543210, Height: 165.2, Married: false},
		})
		users := repo.GetUsersById(context.Background(), []int{1, 2})
		assert.Len(t, users, 2)
		assert.Equal(t, User{ID: 1, FName: "John", City: "New York", Phone: 1234567890, Height: 180.5, Married: true}, users[0])
		assert.Equal(t, User{ID: 2, FName: "Jane", City: "Los Angeles", Phone: 9876543210, Height: 165.2, Married: false}, users[1])
	})

	t.Run("Some users not found", func(t *testing.T) {
		repo := NewRepository(UserDB{
			1: {ID: 1, FName: "John", City: "New York", Phone: 1234567890, Height: 180.5, Married: true},
		})
		users := repo.GetUsersById(context.Background(), []int{1, 2})
		assert.Len(t, users, 1)
		assert.Equal(t, User{ID: 1, FName: "John", City: "New York", Phone: 1234567890, Height: 180.5, Married: true}, users[0])
	})
}

func TestUserRepository_SearchUsers(t *testing.T) {
	repo := NewRepository(UserDB{
		1: {ID: 1, FName: "John", City: "New York", Phone: 1234567890, Height: 180.5, Married: true},
		2: {ID: 2, FName: "Jane", City: "Los Angeles", Phone: 9876543210, Height: 165.2, Married: false},
		3: {ID: 3, FName: "Bob", City: "Chicago", Phone: 5555555555, Height: 175.0, Married: true},
	})

	t.Run("Search by ID", func(t *testing.T) {
		users, _ := repo.SearchUsers(context.Background(), UsersSearchRequest{ID: 2})
		assert.Len(t, users, 1)
		assert.Equal(t, User{ID: 2, FName: "Jane", City: "Los Angeles", Phone: 9876543210, Height: 165.2, Married: false}, users[0])
	})

	t.Run("Search by first name", func(t *testing.T) {
		users, _ := repo.SearchUsers(context.Background(), UsersSearchRequest{FName: "John"})
		assert.Len(t, users, 1)
		assert.Equal(t, User{ID: 1, FName: "John", City: "New York", Phone: 1234567890, Height: 180.5, Married: true}, users[0])
	})

	t.Run("Search by city", func(t *testing.T) {
		users, _ := repo.SearchUsers(context.Background(), UsersSearchRequest{City: "Chicago"})
		assert.Len(t, users, 1)
		assert.Equal(t, User{ID: 3, FName: "Bob", City: "Chicago", Phone: 5555555555, Height: 175.0, Married: true}, users[0])
	})

	t.Run("Search by phone", func(t *testing.T) {
		users, _ := repo.SearchUsers(context.Background(), UsersSearchRequest{Phone: 9876543210})
		assert.Len(t, users, 1)
		assert.Equal(t, User{ID: 2, FName: "Jane", City: "Los Angeles", Phone: 9876543210, Height: 165.2, Married: false}, users[0])
	})

	t.Run("Search by married status", func(t *testing.T) {
		users, _ := repo.SearchUsers(context.Background(), UsersSearchRequest{Married: true, FindMarried: true})
		assert.Len(t, users, 2)
		assert.Equal(t, User{ID: 1, FName: "John", City: "New York", Phone: 1234567890, Height: 180.5, Married: true}, users[0])
		assert.Equal(t, User{ID: 3, FName: "Bob", City: "Chicago", Phone: 5555555555, Height: 175.0, Married: true}, users[1])
	})

	t.Run("Search by married status if married is false", func(t *testing.T) {
		users, _ := repo.SearchUsers(context.Background(), UsersSearchRequest{Married: false, FindMarried: true})
		assert.Len(t, users, 1)
		assert.Equal(t, User{ID: 2, FName: "Jane", City: "Los Angeles", Phone: 9876543210, Height: 165.2, Married: false}, users[0])
	})

	t.Run("Search by invalid request", func(t *testing.T) {
		_, err := repo.SearchUsers(context.Background(), UsersSearchRequest{})
		assert.ErrorIs(t, err, utility.ErrInvalidSearchRequest)
	})
}

func TestUserService(t *testing.T) {
	repo := NewRepository(UserDB{
		1: {ID: 1, FName: "John", City: "New York", Phone: 1234567890, Height: 180.5, Married: true},
		2: {ID: 2, FName: "Jane", City: "Los Angeles", Phone: 9876543210, Height: 165.2, Married: false},
	})
	service := NewService(repo)

	t.Run("GetUserByID", func(t *testing.T) {
		resp, err := service.GetUserByID(context.Background(), &pb.UserIDRequest{Id: 1})
		assert.NoError(t, err)
		assert.Equal(t, &pb.User{
			Id:      1,
			Fname:   "John",
			City:    "New York",
			Phone:   1234567890,
			Height:  180.5,
			Married: true,
		}, resp.User)
	})

	t.Run("GetUsersByIDs", func(t *testing.T) {
		resp, err := service.GetUsersByIDs(context.Background(), &pb.UserIDsRequest{Ids: []int32{1, 2}})
		assert.NoError(t, err)
		assert.Len(t, resp.Users, 2)
		assert.Equal(t, &pb.User{
			Id:      1,
			Fname:   "John",
			City:    "New York",
			Phone:   1234567890,
			Height:  180.5,
			Married: true,
		}, resp.Users[0])
		assert.Equal(t, &pb.User{
			Id:      2,
			Fname:   "Jane",
			City:    "Los Angeles",
			Phone:   9876543210,
			Height:  165.2,
			Married: false,
		}, resp.Users[1])
	})

	t.Run("SearchUsers", func(t *testing.T) {
		resp, err := service.SearchUsers(context.Background(), &pb.SearchRequest{
			Id:      2,
			Fname:   "Jane",
			City:    "Los Angeles",
			Phone:   9876543210,
			Married: false,
		})
		assert.NoError(t, err)
		assert.Len(t, resp.Users, 1)
		assert.Equal(t, &pb.User{
			Id:      2,
			Fname:   "Jane",
			City:    "Los Angeles",
			Phone:   9876543210,
			Height:  165.2,
			Married: false,
		}, resp.Users[0])
	})
}

func TestUserRepository_AddUser(t *testing.T) {
	t.Run("Add new user", func(t *testing.T) {
		repo := NewRepository(UserDB{})
		user := User{ID: 1, FName: "John", City: "New York", Phone: 1234567890, Height: 180.5, Married: true}
		savedUser, err := repo.AddUser(context.Background(), user)
		assert.NoError(t, err)
		assert.Equal(t, user, savedUser)
	})

	t.Run("Add user with existing ID", func(t *testing.T) {
		repo := NewRepository(UserDB{
			1: {ID: 1, FName: "John", City: "New York", Phone: 1234567890, Height: 180.5, Married: true},
		})
		user := User{ID: 1, FName: "Jane", City: "Los Angeles", Phone: 9876543210, Height: 165.2, Married: false}
		_, err := repo.AddUser(context.Background(), user)
		assert.ErrorIs(t, err, utility.ErrUserIdAlreadyExists)
	})

	t.Run("Add new user with valid data", func(t *testing.T) {
		repo := NewRepository(UserDB{})
		user := User{ID: 1, FName: "John", City: "New York", Phone: 1234567890, Height: 180.5, Married: true}
		savedUser, err := repo.AddUser(context.Background(), user)
		assert.NoError(t, err)
		assert.Equal(t, user, savedUser)
	})

	t.Run("Add user with invalid ID", func(t *testing.T) {
		repo := NewRepository(UserDB{})
		user := User{ID: 0, FName: "John", City: "New York", Phone: 1234567890, Height: 180.5, Married: true}
		_, err := repo.AddUser(context.Background(), user)
		assert.ErrorIs(t, err, utility.ErrInvalidIdInput)
	})

	t.Run("Add user with empty city", func(t *testing.T) {
		repo := NewRepository(UserDB{})
		user := User{ID: 1, FName: "John", City: "", Phone: 1234567890, Height: 180.5, Married: true}
		_, err := repo.AddUser(context.Background(), user)
		assert.ErrorIs(t, err, utility.ErrInvalidCityInput)
	})

	t.Run("Add user with empty first name", func(t *testing.T) {
		repo := NewRepository(UserDB{})
		user := User{ID: 1, FName: "", City: "New York", Phone: 1234567890, Height: 180.5, Married: true}
		_, err := repo.AddUser(context.Background(), user)
		assert.ErrorIs(t, err, utility.ErrInvalidFNameInput)
	})

	t.Run("Add user with invalid height", func(t *testing.T) {
		repo := NewRepository(UserDB{})
		user := User{ID: 1, FName: "John", City: "New York", Phone: 1234567890, Height: 0, Married: true}
		_, err := repo.AddUser(context.Background(), user)
		assert.ErrorIs(t, err, utility.ErrInvalidHeightInput)
	})

	t.Run("Add user with invalid phone", func(t *testing.T) {
		repo := NewRepository(UserDB{})
		user := User{ID: 1, FName: "John", City: "New York", Phone: 0, Height: 180.5, Married: true}
		_, err := repo.AddUser(context.Background(), user)
		assert.ErrorIs(t, err, utility.ErrInvalidPhoneInput)
	})

	t.Run("Add second user with existing ID", func(t *testing.T) {
		repo := NewRepository(UserDB{
			1: {ID: 1, FName: "John", City: "New York", Phone: 1234567890, Height: 180.5, Married: true},
		})
		user := User{ID: 1, FName: "Jane", City: "Los Angeles", Phone: 9876543210, Height: 165.2, Married: false}
		_, err := repo.AddUser(context.Background(), user)
		assert.ErrorIs(t, err, utility.ErrUserIdAlreadyExists)
	})
}

func TestUserRepository_ListUsers(t *testing.T) {
	t.Run("List all users", func(t *testing.T) {
		repo := NewRepository(UserDB{
			1: {ID: 1, FName: "John", City: "New York", Phone: 1234567890, Height: 180.5, Married: true},
			2: {ID: 2, FName: "Jane", City: "Los Angeles", Phone: 9876543210, Height: 165.2, Married: false},
			3: {ID: 3, FName: "Bob", City: "Chicago", Phone: 5555555555, Height: 175.0, Married: true},
		})
		users := repo.ListUsers(context.Background(), 10, 1)
		assert.Len(t, users, 3)
		assert.Equal(t, User{ID: 1, FName: "John", City: "New York", Phone: 1234567890, Height: 180.5, Married: true}, users[0])
		assert.Equal(t, User{ID: 2, FName: "Jane", City: "Los Angeles", Phone: 9876543210, Height: 165.2, Married: false}, users[1])
		assert.Equal(t, User{ID: 3, FName: "Bob", City: "Chicago", Phone: 5555555555, Height: 175.0, Married: true}, users[2])
	})

	t.Run("List users with pagination", func(t *testing.T) {
		repo := NewRepository(UserDB{
			1: {ID: 1, FName: "John", City: "New York", Phone: 1234567890, Height: 180.5, Married: true},
			2: {ID: 2, FName: "Jane", City: "Los Angeles", Phone: 9876543210, Height: 165.2, Married: false},
			3: {ID: 3, FName: "Bob", City: "Chicago", Phone: 5555555555, Height: 175.0, Married: true},
		})
		users := repo.ListUsers(context.Background(), 2, 1)
		assert.Len(t, users, 1)
		assert.Equal(t, User{ID: 3, FName: "Bob", City: "Chicago", Phone: 5555555555, Height: 175.0, Married: true}, users[0])
	})
}

func TestUserService_AddUser(t *testing.T) {
	repo := NewRepository(UserDB{})
	service := NewService(repo)

	t.Run("Add new user", func(t *testing.T) {
		resp, err := service.AddUser(context.Background(), &pb.User{
			Id:      1,
			Fname:   "John",
			City:    "New York",
			Phone:   1234567890,
			Height:  180.5,
			Married: true,
		})
		assert.NoError(t, err)
		assert.Equal(t, &pb.User{
			Id:      1,
			Fname:   "John",
			City:    "New York",
			Phone:   1234567890,
			Height:  180.5,
			Married: true,
		}, resp.User)
	})

	t.Run("Add user with existing ID", func(t *testing.T) {
		_, err := service.AddUser(context.Background(), &pb.User{
			Id:      1,
			Fname:   "Jane",
			City:    "Los Angeles",
			Phone:   9876543210,
			Height:  165.2,
			Married: false,
		})
		assert.ErrorIs(t, err, status.Errorf(codes.AlreadyExists, utility.ErrUserIdAlreadyExists.Error()))
	})

	t.Run("Add new user with valid data", func(t *testing.T) {
		resp, err := service.AddUser(context.Background(), &pb.User{
			Id:      5,
			Fname:   "Apple",
			City:    "New York",
			Phone:   4353234562,
			Height:  180.5,
			Married: true,
		})
		assert.NoError(t, err)
		assert.Equal(t, &pb.User{
			Id:      5,
			Fname:   "Apple",
			City:    "New York",
			Phone:   4353234562,
			Height:  180.5,
			Married: true,
		}, resp.User)
	})

	t.Run("Add user with invalid ID", func(t *testing.T) {
		_, err := service.AddUser(context.Background(), &pb.User{
			Id:      0,
			Fname:   "John",
			City:    "New York",
			Phone:   1234567890,
			Height:  180.5,
			Married: true,
		})
		assert.ErrorIs(t, err, status.Errorf(codes.InvalidArgument, utility.ErrInvalidIdInput.Error()))
	})

	t.Run("Add user with empty city", func(t *testing.T) {
		_, err := service.AddUser(context.Background(), &pb.User{
			Id:      1,
			Fname:   "Maxx",
			City:    "",
			Phone:   1234567890,
			Height:  180.5,
			Married: true,
		})
		assert.ErrorIs(t, err, status.Errorf(codes.InvalidArgument, utility.ErrInvalidCityInput.Error()))
	})

	t.Run("Add user with empty first name", func(t *testing.T) {
		_, err := service.AddUser(context.Background(), &pb.User{
			Id:      1,
			Fname:   "",
			City:    "New York",
			Phone:   1234567890,
			Height:  180.5,
			Married: true,
		})
		assert.ErrorIs(t, err, status.Errorf(codes.InvalidArgument, utility.ErrInvalidFNameInput.Error()))
	})

	t.Run("Add user with invalid height", func(t *testing.T) {
		_, err := service.AddUser(context.Background(), &pb.User{
			Id:      1,
			Fname:   "John",
			City:    "New York",
			Phone:   1234567890,
			Height:  0,
			Married: true,
		})
		assert.ErrorIs(t, err, status.Errorf(codes.InvalidArgument, utility.ErrInvalidHeightInput.Error()))
	})

	t.Run("Add user with invalid phone", func(t *testing.T) {
		_, err := service.AddUser(context.Background(), &pb.User{
			Id:      1,
			Fname:   "John",
			City:    "New York",
			Phone:   0,
			Height:  180.5,
			Married: true,
		})
		assert.ErrorIs(t, err, status.Errorf(codes.InvalidArgument, utility.ErrInvalidPhoneInput.Error()))
	})

	t.Run("Add second user with existing ID", func(t *testing.T) {
		_, err := service.AddUser(context.Background(), &pb.User{
			Id:      1,
			Fname:   "Jane",
			City:    "Los Angeles",
			Phone:   9876543210,
			Height:  165.2,
			Married: false,
		})
		assert.ErrorIs(t, err, status.Errorf(codes.AlreadyExists, utility.ErrUserIdAlreadyExists.Error()))
	})

}

func TestUserService_ListUsers(t *testing.T) {
	repo := NewRepository(UserDB{
		1: {ID: 1, FName: "John", City: "New York", Phone: 1234567890, Height: 180.5, Married: true},
		2: {ID: 2, FName: "Jane", City: "Los Angeles", Phone: 9876543210, Height: 165.2, Married: false},
		3: {ID: 3, FName: "Bob", City: "Chicago", Phone: 5555555555, Height: 175.0, Married: true},
	})
	service := NewService(repo)

	t.Run("List all users", func(t *testing.T) {
		resp, err := service.ListUsers(context.Background(), &pb.ListUsersRequest{
			Page:     0,
			PageSize: 10,
		})
		assert.NoError(t, err)
		assert.Len(t, resp.Users, 3)
		assert.Equal(t, &pb.User{
			Id:      1,
			Fname:   "John",
			City:    "New York",
			Phone:   1234567890,
			Height:  180.5,
			Married: true,
		}, resp.Users[0])
		assert.Equal(t, &pb.User{
			Id:      2,
			Fname:   "Jane",
			City:    "Los Angeles",
			Phone:   9876543210,
			Height:  165.2,
			Married: false,
		}, resp.Users[1])
		assert.Equal(t, &pb.User{
			Id:      3,
			Fname:   "Bob",
			City:    "Chicago",
			Phone:   5555555555,
			Height:  175.0,
			Married: true,
		}, resp.Users[2])
	})

	t.Run("List users with pagination", func(t *testing.T) {
		resp, err := service.ListUsers(context.Background(), &pb.ListUsersRequest{
			Page:     1,
			PageSize: 2,
		})
		assert.NoError(t, err)
		assert.Len(t, resp.Users, 1)
		assert.Equal(t, &pb.User{
			Id:      3,
			Fname:   "Bob",
			City:    "Chicago",
			Phone:   5555555555,
			Height:  175.0,
			Married: true,
		}, resp.Users[0])
	})
}
