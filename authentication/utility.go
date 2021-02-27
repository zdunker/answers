package authentication

import "golang.org/x/crypto/bcrypt"

func validateSignupForm(account Account) error {
	if account.UserName == "" {
		return errUserNameEmpty
	}
	if account.EmailAddress == "" {
		return errEmailEmpty
	}
	if account.Password == "" {
		return errPasswordEmpty
	}
	return nil
}

func hashPassword(plainPassword string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(plainPassword), passwordHashCost)
	return string(hashedPassword), err
}
