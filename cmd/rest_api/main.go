package main

import (
	"context"
	"github.com/blessium/metricsgo/internal/api"
	"github.com/blessium/metricsgo/internal/api/books"
	"github.com/blessium/metricsgo/internal/db"
	"log"
	"net/http"
    "fmt"
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

    fmt.Println("Client mongodb got")
	mongoDB := mongoClient.Database("api_server")
    mongoCollection := mongoDB.Collection("Books")
    if err := db.CreateIndex(mongoCollection, "ISBN", true); err != nil {
        panic(err.Error())
    }

	booksRepo := books.GetMongoRepository(mongoCollection)

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

    fmt.Println("Probably server started")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
