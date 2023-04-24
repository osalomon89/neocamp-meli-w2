package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	ctrl "github.com/osalomon/meli-items-v2/internal/controller"
	"github.com/osalomon/meli-items-v2/internal/repository"
	"github.com/osalomon/meli-items-v2/internal/service"
)

func main() {
	r := gin.Default()

	conn, err := repository.GetConnectionDB()
	if err != nil {
		panic(err)
	}

	itemRepository := repository.NewMySQLItemRepository(conn)
	//itemUsecase := service.NewItemUsecase(itemRepository)
	itemUsecase := service.ItemService{
		Repo: itemRepository,
	}

	itemController := ctrl.NewItemController(itemUsecase)

	r.GET("v1/items", itemController.GetItems)
	r.POST("v1/items", itemController.AddItem)
	r.GET("v1/items/:id", itemController.GetItem)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.Run(":8000")
}
