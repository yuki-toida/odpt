package shared

import (
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
