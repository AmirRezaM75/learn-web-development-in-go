package views

import (
	"gallery/resources"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

func Render(w http.ResponseWriter, data any, paths ...string) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	var patterns []string

	for _, path := range paths {
		path = filepath.Join("views", path)
		patterns = append(patterns, path)
	}

	t, err := template.ParseFS(resources.FS, patterns...)

	if err != nil {
		log.Printf("parsing template %v", err)
		http.Error(w, "There was an error parsing the template", http.StatusInternalServerError)
		return
	}

	_ = t.Execute(w, data)
}
