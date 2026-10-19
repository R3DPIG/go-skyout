package database

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	gormmysql "gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"

	"github.com/R3DPIG/go-skyout/config"
)

// MustNewMySQL 初始化 GORM MySQL 连接
func MustNewMySQL(cfg config.MySQLConfig, mode string) *gorm.DB {
	level := gormlogger.Warn
	if mode == "debug" {
		level = gormlogger.Info
	}
	db, err := gorm.Open(gormmysql.Open(cfg.DSN()), &gorm.Config{
		Logger: gormlogger.Default.LogMode(level),
	})
	if err != nil {
		panic("init mysql failed: " + err.Error())
	}
	sqlDB, err := db.DB()
	if err != nil {
		panic("get sql db failed: " + err.Error())
	}
	sqlDB.SetMaxOpenConns(cfg.MaxOpenConns)
	sqlDB.SetMaxIdleConns(cfg.MaxIdleConns)
	sqlDB.SetConnMaxLifetime(time.Hour)
	return db
}

// MustNewRedis 初始化 Redis 客户端
func MustNewRedis(cfg config.RedisConfig) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.Addr,
		Password: cfg.Password,
		DB:       cfg.DB,
	})
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	if err := rdb.Ping(ctx).Err(); err != nil {
		panic("init redis failed: " + err.Error())
	}
	return rdb
}

// LoadViperConfig 读取 yaml 配置为 viper 实例
func LoadViperConfig(path string) *viper.Viper {
	v := viper.New()
	v.SetConfigFile(path)
	if err := v.ReadInConfig(); err != nil {
		panic("read config failed: " + err.Error())
	}
	return v
}
