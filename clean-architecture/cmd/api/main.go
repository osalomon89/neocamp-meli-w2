package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mercadolibre/inventory/internal/infrastructure/controller"
	"github.com/mercadolibre/inventory/internal/infrastructure/repository"
	"github.com/mercadolibre/inventory/internal/usecase"
)

const port = ":9000"

func main() {
	r := gin.Default()

	repo := repository.NewBookRepository()
	usecase := usecase.NewBookUsecase(repo)
	ctrl := controller.NewBookController(usecase)

	r.GET("/ping", controller.Pong)

	r.GET("/api/v1/books", ctrl.GetBooks)
	r.POST("/api/v1/books", ctrl.AddBook)
	r.GET("/api/v1/books/:id", ctrl.GetBook)

	if err := r.Run(port); err != nil {
		panic(err)
	}
}
