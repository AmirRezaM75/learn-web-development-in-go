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
	return

	err = r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	fmt.Fprintln(w, r.PostForm.Get("email"))
	fmt.Fprintln(w, r.PostFormValue("email"))
	fmt.Fprintln(w, r.FormValue("email"))
}
