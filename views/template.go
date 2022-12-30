package views

import (
	"html/template"
	"log"
	"net/http"
)

func Render(w http.ResponseWriter, path string) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	t, err := template.ParseFiles(path)

	if err != nil {
		log.Printf("parsing template %v", err)
		http.Error(w, "There was an error parsing the template", http.StatusInternalServerError)
		return
	}

	_ = t.Execute(w, nil)
}
