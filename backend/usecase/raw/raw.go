package raw

import (
	"encoding/json"
	"io/ioutil"

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

func handleJSON(filename string, results interface{}) error {
	filepath := "./json/" + filename + ".json"
	raw, err := ioutil.ReadFile(filepath)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(raw, results); err != nil {
		return err
	}
	return nil
}

func handleBy(path string, results interface{}, args map[string]string) error {
	c := http.NewOdptClient()
	if err := c.GetBy(path, results, args); err != nil {
		return err
	}
	return nil
}

// GetCalendar func
func (*UseCase) GetCalendar() ([]*raw.Calendar, error) {
	results := []*raw.Calendar{}
	err := handle("odpt:Calendar.json", &results)
	return results, err
}

// GetOperator func
func (*UseCase) GetOperator() ([]*raw.Operator, error) {
	results := []*raw.Operator{}
	err := handle("odpt:Operator.json", &results)
	return results, err
}

// GetPassengerSurvey func
func (*UseCase) GetPassengerSurvey() ([]*raw.PassengerSurvey, error) {
	results := []*raw.PassengerSurvey{}
	err := handle("odpt:PassengerSurvey.json", &results)
	return results, err
}

// GetRailDirection func
func (*UseCase) GetRailDirection() ([]*raw.RailDirection, error) {
	results := []*raw.RailDirection{}
	err := handle("odpt:RailDirection.json", &results)
	return results, err
}

// GetRailway func
func (*UseCase) GetRailway() ([]*raw.Railway, error) {
	results := []*raw.Railway{}
	err := handle("odpt:Railway.json", &results)
	return results, err
}

// GetRailwayFare func
func (*UseCase) GetRailwayFare() ([]*raw.RailwayFare, error) {
	results := []*raw.RailwayFare{}
	err := handle("odpt:RailwayFare.json", &results)
	return results, err
}

// GetStation func
func (*UseCase) GetStation() ([]*raw.Station, error) {
	results := []*raw.Station{}
	err := handle("odpt:Station.json", &results)
	return results, err
}

// GetStationTimetable func
func (*UseCase) GetStationTimetable() ([]*raw.StationTimetable, error) {
	results := []*raw.StationTimetable{}
	err := handle("odpt:StationTimetable.json", &results)
	return results, err
}
