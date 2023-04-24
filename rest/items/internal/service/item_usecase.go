package service

import (
	"github.com/osalomon/meli-items-v2/internal/entity"
	"github.com/osalomon/meli-items-v2/internal/repository"
)

type IItemService interface {
	GetAllItems() ([]entity.Item, error)
	GetItemByID(id int) (entity.Item, error)
	AddItem(item entity.Item) (entity.Item, error)
}

type ItemService struct {
	Repo repository.ItemRepository
}

func (u ItemService) GetAllItems() ([]entity.Item, error) {
	return u.Repo.GetItems()
}

func (u ItemService) GetItemByID(id int) (entity.Item, error) {
	return u.Repo.GetItemByID(uint(id))
}

func (u ItemService) AddItem(item entity.Item) (entity.Item, error) {
	err := u.Repo.AddItem(&item)
	return item, err
}
