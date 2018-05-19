package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"github.com/gorilla/handlers"
	"github.com/amazonparser/handler"
)


func main() {

	r := mux.NewRouter()
	r.HandleFunc("/api/sendurls", handler.UrlHandler).Methods("POST")
	r.HandleFunc("/api/requests", handler.AllRequestsHandler).Methods("GET")
	r.HandleFunc("/", handler.InfoHandler).Methods("GET")
	r.HandleFunc("/api/requests/{id}", handler.GetRequestHandler).Methods("GET")
	http.ListenAndServe(":7755", handlers.CORS()(r))
}

