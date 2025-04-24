package config

import (
	"database/sql"
	"fmt"
	"net/url"

	_ "github.com/lib/pq"
)

func ConnectDb(cfg *Config) (*sql.DB, error) {
	escapePassword := url.QueryEscape(cfg.Password)
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", cfg.Host, cfg.Port, cfg.User, escapePassword, cfg.Name)
	db, err := sql.Open(cfg.Driver, dsn)

	if err != nil {
		return nil, fmt.Errorf("gagal memanggil sql.Open: %w", err)
	}

	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, fmt.Errorf("gagal melakukan ping ke database: %w", err)
	}

	return db, nil
}
