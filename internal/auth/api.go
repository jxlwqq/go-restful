package auth

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/jxlwqq/go-restful/internal/config"
	"github.com/jxlwqq/go-restful/internal/response"
	"github.com/jxlwqq/go-restful/pkg/database"
	"github.com/jxlwqq/go-restful/pkg/log"
	"net/http"
)

func RegisterHandlers(r *mux.Router, db *database.DB, logger *log.Logger, cfg *config.Config, authMiddleware Middleware) {
	authService := NewService(cfg.JWTSigningKey, cfg.JWTExpiration, db, logger)
	res := resource{authService, logger}
	r.HandleFunc("/auth/login", res.Login).Methods(http.MethodPost)
	s := r.PathPrefix("").Subrouter()
	s.HandleFunc("/auth/me", res.Me).Methods(http.MethodGet)
	s.HandleFunc("/auth/logout", res.Logout).Methods(http.MethodPost)
	s.HandleFunc("/auth/refresh", res.Refresh).Methods(http.MethodPost)
	s.Use(authMiddleware.Handler)
}

type resource struct {
	service Service
	logger  *log.Logger
}

func (res resource) Login(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Mobile string `json:"mobile"`
		Code   string `json:"code"`
	}

	_ = json.NewDecoder(r.Body).Decode(&req)

	res.logger.With("mobile", req.Mobile, "code", req.Code).Info()
	token, _ := res.service.Login(req.Mobile, req.Code)

	response.Write(w, struct {
		Token string `json:"token"`
	}{token}, http.StatusOK)
}

func (res resource) Me(w http.ResponseWriter, r *http.Request) {
	id := r.Header.Get("id")
	user, _ := res.service.Me(id)
	response.Write(w, user, http.StatusOK)
}

func (res resource) Logout(w http.ResponseWriter, r *http.Request) {
	// todo
}

func (res resource) Refresh(w http.ResponseWriter, r *http.Request) {
	// todo
}
