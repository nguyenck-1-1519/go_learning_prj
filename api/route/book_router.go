package route

import (
	"time"

	"example.com/go-learning-prj/api/controller"
	"example.com/go-learning-prj/repository"
	"example.com/go-learning-prj/usecase"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewBookRouter(timeout time.Duration, db *gorm.DB, group *gin.RouterGroup) {
	br := repository.NewBooksRepository(db)
	bc := &controller.BookController{
		BookUseCase: usecase.NewBookUseCase(br, timeout),
	}
	group.GET("/books", bc.GetPaginationOfBooks)
}
