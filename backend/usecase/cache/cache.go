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

func (u *UseCase) GetOperators() []master.OperatorMaster {
	return u.cache.Get(cache.Operators).([]master.OperatorMaster)
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

func (u *UseCase) GetStationTimetables() []master.StationTimetableMaster {
	return u.cache.Get(cache.StationTimetables).([]master.StationTimetableMaster)
}
