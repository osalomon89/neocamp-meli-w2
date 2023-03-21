package domain

type BookRepository interface {
	GetAllBooks() []Book
	SaveBook(book Book) Book
}
