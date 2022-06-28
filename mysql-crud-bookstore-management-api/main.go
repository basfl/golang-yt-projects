package main

import (
	"fmt"
	"go-bookstore/pkgs/db"
	"go-bookstore/pkgs/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	db.Connect()
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	fmt.Println("server started .....")
	log.Fatal(http.ListenAndServe(":8080", r))

}
