package main

import (
	"fmt"
	"gallery/resources"
	"gallery/views"
	"github.com/go-chi/chi/v5"
	"net/http"
	"path/filepath"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	path := filepath.Join("views", "home.html")
	views.Render(w, resources.FS, path)
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	path := filepath.Join("views", "contact.html")
	views.Render(w, resources.FS, path)
}

func main() {
	r := chi.NewRouter()
	r.Get("/", homeHandler)
	r.Get("/contact", contactHandler)
	fmt.Println("Listening on port 3000")
	_ = http.ListenAndServe(":3000", r)
}
