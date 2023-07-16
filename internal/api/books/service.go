package books

import (
	"context"
	"time"
    "fmt"
)

type Book struct {
	ISBN      string
	Title     string
	Author    string
	Published time.Time
	Pages     uint
}

type IService interface {
	Create(ctx context.Context, book Book) (Book, error)
	Update(ctx context.Context, book Book) (Book, error)
	Delete(ctx context.Context, isbn string) error
	Get(ctx context.Context, isbn string) (Book, error)
}

type Service struct {
	bookRepository IRepository
}

func GetService(b IRepository) Service {
	return Service{
		bookRepository: b,
	}
}

func (s Service) Create(ctx context.Context, book Book) (Book, error) {

	dbModel := BookDB{
		ISBN:   book.ISBN,
		Title:  book.Title,
		Pages:  book.Pages,
		Author: book.Author,
	}
    dbModel.SetPublished(book.Published)

    fmt.Println("Probably repository called")
    _, err := s.bookRepository.Create(ctx, dbModel)
	if err != nil {
		return book, err
	}

	return book, nil
}

func (s Service) Update(ctx context.Context, book Book) (Book, error) {
	return book, nil
}

func (s Service) Delete(ctx context.Context, isbn string) error {
	return nil
}

func (s Service) Get(ctx context.Context, isbn string) (Book, error) {
	book := Book{}
	return book, nil
}
