package main

import (
	"fmt"
	"net/http"
	"runtime"

	"github.com/bamzi/jobrunner"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/yuki-toida/refodpt/backend/config"
	"github.com/yuki-toida/refodpt/backend/domain/model/master"
	"github.com/yuki-toida/refodpt/backend/domain/model/tran"
	"github.com/yuki-toida/refodpt/backend/infrastructure/cache"
	"github.com/yuki-toida/refodpt/backend/infrastructure/repository"
	"github.com/yuki-toida/refodpt/backend/interface/handler"
	"github.com/yuki-toida/refodpt/backend/interface/job/importer"
	"github.com/yuki-toida/refodpt/backend/interface/registry"
)

func init() {
	config.Init()
}

func main() {
	db := connectDB()
	cc := cache.NewCache()
	cc.Init(db)

	reg := registry.NewRegistry(repository.NewRepository(db, cc))
	ha := handler.NewHandler(reg)

	i := importer.NewImporter(reg)
	i.Run()
	// jobrunner.Start()
	// jobrunner.Schedule("@every 1m", importer.NewImporter(reg))

	router := gin.Default()
	router.Use(cors.Default())

	router.GET("/", func(c *gin.Context) { c.String(http.StatusOK, "OK") })
	router.GET("/healthz", func(c *gin.Context) { c.String(http.StatusOK, "OK") })
	router.GET("/jobrunner", func(c *gin.Context) { c.JSON(http.StatusOK, jobrunner.StatusJson()) })

	train := router.Group("/train")
	{
		train.GET("/railways", func(c *gin.Context) { ha.GetRailwayMasters(c) })
		train.GET("/railways/:sameAs", func(c *gin.Context) { ha.GetRailwayMaster(c) })
		train.GET("/stations/:sameAs", func(c *gin.Context) { ha.GetStationMaster(c) })
		train.GET("/passengerSurveys/:sameAs", func(c *gin.Context) { ha.GetPassengerSurveyMaster(c) })
		train.GET("/stationTimetables/:sameAs", func(c *gin.Context) { ha.GetStationTimetableMaster(c) })
		train.GET("/stationTimetableObjects/:id", func(c *gin.Context) { ha.GetStationTimetableObjectMaster(c) })
		train.GET("/trains", func(c *gin.Context) { ha.GetTrains(c) })
		train.GET("/trains/:sameAs", func(c *gin.Context) { ha.GetTrain(c) })
		train.GET("/trainTimetables/:trainSameAs", func(c *gin.Context) { ha.GetTrainTimetable(c) })
		train.GET("/trainInformations", func(c *gin.Context) { ha.GetTrainInformations(c) })
		train.GET("/trainInformations/:sameAs", func(c *gin.Context) { ha.GetTrainInformation(c) })
	}

	var mem runtime.MemStats
	runtime.ReadMemStats(&mem)
	fmt.Printf("Alloc:%v TotalAlloc:%v HeapAlloc:%v HeapSys:%v", mem.Alloc, mem.TotalAlloc, mem.HeapAlloc, mem.HeapSys)

	router.Run(":" + config.Config.Server.Port)
}

func connectDB() *gorm.DB {
	conn := "root:zaqroot@tcp(" + config.Config.DB.Host + ":" + config.Config.DB.Port + ")/" + config.Config.DB.Name + "?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open("mysql", conn)
	if err != nil {
		panic(err.Error())
	}
	db.DB().SetMaxIdleConns(config.Config.DB.Pool)
	db.LogMode(true)
	db.AutoMigrate(
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
		master.StationTimetableMaster{},
		master.StationTimetableMasterObject{},
		master.StationTimetableMasterObjectDestinationStation{},
		master.StationTimetableMasterObjectOriginStation{},
		master.StationTimetableMasterObjectTrainName{},
		master.StationTimetableMasterObjectViaRailway{},
		master.StationTimetableMasterObjectViaStation{},
		master.TrainTimetableMaster{},
		master.TrainTimetableMasterObject{},
		master.TrainTimetableMasterDestinationStation{},
		master.TrainTimetableMasterOriginStation{},
		master.TrainTimetableMasterNext{},
		master.TrainTimetableMasterPrevious{},
		master.TrainTimetableMasterTrainName{},
		master.TrainTimetableMasterViaRailway{},
		master.TrainTimetableMasterViaStation{},
		master.TrainTypeMaster{},
		tran.TrainTran{},
		tran.TrainTranDestinationStation{},
		tran.TrainTranOriginStation{},
		tran.TrainTranTrainName{},
		tran.TrainTranViaRailway{},
		tran.TrainTranViaStation{},
		tran.TrainInformationTran{},
		tran.TrainInformationTranRailway{},
	)
	return db
}
