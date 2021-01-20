package auth

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/jxlwqq/go-restful/internal/entity"
	"github.com/jxlwqq/go-restful/pkg/database"
	"github.com/jxlwqq/go-restful/pkg/log"
	"time"
)

type Service interface {
	Login(args ...interface{}) (string, error)
	Me(id string) (Identity, error)
	generateToken(Identity) (string, error)
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
	// todo 检测code是否存在或过期
	user := entity.User{
		Mobile: mobile,
	}
	s.db.FirstOrCreate(&user)
	return s.generateToken(user)
}

func (s service) Me(id string) (Identity, error) {
	user := entity.User{}
	err := s.db.Find(&user, id).Error
	return user, err
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
