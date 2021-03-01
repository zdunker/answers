package authentication

import (
	"fmt"
	"net/http"
)

// EndpointSignup is the endpoint providing registration
func EndpointSignup(w http.ResponseWriter, r *http.Request) {
	service := getAuthenticationService()
	if occupied, err := service.isUserNameOccupied(r.FormValue(fieldUserName)); err != nil {
		w.Write([]byte(err.Error()))
		return
	} else if occupied {
		w.Write([]byte(errMsgUserNameOccupied))
		return
	}
	account, err := newAccount(r.FormValue(fieldUserName), r.FormValue(fieldPassword), r.FormValue(fieldEmailAddress))
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	var resp string
	if err := service.signup(account); err != nil {
		resp = fmt.Sprintf("Failed to signup: %s \n", err.Error())
	} else {
		resp = "Signup successful"
	}
	w.Write([]byte(resp))
}
