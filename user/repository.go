package user

import (
	"context"
	"sort"

	"github.com/kunal768/go-grpc-tc/utility"
)

type Repository interface {
	AddUser(ctx context.Context, user User) (User, error)
	GetUserById(ctx context.Context, Id int) (User, error)
	GetUsersById(ctx context.Context, Ids []int) []User
	SearchUsers(ctx context.Context, data UsersSearchRequest) ([]User, error)
	ListUsers(ctx context.Context, pageSize int, page int) []User
}

type repo struct {
	db UserDB
}

func NewRepository(db UserDB) Repository {
	return &repo{
		db: db,
	}
}

func (r repo) AddUser(ctx context.Context, user User) (User, error) {
	if user.ID == 0 {
		return User{}, utility.ErrInvalidIdInput
	}

	if user.City == "" {
		return User{}, utility.ErrInvalidCityInput
	}

	if user.FName == "" {
		return User{}, utility.ErrInvalidFNameInput
	}

	if int(user.Height) == 0 {
		return User{}, utility.ErrInvalidHeightInput
	}

	if user.Phone == 0 {
		return User{}, utility.ErrInvalidPhoneInput
	}

	if _, exists := r.db[user.ID]; exists {
		return User{}, utility.ErrUserIdAlreadyExists
	}

	r.db[user.ID] = user
	return user, nil
}

func (r repo) GetUserById(ctx context.Context, Id int) (User, error) {
	if Id == 0 {
		return User{}, utility.ErrInvalidIdInput
	}
	user, found := r.db[UserId(Id)]
	if !found {
		return User{}, utility.ErrUserNotFound
	}
	return user, nil
}

func (r repo) GetUsersById(ctx context.Context, Ids []int) []User {
	ans := []User{}
	for _, id := range Ids {
		user, err := r.GetUserById(ctx, id)
		if err == nil {
			ans = append(ans, user)
		}
	}
	return ans
}

func (r repo) SearchUsers(ctx context.Context, data UsersSearchRequest) ([]User, error) {
	if data.FName == "" && data.City == "" && data.Phone == 0 && data.Height == 0 && !data.Married && data.ID == 0 && !data.FindMarried {
		return nil, utility.ErrInvalidSearchRequest
	}

	ans := []User{}
	for _, user := range r.db {
		match := true
		if data.ID != 0 && user.ID != UserId(data.ID) {
			match = false
		}

		if data.FName != "" && user.FName != data.FName {
			match = false
		}

		if data.City != "" && user.City != data.City {
			match = false
		}

		if data.Phone != 0 && user.Phone != data.Phone {
			match = false
		}

		if data.FindMarried && data.Married != user.Married {
			match = false
		}

		if match {
			ans = append(ans, user)
		}

	}

	return ans, nil
}

func (r repo) ListUsers(ctx context.Context, pageSize int, page int) []User {
	if pageSize <= 0 {
		pageSize = len(r.db)
	}

	users := make([]User, 0, len(r.db))
	for _, user := range r.db {
		users = append(users, user)
	}

	// sort according to user IDs
	sort.Slice(users, func(i, j int) bool {
		return users[i].ID < users[j].ID
	})

	start := page * pageSize
	end := start + pageSize

	if start >= len(users) {
		return users
	}

	if end > len(users) {
		end = len(users)
	}

	return users[start:end]
}
