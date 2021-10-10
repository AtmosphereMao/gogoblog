package middlewares

import (
	"myblog/lib/session"
	"net/http"
)

func StartSession(next http.Handler) http.Handler{
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request){
			session.StartSession(w, r)
			next.ServeHTTP(w, r)
	})
}