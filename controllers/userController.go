package controllers

import (
	"database/sql"
	"fmt"
	"gallery/models"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

type UserController struct {
	UserService models.UserService
}

func (uc UserController) Store(w http.ResponseWriter, r *http.Request) {
	_, err := uc.UserService.DB.Exec(`
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

	user, err := uc.UserService.Create(email, password)

	if err != nil {
		http.Error(w, "Something goes wrong.", 500)
		return
	}

	_, _ = fmt.Fprintln(w, "User has been created successfully with id", user.Id)
}

func (uc UserController) Login(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	password := r.FormValue("password")

	user, err := uc.UserService.Login(email, password)

	if err == sql.ErrNoRows {
		http.NotFound(w, r)
		return
	}

	if err != nil {
		w.WriteHeader(401)
		_, _ = fmt.Fprintln(w, "Bad credentials")
		return
	}
	cookie := http.Cookie{
		Name:     "session",
		Value:    strconv.Itoa(user.Id),
		Path:     "/",
		HttpOnly: true,
	}
	http.SetCookie(w, &cookie)

	_, _ = fmt.Fprintln(w, "You have logged in successfully")
}

func (uc UserController) Show(w http.ResponseWriter, r *http.Request) {
	userId := chi.URLParam(r, "userId")

	user, err := uc.UserService.GetUserById(userId)

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
	users, err := uc.UserService.Get(10)

	if err != nil {
		http.Error(w, "Something goes wrong!", 500)
		return
	}

	_, _ = fmt.Fprintln(w, users)
}

func (uc UserController) Me(w http.ResponseWriter, r *http.Request) {
	session, err := r.Cookie("session")
	if err != nil {
		http.Error(w, "Unauthenticated.", 401)
		return
	}

	fmt.Fprintln(w, session.Value)
}
