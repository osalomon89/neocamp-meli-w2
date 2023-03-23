package db

import "github.com/teamcubation/neocamp-meli/clean-architecture/internal/domain"

func GetConnectionDB() []domain.Book {
	db := []domain.Book{
		{
			ID:     1,
			Title:  "Dune",
			Price:  1965,
			Author: "Frank Herbert",
		},
		{
			ID:     2,
			Title:  "Cita con Rama",
			Price:  1974,
			Author: "Arthur C. Clarke",
		},
		{
			ID:     3,
			Title:  "Un guijarro en el cielo",
			Price:  500,
			Author: "Isaac Asimov",
		},
	}

	return db
}
