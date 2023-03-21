package repositories

//go:generate mockgen -source=./user_repository.go -destination=../mocks/user_repository_mock.go -package=mocks
type UserRepository interface {
	SaveUser(name, lastname string, age int) error
	GetUserByID(itemID uint) error
}
