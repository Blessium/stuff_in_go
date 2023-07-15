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

	booksRepo := books.GetMongoRepository(mongoDB.Collection("Books"))

	booksService := books.GetService(booksRepo)
	bookHandler := books.GetHandler(booksService)

	middleware := []api.Middleware{
		api.LogMiddleware,
	}

	endpoints := map[string]http.Handler{
		"/books": bookHandler,
	}

	for endpoints, f := range endpoints {
		http.HandleFunc(endpoints, api.MultipleMiddleware(f, middleware...))
	}

	log.Fatal(http.ListenAndServe(":8080", nil))
}
