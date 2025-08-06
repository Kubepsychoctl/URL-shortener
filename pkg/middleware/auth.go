package middleware

import (
	"app/url-shorter/configs"
	"app/url-shorter/pkg/jwt"
	"fmt"
	"net/http"
	"strings"
)

func IsAuthed(next http.Handler, config *configs.Config) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authedHeader := r.Header.Get("Authorization")
		token := strings.TrimPrefix(authedHeader, "Bearer ")
		isValid, data := jwt.NewJwt(config.Auth.Secret).Parse(token)
		fmt.Println(isValid)
		fmt.Println(data)
		next.ServeHTTP(w, r)
	})
}
