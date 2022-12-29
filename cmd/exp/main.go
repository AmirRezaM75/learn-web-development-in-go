package main

import (
	"html/template"
	"os"
)

type User struct {
	Name string
	Meta UserMeta
}

type UserMeta struct {
	Age int
}

func main() {
	user := User{
		Name: "Twitch",
		Meta: UserMeta{
			Age: 20,
		},
	}
	t, err := template.ParseFiles("index.html")
	if err != nil {
		panic(err)
	}
	_ = t.Execute(os.Stdout, user)
}
