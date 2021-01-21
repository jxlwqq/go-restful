package main

import (
	"github.com/jxlwqq/go-restful/internal/config"
	"github.com/jxlwqq/go-restful/internal/entity"
	"github.com/jxlwqq/go-restful/pkg/database"
	"gorm.io/gorm"
)

func main() {
	cfg, _ := config.Load("./configs/.env")
	db, _ := database.New(cfg.DSN, &gorm.Config{})
	_ = db.AutoMigrate(
		entity.User{},
		entity.Post{},
	)
}
