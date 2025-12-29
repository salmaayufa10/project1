package config

import "os"

type Config struct {
	DatabaseURL string
	Port        string
}

func Load() Config {
	cfg := Config{
		DatabaseURL: os.Getenv("DATABASE_URL"),
		Port:        os.Getenv("PORT"),
	}
	if cfg.Port == "" {
		cfg.Port = "8080"
	}

	if cfg.DatabaseURL == "" {
		cfg.DatabaseURL = "postgresql://postgres:xYcDZeLgqypCZidoBTrLBPbakzFRgheA@tramway.proxy.rlwy.net:14393/railway?sslmode=require"
	}
	return cfg
}
