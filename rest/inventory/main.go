package main

import (
	"log"

	"github.com/gin-gonic/gin"
	ctrl "github.com/mercadolibre/inventory/internal/infrastructure/controller"
	"github.com/mercadolibre/inventory/internal/infrastructure/repository"
	"github.com/mercadolibre/inventory/internal/usecase"
)

const port = ":9000"

func main() {
	r := gin.Default()

	boorRepository := repository.NewBookRepository()
	bookUsecase := usecase.NewBookUsecase(boorRepository)
	bookController := ctrl.NewBookController(bookUsecase)

	r.GET("/ping", ctrl.Pong)

	r.GET("/api/v1/books", bookController.GetBooks)
	r.POST("/api/v1/books", bookController.AddBook)
	r.GET("/api/v1/books/:id", bookController.GetBook)

	r.Run(port)

	if err := r.Run(port); err != nil {
		log.Fatalln(err)
	}
}
