package domain

import (
	"context"

	"example.com/go-learning-prj/domain/entity"
)

type BookUseCase interface {
	GetListBooks(c context.Context, page int, limit int) ([]entity.Book, entity.PaginationData, error)
}
