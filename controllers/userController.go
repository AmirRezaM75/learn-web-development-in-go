package controllers

import (
	"database/sql"
	"fmt"
	_ "github.com/jackc/pgx/v5/stdlib"
	"net/http"
)

type UserController struct{}

type PostgresConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Database string
	SSLMode  bool
}

func (c PostgresConfig) toString() string {
	SSLMode := "disable"
	if c.SSLMode {
		SSLMode = "enable"
	}
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		c.Host, c.Port, c.User, c.Password, c.Database, SSLMode)
}

func (uc UserController) Store(w http.ResponseWriter, r *http.Request) {
	config := PostgresConfig{
		Host:     "localhost",
		Port:     5432,
		User:     "root",
		Password: "root",
		Database: "unsplash",
		SSLMode:  false,
	}
	db, err := sql.Open("pgx", config.toString())
	if err != nil {
		panic(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Fprintln(w, "Connected!")

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			name VARCHAR(128) NOT NULL,
			email TEXT UNIQUE NOT NULL,
		    password TEXT NOT NULL
		)
	`)

	if err != nil {
		panic(err)
	}

	email := r.FormValue("email")
	name := r.FormValue("name")
	password := r.FormValue("password")

	_, err = db.Exec(`
		INSERT INTO users (email, name, password)
		VALUES ($1, $2, $3)
	`, email, name, password)

	if err != nil {
		panic(err)
	}
}
