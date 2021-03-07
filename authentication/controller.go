package authentication

import (
	"answers/base"
	"errors"
	"fmt"

	"github.com/go-pg/pg/v10"
)

type authenticationService struct {
	db *pg.DB
}

func getAuthenticationService() authenticationService {
	return authenticationService{
		db: base.GetPostgresDB(),
	}
}

func newAccount(userName, password, emailAddress string) (account Account, err error) {
	hashedPassword, err := hashPassword(password)
	if err != nil {
		return
	}
	account = Account{
		UserName:     userName,
		UserID:       base.GetIDGenerator().Generate().Base58(),
		Password:     hashedPassword,
		EmailAddress: emailAddress,
		MemberSince:  base.GetServerTime(),
		LastLogin:    base.GetServerTime(),
	}
	return
}

func (service authenticationService) isUserNameOccupied(userName string) (bool, error) {
	dbHandler := getPostgresAccountHandler(service.db)
	return dbHandler.UserNameOccupied(userName)
}

func (service authenticationService) signup(userName, password, emailAddress string) ([]byte, error) {
	account, err := newAccount(userName, password, emailAddress)
	if err != nil {
		return nil, err
	}
	if err := service.validateSignupForm(account); err != nil {
		return nil, err
	}

	if createAccountErr := service.createAccount(account); createAccountErr != nil {
		err = fmt.Errorf("failed to signup: %w", createAccountErr)
		return nil, err
	}
	return []byte("Signup successful"), nil
}

func (service authenticationService) createAccount(account Account) error {
	dbHandler := getPostgresAccountHandler(service.db)
	return dbHandler.Signup(account)
}

func (service authenticationService) validateSignupForm(account Account) error {
	if err := validateSignupForm(account); err != nil {
		return err
	}

	dbHandler := getPostgresAccountHandler(service.db)
	if occupied, err := dbHandler.UserNameOccupied(account.UserName); err != nil {
		return err
	} else if occupied {
		return errors.New(errMsgUserNameOccupied)
	}

	if exists, err := dbHandler.EmailExists(account.EmailAddress); err != nil {
		return err
	} else if exists {
		return errors.New(errMsgEmailAlreadyExists)
	}
	return nil
}
