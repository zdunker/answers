package authentication

import "golang.org/x/crypto/bcrypt"

func hashPassword(plainPassword string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(plainPassword), passwordHashCost)
	return string(hashedPassword), err
}
