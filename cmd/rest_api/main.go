package main

import (
	"context"
	"github.com/blessium/metricsgo/internal/api"
	"github.com/blessium/metricsgo/internal/api/books"
	"github.com/blessium/metricsgo/internal/db"
	"log"
	"net/http"
)

func main() {

	mongoClient, err := db.GetMongoClient()
	if err != nil {
		panic(err.Error())
	}
	defer func() {
		if err = mongoClient.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	mongoDB := mongoClient.Database("api_server")
	mongoCollection := mongoDB.Collection("Books")
	if err := db.CreateIndex(mongoCollection, "ISBN", true); err != nil {
		panic(err.Error())
	}

	booksRepo := books.GetMongoRepository(mongoCollection)

	booksService := books.GetService(booksRepo)
	bookHandler := books.GetHandler(booksService)
    bookSigleHandler := books.GetSingleHandler(booksService)

	middleware := []api.Middleware{
		api.LogMiddleware,
	}

	endpoints := map[string]http.Handler{
		"/books": bookHandler,
        "/books/": bookSigleHandler,
	}

	for endpoints, f := range endpoints {
		http.HandleFunc(endpoints, api.MultipleMiddleware(f, middleware...))
	}

	log.Fatal(http.ListenAndServe(":8080", nil))
}
