package main

import (
	"net/http"
)

func serverStart() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /user/{id}", handleFunc)
	http.ListenAndServe(":8081", mux)
}

func handleFunc(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	w.Write([]byte(id))
}
