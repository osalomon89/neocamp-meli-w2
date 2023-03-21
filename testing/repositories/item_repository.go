package repositories

//go:generate mockgen -source=./item_repository.go -destination=../mocks/item_repository_mock.go -package=mocks
type ItemRepository interface {
	SaveItem(name string, stock int) error
	GetItemByID(itemID uint) error
}

/*type itemRepository struct {
	conn *sql.DB
}

func NewItemRepository(db *sql.DB) ItemRepository {
	return &itemRepository{
		conn: db,
	}
}

func (repo *itemRepository) GetItemByID(id uint) error {
	row := repo.conn.QueryRow("SELECT * FROM items WHERE id=?", id)
	if row.Err() != nil {
		return fmt.Errorf("error getting item: %w", row.Err())
	}

	return nil
}

func (repo *itemRepository) SaveItem(name string, stock int) error {
	createdAt := time.Now()

	_, err := repo.conn.Exec(`INSERT INTO items
		(name, stock, created_at, updated_at)
		VALUES(?,?,?,?)`, name, stock, createdAt, createdAt)

	if err != nil {
		return fmt.Errorf("error saving item: %w", err)
	}

	return nil
}*/
