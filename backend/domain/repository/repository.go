package repository

// Repository type
type Repository interface {
	First(target interface{}, query interface{})
	FindAll(targets interface{})
	Find(targets interface{}, query interface{})
	Create(target interface{})
	Update(target interface{})
	Delete(target interface{})
}
