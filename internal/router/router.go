package router

import (
	"github.com/gorilla/mux"
	"github.com/jxlwqq/go-restful/internal/auth"
	"github.com/jxlwqq/go-restful/internal/config"
	"github.com/jxlwqq/go-restful/internal/post"
	"github.com/jxlwqq/go-restful/pkg/database"
	"github.com/jxlwqq/go-restful/pkg/log"
)

func BuildHandlers(r *mux.Router, db *database.DB, logger *log.Logger, cfg *config.Config) {
	authMiddleware := auth.NewMiddleware(cfg.JWTSigningKey)
	post.RegisterHandlers(r.PathPrefix("").Subrouter(), db, logger, authMiddleware)
	auth.RegisterHandlers(r.PathPrefix("").Subrouter(), db, logger, cfg)
}
