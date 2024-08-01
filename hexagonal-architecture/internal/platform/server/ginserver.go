package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mercadolibre/inventory/internal/adapter"
)

type httpServer struct {
	router      *gin.Engine
	bookHandler adapter.BookHandler
}

func NewHTTPServer(bookHandler adapter.BookHandler) *httpServer {
	return &httpServer{
		router:      gin.Default(),
		bookHandler: bookHandler,
	}
}

func (srv *httpServer) RegisterRouter() {
	srv.router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"error": false,
			"data":  "pong",
		})
	})

	basePath := "/api/v1"
	publicRouter := srv.router.Group(basePath)

	publicRouter.GET("/books", srv.bookHandler.GetBooks)
	publicRouter.POST("/books", srv.bookHandler.AddBook)
	publicRouter.GET("/books/:id", srv.bookHandler.GetBookByID)
	publicRouter.PATCH("/books/:id", srv.bookHandler.UpdateBook)
}

func (srv *httpServer) Run(port string) error {
	return srv.router.Run(":" + port)
}
