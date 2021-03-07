package authentication

import (
	"time"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
)

// AccountModel is the only interface which represents the account model layer,
// it should provide funcs which communicate with controller layer, controller layer
// shouldn't worry about while switch db instances, when implementing new db handlers
// on account model, those handlers should provide the funcs in AccountModel interface
type AccountModel interface {
	UserNameOccupied(userName string) (bool, error)
	EmailExists(emailAddress string) (bool, error)
	PhoneNumberExists(phoneNumber string) (bool, error)
	Signup(account Account) error
}

// PostgresAccountHandler name is self-explained
type PostgresAccountHandler struct {
	db orm.DB
}

func getPostgresAccountHandler(db *pg.DB) AccountModel {
	return PostgresAccountHandler{db}
}

// UserNameOccupied -- check if username is occupied using postgresDB handler
func (postgresDB PostgresAccountHandler) UserNameOccupied(userName string) (bool, error) {
	return postgresDB.db.Model((*Account)(nil)).Where(fieldUserName+" = ?", userName).Exists()
}

// EmailExists -- check if email is in use using postgresDB handler
func (postgresDB PostgresAccountHandler) EmailExists(emailAddress string) (bool, error) {
	return postgresDB.db.Model((*Account)(nil)).Where(fieldEmailAddress+" = ?", emailAddress).Exists()
}

// PhoneNumberExists -- check if phone number is in use using postgresDB handler
func (postgresDB PostgresAccountHandler) PhoneNumberExists(phoneNumber string) (bool, error) {
	return postgresDB.db.Model((*Account)(nil)).Where(fieldPhoneNumber+" = ?", phoneNumber).Exists()
}

// Signup -- signup new user with input account
func (postgresDB PostgresAccountHandler) Signup(account Account) error {
	if err := validateSignupForm(account); err != nil {
		return err
	}
	_, err := postgresDB.db.Model(&account).Insert()
	return err
}

// Account -- abstract data model of account, currently using postgres tags,
// hoping if switching to other db, this struct is still flexible enough to function with just modify tags
type Account struct {
	tableName struct{} `pg:"user_auth"`

	UserName     string    `pg:"user_name"`
	UserID       string    `pg:"user_id"`
	Password     string    `pg:"password"`
	EmailAddress string    `pg:"email_address"`
	MemberSince  time.Time `pg:"member_since"`
	LastLogin    time.Time `pg:"last_login"`
}
