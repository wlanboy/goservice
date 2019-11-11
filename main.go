package main

import (
	"fmt"
	"net/http"
	"os"

	controller "./controller"
	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/api/v1/event", controller.PostCreate).Methods("POST")
	router.HandleFunc("/api/v1/event", controller.GetAll).Methods("GET")
	router.HandleFunc("/api/v1/event/{id}", controller.GetByID).Methods("GET")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	fmt.Println(port)

	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		fmt.Print(err)
	}
}
