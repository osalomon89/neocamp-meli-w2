package presenter

import "github.com/teamcubation/pod/internal/inventory/book"

type jsonBook struct {
	ID     string `json:"id"`
	Author string `json:"author"`
	Title  string `json:"title"`
	Price  int    `json:"price"`
}

func Book(b book.Book) *jsonBook {
	toReturn := &jsonBook{
		ID:     b.ID,
		Author: b.Author,
		Title:  b.Title,
		Price:  b.Price,
	}

	return toReturn
}

func Books(books []book.Book) []jsonBook {
	var jsonBooks []jsonBook

	for _, b := range books {
		jsonBooks = append(jsonBooks, jsonBook{
			ID:     b.ID,
			Author: b.Author,
			Title:  b.Title,
			Price:  b.Price,
		})
	}

	return jsonBooks
}

type ApiError struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}
