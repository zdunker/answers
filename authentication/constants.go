package authentication

import "errors"

var (
	errUserNameEmpty error = errors.New("user name is empty")
	errPasswordEmpty error = errors.New("password is empty")
	errEmailEmpty    error = errors.New("email address is empty")
)

const (
	fieldUserName     string = "user_name"
	fieldPassword     string = "password"
	fieldEmailAddress string = "email_address"

	errMsgUserNameOccupied string = "username is occupied"

	passwordHashCost int = 8
)
