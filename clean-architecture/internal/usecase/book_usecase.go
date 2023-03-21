package usecase

import (
	"github.com/teamcubation/neocamp-meli/clean-architecture/internal/domain"
)

type BookUsecase interface {
	GetAllBooks() []domain.Book
	SaveBook(book domain.Book) domain.Book
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
	return u.repo.GetAllBooks()
}

func (u *bookUsecase) SaveBook(book domain.Book) domain.Book {
	return u.repo.SaveBook(book)
}
