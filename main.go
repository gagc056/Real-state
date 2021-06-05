package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	var r := mux.NewRouter()
	http.HandleFunc("/", home)
	http.ListenAndServe(":3000", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}
