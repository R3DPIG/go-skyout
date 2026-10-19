package config

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoadOverrides(t *testing.T) {
	dir := t.TempDir()
	base := filepath.Join(dir, "config.yaml")
	local := filepath.Join(dir, "config.local.yaml")
	if err := os.WriteFile(base, []byte("app:\n  port: 8080\nmysql:\n  password: default\njwt:\n  secret: default\n"), 0600); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(local, []byte("mysql:\n  password: local\njwt:\n  secret: local\n"), 0600); err != nil {
		t.Fatal(err)
	}
	t.Setenv("APP_PORT", "9090")
	config, err := Load(base)
	if err != nil {
		t.Fatal(err)
	}
	if config.App.Port != 9090 || config.MySQL.Password != "local" || config.JWT.Secret != "local" {
		t.Fatalf("unexpected merged config: port=%d, local mysql override=%t, local jwt override=%t",
			config.App.Port, config.MySQL.Password == "local", config.JWT.Secret == "local")
	}
}
