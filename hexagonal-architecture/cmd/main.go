package main

import (
	"log"

	"github.com/mercadolibre/inventory/cmd/bootstrap"
)

func main() {
	if err := bootstrap.Run(); err != nil {
		log.Fatal(err)
	}
}

// const port = ":9000"

// func main() {
// 	r := gin.Default()

// 	repo := repository.NewBookRepository()
// 	usecase := usecase.NewBookUsecase(repo)
// 	ctrl := controller.NewBookController(usecase)

// 	r.GET("/ping", controller.Pong)

// 	r.GET("/api/v1/books", ctrl.GetBooks)
// 	r.POST("/api/v1/books", ctrl.AddBook)
// 	r.GET("/api/v1/books/:id", ctrl.GetBook)

// 	if err := r.Run(port); err != nil {
// 		panic(err)
// 	}
// }
