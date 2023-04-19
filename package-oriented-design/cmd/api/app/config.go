package app

import (
	"fmt"

	"github.com/teamcubation/pod/internal/inventory/book"
	"github.com/teamcubation/pod/internal/platform/environment"
	"github.com/teamcubation/pod/internal/platform/mysql"
)

type Dependencies struct {
	BookRepository book.Repository
}

func BuildDependencies(env environment.Environment) (*Dependencies, error) {
	//localDb := memorydb.New()
	mysqlConn, err := mysql.GetConnectionDB()
	if err != nil {
		return nil, fmt.Errorf("error connecting to DB: %w", err)
	}

	devRepo := book.NewMySQLRepo(mysqlConn)
	//devRepo := book.NewLocalRepo(mysqlConn)

	return &Dependencies{
		BookRepository: devRepo,
	}, nil
}
