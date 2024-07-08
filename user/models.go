package user

type UserId int

type User struct {
	ID      UserId
	FName   string
	City    string
	Phone   int64
	Height  float64
	Married bool
}

type UserDB map[UserId]User
