package memory

import (
	"github.com/teamcubation/neocamp-meli/clean-architecture/internal/domain"
)

type bookRepository struct {
	db []domain.Book
}

func NewBookRepository(db []domain.Book) domain.BookRepository {
	return &bookRepository{
		db: db,
	}
}

func (r *bookRepository) GetAllBooks() []domain.Book {
	return r.db
}

func (r *bookRepository) SaveBook(book domain.Book) domain.Book {
	r.db = append(r.db, book)

	return book
}
