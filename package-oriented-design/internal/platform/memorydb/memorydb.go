package memorydb

type Book struct {
	ID     string
	Author string
	Title  string
	Price  int
}

var db []Book = []Book{
	{
		ID:     "1",
		Title:  "Dune",
		Price:  1965,
		Author: "Frank Herbert",
	},
	{
		ID:     "2",
		Title:  "Cita con Rama",
		Price:  1974,
		Author: "Arthur C. Clarke",
	},
	{
		ID:     "3",
		Title:  "Un guijarro en el cielo",
		Price:  500,
		Author: "Isaac Asimov",
	},
}

type LocalDB struct {
	storage []Book
}

func New() *LocalDB {
	return &LocalDB{storage: db}
}

func (db *LocalDB) SaveItem(item Book) {
	db.storage = append(db.storage, item)
}

func (db *LocalDB) GetAll() []Book {
	return db.storage
}
