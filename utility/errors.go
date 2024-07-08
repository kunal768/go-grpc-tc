package utility

import "errors"

var (
	ErrUserNotFound         = errors.New("user not found in db")
	ErrUserIdAlreadyExists  = errors.New("user with this Id already exists")
	ErrInvalidSearchRequest = errors.New("invalid search request")
	ErrInvalidHeightInput   = errors.New("invalid height input")
	ErrInvalidFNameInput    = errors.New("invalid first name input")
	ErrInvalidCityInput     = errors.New("invalid city input")
	ErrInvalidPhoneInput    = errors.New("invalid phone number input")
	ErrInvalidIdInput       = errors.New("invalid user ID input")
)
