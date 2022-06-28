package models

import (
	"go-bookstore/pkgs/db"

	"github.com/jinzhu/gorm"
)

var database *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:""json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {

	db.Connect()
	database = db.GetDB()
	database.AutoMigrate(&Book{})

}
func (book *Book) CreateBook() *Book {
	database.NewRecord(book)
	database.Create(book)
	return book
}
func GetAllBooks() []Book {
	var Books []Book
	database.Find(&Books)
	return Books
}

func GetBookById(Id int64) (*Book, *gorm.DB) {
	var getBook Book
	database := database.Where("ID=?", Id).Find(&getBook)
	return &getBook, database
}

func DeleteBook(ID int64) Book {
	var book Book
	database.Where("ID=?", ID).Delete(book)
	return book
}
