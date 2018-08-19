package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/yuki-toida/refodpt/backend/infrastructure/cache"
)

type Repository struct {
	DB    *gorm.DB
	Cache *cache.Cache
}

func NewRepository(db *gorm.DB, c *cache.Cache) *Repository {
	return &Repository{
		DB:    db,
		Cache: c,
	}
}
