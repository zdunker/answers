package authentication

import (
	"answers/base"
	"fmt"
	"net/http"
	"time"

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

func EndpointSignup(w http.ResponseWriter, r *http.Request) {
	now := time.Now()
	account := Account{
		UserName:     r.FormValue(fieldUserName),
		UserID:       base.GetIDGenerator().Generate().Base58(),
		Password:     r.FormValue(fieldPassword),
		EmailAddress: r.FormValue(fieldEmailAddress),
		MemberSince:  now,
		LastLogin:    now,
	}
	service := getAuthenticationService()
	var resp string
	if err := account.signup(service.db); err != nil {
		resp = fmt.Sprintf("Failed to signup: %s \n", err.Error())
	} else {
		resp = "Signup successful"
	}
	w.Write([]byte(resp))
}
