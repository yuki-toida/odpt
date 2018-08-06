package train

import (
	"github.com/yuki-toida/refodpt/backend/domain/model/train"
	"github.com/yuki-toida/refodpt/backend/infrastructure/http"
)

// UseCase type
type UseCase struct{}

// NewUseCase func
func NewUseCase() *UseCase {
	return &UseCase{}
}

// GetPassengerSurvey func ※OVER1000
func (*UseCase) GetPassengerSurvey() ([]*train.PassengerSurvey, error) {
	c := http.NewOdptClient()
	results := []*train.PassengerSurvey{}
	if err := c.Get("odpt:PassengerSurvey", &results); err != nil {
		return nil, err
	}
	return results, nil
}

// GetRailDirection func
func (*UseCase) GetRailDirection() ([]*train.RailDirection, error) {
	c := http.NewOdptClient()
	results := []*train.RailDirection{}
	if err := c.Get("odpt:RailDirection", &results); err != nil {
		return nil, err
	}
	return results, nil
}

// GetRailway func
func (*UseCase) GetRailway() ([]*train.Railway, error) {
	c := http.NewOdptClient()
	results := []*train.Railway{}
	if err := c.Get("odpt:Railway", &results); err != nil {
		return nil, err
	}
	return results, nil
}
