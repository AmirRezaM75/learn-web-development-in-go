package controllers

import (
	"fmt"
	"net/http"
)

type UserController struct{}

func (uc UserController) Store(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	fmt.Fprintln(w, r.PostForm.Get("email"))
	fmt.Fprintln(w, r.PostFormValue("email"))
	fmt.Fprintln(w, r.FormValue("email"))
}
