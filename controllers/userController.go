package controllers

import (
	"database/sql"
	"fmt"
	_ "github.com/jackc/pgx/v5/stdlib"
	"net/http"
)

type UserController struct{}

func (uc UserController) Store(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("pgx", "host=localhost port=5432 user=root password=root dbname=unsplash sslmode=disable")
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
