package book

import (
	"github.com/google/uuid"
	"github.com/teamcubation/pod/internal/platform/memorydb"
)

type localRepo struct {
	db *memorydb.LocalDB
}

func NewLocalRepo(db *memorydb.LocalDB) Repository {
	return &localRepo{db: db}
}

func (l localRepo) Save(d *Book) error {
	if d.ID == "" {
		id := uuid.New()
		d.ID = id.String()
	}

	book := memorydb.Book{
		ID:     d.ID,
		Author: d.Author,
		Title:  d.Title,
		Price:  d.Price,
	}

	l.db.SaveItem(book)
	return nil
}

func (l localRepo) GetAll() ([]Book, error) {
	var bookSlice []Book

	items := l.db.GetAll()

	for _, b := range items {
		bookSlice = append(bookSlice, Book{
			ID:     b.ID,
			Author: b.Author,
			Title:  b.Title,
			Price:  b.Price,
		})
	}

	return bookSlice, nil
}
