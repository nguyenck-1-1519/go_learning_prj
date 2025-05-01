package usecase

import (
	"context"
	"time"

	"example.com/go-learning-prj/domain"
	"example.com/go-learning-prj/domain/entity"
	"example.com/go-learning-prj/repository"
)

type bookUseCase struct {
	booksRepository repository.BooksRepository
	contextTimeout  time.Duration
}

func NewBookUseCase(booksRepository repository.BooksRepository, timeout time.Duration) domain.BookUseCase {
	return &bookUseCase{
		booksRepository: booksRepository,
		contextTimeout:  timeout,
	}
}

func (bc bookUseCase) GetListBooks(c context.Context, page int, limit int) ([]entity.Book, entity.PaginationData, error) {
	ctx, cancel := context.WithTimeout(c, bc.contextTimeout)
	defer cancel()
	return bc.booksRepository.GetPaginationOfBooks(ctx, page, limit)
}
