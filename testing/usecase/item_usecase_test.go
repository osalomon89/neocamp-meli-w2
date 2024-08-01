package usecase

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/teamcubation/neocamp-meli/testing/repository"
)

//mockgen -source=./repository/item_repository.go -destination=./mocks/item_repository_mock.go -package=mocks

// SaveItem(name string, stock int) error
// GetItemByID(itemID uint) error

type itemRepositoryMock struct {
	err error
}

func (repo itemRepositoryMock) SaveItem(name string, stock int) error {
	return repo.err
}

func (repo itemRepositoryMock) GetItemByID(itemID uint) error {
	return repo.err
}

func Test_itemUsecase_CreateItem_Manual(t *testing.T) {
	type args struct {
		name  string
		stock int
	}

	tests := []struct {
		name        string
		repo        repository.ItemRepository
		args        args
		wantedError error
	}{
		{
			name: "Test with error in name",
			repo: nil,
			args: args{
				name:  "",
				stock: 5,
			},
			wantedError: fmt.Errorf("item name could not be empty"),
		},
		{
			name: "Test with error in stock",
			repo: nil,
			args: args{
				name:  "Tablet",
				stock: 0,
			},
			wantedError: fmt.Errorf("stock could not be zero"),
		},
		{
			name: "Test Ok",
			repo: itemRepositoryMock{
				err: nil,
			},
			args: args{
				name:  "Tablet",
				stock: 10,
			},
			wantedError: nil,
		},
		{
			name: "Test error in repository",
			repo: itemRepositoryMock{
				err: errors.New("mysql error"),
			},
			args: args{
				name:  "Tablet",
				stock: 10,
			},
			wantedError: fmt.Errorf("error in repository: %w",
				errors.New("mysql error")),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc := NewItemUsecase(tt.repo)

			err := svc.CreateItem(tt.args.name, tt.args.stock)
			if tt.wantedError != nil {
				assert.NotNil(t, err, "error cannot be nil")
				assert.Equal(t, tt.wantedError, err, "they should be equals")
				return
			}

			assert.Nil(t, err)
		})
	}
}

func TestCreateItemErro(t *testing.T) {
	repo := itemRepositoryMock{
		err: errors.New("mysql error"),
	}

	svc := NewItemUsecase(repo)
	err := svc.CreateItem("Tablet", 10)

	assert.NotNil(t, err, "error cannot be nil")
	assert.Equal(t, errors.New("mysql error"), err, "they should be equals")
}

// func Test_itemService_CreateItem(t *testing.T) {
// 	assert := assert.New(t)

// 	type args struct {
// 		name  string
// 		stock int
// 	}

// 	tests := []struct {
// 		name      string
// 		args      args
// 		repoError error
// 		repoTimes int
// 		wantedErr error
// 	}{
// 		{
// 			name:      "Should work correctly",
// 			wantedErr: nil,
// 			args: args{
// 				name:  "tablet",
// 				stock: 10,
// 			},
// 			repoError: nil,
// 			repoTimes: 1,
// 		},
// 		{
// 			name:      "Should return error when item name is empty",
// 			wantedErr: fmt.Errorf("item name could not be empty"),
// 			args: args{
// 				name:  "",
// 				stock: 10,
// 			},
// 			repoError: nil,
// 			repoTimes: 0,
// 		},
// 		{
// 			name:      "Should return error when item stock is zero",
// 			wantedErr: fmt.Errorf("stock could not be zero"),
// 			args: args{
// 				name:  "tablet",
// 				stock: 0,
// 			},
// 			repoError: nil,
// 			repoTimes: 0,
// 		},
// 		{
// 			name:      "Should return error when repository returns an error",
// 			wantedErr: fmt.Errorf("error in repository: %w", errors.New("the repository error")),
// 			args: args{
// 				name:  "tablet",
// 				stock: 10,
// 			},
// 			repoError: errors.New("the repository error"),
// 			repoTimes: 1,
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			ctrl := gomock.NewController(t)
// 			defer ctrl.Finish()

// 			repositoryMock := mocks.NewMockItemRepository(ctrl)

// 			repositoryMock.EXPECT().
// 				SaveItem(tt.name, tt.args.stock).
// 				Return(tt.repoError).
// 				Times(tt.repoTimes)

// 			svc := NewItemUsecase(repositoryMock)

// 			err := svc.CreateItem(tt.args.name, tt.args.stock)
// 			if tt.wantedErr != nil {
// 				assert.NotNil(err)
// 				assert.Equal(tt.wantedErr.Error(), err.Error(), "Error message is not the expected")
// 				return
// 			}

// 			assert.Nil(err)
// 		})
// 	}
// }
