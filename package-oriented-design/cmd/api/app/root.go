package app

import (
	"github.com/gin-gonic/gin"
	"github.com/teamcubation/pod/cmd/api/app/handler"
	"github.com/teamcubation/pod/internal/inventory"
)

func Build(dep *Dependencies) *gin.Engine {
	router := gin.Default()

	// use cases
	bookUseCase := inventory.NewUseCaseCRUD(dep.BookRepository)

	// controller adapters
	booksHandler := handler.NewCRUDHandler(bookUseCase)

	basePath := "/api/v1/inventory"
	r := router.Group(basePath)

	r.GET("/books", booksHandler.HandleRead)
	r.POST("/books", booksHandler.HandleCreate)

	return router
}
