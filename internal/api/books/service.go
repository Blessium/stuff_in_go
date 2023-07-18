package books

import (
	"context"
	"time"
)

type Book struct {
	ISBN      string
	Title     string
	Author    string
	Published time.Time
	Pages     uint
}

func (b Book) ToDB() BookDB {
	dbModel := BookDB{
		ISBN:   b.ISBN,
		Title:  b.Title,
		Pages:  b.Pages,
		Author: b.Author,
	}
	dbModel.SetPublished(b.Published)
	return dbModel
}

func BookFromDB(b BookDB) Book {
	return Book{
		ISBN:      b.ISBN,
		Title:     b.Title,
		Pages:     b.Pages,
		Author:    b.Author,
		Published: b.Published.Time(),
	}
}

type IService interface {
	Create(ctx context.Context, book Book) (Book, error)
	Update(ctx context.Context, book Book) (Book, error)
	Delete(ctx context.Context, isbn string) error
	Get(ctx context.Context, isbn string) (Book, error)
	GetAll(ctx context.Context) ([]Book, error)
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
	dbModel := book.ToDB()
	_, err := s.bookRepository.Create(ctx, dbModel)
	if err != nil {
		return book, err
	}

	return book, nil
}

func (s Service) Update(ctx context.Context, book Book) (Book, error) {
	dbModel := book.ToDB()
	_, err := s.bookRepository.Update(ctx, dbModel)
	if err != nil {
		return book, err
	}
	return book, nil
}

func (s Service) Delete(ctx context.Context, isbn string) error {
	err := s.bookRepository.Delete(ctx, isbn)
	if err != nil {
		return err
	}
	return nil
}

func (s Service) Get(ctx context.Context, isbn string) (Book, error) {
	book := Book{}
	modelDB, err := s.bookRepository.Get(ctx, isbn)
	if err != nil {
		return book, err
	}
	return BookFromDB(modelDB), nil
}

func(s Service) GetAll(ctx context.Context) ([]Book, error) {
    booksDB, err := s.bookRepository.GetAll(ctx)
    if err != nil {
        return []Book{}, err
    }

    var books []Book
    for _, bookDB := range booksDB {
        books = append(books, BookFromDB(bookDB))
    }

    return books, nil
}
