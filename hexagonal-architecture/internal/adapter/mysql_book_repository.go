package adapter

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/mercadolibre/inventory/internal/core/domain"
)

var operationsByColumn = map[string]string{
	"author":     "LIKE",
	"title":      "LIKE",
	"price":      "",
	"isbn":       "=",
	"stock":      "",
	"updated_at": "",
}

type bookRepository struct {
	conn *sqlx.DB
}

func NewBookRepository(db *sqlx.DB) *bookRepository {
	return &bookRepository{
		conn: db,
	}
}

func (repo *bookRepository) GetBookByID(ctx context.Context, id uint) (*domain.Book, error) {
	book := new(domain.Book)
	err := repo.conn.Get(book, "SELECT * FROM books WHERE id=?", id)
	if err != nil {
		return nil, fmt.Errorf("error getting book: %w", err)
	}

	return book, nil
}

func (repo *bookRepository) GetBooks(ctx context.Context, params map[string]interface{}) ([]domain.Book, error) {
	books := []domain.Book{}
	limit := params["limit"]
	offset := params["offset"]
	delete(params, "limit")
	delete(params, "offset")

	whereQuery := ""
	setParamsSlice, setValues := repo.getStatementParams(params)
	if len(setParamsSlice) > 0 {
		setParams := strings.Join(setParamsSlice, " AND ")
		whereQuery = fmt.Sprintf(" WHERE %s", setParams)
	}

	queryLimit := repo.getLimitOffsetStatement(limit.(float64), offset.(float64))

	query := fmt.Sprintf("SELECT * FROM books%s %s", whereQuery, queryLimit)

	err := repo.conn.Select(&books, query, setValues...)
	if err != nil {
		return nil, fmt.Errorf("error getting all books: %w", err)
	}

	return books, nil
}

func (repo *bookRepository) CreateBook(ctx context.Context, book *domain.Book) error {
	createdAt := time.Now()

	result, err := repo.conn.Exec(`INSERT INTO books 
		(title, author, price, stock, isbn, created_at, updated_at) 
		VALUES(?,?,?,?,?,?, ?)`, book.Title, book.Author, book.Price, book.Stock, book.Isbn, createdAt, createdAt)

	if err != nil {
		return fmt.Errorf("error inserting book: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("error saving book: %w", err)
	}

	book.ID = uint(id)

	return nil
}

func (repo *bookRepository) UpdateBook(ctx context.Context, book *domain.Book, params map[string]interface{}) error {
	updateAt := time.Now()

	params["updated_at"] = updateAt
	ja, err := json.Marshal(params)
	if err != nil {
		return fmt.Errorf("error marshalling book: %w", err)
	}

	err = json.Unmarshal(ja, book)
	if err != nil {
		return fmt.Errorf("error unmarshalling book: %w", err)
	}

	setParamsSlice, setValues := repo.getStatementParams(params)
	setParams := strings.Join(setParamsSlice, ",")

	query := fmt.Sprintf("UPDATE books SET %s WHERE id=?", setParams)
	setValues = append(setValues, book.ID)

	_, err = repo.conn.Exec(query, setValues...)
	if err != nil {
		return fmt.Errorf("error updating item: %w", err)
	}

	return nil
}

func (repo *bookRepository) getStatementParams(params map[string]interface{}) ([]string, []interface{}) {
	var setParams []string
	var setValues []interface{}

	for key, val := range params {
		_, ok := operationsByColumn[key]
		if !ok || val == nil || val == "" || val == 0 {
			continue
		}
		setParams = append(setParams, fmt.Sprintf("%s=?", key))
		setValues = append(setValues, val)
	}

	return setParams, setValues
}

func (repo *bookRepository) getLimitOffsetStatement(limit, offset float64) string {
	queryLimit := 100
	queryOffset := 0
	if limit > 0 && limit < 1000 {
		queryLimit = int(limit)
	}

	if offset > 0 {
		queryOffset = int(offset)
	}

	return fmt.Sprintf("LIMIT %d OFFSET %d", queryLimit, queryOffset)
}
