package app

import (
	"github.com/teamcubation/pod/internal/inventory/book"
	"github.com/teamcubation/pod/internal/platform/environment"
	"github.com/teamcubation/pod/internal/platform/memorydb"
)

type Dependencies struct {
	BookRepository book.Repository
}

func BuildDependencies(env environment.Environment) (*Dependencies, error) {
	localDb := memorydb.New()

	devRepo := book.NewLocalRepo(localDb)

	return &Dependencies{
		BookRepository: devRepo,
	}, nil
}
