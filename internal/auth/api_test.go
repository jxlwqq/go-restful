package auth

import (
	"bytes"
	"encoding/json"
	"github.com/jxlwqq/go-restful/internal/config"
	"github.com/jxlwqq/go-restful/pkg/database"
	"github.com/jxlwqq/go-restful/pkg/log"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestResource_Login(t *testing.T) {
	cfg, _ := config.Load("../../configs/.env")
	logger := log.New()
	db, _ := database.New(cfg.DSN, &gorm.Config{})
	svc := NewService(cfg.JWTSigningKey, cfg.JWTExpiration, db, logger)
	resource := resource{svc, logger}
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "/auth/login", bytes.NewBufferString(`{"mobile":"demo","code":"1234"}`))
	resource.Login(recorder, request)
	assert.Equal(t, recorder.Code, http.StatusOK)
	res := struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
		Data    struct {
			Token string `json:"token"`
		} `json:"data"`
	}{}
	_ = json.NewDecoder(recorder.Body).Decode(&res)
	assert.Equal(t, res.Code, http.StatusOK)
	assert.NotEmpty(t, res.Data.Token)
}
