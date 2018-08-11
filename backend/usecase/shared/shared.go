package shared

import (
	"github.com/yuki-toida/refodpt/backend/domain/model/master"
	"github.com/yuki-toida/refodpt/backend/infrastructure/repository"
)

// UseCase type
type UseCase struct {
	Repository *repository.Repository
}

// NewUseCase func
func NewUseCase(r *repository.Repository) *UseCase {
	return &UseCase{
		Repository: r,
	}
}

// GetCategoryMasters func
func (u *UseCase) GetCategoryMasters() []master.CategoryMaster {
	var rows []master.CategoryMaster
	u.Repository.DB.Find(&rows)
	return rows
}
