package domain

import (
	"errors"
	"time"
)

type Book struct {
	ID        uint      `json:"id"`
	Author    string    `json:"author" validate:"required"`
	Title     string    `json:"title"`
	Price     int       `json:"price"`
	Isbn      string    `json:"isbn"`
	Stock     int       `json:"stock"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

var ErrBookConflict = errors.New("book already exists")
var ErrBookNotFound = errors.New("book not found")
