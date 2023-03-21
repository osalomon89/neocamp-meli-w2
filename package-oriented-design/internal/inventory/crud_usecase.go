package inventory

import (
	"fmt"

	"github.com/teamcubation/pod/internal/inventory/book"
)

type UseCaseCRUD interface {
	GetAllBooks() ([]book.Book, error)
	SaveBook(author, title string, price int) (book.Book, error)
}

type useCaseCRUD struct {
	repo book.Repository
}

func NewUseCaseCRUD(repo book.Repository) *useCaseCRUD {
	return &useCaseCRUD{repo: repo}
}

func (c useCaseCRUD) SaveBook(author, title string, price int) (book.Book, error) {
	dev := book.Book{
		Author: author,
		Title:  title,
		Price:  price,
	}

	if err := c.repo.Save(&dev); err != nil {
		return dev, err
	}

	return dev, nil
}

func (c useCaseCRUD) GetAllBooks() ([]book.Book, error) {
	dev, err := c.repo.GetAll()
	if err != nil {
		return nil, err
	}

	if dev == nil {
		return nil, fmt.Errorf("books not found")
	}

	return dev, nil
}
