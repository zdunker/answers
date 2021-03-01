package authentication

import (
	"answers/base"

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
	return isUserNameOccupied(service.db, userName)
}

func (service authenticationService) signup(account Account) error {
	return account.signup(service.db)
}
