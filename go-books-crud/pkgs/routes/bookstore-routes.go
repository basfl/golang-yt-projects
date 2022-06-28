package routes

import (
	"go-books-crud/pkgs/controllers"

	"github.com/gorilla/mux"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/books", controllers.Getbooks).Methods("GET")
	router.HandleFunc("/books/{id}", controllers.Getbook).Methods("GET")
	router.HandleFunc("/books", controllers.Createbooks).Methods("POST")
	router.HandleFunc("/books/{id}", controllers.Updatebooks).Methods("PUT")
	router.HandleFunc("/books/{id}", controllers.Deletebooks).Methods("DELETE")

}
