package entity

type Book struct {
	BookId    int `gorm:"primaryKey"`
	Name      string
	Stock     int
	Price     float64
	Category  Category  `gorm:"foreignKey:category_id;references:categories"`
	Publisher Publisher `gorm:"foreignKey:publisher_id;references:publishers"`
}

type PaginationData struct {
	TotalItems  int64 `json:"total_items" binding:"required,gte=0"`
	CurrentPage int   `json:"current_page" binding:"required,gte=0"`
	PageSize    int   `json:"page_size" binding:"required,gte=0"`
	TotalPages  int64 `json:"total_pages" binding:"required,gte=0"`
}

type BaseResponse struct {
	Data    any            `json:"data,omitempty"`
	Error   *error         `json:"error,omitempty"`
	Message string         `json:"message" binding:"required"`
	Page    PaginationData `json:"page,omitzero"`
	Status  Status         `json:"status" binding:"required"`
}

type Status string

const (
	StatusOK    = "OK"
	StatusError = "Error"
)
