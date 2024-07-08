package db

import "github.com/kunal768/go-grpc-tc/user"

func InitDb() user.UserDB {
	db := user.UserDB{
		1: {ID: 1, FName: "John", City: "New York", Phone: 1234567890, Height: 180.5, Married: true},
		2: {ID: 2, FName: "Jane", City: "Los Angeles", Phone: 9876543210, Height: 165.2, Married: false},
	}
	return db
}
