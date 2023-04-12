package domain

type Book struct {
	ID     int    `json:"id"`
	Author string `json:"author"`
	Title  string `json:"title"`
	Price  int    `json:"price"`
}

type BookRepository interface {
	GetBooks() []Book
	GetBook(id int) *Book
	AddBook(book Book) *Book
}
