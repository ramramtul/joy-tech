package main

import "joy-tech/pkg/book"

type Handlers struct {
	bookHandler *book.BookHandler
}

func InitHandlers() *Handlers {
	return &Handlers{
		bookHandler: book.NewBookHandler(),
	}
}
