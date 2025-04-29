package main

import (
    "net/http"
    "computersciencestudy/api"
    "github.com/gorilla/mux"
)

func main() {
    r := mux.NewRouter()
    r.HandleFunc("/api/posts", api.Handler).Methods("GET")
    r.HandleFunc("/api/posts/{slug}", api.Handler).Methods("GET")
    r.PathPrefix("/").Handler(http.FileServer(http.Dir("public")))
    http.ListenAndServe(":8080", r)
}
