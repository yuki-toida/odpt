package main

import (
	"net/http"

	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/bamzi/jobrunner"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	"github.com/yuki-toida/refodpt/backend/config"
	"github.com/yuki-toida/refodpt/backend/domain/model"
	"github.com/yuki-toida/refodpt/backend/infrastructure/repository"
	"github.com/yuki-toida/refodpt/backend/interface/handler"
	"github.com/yuki-toida/refodpt/backend/interface/job/importer"
	"github.com/yuki-toida/refodpt/backend/interface/registry"
)

func init() {
	config.Init()
}

func main() {
	conn := "root:zaqroot@tcp(" + config.Config.Db.Host + ":" + config.Config.Db.Port + ")/" + config.Config.Db.Name + "?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open("mysql", conn)
	if err != nil {
		panic(err.Error())
	}
	// db.LogMode(true)
	db.AutoMigrate(
		model.Category{},
		model.Calendar{},
		model.CalendarDay{},
		model.Operator{},
		model.PassengerSurvey{},
		model.PassengerSurveyObject{},
		model.PassengerSurveyRailway{},
		model.PassengerSurveyStation{},
		model.RailDirection{},
		model.Railway{},
		model.RailwayStationOrder{},
		model.RailwayFare{},
		model.Station{},
		model.StationConnectingRailway{},
		model.StationExit{},
		model.StationPassengerSurvey{},
		model.StationStationTimetable{},
		model.StationTimetable{},
		model.StationTimetableObject{},
		model.StationTimetableObjectDestinationStation{},
		model.StationTimetableObjectOriginStation{},
		model.StationTimetableObjectViaRailway{},
		model.StationTimetableObjectViaStation{},
	)

	cr := repository.NewRepository(db)

	registry := registry.NewRegistry(cr)
	handler := handler.NewHandler(registry)

	i := importer.NewImporter(registry)
	i.Run()
	// jobrunner.Start()
	// jobrunner.Schedule("@every 1m", importer.NewImporter(registry))

	router := gin.Default()
	router.Use(cors.Default())

	router.GET("/healthz", func(c *gin.Context) { c.String(http.StatusOK, "OK") })
	router.GET("/jobrunner", func(c *gin.Context) { c.JSON(http.StatusOK, jobrunner.StatusJson()) })

	raw := router.Group("/raw")
	{
		raw.GET("/calendar", func(c *gin.Context) { handler.GetCalendar(c) })
		raw.GET("/operator", func(c *gin.Context) { handler.GetOperator(c) })
		raw.GET("/passengerSurvey", func(c *gin.Context) { handler.GetPassengerSurvey(c) })
		raw.GET("/railDirection", func(c *gin.Context) { handler.GetRailDirection(c) })
		raw.GET("/railway", func(c *gin.Context) { handler.GetRailway(c) })
		raw.GET("/railwayFare", func(c *gin.Context) { handler.GetRailwayFare(c) })
		raw.GET("/station", func(c *gin.Context) { handler.GetStation(c) })
		raw.GET("/stationTimetable", func(c *gin.Context) { handler.GetStationTimetable(c) })
	}

	router.Run(":" + config.Config.Server.Port)
}
