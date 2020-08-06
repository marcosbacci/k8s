package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", greeting)
	http.ListenAndServe(":8000", nil)
}

func greeting(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<b>Code.education Rocks!</b>")
}
