package http

import (
	"log"

	"github.com/gin-gonic/gin"
	ctrl "github.com/teamcubation/neocamp-meli/clean-architecture/internal/infrastructure/http/controller"
)

type HTTPServer interface {
	RegisterRouter()
	Run(port string) error
}

type httpServer struct {
	router   *gin.Engine
	bookCtrl ctrl.BookController
}

func NewHTTPServer(bookCtrl ctrl.BookController) HTTPServer {
	ginrouter := gin.Default()

	return &httpServer{
		router:   ginrouter,
		bookCtrl: bookCtrl,
	}
}

func (srv *httpServer) RegisterRouter() {
	basePath := "/api/v1/books"
	r := srv.router.Group(basePath)

	r.GET("/", srv.bookCtrl.GetBooks)
	r.POST("/", srv.bookCtrl.SaveBook)
}

func (srv *httpServer) Run(port string) error {
	log.Println("Server listening on port", port)

	return srv.router.Run(":" + port)
}
