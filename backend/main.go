package main

import (
	"net/http"

	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/bamzi/jobrunner"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	"github.com/yuki-toida/refodpt/backend/config"
	"github.com/yuki-toida/refodpt/backend/domain/model/master"
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
	db.LogMode(true)
	db.AutoMigrate(
		master.CategoryMaster{},
		master.CalendarMaster{},
		master.CalendarMasterDay{},
		master.OperatorMaster{},
		master.PassengerSurveyMaster{},
		master.PassengerSurveyMasterObject{},
		master.PassengerSurveyMasterRailway{},
		master.PassengerSurveyMasterStation{},
		master.RailDirectionMaster{},
		master.RailwayMaster{},
		master.RailwayMasterStationOrder{},
		master.RailwayFareMaster{},
		master.StationMaster{},
		master.StationMasterConnectingRailway{},
		master.StationMasterExit{},
		master.StationMasterPassengerSurvey{},
		master.StationMasterTimetable{},
		// master.StationTimetable{},
		// master.StationTimetableObject{},
		// master.StationTimetableObjectDestinationStation{},
		// master.StationTimetableObjectOriginStation{},
		// master.StationTimetableObjectViaRailway{},
		// master.StationTimetableObjectViaStation{},
		// master.TrainTimetable{},
		// master.TrainTimetableObject{},
		// master.TrainTimetableDestinationStation{},
		// master.TrainTimetableOriginStation{},
		// master.TrainTimetableNext{},
		// master.TrainTimetablePrevious{},
		// master.TrainTimetableViaRailway{},
		// master.TrainTimetableViaStation{},
		// master.TrainType{},
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
	shared := router.Group("/shared")
	{
		shared.GET("/categories", func(c *gin.Context) { handler.GetCategoryMasters(c) })
	}
	train := router.Group("/train")
	{
		train.GET("/railways", func(c *gin.Context) { handler.GetRailwayMasters(c) })
		train.GET("/railways/:sameAs", func(c *gin.Context) { handler.GetRailwayMaster(c) })
		train.GET("/stations/:sameAs", func(c *gin.Context) { handler.GetStationMaster(c) })
		train.GET("/passengerSurveys/:sameAs", func(c *gin.Context) { handler.GetPassengerSurveyMaster(c) })
	}

	router.Run(":" + config.Config.Server.Port)
}
