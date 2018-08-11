package registry

import "github.com/yuki-toida/refodpt/backend/infrastructure/repository"

// Registry type
type Registry struct {
	Repository *repository.Repository
}

// NewRegistry func
func NewRegistry(r *repository.Repository) *Registry {
	return &Registry{
		Repository: r,
	}
}
