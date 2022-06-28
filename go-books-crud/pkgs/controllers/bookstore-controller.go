package controllers

import (
	"encoding/json"
	"fmt"
	"go-books-crud/pkgs/models"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func Getbooks(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(models.Books)
}

func Getbook(res http.ResponseWriter, req *http.Request) {

	params := mux.Vars(req)

	for index, elem := range models.Books {
		if elem.ID == params["id"] {
			res.Header().Set("Content-Type", "application/json")
			json.NewEncoder(res).Encode(models.Books[index])
		}
	}
}

func Createbooks(res http.ResponseWriter, req *http.Request) {
	var newBook models.Book
	newBook.ID = strconv.Itoa(rand.Intn(10000000))
	if err := json.NewDecoder(req.Body).Decode(&newBook); err != nil {
		log.Fatal(err)
	}
	models.Books = append(models.Books, newBook)
	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(models.Books)
}

func Updatebooks(res http.ResponseWriter, req *http.Request) {

	params := mux.Vars(req)
	var updatedBook models.Book
	for index, elem := range models.Books {
		if elem.ID == params["id"] {
			if err := json.NewDecoder(req.Body).Decode(&updatedBook); err != nil {
				log.Fatal(err)
			}
			models.Books[index] = updatedBook
			res.Header().Set("Content-Type", "application/json")
			json.NewEncoder(res).Encode(models.Books[index])
		}
	}

	fmt.Println("running update")
}

func Deletebooks(res http.ResponseWriter, req *http.Request) {

	params := mux.Vars(req)
	res.Header().Set("Content-Type", "application/json")
	models.Deletebook(params["id"])
	json.NewEncoder(res).Encode(models.Books)

}
