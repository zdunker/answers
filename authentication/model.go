package authentication

import (
	"time"

	"github.com/go-pg/pg/v10/orm"
)

// type Account interface {
// 	signup() error
// }

type Account struct {
	tableName struct{} `pg:"user_auth"`

	UserName     string
	UserID       string
	Password     string
	EmailAddress string
	MemberSince  time.Time
	LastLogin    time.Time
}

func (account Account) signup(db orm.DB) error {
	if err := validateSignupForm(account); err != nil {
		return err
	}
	_, err := db.Model(&account).Insert()
	return err
}
