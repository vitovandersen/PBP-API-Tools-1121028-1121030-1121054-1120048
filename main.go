package main

import (
	"GoTools/controller"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	if (!controller.CheckUrls([]string{"https://mail.google.com/", "https://www.google.com/"})) {
		log.Println("Server is down!")
		return
	}
	router := mux.NewRouter()
	router.HandleFunc("/users/list", controller.GetUserListEndpoint).Methods("GET")
	
	router.HandleFunc("/users/login", controller.LoginEndpoint).Methods("POST")
	router.HandleFunc("/users/logout", controller.LogoutEndpoint).Methods("POST")

	http.Handle("/", router)
	log.Println("Connected to port 8081")
	log.Fatal(http.ListenAndServe(":8081", router))
}