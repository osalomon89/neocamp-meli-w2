package book

import (
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
)

type bookDAO struct {
	ID        uint
	Author    string
	Title     string
	Price     int
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type mysqlRepo struct {
	mysql *sqlx.DB
}

func NewMySQLRepo(conn *sqlx.DB) Repository {
	return &mysqlRepo{
		mysql: conn,
	}
}

func (r *mysqlRepo) Save(b *Book) error {
	createdAt := time.Now()
	result, err := r.mysql.Exec(`INSERT INTO books 
		(author, title, price, created_at, updated_at) 
		VALUES(?,?,?,?,?)`, b.Author, b.Title, b.Price, createdAt, createdAt)

	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	b.ID = fmt.Sprint(id)

	return nil
}

func (r *mysqlRepo) GetAll() ([]Book, error) {
	var bookSlice []Book
	var books []bookDAO

	err := r.mysql.Select(&books, `SELECT * FROM books`)
	if err != nil {
		return bookSlice, err
	}

	for _, b := range books {
		bookSlice = append(bookSlice, Book{
			ID:     fmt.Sprint(b.ID),
			Author: b.Author,
			Title:  b.Title,
			Price:  b.Price,
		})
	}

	return bookSlice, nil
}
