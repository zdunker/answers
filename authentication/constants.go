package authentication

import "errors"

var (
	errUserNameEmpty = errors.New("user name is empty")
	errPasswordEmpty = errors.New("password is empty")
	errEmailEmpty    = errors.New("email address is empty")
)

const (
	fieldUserName     string = "user_name"
	fieldPassword     string = "password"
	fieldEmailAddress string = "email_address"

	passwordHashCost int = 8
)
