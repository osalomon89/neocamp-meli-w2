package main

import (
	ctrl "github.com/teamcubation/neocamp-meli/clean-architecture/internal/adapter/controller"
	repo "github.com/teamcubation/neocamp-meli/clean-architecture/internal/adapter/repository"
	"github.com/teamcubation/neocamp-meli/clean-architecture/internal/infra/db"
	"github.com/teamcubation/neocamp-meli/clean-architecture/internal/infra/web"
	"github.com/teamcubation/neocamp-meli/clean-architecture/internal/usecase"
)

const port = "9000"

func main() {
	conn := db.GetConnectionDB()
	repository := repo.NewBookRepository(conn)
	usecase := usecase.NewBookUsecase(repository)
	controller := ctrl.NewBookController(usecase)
	server := web.NewHTTPServer(controller)

	server.RegisterRouter()

	if err := server.Run(port); err != nil {
		panic(err.Error())
	}
}
