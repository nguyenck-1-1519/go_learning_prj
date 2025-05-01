package controller

import (
	"net/http"
	"strconv"

	"example.com/go-learning-prj/domain"
	"example.com/go-learning-prj/domain/entity"
	"github.com/gin-gonic/gin"
)

type BookController struct {
	BookUseCase domain.BookUseCase
}

func (bc BookController) GetPaginationOfBooks(c *gin.Context) {
	pageStr := c.Query("page")
	limitStr := c.Query("limit")

	page := 0
	limit := 10

	// Get & assign default value of page
	if pageStr != "" {
		pageConvert, err := strconv.Atoi(pageStr)
		if err != nil || page < 0 {
			throwBadRequestWithError(&err, c, "Invalid Parameter")
			return
		}
		page = pageConvert
	}

	// Get & assign default value of limit
	if limitStr != "" {
		limitConvert, err := strconv.Atoi(limitStr)
		if err != nil || limit <= 0 {
			throwBadRequestWithError(&err, c, "Invalid Parameter")
			return
		}
		limit = limitConvert
	}

	books, pageInfo, err := bc.BookUseCase.GetListBooks(c, page, limit)
	if err != nil {
		throwBadRequestWithError(&err, c, "")
		return
	}

	startIndex := page * limit
	if startIndex >= len(books) {
		c.IndentedJSON(http.StatusOK, entity.BaseResponse{
			Status:  entity.StatusOK,
			Data:    []entity.Book{},
			Message: "OK",
			Page:    pageInfo,
		})
		return
	}

	endIndex := min(startIndex+limit, len(books))
	resultBooks := books[startIndex:endIndex]

	c.IndentedJSON(http.StatusOK, entity.BaseResponse{
		Status:  entity.StatusOK,
		Data:    resultBooks,
		Message: "OK",
		Page:    pageInfo,
	})
}

func throwBadRequestWithError(err *error, c *gin.Context, m string) {
	message := m
	if m == "" {
		message = "Bad Request"
	}

	c.IndentedJSON(http.StatusBadRequest, entity.BaseResponse{
		Status:  entity.StatusError,
		Message: message,
		Error:   err,
	})
}
