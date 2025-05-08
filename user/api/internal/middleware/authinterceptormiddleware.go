package middleware

import "net/http"

type AuthInterceptorMiddleware struct {
}

func NewAuthInterceptorMiddleware() *AuthInterceptorMiddleware {
	return &AuthInterceptorMiddleware{}
}

func (m *AuthInterceptorMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Authorization") != "123456" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		next(w, r)
	}
}
