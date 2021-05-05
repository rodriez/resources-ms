package main

import (
	"log"
	"net/http"
	"resources-ms/handlers"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.Use(handlers.BasicAuthMiddleware)

	initRoutes(router)

	log.Println("Listening por :8080")

	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
}

func initRoutes(router *mux.Router) {
	router.HandleFunc("/resource", handlers.CreateResource).Methods("POST")
	router.HandleFunc("/resource/{id}", handlers.FindResource).Methods("GET")
	router.HandleFunc("/resource/{id}", handlers.UpdateResource).Methods("PUT")
	router.HandleFunc("/resource/{id}", handlers.DeleteResource).Methods("DELETE")
}
