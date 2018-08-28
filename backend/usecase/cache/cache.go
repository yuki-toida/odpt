package cache

import (
	"errors"

	"github.com/yuki-toida/refodpt/backend/domain/model/master"
	"github.com/yuki-toida/refodpt/backend/infrastructure/cache"
)

type UseCase struct {
	cache *cache.Cache
}

func NewUseCase(c *cache.Cache) *UseCase {
	return &UseCase{
		cache: c,
	}
}

func notFound() error {
	return errors.New("CacheNotFound")
}

func (u *UseCase) GetCalendars() []master.CalendarMaster {
	return u.cache.Get(cache.Calendars).([]master.CalendarMaster)
}

func (u *UseCase) GetCalendar(sameAs string) (master.CalendarMaster, error) {
	for _, v := range u.GetCalendars() {
		if v.SameAs == sameAs {
			return v, nil
		}
	}
	return master.CalendarMaster{}, notFound()
}

func (u *UseCase) GetOperators() []master.OperatorMaster {
	return u.cache.Get(cache.Operators).([]master.OperatorMaster)
}

func (u *UseCase) GetOperator(sameAs string) (master.OperatorMaster, error) {
	for _, v := range u.GetOperators() {
		if v.SameAs == sameAs {
			return v, nil
		}
	}
	return master.OperatorMaster{}, notFound()
}

func (u *UseCase) GetPassengerSurveys() []master.PassengerSurveyMaster {
	return u.cache.Get(cache.PassengerSurveys).([]master.PassengerSurveyMaster)
}

func (u *UseCase) GetPassengerSurvey(sameAs string) (master.PassengerSurveyMaster, error) {
	for _, v := range u.GetPassengerSurveys() {
		if v.SameAs == sameAs {
			return v, nil
		}
	}
	return master.PassengerSurveyMaster{}, notFound()
}

func (u *UseCase) GetRailDirections() []master.RailDirectionMaster {
	return u.cache.Get(cache.RailDirection).([]master.RailDirectionMaster)
}

func (u *UseCase) GetRailDirection(sameAs string) (master.RailDirectionMaster, error) {
	for _, v := range u.GetRailDirections() {
		if v.SameAs == sameAs {
			return v, nil
		}
	}
	return master.RailDirectionMaster{}, notFound()
}

func (u *UseCase) GetRailways() []master.RailwayMaster {
	return u.cache.Get(cache.Railways).([]master.RailwayMaster)
}

func (u *UseCase) GetRailway(sameAs string) (master.RailwayMaster, error) {
	for _, v := range u.GetRailways() {
		if v.SameAs == sameAs {
			return v, nil
		}
	}
	return master.RailwayMaster{}, notFound()
}

func (u *UseCase) GetStations() []master.StationMaster {
	return u.cache.Get(cache.Stations).([]master.StationMaster)
}

func (u *UseCase) GetStation(sameAs string) (master.StationMaster, error) {
	for _, v := range u.GetStations() {
		if v.SameAs == sameAs {
			return v, nil
		}
	}
	return master.StationMaster{}, notFound()
}

func (u *UseCase) GetStationTimetableObjectTrainNames() []master.StationTimetableMasterObjectTrainName {
	return u.cache.Get(cache.StationTimetableObjectTrainNames).([]master.StationTimetableMasterObjectTrainName)
}

func (u *UseCase) GetStationTimetableObjectTrainNamesByObjectID(objectID int) []master.StationTimetableMasterObjectTrainName {
	rows := []master.StationTimetableMasterObjectTrainName{}
	for _, v := range u.GetStationTimetableObjectTrainNames() {
		if v.StationTimetableMasterObjectID == objectID {
			rows = append(rows, v)
		}
	}
	return rows
}

func (u *UseCase) GetStationTimetableObjectViaRailways() []master.StationTimetableMasterObjectViaRailway {
	return u.cache.Get(cache.StationTimetableObjectViaRailways).([]master.StationTimetableMasterObjectViaRailway)
}

func (u *UseCase) GetStationTimetableObjectViaRailwaysByObjectID(objectID int) []master.StationTimetableMasterObjectViaRailway {
	rows := []master.StationTimetableMasterObjectViaRailway{}
	for _, v := range u.GetStationTimetableObjectViaRailways() {
		if v.StationTimetableMasterObjectID == objectID {
			rows = append(rows, v)
		}
	}
	return rows
}

func (u *UseCase) GetTrainTimetableNexts() []master.TrainTimetableMasterNext {
	return u.cache.Get(cache.TrainTimetableNexts).([]master.TrainTimetableMasterNext)
}

func (u *UseCase) GetTrainTimetableNextsByID(ID string) []master.TrainTimetableMasterNext {
	rows := []master.TrainTimetableMasterNext{}
	for _, v := range u.GetTrainTimetableNexts() {
		if v.TrainTimetableMasterID == ID {
			rows = append(rows, v)
		}
	}
	return rows
}

func (u *UseCase) GetTrainTimetablePrevious() []master.TrainTimetableMasterPrevious {
	return u.cache.Get(cache.TrainTimetablePrevious).([]master.TrainTimetableMasterPrevious)
}

func (u *UseCase) GetTrainTimetablePreviousByID(ID string) []master.TrainTimetableMasterPrevious {
	rows := []master.TrainTimetableMasterPrevious{}
	for _, v := range u.GetTrainTimetablePrevious() {
		if v.TrainTimetableMasterID == ID {
			rows = append(rows, v)
		}
	}
	return rows
}

func (u *UseCase) GetTrainTypes() []master.TrainTypeMaster {
	return u.cache.Get(cache.TrainTypes).([]master.TrainTypeMaster)
}

func (u *UseCase) GetTrainType(sameAs string) (master.TrainTypeMaster, error) {
	for _, v := range u.GetTrainTypes() {
		if v.SameAs == sameAs {
			return v, nil
		}
	}
	return master.TrainTypeMaster{}, notFound()
}
