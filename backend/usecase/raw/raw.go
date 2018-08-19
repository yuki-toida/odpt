package raw

import (
	"github.com/yuki-toida/refodpt/backend/domain/model/raw"
	"github.com/yuki-toida/refodpt/backend/infrastructure/http"
)

type UseCase struct {
}

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

func handleBy(path string, results interface{}, args map[string]string) error {
	c := http.NewOdptClient()
	if err := c.GetBy(path, results, args); err != nil {
		return err
	}
	return nil
}

func (u *UseCase) GetCalendars() ([]raw.Calendar, error) {
	results := []raw.Calendar{}
	err := handle("odpt:Calendar.json", &results)
	return results, err
}

func (u *UseCase) GetOperators() ([]raw.Operator, error) {
	results := []raw.Operator{}
	err := handle("odpt:Operator.json", &results)
	return results, err
}

func (u *UseCase) GetPassengerSurveys() ([]raw.PassengerSurvey, error) {
	results := []raw.PassengerSurvey{}
	err := handle("odpt:PassengerSurvey.json", &results)
	return results, err
}

func (u *UseCase) GetRailDirections() ([]raw.RailDirection, error) {
	results := []raw.RailDirection{}
	err := handle("odpt:RailDirection.json", &results)
	return results, err
}

func (u *UseCase) GetRailways() ([]raw.Railway, error) {
	results := []raw.Railway{}
	err := handle("odpt:Railway.json", &results)
	return results, err
}

func (u *UseCase) GetRailwayFares() ([]raw.RailwayFare, error) {
	results := []raw.RailwayFare{}
	err := handle("odpt:RailwayFare.json", &results)
	return results, err
}

func (u *UseCase) GetStations() ([]raw.Station, error) {
	results := []raw.Station{}
	err := handle("odpt:Station.json", &results)
	return results, err
}

func (u *UseCase) GetStationTimetables() ([]raw.StationTimetable, error) {
	results := []raw.StationTimetable{}
	err := handle("odpt:StationTimetable.json", &results)
	return results, err
}

func (u *UseCase) GetTrains() ([]raw.Train, error) {
	results := []raw.Train{}
	err := handle("odpt:Train", &results)
	return results, err
}

func (u *UseCase) GetTrainInformations() ([]raw.TrainInformation, error) {
	results := []raw.TrainInformation{}
	err := handle("odpt:TrainInformation", &results)
	return results, err
}

func (u *UseCase) GetTrainTimetables() ([]raw.TrainTimetable, error) {
	results := []raw.TrainTimetable{}
	err := handle("odpt:TrainTimetable.json", &results)
	return results, err
}

func (u *UseCase) GetTrainTypes() ([]raw.TrainType, error) {
	results := []raw.TrainType{}
	err := handle("odpt:TrainType.json", &results)
	return results, err
}
