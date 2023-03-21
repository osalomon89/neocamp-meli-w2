package main

import (
	"github.com/teamcubation/neocamp-meli/clean-architecture/internal/domain"
	"github.com/teamcubation/neocamp-meli/clean-architecture/internal/infrastructure/http"
	ctrl "github.com/teamcubation/neocamp-meli/clean-architecture/internal/infrastructure/http/controller"
	repo "github.com/teamcubation/neocamp-meli/clean-architecture/internal/infrastructure/repository/memory"
	"github.com/teamcubation/neocamp-meli/clean-architecture/internal/usecase"
)

const port = "9000"

var db []domain.Book = []domain.Book{
	{
		ID:     1,
		Title:  "Dune",
		Price:  1965,
		Author: "Frank Herbert",
	},
	{
		ID:     2,
		Title:  "Cita con Rama",
		Price:  1974,
		Author: "Arthur C. Clarke",
	},
	{
		ID:     3,
		Title:  "Un guijarro en el cielo",
		Price:  500,
		Author: "Isaac Asimov",
	},
}

func main() {
	repository := repo.NewBookRepository(db)
	usecase := usecase.NewBookUsecase(repository)
	handler := ctrl.NewBookHandler(usecase)
	server := http.NewHTTPServer(handler)

	server.RegisterRouter()

	if err := server.Run(port); err != nil {
		panic(err.Error())
	}
}
