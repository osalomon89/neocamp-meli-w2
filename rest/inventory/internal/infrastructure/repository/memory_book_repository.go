package repository

import "github.com/mercadolibre/inventory/internal/domain"

type bookRepository struct {
	db []domain.Book
}

func NewBookRepository() domain.BookRepository {
	return &bookRepository{}
}

func (r *bookRepository) GetBooks() []domain.Book {
	return r.db
}

func (r *bookRepository) GetBook(id int) *domain.Book {
	for _, book := range r.db {
		if book.ID == id {
			return &book
		}
	}

	return nil
}

func (r *bookRepository) AddBook(book domain.Book) *domain.Book {
	r.db = append(r.db, book)

	return &book
}
