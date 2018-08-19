package registry

import "github.com/yuki-toida/refodpt/backend/infrastructure/repository"

type Registry struct {
	Repository *repository.Repository
}

func NewRegistry(r *repository.Repository) *Registry {
	return &Registry{
		Repository: r,
	}
}
