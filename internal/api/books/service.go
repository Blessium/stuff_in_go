package books

import (
	"context"
)

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
    return book, nil
}

func (s Service) Update(ctx context.Context, book Book) (Book, error) {
    return book, nil
}

func (s Service) Delete(ctx context.Context, isbn string) error {
    return nil
}

func (s Service) Get(ctx context.Context, isbn string) (Book, error) {
    book := Book {}
    return book, nil
}
