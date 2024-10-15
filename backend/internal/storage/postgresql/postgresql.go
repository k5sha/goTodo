package postgresql

import (
	"database/sql"
	"fmt"
	"github.com/k5sha/golang-todo-example/internal/config"
	_ "github.com/lib/pq"
	"log/slog"
	"time"
)

type Storage struct {
	DB *sql.DB
}

func New(cfg config.Database, log *slog.Logger) (*Storage, error) {
	const op = "storage.postgresql.New"

	var psqlInfo string
	if cfg.DatabaseURL != "" {
		psqlInfo = cfg.DatabaseURL
	} else {
		psqlInfo = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
			cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DatabaseName, sslMode(cfg.SSL))
	}

	var db *sql.DB
	var err error

	for attempts := 0; attempts < 5; attempts++ {
		db, err = sql.Open("postgres", psqlInfo)
		if err != nil {
			log.Warn("Failed to open database connection", "attempt", attempts+1, "error", err)
			time.Sleep(5 * time.Second)
			continue
		}

		db.SetMaxOpenConns(cfg.Pool.MaxConn)
		db.SetMaxIdleConns(cfg.Pool.MaxIdleConn)
		db.SetConnMaxLifetime(cfg.Pool.MaxLiveTime)

		if err = db.Ping(); err == nil {
			log.Info("Successfully connected to PostgreSQL database")
			break
		}

		log.Warn("Failed to ping database", "attempt", attempts+1, "error", err)
		time.Sleep(5 * time.Second)
	}

	if err != nil {
		return nil, fmt.Errorf("%s: unable to connect to database after 5 attempts: %w", op, err)
	}

	// Создание таблицы todos, если она не существует
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS todos (
			id SERIAL PRIMARY KEY,
			title TEXT NOT NULL,
			completed BOOLEAN DEFAULT FALSE,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		);
	`)
	if err != nil {
		return nil, fmt.Errorf("%s: failed to create todos table: %w", op, err)
	}

	return &Storage{DB: db}, nil
}

func sslMode(ssl bool) string {
	if ssl {
		return "require"
	}
	return "disable"
}
