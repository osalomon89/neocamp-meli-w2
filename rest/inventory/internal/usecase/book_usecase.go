package usecase

import (
	"errors"
	"fmt"

	"github.com/mercadolibre/inventory/internal/entity"
)

type BookUsecase interface {
	GetAllBooks() []entity.Book
	GetBookByID(id int) (entity.Book, error)
	AddBook(book entity.Book) (entity.Book, error)
}

type bookUsecase struct {
	repo entity.BookRepository
}

func NewBookUsecase(repo entity.BookRepository) BookUsecase {
	return &bookUsecase{
		repo: repo,
	}
}

func (u *bookUsecase) GetAllBooks() []entity.Book {
	return u.repo.GetBooks()
}

func (u *bookUsecase) GetBookByID(id int) (entity.Book, error) {
	book, err := u.repo.GetBook(uint(id))
	if err != nil {
		return entity.Book{}, fmt.Errorf("error in repository: %w", err)
	}

	return book, nil
}

func (u *bookUsecase) AddBook(book entity.Book) (entity.Book, error) {
	books := u.repo.GetBooks()
	for _, b := range books {
		if b.Code == book.Code {
			return entity.Book{}, errors.New("book already exist")
		}
	}

	err := u.repo.AddBook(&book)
	if err != nil {
		return entity.Book{}, fmt.Errorf("error in repository: %w", err)
	}

	return book, nil
}
