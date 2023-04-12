package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mercadolibre/inventory/internal/domain"
	"github.com/mercadolibre/inventory/internal/usecase"
)

type BookController struct {
	bookUsecase usecase.BookUsecase
}

type responseInfo struct {
	Error bool        `json:"error"`
	Data  interface{} `json:"data"`
}

func NewBookController(usecase usecase.BookUsecase) *BookController {
	return &BookController{
		bookUsecase: usecase,
	}
}

func Pong(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"error": false,
		"data":  "pong",
	})
}

func (ctrl *BookController) GetBooks(c *gin.Context) {
	c.JSON(http.StatusOK, responseInfo{
		Error: false,
		Data:  ctrl.bookUsecase.GetAllBooks(),
	})
}

func (ctrl *BookController) GetBook(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, responseInfo{
			Error: true,
			Data:  fmt.Sprintf("invalid param: %s", err.Error()),
		})
		return
	}

	book := ctrl.bookUsecase.GetBookByID(id)
	if book == nil {
		c.JSON(http.StatusNotFound, responseInfo{
			Error: true,
			Data:  "book not found",
		})
	}

	c.JSON(http.StatusOK, responseInfo{
		Error: false,
		Data:  book,
	})
}

func (ctrl *BookController) AddBook(c *gin.Context) {
	request := c.Request
	body := request.Body

	var book domain.Book
	err := json.NewDecoder(body).Decode(&book)
	if err != nil {
		c.JSON(http.StatusBadRequest, responseInfo{
			Error: true,
			Data:  fmt.Sprintf("invalid json: %s", err.Error()),
		})
		return
	}

	result, err := ctrl.bookUsecase.AddBook(book)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responseInfo{
			Error: true,
			Data:  fmt.Sprintf("error saving book: %s", err.Error()),
		})
		return
	}

	c.JSON(http.StatusOK, responseInfo{
		Error: false,
		Data:  result,
	})
}
