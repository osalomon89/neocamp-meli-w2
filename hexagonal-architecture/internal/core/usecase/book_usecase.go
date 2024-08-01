package usecase

import (
	"context"
	"fmt"

	"github.com/mercadolibre/inventory/internal/core/domain"
	"github.com/mercadolibre/inventory/internal/core/ports"
)

type bookUsecase struct {
	repo ports.BookRepository
}

func NewBookUsecase(repo ports.BookRepository) *bookUsecase {
	return &bookUsecase{
		repo: repo,
	}
}

func (u *bookUsecase) GetBooks(ctx context.Context, params map[string]interface{}) ([]domain.Book, error) {
	return u.repo.GetBooks(ctx, params)
}

func (u *bookUsecase) GetBookByID(ctx context.Context, id uint) (*domain.Book, error) {
	result, err := u.getBookByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if result == nil {
		return nil, domain.ErrBookNotFound
	}

	return result, nil
}

func (u *bookUsecase) AddBook(ctx context.Context, book domain.Book) (*domain.Book, error) {
	result, err := u.getBookByID(ctx, book.ID)
	if err != nil {
		return nil, err
	}

	if result != nil {
		return nil, domain.ErrBookConflict
	}

	if err := u.repo.CreateBook(ctx, &book); err != nil {
		return nil, fmt.Errorf("error creating book: %w", err)
	}

	return &book, nil
}

func (u *bookUsecase) UpdateBook(ctx context.Context, id uint, params map[string]interface{}) (*domain.Book, error) {
	book, err := u.getBookByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if book == nil {
		return nil, domain.ErrBookNotFound
	}

	if err := u.repo.UpdateBook(ctx, book, params); err != nil {
		return nil, fmt.Errorf("error updating book: %w", err)
	}

	return book, nil
}

func (u *bookUsecase) getBookByID(ctx context.Context, id uint) (*domain.Book, error) {
	result, err := u.repo.GetBookByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("error getting book: %w", err)
	}

	return result, nil
}
