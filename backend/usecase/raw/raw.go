package raw

import (
	"github.com/yuki-toida/refodpt/backend/domain/model/raw"
	"github.com/yuki-toida/refodpt/backend/infrastructure/http"
)

// UseCase type
type UseCase struct {
}

// NewUseCase func
func NewUseCase() *UseCase {
	return &UseCase{}
}

func handle(path string, results interface{}) error {
	c := http.NewOdptClient()
	if err := c.Get(path, results); err != nil {
		return err
	}
	return nil
}

// GetCalendar func
func (*UseCase) GetCalendar() ([]*raw.Calendar, error) {
	results := []*raw.Calendar{}
	err := handle("odpt:Calendar", &results)
	return results, err
}

// GetOperator func
func (*UseCase) GetOperator() ([]*raw.Operator, error) {
	results := []*raw.Operator{}
	err := handle("odpt:Operator", &results)
	return results, err
}

// GetPassengerSurvey func ※OVER1000
func (*UseCase) GetPassengerSurvey() ([]*raw.PassengerSurvey, error) {
	results := []*raw.PassengerSurvey{}
	err := handle("odpt:PassengerSurvey", &results)
	return results, err
}

// GetRailDirection func
func (*UseCase) GetRailDirection() ([]*raw.RailDirection, error) {
	results := []*raw.RailDirection{}
	err := handle("odpt:RailDirection", &results)
	return results, err
}

// GetRailway func
func (*UseCase) GetRailway() ([]*raw.Railway, error) {
	results := []*raw.Railway{}
	err := handle("odpt:Railway", &results)
	return results, err
}
