package books

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/blessium/metricsgo/internal/api"
)

const timeLayout = "01/02/2006 15:05"

// For bulk crud
type Handler struct {
	service IService
}

// For single book crud
type SingleHandler struct {
	service IService
}

type BookFullRequest struct {
	ISBN      string `json:"isbn"`
	Title     string `json:"title"`
	Author    string `json:"author"`
	Published string `json:"published"`
	Pages     uint   `json:"pages"`
}

func (b BookFullRequest) Validate() error {
	if b.ISBN == "" {
		return errors.New("\"isbn\" field required")
	}

	if b.Title == "" {
		return errors.New("\"title\" field required")
	}

	if b.Author == "" {
		return errors.New("\"author\" field required")
	}

	if b.Published == "" {
		return errors.New("\"published\" field required")
	}

	if b.Pages == 0 {
		return errors.New("\"pages\" field required")
	}

	return nil
}

func BookFromService(b Book) BookFullRequest {
	return BookFullRequest{
		ISBN:      b.ISBN,
		Title:     b.Title,
		Author:    b.Author,
		Published: b.Published.Format(timeLayout),
		Pages:     b.Pages,
	}
}

func (b BookFullRequest) ToService() (Book, error) {
	p, err := time.Parse(timeLayout, b.Published)
	if err != nil {
		return Book{}, errors.New("Wrong published date format (DD/MM/YYYY HH:MM)")
	}

	return Book{
		ISBN:      b.ISBN,
		Title:     b.Title,
		Author:    b.Author,
		Published: p,
		Pages:     b.Pages,
	}, nil
}

type BookUpdateRequest struct {
	Title     string `json:"title"`
	Author    string `json:"author"`
	Published string `json:"published"`
	Pages     uint   `json:"pages"`
}

func (b BookUpdateRequest) ToService() (Book, error) {
	p, err := time.Parse(timeLayout, b.Published)
	if err != nil {
		return Book{}, errors.New("Wrong published date format (DD/MM/YYYY HH:MM)")
	}

	return Book{
		Title:     b.Title,
		Author:    b.Author,
		Published: p,
		Pages:     b.Pages,
	}, nil
}

type HttpResponse struct {
	StatusCode uint        `json:"status_code"`
	Data       interface{} `json:"data"`
}

type HttpErrorResponse struct {
	ErrorCode uint   `json:"error_code"`
	Msg       string `json:"message"`
}

func GetHandler(service IService) Handler {
	return Handler{
		service: service,
	}
}

func GetSingleHandler(service IService) SingleHandler {
	return SingleHandler{
		service: service,
	}
}

func (b Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	switch r.Method {
	case "GET":
		{
			books, err := b.service.GetAll(context.TODO())
			if err != nil {
				w.Write([]byte(err.Error()))
				return
			}

			var booksFull []BookFullRequest
			for _, book := range books {
				booksFull = append(booksFull, BookFromService(book))
			}

			resp, err := json.Marshal(booksFull)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(err.Error()))
			}
			w.Write([]byte(resp))
		}
	case "POST":
		{
			req := &BookFullRequest{}
			if err := api.DecodeJSONBody(w, r, req); err != nil {
				w.Write([]byte(err.Error()))
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			if err := req.Validate(); err != nil {
				w.Write([]byte(err.Error()))
				return
			}

			book, err := req.ToService()
			if err != nil {
				w.Write([]byte(err.Error()))
				return
			}

			if _, err := b.service.Create(context.TODO(), book); err != nil {
				w.Write([]byte(err.Error()))
				return
			}

			json, _ := json.Marshal(req)
			w.Write(json)
		}
	default:
		{
			w.Write([]byte("Method not allowed"))
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	}
}

func (h SingleHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	urlPath := r.URL.RequestURI()
	bookISBN := urlPath[len("/books/"):]

	if len(bookISBN) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("/books/{isbn} path param required"))
		return
	}

	switch r.Method {
	case "GET":
		{
			book, err := h.service.Get(context.TODO(), bookISBN)
			if err != nil {
				w.Write([]byte(err.Error()))
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			json, _ := json.Marshal(BookFromService(book))
			w.Write(json)
		}
	case "PUT":
		{
			req := &BookUpdateRequest{}
			if err := api.DecodeJSONBody(w, r, req); err != nil {
				w.Write([]byte(err.Error()))
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			book, err := req.ToService()
			if err != nil {
				w.Write([]byte(err.Error()))
				return
			}

			book.ISBN = bookISBN

			book, err = h.service.Update(context.TODO(), book)
			if err != nil {
				w.Write([]byte(err.Error()))
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			json, _ := json.Marshal(BookFromService(book))
			w.Write(json)
		}
	case "POST":
		{
			w.Write([]byte("Method not allowed"))
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	case "DELETE":
		{
			if err := h.service.Delete(context.TODO(), bookISBN); err != nil {
				w.Write([]byte(err.Error()))
			}
			w.WriteHeader(http.StatusNoContent)
		}
	default:
		{
			w.Write([]byte("Method not allowed"))
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	}

}
