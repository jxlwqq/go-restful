package auth

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"strings"
)

type Middleware interface {
	Handler(next http.Handler) http.Handler
	verifyToken(string) (jwt.Claims, error)
}

type middleware struct {
	signingKey string
}

func NewMiddleware(signingKey string) Middleware {
	return middleware{signingKey}
}

func (m middleware) Handler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if len(token) == 0 {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintf(w, "401")
			return
		}

		token = strings.Replace(token, "Bearer ", "", 1)
		claims, err := m.verifyToken(token)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintf(w, "401")
			return
		}

		name := claims.(jwt.MapClaims)["name"].(string)
		id := claims.(jwt.MapClaims)["id"].(string)

		r.Header.Set("id", id)
		r.Header.Set("name", name)

		next.ServeHTTP(w, r)
	})
}

func (m middleware) verifyToken(t string) (jwt.Claims, error) {
	token, err := jwt.Parse(t, func(token *jwt.Token) (interface{}, error) {
		return m.signingKey, nil
	})
	if token == nil {
		return nil, err
	}

	return token.Claims, nil
}
