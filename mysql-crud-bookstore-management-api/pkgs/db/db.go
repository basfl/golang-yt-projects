package db

import (
	"fmt"
	"go-bookstore/pkgs/config"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func Connect() {
	//since the .env file is under the project path we use "."
	config, err := config.Loadconfig(".")
	if err != nil {
		log.Fatal("can not load configuration", err)
	}
	d, err := gorm.Open("mysql", config.DBDriver)
	if err != nil {
		panic(err)
	}
	db = d
	fmt.Println("connected to mysql successfully  !!!")
}

func GetDB() *gorm.DB {
	return db
}
