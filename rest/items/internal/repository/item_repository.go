package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/osalomon/meli-items-v2/internal/entity"
)

type ItemRepository interface {
	GetItems() ([]entity.Item, error)
	GetItemByID(id uint) (entity.Item, error)
	AddItem(item *entity.Item) error
}

type mysqlItemRepository struct {
	conn *sqlx.DB
}

func NewMySQLItemRepository(db *sqlx.DB) *mysqlItemRepository {
	return &mysqlItemRepository{
		conn: db,
	}
}

func (r *mysqlItemRepository) GetItemByID(id uint) (entity.Item, error) {
	return entity.Item{}, nil
}

func (r *mysqlItemRepository) AddItem(item *entity.Item) error {
	return nil
}

func (r *mysqlItemRepository) GetItems() ([]entity.Item, error) {
	return nil, nil
}
