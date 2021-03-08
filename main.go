package main

import (
	"answers/authentication"
	"answers/base"
	"answers/engine/middleware"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	if err := base.Init(); err != nil {
		panic(err)
	}
	router := mux.NewRouter()
	router.Use(middleware.Timer)
	router.HandleFunc("/v1/signup", authentication.EndpointSignup).Methods(http.MethodPost)
	http.ListenAndServe("127.0.0.1:8080", router)
}
