package models

import (
	"database/sql"
	"fmt"
	_ "github.com/jackc/pgx/v5/stdlib"
)

type PostgresConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Database string
	SSLMode  bool
}

func DefaultPostgresConfig() PostgresConfig {
	return PostgresConfig{
		Host:     "localhost",
		Port:     5432,
		User:     "root",
		Password: "root",
		Database: "unsplash",
		SSLMode:  false,
	}
}

func (c PostgresConfig) toString() string {
	SSLMode := "disable"
	if c.SSLMode {
		SSLMode = "enable"
	}
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		c.Host, c.Port, c.User, c.Password, c.Database, SSLMode)
}

func Open(config PostgresConfig) (*sql.DB, error) {
	db, err := sql.Open("pgx", config.toString())

	if err != nil {
		return nil, fmt.Errorf("open postgres failed %w", err)
	}

	err = db.Ping()

	if err != nil {
		return nil, fmt.Errorf("ping postgres failed %w", err)
	}

	fmt.Println("Database connected.")

	return db, nil
}
