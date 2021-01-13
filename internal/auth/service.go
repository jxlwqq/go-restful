package auth

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/jxlwqq/go-restful/internal/entity"
	"github.com/jxlwqq/go-restful/pkg/database"
	"github.com/jxlwqq/go-restful/pkg/log"
	"time"
)

type Service interface {
	Login(args ...interface{}) (string, error)
	generateToken(identity Identity) (string, error)
}

func NewService(signingKey string, tokenExpiration int, db *database.DB, logger *log.Logger) Service {
	return service{signingKey, tokenExpiration, db, logger}
}

type service struct {
	signingKey      string
	tokenExpiration int
	db              *database.DB
	logger          *log.Logger
}

func (s service) Login(args ...interface{}) (string, error) {
	mobile := args[0].(string)
	code := args[1].(string)
	fmt.Print(mobile)
	if mobile == "demo" && code == "1234" {
		user := entity.User{ID: 100, Name: "demo"}
		return s.generateToken(user)
	}

	return "", nil
}

type Identity interface {
	// GetID returns the user ID.
	GetID() string
	// GetName returns the user name.
	GetName() string
}

func (s service) generateToken(identity Identity) (string, error) {
	return jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":   identity.GetID(),
		"name": identity.GetName(),
		"exp":  time.Now().Add(time.Duration(s.tokenExpiration) * time.Hour).Unix(),
	}).SignedString([]byte(s.signingKey))
}
