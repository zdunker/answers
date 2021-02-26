package authentication

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
