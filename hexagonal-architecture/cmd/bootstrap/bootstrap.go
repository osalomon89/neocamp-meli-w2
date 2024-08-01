package bootstrap

import (
	"log"

	"github.com/mercadolibre/inventory/internal/adapter"
	"github.com/mercadolibre/inventory/internal/core/usecase"
	"github.com/mercadolibre/inventory/internal/platform/config"
	mysqlconn "github.com/mercadolibre/inventory/internal/platform/mysql"
	"github.com/mercadolibre/inventory/internal/platform/server"
)

func Run() error {
	err := config.LoadConfigs()
	if err != nil {
		return err
	}

	db, err := mysqlconn.GetConnectionDB()
	if err != nil {
		log.Fatalf("Database configuration failed: %v", err)
	}

	bookRepository := adapter.NewBookRepository(db)
	bookUsecase := usecase.NewBookUsecase(bookRepository)

	srv := server.NewHTTPServer(adapter.NewBookHandler(bookUsecase))

	return srv.Run("8080")
}
