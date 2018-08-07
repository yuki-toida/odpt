package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/yuki-toida/refodpt/backend/domain/repository"
)

type repositoryImpl struct {
	db *gorm.DB
}

// NewRepository func
func NewRepository(db *gorm.DB) repository.Repository {
	return &repositoryImpl{
		db: db,
	}
}

func (r *repositoryImpl) First(target interface{}, query interface{}) {
	r.db.Where(query).First(target)
}

func (r *repositoryImpl) FindAll(targets interface{}) {
	r.db.Find(targets)
}

func (r *repositoryImpl) Find(targets interface{}, query interface{}) {
	r.db.Where(query).Find(targets)
}

func (r *repositoryImpl) Create(target interface{}) {
	r.db.Create(target)
}

func (r *repositoryImpl) Update(target interface{}) {
	r.db.Save(target)
}

func (r *repositoryImpl) Delete(target interface{}) {
	r.db.Delete(target)
}
