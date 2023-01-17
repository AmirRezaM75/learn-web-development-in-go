package main

import (
	"database/sql"
	"fmt"
	"gallery/controllers"
	"gallery/models"
	"gallery/views"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func main() {
	db := prepareDatabase()

	userService := models.UserService{
		DB: db,
	}

	defer db.Close()

	sc := controllers.StaticController{
		View: views.View{},
	}

	uc := controllers.UserController{
		UserService: userService,
	}

	r := chi.NewRouter()
	r.Get("/", sc.Home)
	r.Get("/contact", sc.Contact)
	r.Get("/faq", sc.Faq)
	r.Post("/register", uc.Store)
	r.Post("/login", uc.Login)
	r.Get("/users", uc.Index)
	r.Get("/users/{userId}", uc.Show)
	r.Get("/users/me", uc.Me)
	fmt.Println("Listening on port 3000")
	_ = http.ListenAndServe(":3000", r)
}

func prepareDatabase() *sql.DB {
	config := models.DefaultPostgresConfig()

	db, err := models.Open(config)

	if err != nil {
		panic(err)
	}

	return db
}
