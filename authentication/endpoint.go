package authentication

import (
	"net/http"
)

// EndpointSignup is the endpoint providing registration
func EndpointSignup(w http.ResponseWriter, r *http.Request) {
	service := getAuthenticationService()
	resp, err := service.signup(r.FormValue(fieldUserName), r.FormValue(fieldPassword), r.FormValue(fieldEmailAddress))
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	w.Write([]byte(resp))
}
