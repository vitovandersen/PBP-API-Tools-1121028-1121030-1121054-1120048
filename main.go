package main

import (
	"GoTools/controller"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()


	http.Handle("/", router)
	log.Println("Connected to port 8081")
	log.Fatal(http.ListenAndServe(":8081", router))
}