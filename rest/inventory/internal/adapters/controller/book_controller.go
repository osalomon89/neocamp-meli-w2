package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mercadolibre/inventory/internal/adapters/controller/presenter"
	"github.com/mercadolibre/inventory/internal/entity"
	"github.com/mercadolibre/inventory/internal/usecase"
)

type BookController struct {
	bookUsecase usecase.BookUsecase
}

func NewBookController(bookUsecase usecase.BookUsecase) BookController {
	return BookController{
		bookUsecase: bookUsecase,
	}
}

func Pong(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"error": false,
		"data":  "pong",
	})
}

func (ctrl BookController) GetBooks(c *gin.Context) {
	c.JSON(http.StatusOK, presenter.BooksResponse{
		Error: false,
		Data:  presenter.Books(ctrl.bookUsecase.GetAllBooks()),
	})
}

type bookRequestDTO struct {
	Code   string `json:"code" binding:"required"`
	Author string `json:"author" binding:"required"`
	Title  string `json:"title" binding:"required"`
	Price  int    `json:"price" binding:"required"`
	Stock  int    `json:"stock" binding:"required"`
}

func (ctrl BookController) AddBook(c *gin.Context) {
	var bookRequest bookRequestDTO

	if err := c.ShouldBindJSON(&bookRequest); err != nil {
		c.JSON(http.StatusBadRequest, presenter.ApiError{
			StatusCode: http.StatusBadRequest,
			Message:    fmt.Sprintf("invalid json: %s", err.Error()),
		})
		return
	}

	book := entity.Book{
		Code:   bookRequest.Code,
		Author: bookRequest.Author,
		Title:  bookRequest.Title,
		Price:  bookRequest.Price,
		Stock:  bookRequest.Stock,
	}

	result, err := ctrl.bookUsecase.AddBook(book)
	if err != nil {
		c.JSON(http.StatusInternalServerError, presenter.ApiError{
			StatusCode: http.StatusInternalServerError,
			Message:    fmt.Sprintf("error saving book: %s", err.Error()),
		})
		return
	}

	c.JSON(http.StatusOK, presenter.BookResponse{
		Error: false,
		Data:  presenter.Book(result),
	})
}

func (ctrl *BookController) GetBook(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, presenter.ApiError{
			StatusCode: http.StatusBadRequest,
			Message:    fmt.Sprintf("invalid param: %s", err.Error()),
		})
		return
	}

	book, err := ctrl.bookUsecase.GetBookByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, presenter.ApiError{
			StatusCode: http.StatusInternalServerError,
			Message:    fmt.Errorf("error getting book: %w", err).Error(),
		})
	}

	c.JSON(http.StatusOK, presenter.BookResponse{
		Error: false,
		Data:  presenter.Book(book),
	})
}
