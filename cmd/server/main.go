package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jxlwqq/go-restful/internal/config"
	"github.com/jxlwqq/go-restful/internal/router"
	"github.com/jxlwqq/go-restful/pkg/database"
	"github.com/jxlwqq/go-restful/pkg/log"
	"gorm.io/gorm"
	"net/http"
	"os"
)

func main()  {

	cfg, err := config.Load()
	if cfg == nil {
		os.Exit(-1)
	}

	logger := log.New()

	db ,err := database.New(cfg.DSN, &gorm.Config{})
	if err != nil {
		os.Exit(-1)
	}

	r := mux.NewRouter().StrictSlash(true)
	router.BuildHandlers(r, db, logger, cfg)

	addr := fmt.Sprintf(":%v", cfg.ServerPort)
	_ = http.ListenAndServe(addr, r)
}

