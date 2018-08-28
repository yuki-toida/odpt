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
	RailDirection
	Railways
	Stations
	StationTimetableObjectTrainNames
	StationTimetableObjectViaRailways
	TrainTimetableNexts
	TrainTimetablePrevious
	TrainTypes
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
	cached, _ := c.cache.Get(strconv.Itoa(key))
	return cached
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

	var railDirection []master.RailDirectionMaster
	db.Find(&railDirection)
	c.Set(RailDirection, railDirection)

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
		Find(&stations)
	c.Set(Stations, stations)

	var stationTimetableObjectTrainNames []master.StationTimetableMasterObjectTrainName
	db.Find(&stationTimetableObjectTrainNames)
	c.Set(StationTimetableObjectTrainNames, stationTimetableObjectTrainNames)

	var trainTimetableNexts []master.TrainTimetableMasterNext
	db.Find(&trainTimetableNexts)
	c.Set(TrainTimetableNexts, trainTimetableNexts)

	var trainTimetablePrevious []master.TrainTimetableMasterPrevious
	db.Find(&trainTimetablePrevious)
	c.Set(TrainTimetablePrevious, trainTimetablePrevious)

	var trainTypes []master.TrainTypeMaster
	db.Preload("Operator").Find(&trainTypes)
	c.Set(TrainTypes, trainTypes)
}
