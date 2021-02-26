package authentication

import "errors"

var (
	errUserNameEmpty = errors.New("user name is empty")
	errPasswordEmpty = errors.New("password is empty")
	errEmailEmpty    = errors.New("email address is empty")
)

const (
	fieldUserName     = "user_name"
	fieldPassword     = "password"
	fieldEmailAddress = "email_address"
)
