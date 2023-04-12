package usecase

import (
	"errors"

	"github.com/mercadolibre/inventory/internal/domain"
)

type BookUsecase interface {
	GetAllBooks() []domain.Book
	GetBookByID(id int) *domain.Book
	AddBook(book domain.Book) (*domain.Book, error)
}

type bookUsecase struct {
	repo domain.BookRepository
}

func NewBookUsecase(repo domain.BookRepository) BookUsecase {
	return &bookUsecase{
		repo: repo,
	}
}

func (u *bookUsecase) GetAllBooks() []domain.Book {
	return u.repo.GetBooks()
}

func (u *bookUsecase) GetBookByID(id int) *domain.Book {
	return u.repo.GetBook(id)
}

func (u *bookUsecase) AddBook(book domain.Book) (*domain.Book, error) {
	books := u.repo.GetBooks()
	for _, b := range books {
		if b.ID == book.ID {
			return nil, errors.New("book already exist")
		}
	}

	result := u.repo.AddBook(book)

	return result, nil
}
