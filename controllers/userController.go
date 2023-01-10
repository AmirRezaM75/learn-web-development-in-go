package controllers

import (
	"database/sql"
	"fmt"
	"gallery/models"
	"github.com/go-chi/chi/v5"
	"net/http"
)

type UserController struct{}

func (uc UserController) Store(w http.ResponseWriter, r *http.Request) {
	db := prepareDatabase()

	userService := models.UserService{
		DB: db,
	}

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			http.Error(w, "Something goes wrong.", 500)
		}
	}(db)

	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			email TEXT UNIQUE NOT NULL,
		    password TEXT NOT NULL
		)
	`)

	if err != nil {
		panic(err)
	}

	email := r.FormValue("email")
	password := r.FormValue("password")

	user, err := userService.Create(email, password)

	if err != nil {
		http.Error(w, "Something goes wrong.", 500)
		return
	}

	_, _ = fmt.Fprintln(w, "User has been created successfully with id", user.Id)
}

func (uc UserController) Login(w http.ResponseWriter, r *http.Request) {
	db := prepareDatabase()

	userService := models.UserService{
		DB: db,
	}

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			http.Error(w, "Something goes wrong.", 500)
		}
	}(db)

	email := r.FormValue("email")
	password := r.FormValue("password")

	err := userService.Login(email, password)

	if err == sql.ErrNoRows {
		http.NotFound(w, r)
		return
	}

	if err != nil {
		w.WriteHeader(401)
		_, _ = fmt.Fprintln(w, "Bad credentials")
		return
	}

	_, _ = fmt.Fprintln(w, "You have logged in successfully")
}

func (uc UserController) Show(w http.ResponseWriter, r *http.Request) {
	db := prepareDatabase()

	userService := models.UserService{
		DB: db,
	}

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			http.Error(w, "Something goes wrong.", 500)
		}
	}(db)

	userId := chi.URLParam(r, "userId")

	user, err := userService.GetUserById(userId)

	if err == sql.ErrNoRows {
		http.NotFound(w, r)
		return
	}

	if err != nil {
		http.Error(w, "Something goes wrong", 500)
		return
	}

	_, _ = fmt.Fprintln(w, user.Email)
}

func (uc UserController) Index(w http.ResponseWriter, _ *http.Request) {

	db := prepareDatabase()

	userService := models.UserService{
		DB: db,
	}

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			http.Error(w, "Something goes wrong.", 500)
		}
	}(db)

	users, err := userService.Get(10)

	if err != nil {
		http.Error(w, "Something goes wrong!", 500)
		return
	}

	_, _ = fmt.Fprintln(w, users)
}

func prepareDatabase() *sql.DB {
	config := models.DefaultPostgresConfig()

	db, err := models.Open(config)

	if err != nil {
		panic(err)
	}

	return db
}
