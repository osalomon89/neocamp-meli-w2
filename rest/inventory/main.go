package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

const port = ":9000"

type Book struct {
	ID     int    `json:"id"`
	Author string `json:"author"`
	Title  string `json:"title"`
	Price  int    `json:"price"`
}

var db []Book

func main() {
	b1 := Book{
		ID:     1,
		Title:  "Dune",
		Price:  1965,
		Author: "Frank Herbert",
	}

	b2 := Book{
		ID:     2,
		Title:  "Cita con Rama",
		Price:  1974,
		Author: "Arthur C. Clarke",
	}

	b3 := Book{
		ID:     3,
		Title:  "Un guijarro en el cielo",
		Price:  500,
		Author: "Isaac Asimov",
	}

	db = append(db, b1, b2, b3)

	r := gin.Default()

	r.GET("/ping", pong)

	r.GET("/api/v1/books", getBooks)
	r.POST("/api/v1/books", addBook)
	r.GET("/api/v1/books/:id", getBook)

	r.Run(port)
}

func pong(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"error": false,
		"data":  "pong",
	})
}

type responseInfo struct {
	Error bool        `json:"error"`
	Data  interface{} `json:"data"`
}

func getBooks(c *gin.Context) {
	c.JSON(http.StatusOK, responseInfo{
		Error: false,
		Data:  db,
	})
}

func addBook(c *gin.Context) {
	request := c.Request
	body := request.Body

	var book Book
	err := json.NewDecoder(body).Decode(&book)
	if err != nil {
		c.JSON(http.StatusBadRequest, responseInfo{
			Error: true,
			Data:  fmt.Sprintf("invalid json: %s", err.Error()),
		})
		return
	}

	db = append(db, book)

	c.JSON(http.StatusOK, responseInfo{
		Error: false,
		Data:  book,
	})
}

func getBook(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, responseInfo{
			Error: true,
			Data:  fmt.Sprintf("invalid param: %s", err.Error()),
		})
		return
	}

	for _, v := range db {
		if v.ID == id {
			c.JSON(http.StatusOK, responseInfo{
				Error: false,
				Data:  v,
			})
			return
		}
	}

	c.JSON(http.StatusNotFound, responseInfo{
		Error: true,
		Data:  "book not found",
	})
}
