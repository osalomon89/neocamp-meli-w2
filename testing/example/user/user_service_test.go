package user_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/teamcubation/neocamp-meli/testing/example/user"
	"github.com/teamcubation/neocamp-meli/testing/example/user/mocks"
)

type MockUserRepository struct{}

func (m *MockUserRepository) GetByID(id int) (*user.User, error) {
	if id == 99 {
		return nil, fmt.Errorf("user not found")
	}

	user := &user.User{
		ID:   id,
		Name: "Alice",
	}

	return user, nil
}

func TestGetUserNameEasy(t *testing.T) {
	// Configurar o mock
	mockRepo := &MockUserRepository{}

	// Criar o serviço com o mock
	service := user.NewUserService(mockRepo)

	// Teste para um usuário existente
	name, err := service.GetUserName(1)
	assert.Nil(t, err)
	assert.Equal(t, "Alice", name)
}

func TestGetUserNameErrorID(t *testing.T) {
	// Configurar o mock
	mockRepo := &MockUserRepository{}

	// Criar o serviço com o mock
	service := user.NewUserService(mockRepo)

	// Teste para um usuário existente
	name, err := service.GetUserName(0)
	assert.NotNil(t, err)
	assert.Equal(t, "", name)
	assert.Equal(t, err.Error(), "error")
}

func TestGetUserNameErrorMock(t *testing.T) {
	// Configurar o mock
	mockRepo := &MockUserRepository{}

	// Criar o serviço com o mock
	service := user.NewUserService(mockRepo)

	// Teste para um usuário existente
	name, err := service.GetUserName(99)
	assert.NotNil(t, err)
	assert.Equal(t, "", name)
	assert.Equal(t, err.Error(), "user not found")
}

type MyMockedRepository struct {
	mock.Mock
}

func (m *MyMockedRepository) GetByID(id int) (*user.User, error) {
	args := m.Called(id)
	if args.Get(0) != nil {
		return args.Get(0).(*user.User), args.Error(1)
	}
	return nil, args.Error(1)
}

func TestGetUserNameWithTestifyMocking(t *testing.T) {
	t.Run("Success case", func(t *testing.T) {
		repoMock := new(MyMockedRepository)

		// set up expectations
		repoMock.On("GetByID", 123).Return(&user.User{
			ID:   123,
			Name: "Alice",
		}, nil)

		service := user.NewUserService(repoMock)

		name, err := service.GetUserName(123)
		assert.Nil(t, err)
		assert.Equal(t, "Alice", name)

		repoMock.AssertExpectations(t)
		repoMock.AssertCalled(t, "GetByID", 123)
	})

	t.Run("Error case", func(t *testing.T) {
		repoMock := new(MyMockedRepository)

		// Configura las expectativas
		expectedError := errors.New("user not found")
		repoMock.On("GetByID", 123).Return(nil, expectedError)

		service := user.NewUserService(repoMock)

		name, err := service.GetUserName(123)
		assert.NotNil(t, err)
		assert.Equal(t, "", name)
		assert.Equal(t, expectedError, err)

		// Verifica que las expectativas se hayan cumplido
		repoMock.AssertExpectations(t)
	})
}

func TestGetUserNameWithMockery(t *testing.T) {
	t.Run("Success case", func(t *testing.T) {
		repoMock := new(mocks.UserRepository)

		// set up expectations
		repoMock.On("GetByID", 123).Return(&user.User{
			ID:   123,
			Name: "Alice",
		}, nil)

		service := user.NewUserService(repoMock)

		name, err := service.GetUserName(123)
		assert.Nil(t, err)
		assert.Equal(t, "Alice", name)

		repoMock.AssertExpectations(t)
		repoMock.AssertCalled(t, "GetByID", 123)
	})

	t.Run("Error case", func(t *testing.T) {
		repoMock := new(mocks.UserRepository)

		// Configura las expectativas
		expectedError := errors.New("user not found")
		repoMock.On("GetByID", 123).Return(nil, expectedError)

		service := user.NewUserService(repoMock)

		name, err := service.GetUserName(123)
		assert.NotNil(t, err)
		assert.Equal(t, "", name)
		assert.Equal(t, expectedError, err)

		// Verifica que las expectativas se hayan cumplido
		repoMock.AssertExpectations(t)
	})
}
