package user

import "errors"

type User struct {
	ID   int
	Name string
}

type UserRepository interface {
	GetByID(id int) (*User, error)
}

type UserService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetUserName(id int) (string, error) {
	if id == 0 {
		return "", errors.New("error")
	}

	user, err := s.repo.GetByID(id)
	if err != nil {
		return "", err
	}
	return user.Name, nil
}
