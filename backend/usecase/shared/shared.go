package shared

import (
	"github.com/yuki-toida/refodpt/backend/domain/model/shared"
	"github.com/yuki-toida/refodpt/backend/infrastructure/http"
)

// UseCase type
type UseCase struct{}

// NewUseCase func
func NewUseCase() *UseCase {
	return &UseCase{}
}

// GetCalendar func
func (*UseCase) GetCalendar() ([]*shared.Calendar, error) {
	c := http.NewOdptClient()
	results := []*shared.Calendar{}
	if err := c.Get("odpt:Calendar", &results); err != nil {
		return nil, err
	}
	return results, nil
}

// GetOperator func
func (*UseCase) GetOperator() ([]*shared.Operator, error) {
	c := http.NewOdptClient()
	results := []*shared.Operator{}
	if err := c.Get("odpt:Operator", &results); err != nil {
		return nil, err
	}
	return results, nil
}
