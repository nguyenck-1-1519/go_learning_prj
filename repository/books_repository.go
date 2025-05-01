package repository

import (
	"context"
	"errors"
	"log"

	"example.com/go-learning-prj/domain/entity"
	"gorm.io/gorm"
)

type booksRepository struct {
	database *gorm.DB
}

func NewBooksRepository(db *gorm.DB) BooksRepository {
	return &booksRepository{database: db}
}

type BooksRepository interface {
	GetPaginationOfBooks(c context.Context, page int, limit int) ([]entity.Book, entity.PaginationData, error)
}

func (br booksRepository) GetPaginationOfBooks(c context.Context, page int, limit int) ([]entity.Book, entity.PaginationData, error) {
	// Calculate offset
	offset := max(page*limit, 0)

	books := []entity.Book{}
	// Query list books with offset & limit
	result := br.database.Limit(limit).Offset(offset).Find(&books)
	if result.Error != nil {
		log.Fatal("1", result.Error)
		return nil, entity.PaginationData{}, errors.New("query items from db failed")
	}

	// Query get total item count
	var totalItems int64
	result = br.database.Model(&entity.Book{}).Count(&totalItems)
	if result.Error != nil {
		log.Fatal("2", result.Error)
		return nil, entity.PaginationData{}, errors.New("query total count of db failed")
	}

	var totalPages int64 = 0
	if limit > 0 {
		totalPages = (totalItems + int64(limit) - 1) / int64(limit)
	}

	pageInfo := entity.PaginationData{
		TotalItems:  totalItems,
		CurrentPage: page,
		PageSize:    limit,
		TotalPages:  totalPages,
	}

	return books, pageInfo, nil
}
