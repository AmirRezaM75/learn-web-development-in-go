package main

import (
	"fmt"
	"gallery/controllers"
	"gallery/views"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func main() {
	sc := controllers.StaticController{
		View: views.View{},
	}
	uc := controllers.UserController{}

	r := chi.NewRouter()
	r.Get("/", sc.Home)
	r.Get("/contact", sc.Contact)
	r.Get("/faq", sc.Faq)
	r.Post("/register", uc.Store)
	r.Get("/users", uc.Index)
	r.Get("/users/{userId}", uc.Show)
	fmt.Println("Listening on port 3000")
	_ = http.ListenAndServe(":3000", r)
}
