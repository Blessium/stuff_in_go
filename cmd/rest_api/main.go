package main

import (
    "net/http"
    "log"
    "github.com/blessium/metricsgo/internal/api"
    "github.com/blessium/metricsgo/internal/api/books"
)

func main() {
    
    booksRepo, err := books.GetMongoRepository()
    if err != nil {
        panic(err.Error())
    }
    booksService := books.GetService(booksRepo)
    book := api.GetBookHandler(booksService)

    http.HandleFunc("/books", book.GetBookHandler)

    log.Fatal(http.ListenAndServe(":8080", nil))
}
