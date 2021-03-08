package middleware

import (
	"answers/base"
	"fmt"
	"net/http"
)

func Timer(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		now := base.GetServerTime()
		fmt.Printf("request start time: %v \n", now)
		defer fmt.Printf("request finish time: %v, time cost: %v \n", base.GetServerTime(), base.GetServerTime().Sub(now))

		next.ServeHTTP(w, r)
	})
}
