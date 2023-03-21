package services

import (
	"fmt"

	"github.com/osalomon89/go-testing/repositories"
)

type UserService interface {
	CreateUser(name, lastname string, age int) error
}

type userService struct {
	repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) UserService {
	return &userService{
		repo: repo,
	}
}

func (svc *userService) CreateUser(name, lastname string, age int) error {
	if name == "" || lastname == "" {
		return fmt.Errorf("name or lastname could not be empty")
	}

	if age < 18 {
		return fmt.Errorf("age cannot be less than 18")
	}

	err := svc.repo.SaveUser(name, lastname, age)
	if err != nil {
		return fmt.Errorf("error in repository: %w", err)
	}

	return nil
}

func (svc *userService) GetUserByID(id uint) error {
	if id == 0 {
		return fmt.Errorf("user id cannot be zero")
	}

	err := svc.repo.GetUserByID(id)
	if err != nil {
		return fmt.Errorf("error in repository: %w", err)
	}

	return nil
}
