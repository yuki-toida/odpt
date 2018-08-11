package repository

import (
	"github.com/jinzhu/gorm"
)

// Repository struct
type Repository struct {
	DB *gorm.DB
}

// NewRepository func
func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		DB: db,
	}
}
