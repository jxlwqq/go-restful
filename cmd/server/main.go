package main

import (
	"fmt"
	"github.com/jxlwqq/go-restful/internal/config"
	"github.com/jxlwqq/go-restful/internal/routes"
	"github.com/jxlwqq/go-restful/pkg/database"
	"github.com/jxlwqq/go-restful/pkg/log"
	"github.com/jxlwqq/go-restful/pkg/router"
	"gorm.io/gorm"
	"net/http"
	"os"
)

func main() {

	cfg, _ := config.Load()
	if cfg == nil {
		os.Exit(-1)
	}

	logger := log.New()

	db, err := database.New(cfg.DSN, &gorm.Config{})
	if err != nil {
		os.Exit(-1)
	}

	r := router.New()
	routes.BuildHandlers(r, db, logger, cfg)

	addr := fmt.Sprintf(":%v", cfg.ServerPort)
	_ = http.ListenAndServe(addr, r)
}
