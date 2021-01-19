package main

import (
	"context"
	"fmt"
	"github.com/jxlwqq/go-restful/internal/config"
	"github.com/jxlwqq/go-restful/internal/routes"
	"github.com/jxlwqq/go-restful/pkg/database"
	"github.com/jxlwqq/go-restful/pkg/log"
	"github.com/jxlwqq/go-restful/pkg/router"
	"gorm.io/gorm"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	cfg, _ := config.Load("./configs/.env")
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

	server := &http.Server{
		Addr:    addr,
		Handler: r,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil {
			os.Exit(-1)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logger.Info("Server is shutting down")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		logger.Fatal("Server forced to shutdown:", err)
	}
	logger.Info("Server exiting")
}
