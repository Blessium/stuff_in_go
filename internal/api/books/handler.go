package books

import (
	"net/http"
)

type Handler struct {
	service IService
}

type BookCreateRequest struct {
	ISBN      string `json:"isbn"`
	Title     string `json:"title"`
	Author    string `json:"author"`
	Published string `json:"published"`
	Pages     uint   `json:"pages"`
}

func GetHandler(service IService) Handler {
	return Handler{
		service: service,
	}
}

func (b Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":
		{
			w.Write([]byte("Ciao"))
		}
		// TODO: finish POST Request
	case "POST":
		{
			w.Write([]byte("Ciao"))
		}
	default:
		{
			w.Write([]byte("Message not allowed"))
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	}
}
