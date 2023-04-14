package entity

import "time"

type Book struct {
	ID        uint
	Code      string
	Author    string
	Title     string
	Price     int
	Stock     int
	CreatedAt time.Time
	UpdatedAt time.Time
}

type BookRepository interface {
	GetBooks() []Book
	GetBook(id uint) (Book, error)
	AddBook(book *Book) error
}
