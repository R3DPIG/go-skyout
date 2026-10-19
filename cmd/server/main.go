package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"

	"github.com/R3DPIG/go-skyout/config"
	"github.com/R3DPIG/go-skyout/internal/database"
	"github.com/R3DPIG/go-skyout/internal/model"
	"github.com/R3DPIG/go-skyout/internal/router"
)

func main() {
	cfg, err := config.Load("configs/config.yaml")
	if err != nil {
		log.Fatalf("load config: %v", err)
	}

	gin.SetMode(gin.ReleaseMode)
	if cfg.App.Mode == "debug" {
		gin.SetMode(gin.DebugMode)
	}

	db := database.MustNewMySQL(cfg.MySQL, cfg.App.Mode)
	// 自动建表（脚手架阶段，生产建议用 SQL 迁移脚本）
	if err := db.AutoMigrate(
		&model.Employee{}, &model.Category{}, &model.Dish{}, &model.DishFlavor{},
		&model.Setmeal{}, &model.SetmealDish{}, &model.User{}, &model.AddressBook{},
		&model.ShoppingCart{}, &model.Orders{}, &model.OrderDetail{},
	); err != nil {
		log.Fatalf("auto migrate: %v", err)
	}

	r := router.New(db)
	addr := fmt.Sprintf(":%d", cfg.App.Port)
	log.Printf("[%s] listening on %s (mode=%s)", cfg.App.Name, addr, cfg.App.Mode)
	if err := r.Run(addr); err != nil {
		log.Fatalf("run: %v", err)
	}
}
