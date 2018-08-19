package cache

import (
	"strconv"

	"github.com/jinzhu/gorm"
	"github.com/patrickmn/go-cache"
	"github.com/yuki-toida/refodpt/backend/domain/model/master"
)

// cache keys
const (
	Calendars = iota
	Operators
	PassengerSurveys
	Railways
	Stations
	StationTimetables
)

type Cache struct {
	cache *cache.Cache
}

func NewCache() *Cache {
	return &Cache{
		cache: cache.New(cache.NoExpiration, cache.NoExpiration),
	}
}

func (c *Cache) Set(key int, value interface{}) {
	c.cache.SetDefault(strconv.Itoa(key), value)
}

func (c *Cache) Get(key int) interface{} {
	results, _ := c.cache.Get(strconv.Itoa(key))
	return results
}

func (c *Cache) Init(db *gorm.DB) {
	var calendars []master.CalendarMaster
	db.Preload("Days").Find(&calendars)
	c.Set(Calendars, calendars)

	var operators []master.OperatorMaster
	db.Find(&operators)
	c.Set(Operators, operators)

	var passengerSurveys []master.PassengerSurveyMaster
	db.Preload("Operator").
		Preload("Railways").
		Preload("Railways.Railway").
		Preload("Stations").
		Preload("Stations.Station").
		Preload("Objects").
		Find(&passengerSurveys)
	c.Set(PassengerSurveys, passengerSurveys)

	var railways []master.RailwayMaster
	db.Preload("Operator").Preload("StationOrders").Find(&railways)
	c.Set(Railways, railways)

	var stations []master.StationMaster
	db.Preload("Operator").
		Preload("Railway").
		Preload("ConnectingRailways").
		Preload("ConnectingRailways.Railway").
		Preload("PassengerSurveys").
		Preload("Timetables").
		Preload("Timetables.StationTimetable").
		Find(&stations)
	c.Set(Stations, stations)

	var stationTimetables []master.StationTimetableMaster
	db.Preload("Calendar").Preload("RailDirection").Find(&stationTimetables)
	c.Set(StationTimetables, stationTimetables)

}
