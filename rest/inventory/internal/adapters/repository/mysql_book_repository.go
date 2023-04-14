package repository

import (
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/mercadolibre/inventory/internal/entity"
)

type bookDAO struct {
	ID        uint
	Code      string
	Author    string
	Title     string
	Price     int
	Stock     int
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

func (b bookDAO) toBookDomain() entity.Book {
	return entity.Book{
		ID:        b.ID,
		Code:      b.Code,
		Author:    b.Author,
		Title:     b.Title,
		Price:     b.Price,
		Stock:     b.Stock,
		CreatedAt: b.CreatedAt,
		UpdatedAt: b.UpdatedAt,
	}
}

type mysqlBookRepository struct {
	conn *sqlx.DB
}

func NewMySQLBookRepository(db *sqlx.DB) entity.BookRepository {
	return &mysqlBookRepository{
		conn: db,
	}
}

func (repo *mysqlBookRepository) GetBook(id uint) (entity.Book, error) {
	var book entity.Book

	bookDB := new(bookDAO)
	err := repo.conn.Get(bookDB, "SELECT * FROM books WHERE id=?", id)
	if err != nil {
		return book, fmt.Errorf("error getting book: %w", err)
	}

	return bookDB.toBookDomain(), nil
}

func (repo *mysqlBookRepository) AddBook(book *entity.Book) error {
	createdAt := time.Now()

	result, err := repo.conn.Exec(`INSERT INTO books 
		(code, title, author, price, stock, created_at, updated_at) 
		VALUES(?,?,?,?,?,?,?)`, book.Code, book.Title, book.Author, book.Price, book.Stock, createdAt, createdAt)

	if err != nil {
		return fmt.Errorf("error inserting book: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("error saving book: %w", err)
	}

	book.ID = uint(id)
	book.CreatedAt = createdAt
	book.CreatedAt = createdAt

	return nil
}

func (r *mysqlBookRepository) GetBooks() []entity.Book {
	return nil
}
