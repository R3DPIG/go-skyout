package config

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
)

// AppConfig 应用配置
type AppConfig struct {
	Name string `mapstructure:"name"`
	Port int    `mapstructure:"port"`
	Mode string `mapstructure:"mode"`
}

// MySQLConfig MySQL 配置
type MySQLConfig struct {
	Host         string `mapstructure:"host"`
	Port         int    `mapstructure:"port"`
	User         string `mapstructure:"user"`
	Password     string `mapstructure:"password"`
	Database     string `mapstructure:"database"`
	Charset      string `mapstructure:"charset"`
	MaxOpenConns int    `mapstructure:"max_open_conns"`
	MaxIdleConns int    `mapstructure:"max_idle_conns"`
}

// DSN 生成 GORM 的 MySQL DSN
func (m *MySQLConfig) DSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
		m.User, m.Password, m.Host, m.Port, m.Database, m.Charset)
}

// RedisConfig Redis 配置
type RedisConfig struct {
	Addr     string `mapstructure:"addr"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
}

// JWTConfig JWT 配置
type JWTConfig struct {
	Secret      string `mapstructure:"secret"`
	ExpireHours int    `mapstructure:"expire_hours"`
	Issuer      string `mapstructure:"issuer"`
}

// Config 总配置
type Config struct {
	App   AppConfig   `mapstructure:"app"`
	MySQL MySQLConfig `mapstructure:"mysql"`
	Redis RedisConfig `mapstructure:"redis"`
	JWT   JWTConfig   `mapstructure:"jwt"`
}

// Load 从默认 yaml、本地配置和环境变量加载配置。后面的配置覆盖前面的配置。
func Load(path string) (*Config, error) {
	v := viper.New()
	v.SetConfigFile(path)
	if err := v.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("read config: %w", err)
	}
	localPath := filepath.Join(filepath.Dir(path), "config.local.yaml")
	if _, err := os.Stat(localPath); err == nil {
		v.SetConfigFile(localPath)
		if err := v.MergeInConfig(); err != nil {
			return nil, fmt.Errorf("read local config: %w", err)
		}
	} else if !os.IsNotExist(err) {
		return nil, fmt.Errorf("check local config: %w", err)
	}
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.AutomaticEnv()
	var c Config
	if err := v.Unmarshal(&c); err != nil {
		return nil, fmt.Errorf("unmarshal config: %w", err)
	}
	return &c, nil
}
