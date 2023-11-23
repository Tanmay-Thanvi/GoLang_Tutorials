package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Hello MOD in golang")
	greeter()
	r := mux.NewRouter()
	r.HandleFunc("/", servHome).Methods("GET")

	log.Fatal(http.ListenAndServe(":4000", r)) // better than comma ok syntax 
}

func greeter() {
	fmt.Println("Hello Mod Users!")
}

func servHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Hello Golang</h1>"))
}
