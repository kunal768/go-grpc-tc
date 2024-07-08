package user

type UsersSearchRequest struct {
	ID          int
	FName       string
	City        string
	Phone       int64
	Height      float64
	Married     bool
	FindMarried bool
}

type UserResponse struct {
	User User
}

type UsersResponse struct {
	Users []User
}
