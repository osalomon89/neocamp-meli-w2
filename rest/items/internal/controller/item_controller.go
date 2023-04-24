package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/osalomon/meli-items-v2/internal/entity"
	"github.com/osalomon/meli-items-v2/internal/service"
)

type ItemController struct {
	itemUsecase service.IItemService
}

func NewItemController(itemUsecase service.IItemService) ItemController {
	return ItemController{
		itemUsecase: itemUsecase,
	}
}

func Pong(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"error": false,
		"data":  "pong",
	})
}

func (ctrl ItemController) GetItems(c *gin.Context) {
	books, err := ctrl.itemUsecase.GetAllItems()
	c.JSON(http.StatusInternalServerError, gin.H{
		"StatusCode": http.StatusInternalServerError,
		"Message":    fmt.Sprintf("error getting books: %s", err.Error()),
	})

	c.JSON(http.StatusOK, gin.H{
		"Error": false,
		"Data":  books,
	})
}

type bookRequestDTO struct {
	Code   string `json:"code" binding:"required"`
	Author string `json:"author" binding:"required"`
	Title  string `json:"title" binding:"required"`
	Price  int    `json:"price" binding:"required"`
	Stock  int    `json:"stock" binding:"required"`
}

func (ctrl ItemController) AddItem(c *gin.Context) {
	var bookRequest bookRequestDTO

	if err := c.ShouldBindJSON(&bookRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"StatusCode": http.StatusBadRequest,
			"Message":    fmt.Sprintf("invalid json: %s", err.Error()),
		})
		return
	}

	book := entity.Item{
		Code:  bookRequest.Code,
		Title: bookRequest.Title,
		Price: bookRequest.Price,
		Stock: bookRequest.Stock,
	}

	result, err := ctrl.itemUsecase.AddItem(book)
	if err != nil {
		c.JSON(500, gin.H{
			"StatusCode": 500,
			"Message":    err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Error": false,
		"Data":  result,
	})
}

func (ctrl *ItemController) GetItem(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"StatusCode": http.StatusBadRequest,
			"Message":    fmt.Sprintf("invalid param: %s", err.Error()),
		})
		return
	}

	book, err := ctrl.itemUsecase.GetItemByID(id)
	if err != nil {
		c.JSON(500, gin.H{
			"StatusCode": 500,
			"Message":    err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Error": false,
		"Data":  book,
	})
}
