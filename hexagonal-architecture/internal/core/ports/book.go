package ports

import (
	"context"

	"github.com/mercadolibre/inventory/internal/core/domain"
)

type BookUsecase interface {
	GetBooks(ctx context.Context, params map[string]interface{}) ([]domain.Book, error)
	GetBookByID(ctx context.Context, id uint) (*domain.Book, error)
	AddBook(ctx context.Context, book domain.Book) (*domain.Book, error)
	UpdateBook(ctx context.Context, id uint, params map[string]interface{}) (*domain.Book, error)
}

type BookRepository interface {
	GetBooks(ctx context.Context, params map[string]interface{}) ([]domain.Book, error)
	GetBookByID(ctx context.Context, id uint) (*domain.Book, error)
	CreateBook(ctx context.Context, book *domain.Book) error
	UpdateBook(ctx context.Context, book *domain.Book, params map[string]interface{}) error
}
