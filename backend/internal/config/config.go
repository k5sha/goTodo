package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"os"
	"time"
)

type Config struct {
	Env        string `yaml:"env,omitempty"`
	HttpServer `yaml:"http_server"`
	Database   `yaml:"db"`
}

type HttpServer struct {
	Address     string        `yaml:"address"`
	Timeout     time.Duration `yaml:"timeout"`
	IdleTimeout time.Duration `yaml:"idle_timeout"`
}

type Database struct {
	Host         string `yaml:"host"`
	Port         uint16 `yaml:"port"`
	User         string `yaml:"user"`
	Password     string `yaml:"password"`
	DatabaseName string `yaml:"dbname"`
	SSL          bool   `yaml:"ssl_mode"`
	DatabaseURL  string `yaml:"database_url,omitempty"`
	Pool         `yaml:"pool"`
}

type Pool struct {
	MaxConn     int           `yaml:"max_conn"`
	MaxIdleConn int           `yaml:"max_idle_conn"`
	MaxLiveTime time.Duration `yaml:"max_live_time"`
}

var cfg Config

func MustLoad() *Config {
	configPath := os.Getenv("CONFIG_PATH")

	if configPath == "" {
		log.Fatal("CONFIG_PATH environment variable not set")
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file doesn't exist: %s", configPath)
	}

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("cannot read config %s", err)
	}

	if dbURL := os.Getenv("DATABASE_URL"); dbURL != "" {
		cfg.Database.DatabaseURL = dbURL
	}
	return &cfg
}
