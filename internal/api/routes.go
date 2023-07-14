package api

import (
    "net/http"
    "github.com/blessium/metricsgo/internal/api/books"
)

type BookHandler struct {
    service books.IService
}

func GetBookHandler(service books.IService) BookHandler {
    return BookHandler {
        service: service,
    } 
}

func (b BookHandler) GetBookHandler(w http.ResponseWriter, r *http.Request) {

    w.WriteHeader(http.StatusOK) 
}
