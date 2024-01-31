package main

import (
	"joy-tech/pkg/book"
	"joy-tech/pkg/schedule"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func handleRequests() {

	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/book/list", book.HandleGetBookList).Methods("GET")
	myRouter.HandleFunc("/book/schedule", schedule.HandleSchedulePickUp).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", myRouter))

}
