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
	fieldPhoneNumber  string = "phone_number"

	errMsgUserNameOccupied   string = "username is occupied"
	errMsgEmailAlreadyExists string = "email already in use"
	errMsgPhoneNumberExists  string = "phone number already in use"

	passwordHashCost int = 8
)
