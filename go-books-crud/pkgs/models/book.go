package models

import "fmt"

type Book struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var Books []Book

func init() {

	fmt.Println("This will get called on main initialization")
	Books = append(Books, Book{ID: "1", Isbn: "123", Title: "first book", Author: &Author{Firstname: "john", Lastname: "doe"}})
	Books = append(Books, Book{ID: "2", Isbn: "456", Title: "second book", Author: &Author{Firstname: "john", Lastname: "doe"}})
}

func Deletebook(id string) {
	for index, elem := range Books {
		if elem.ID == id {
			Books = append(Books[:index], Books[index+1:]...)
			break
		}
	}

}
