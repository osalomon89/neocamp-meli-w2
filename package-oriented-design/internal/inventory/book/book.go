package book

type Repository interface {
	Save(book *Book) error
	GetAll() ([]Book, error)
}

type Book struct {
	ID     string
	Author string
	Title  string
	Price  int
}
