package books

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/blessium/metricsgo/internal/api"
)

type Handler struct {
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

func (b BookFullRequest) ToService() (Book, error) {
	p, err := time.Parse("01/02/2006 15:05", b.Published)
	if err != nil {
		return Book{}, errors.New("Wrong published date format (DD/MM/YY HH:MM)")
	}

	return Book{
		ISBN:      b.ISBN,
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

func (b Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	switch r.Method {
	case "GET":
		{
			w.Write([]byte("Ciao"))
		}
		// TODO: finish POST Request
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
			w.Write([]byte("Message not allowed"))
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	}
}
