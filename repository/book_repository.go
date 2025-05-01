package repository

import (
	"context"
	"database/sql"
	"errors"
	"log"

	"example.com/go-learning-prj/domain/entity"
)

type Book entity.Book

var (
	QueryGetBooksWithPagination = "SELECT title, price, stock FROM books LIMIT ? OFFSET ?"
	QueryGetTotalItemCount      = "SELECT COUNT(*) FROM books"
	// QueryGetBookInfo            = "SELECT id, title, author, price, stock FROM books WHERE id = ?"
	// QueryInsertBook             = "INSERT INTO books (title, author, price, stock) VALUES (?, ?, ?, ?)"
	// QueryUpdateBookInfo         = "UPDATE books SET title = ?, author = ?, price = ?, stock = ? WHERE id = ?"
	// QueryDeleteBookWithID       = "DELETE FROM books WHERE id = ?"
)

type bookRepository struct {
	database *sql.DB
}

func NewBookRepository(db *sql.DB) BookRepository {
	return &bookRepository{database: db}
}

type BookRepository interface {
	GetBooksFromDB(ctx context.Context, page int, limit int) ([]Book, entity.PaginationData, error)
}

func (rp bookRepository) GetBooksFromDB(ctx context.Context, page int, limit int) ([]Book, entity.PaginationData, error) {
	db := rp.database

	// Calculate offset
	offset := max(page*limit, 0)

	// Query
	rows, err := db.Query(QueryGetBooksWithPagination, limit, offset)
	if err != nil {
		log.Fatal("query items from db failed", page, limit)
		return nil, entity.PaginationData{}, errors.New("query items from db failed")
	}
	defer rows.Close()

	// access rows & write to books
	var books []Book
	for rows.Next() {
		var book Book
		if err := rows.Scan(&book.Name, &book.Price, &book.Stock); err != nil {
			log.Fatal("scan items db failed")
			return nil, entity.PaginationData{}, errors.New("scan items db failed")
		}
		books = append(books, book)
	}

	// Query get total item count
	var totalItems int
	err = db.QueryRow(QueryGetTotalItemCount).Scan(&totalItems)
	if err != nil {
		log.Fatal("query total count of db failed")
		return nil, entity.PaginationData{}, errors.New("query total count of db failed")
	}

	totalPages := 0
	if limit > 0 {
		totalPages = (totalItems + limit - 1) / limit
	}

	pageInfo := entity.PaginationData{
		TotalItems:  int64(totalItems),
		CurrentPage: page,
		PageSize:    limit,
		TotalPages:  int64(totalPages),
	}

	return books, pageInfo, nil
}
