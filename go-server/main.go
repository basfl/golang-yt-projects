package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formhandler)
	http.HandleFunc("/hello", helloHandler)
	fmt.Println("server started ...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func helloHandler(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/hello" {
		http.Error(res, "404 not found", http.StatusNotFound)
		return
	}
	if req.Method != "GET" {
		http.Error(res, "Method not supported", http.StatusNotFound)
		return
	}
	fmt.Fprintf(res, "hello")
}

func formhandler(res http.ResponseWriter, req *http.Request) {
	if err := req.ParseForm(); err != nil {
		fmt.Fprintf(res, "parseForm(); err: %v", err)
	}
	fmt.Fprintf(res, "POST Request successful !")
	name := req.FormValue("name")
	address := req.FormValue("address")
	fmt.Fprintf(res, "Name = %s\n", name)
	fmt.Fprintf(res, "Address = %s\n", address)

}
