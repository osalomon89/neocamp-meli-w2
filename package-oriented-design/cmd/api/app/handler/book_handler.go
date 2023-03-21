package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/teamcubation/pod/cmd/api/app/handler/presenter"
	"github.com/teamcubation/pod/internal/inventory"
)

type crudHandler struct {
	useCase inventory.UseCaseCRUD
}

func NewCRUDHandler(useCase inventory.UseCaseCRUD) *crudHandler {
	return &crudHandler{useCase: useCase}
}

type createBookRequest struct {
	Author string `json:"author" binding:"required"`
	Title  string `json:"title" binding:"required"`
	Price  int    `json:"price" binding:"required"`
}

func (h crudHandler) HandleCreate(c *gin.Context) {
	var payload createBookRequest
	if err := c.BindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, presenter.ApiError{})
		return
	}

	book, err := h.useCase.SaveBook(payload.Author, payload.Title, payload.Price)
	if err != nil {
		c.JSON(http.StatusInternalServerError, presenter.ApiError{})
		return
	}

	c.JSON(http.StatusCreated, presenter.Book(book))
}

func (h crudHandler) HandleRead(c *gin.Context) {
	books, err := h.useCase.GetAllBooks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, presenter.ApiError{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, presenter.Books(books))
}
