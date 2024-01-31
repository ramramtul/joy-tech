package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func handleRequests(handler *Handlers) {

	myRouter := mux.NewRouter().StrictSlash(true)

	// GET list of books
	myRouter.HandleFunc("/book/list", handler.bookHandler.HandleGetBookList)

	// POST

	log.Fatal(http.ListenAndServe(":8080", myRouter))

}
