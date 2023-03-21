package controller

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/teamcubation/neocamp-meli/clean-architecture/internal/domain"
	"github.com/teamcubation/neocamp-meli/clean-architecture/internal/usecase"
)

type BookController interface {
	GetBooks(ctx *gin.Context)
	SaveBook(ctx *gin.Context)
}

type bookController struct {
	bookUsecase usecase.BookUsecase
}

func NewBookHandler(bookUsecase usecase.BookUsecase) BookController {
	return &bookController{
		bookUsecase: bookUsecase,
	}
}

func (h *bookController) GetBooks(ctx *gin.Context) {
	books := h.bookUsecase.GetAllBooks()

	ctx.JSON(http.StatusOK, gin.H{
		"error": false,
		"data":  books,
	})
}

func (h *bookController) SaveBook(ctx *gin.Context) {
	var book domain.Book
	if err := json.NewDecoder(ctx.Request.Body).Decode(&book); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": false,
			"data":  err.Error(),
		})
		return
	}

	result := h.bookUsecase.SaveBook(book)

	ctx.JSON(http.StatusOK, gin.H{
		"error": false,
		"data":  result,
	})
}
