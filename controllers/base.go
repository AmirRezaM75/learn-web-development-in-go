package controllers

import (
	"gallery/views"
	"net/http"
)

func View() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		views.Render(w, "./home.html")
	}
}
