package main

import (
	"fmt"
	"net/http"
)

func handleFunc(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprint(w, "<h1>Hello World</h1>")
}

func main() {
	http.HandleFunc("/", handleFunc)
	fmt.Println("Listening on port 3000")
	_ = http.ListenAndServe(":3000", nil)
}
