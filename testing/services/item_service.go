package services

import (
	"fmt"

	"github.com/osalomon89/go-testing/repositories"
)

type ItemService interface {
	CreateItem(name string, stock int) error
}

type itemService struct {
	repo repositories.ItemRepository
}

func NewItemService(repo repositories.ItemRepository) ItemService {
	return &itemService{
		repo: repo,
	}
}

func (svc *itemService) CreateItem(name string, stock int) error {
	if name == "" {
		return fmt.Errorf("item name could not be empty")
	}

	if stock == 0 {
		return fmt.Errorf("stock could not be zero")
	}

	err := svc.repo.SaveItem(name, stock)
	if err != nil {
		return fmt.Errorf("error in repository: %w", err)
	}

	return nil
}

func (svc *itemService) GetItemByID(id uint) error {
	if id == 0 {
		return fmt.Errorf("item id cannot be zero")
	}

	err := svc.repo.GetItemByID(id)
	if err != nil {
		return fmt.Errorf("error in repository: %w", err)
	}

	return nil
}
