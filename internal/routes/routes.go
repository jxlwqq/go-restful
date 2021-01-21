package routes

import (
	"github.com/jxlwqq/go-restful/internal/auth"
	"github.com/jxlwqq/go-restful/internal/config"
	"github.com/jxlwqq/go-restful/internal/healthz"
	"github.com/jxlwqq/go-restful/internal/post"
	"github.com/jxlwqq/go-restful/pkg/database"
	"github.com/jxlwqq/go-restful/pkg/log"
	"github.com/jxlwqq/go-restful/pkg/prometheus"
	"github.com/jxlwqq/go-restful/pkg/router"
)

func BuildHandlers(r *router.Router, db *database.DB, logger *log.Logger, cfg *config.Config) {
	r.Use(prometheus.Handler)
	authMiddleware := auth.NewMiddleware(cfg.JWTSigningKey)
	post.RegisterHandlers(r.PathPrefix("").Subrouter(), db, logger, authMiddleware)
	auth.RegisterHandlers(r.PathPrefix("").Subrouter(), db, logger, cfg, authMiddleware)
	healthz.RegisterHandlers(r.PathPrefix("").Subrouter())
	prometheus.RegisterHandlers(r.PathPrefix("").Subrouter())
}
