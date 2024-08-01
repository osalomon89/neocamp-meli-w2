package adapter

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mercadolibre/inventory/internal/core/domain"
	"github.com/mercadolibre/inventory/internal/core/ports"
)

type BookHandler interface {
	GetBooks(c *gin.Context)
	GetBookByID(c *gin.Context)
	AddBook(c *gin.Context)
	UpdateBook(c *gin.Context)
}

type bookHandler struct {
	bookUsecase ports.BookUsecase
}

type createBookReq struct {
	ID        uint      `json:"id"`
	Author    string    `json:"author" validate:"required"`
	Title     string    `json:"title"`
	Price     int       `json:"price"`
	Isbn      string    `json:"isbn"`
	Stock     int       `json:"stock"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type bookRequestQuery struct {
	Author string `form:"author"`
	Title  string `form:"title"`
	Isbn   string `form:"isbn"`
	Limit  int    `form:"limit"`
	Offset int    `form:"offset"`
}

type responseInfo struct {
	Error bool        `json:"error"`
	Data  interface{} `json:"data"`
}

func NewBookHandler(usecase ports.BookUsecase) *bookHandler {
	return &bookHandler{
		bookUsecase: usecase,
	}
}

func (h *bookHandler) GetBooks(c *gin.Context) {
	bookRequestQuery := new(bookRequestQuery)
	err := c.ShouldBind(bookRequestQuery)
	if err != nil {
		c.JSON(http.StatusBadRequest, responseInfo{
			Error: true,
			Data:  "invalid params",
		})
		return
	}

	bookRequestQueryString, err := json.Marshal(bookRequestQuery)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responseInfo{
			Error: true,
			Data:  "Unexpected server error",
		})
		return
	}

	var params map[string]interface{}
	if err := json.Unmarshal(bookRequestQueryString, &params); err != nil {
		c.JSON(http.StatusInternalServerError, responseInfo{
			Error: true,
			Data:  "Unexpected server error",
		})
		return
	}

	books, err := h.bookUsecase.GetBooks(c, params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responseInfo{
			Error: true,
			Data:  "Unexpected server error",
		})
		return
	}

	c.JSON(http.StatusOK, responseInfo{
		Error: false,
		Data:  books,
	})
}

func (h *bookHandler) GetBookByID(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, responseInfo{
			Error: true,
			Data:  fmt.Sprintf("invalid param: %s", err.Error()),
		})
		return
	}

	book, err := h.bookUsecase.GetBookByID(c, uint(id))
	if err != nil {
		if err == domain.ErrBookNotFound {
			c.JSON(http.StatusNotFound, responseInfo{
				Error: true,
				Data:  fmt.Sprintf("book not found. ID: %d", id),
			})
			return
		}
		c.JSON(http.StatusInternalServerError, responseInfo{
			Error: true,
			Data:  "Unexpected server error",
		})
		return
	}

	c.JSON(http.StatusOK, responseInfo{
		Error: false,
		Data:  book,
	})
}

func (h *bookHandler) AddBook(c *gin.Context) {
	request := c.Request
	body := request.Body

	var book createBookReq
	err := json.NewDecoder(body).Decode(&book)
	if err != nil {
		c.JSON(http.StatusBadRequest, responseInfo{
			Error: true,
			Data:  fmt.Sprintf("invalid json: %s", err.Error()),
		})
		return
	}

	result, err := h.bookUsecase.AddBook(c, domain.Book(book))
	if err != nil {
		c.JSON(http.StatusInternalServerError, responseInfo{
			Error: true,
			Data:  fmt.Sprintf("error saving book: %s", err.Error()),
		})
		return
	}

	c.JSON(http.StatusCreated, responseInfo{
		Error: false,
		Data:  result,
	})
}

func (h *bookHandler) UpdateBook(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, responseInfo{
			Error: true,
			Data:  fmt.Sprintf("invalid param: %s", err.Error()),
		})
		return
	}

	bookRequestBody := make(map[string]interface{})
	err = json.NewDecoder(c.Request.Body).Decode(&bookRequestBody)
	if err != nil {
		c.JSON(http.StatusBadRequest, responseInfo{
			Error: true,
			Data:  "invalid params",
		})
		return
	}

	book, err := h.bookUsecase.UpdateBook(c, uint(id), bookRequestBody)
	if err != nil {
		if err == domain.ErrBookNotFound {
			c.JSON(http.StatusNotFound, responseInfo{
				Error: true,
				Data:  fmt.Sprintf("book not found. ID: %d", id),
			})
			return
		}
		c.JSON(http.StatusInternalServerError, responseInfo{
			Error: true,
			Data:  "Unexpected server error",
		})
		return
	}

	c.JSON(http.StatusOK, responseInfo{
		Error: false,
		Data:  book,
	})
}
