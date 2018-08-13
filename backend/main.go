package main

import (
	"net/http"

	"github.com/yuki-toida/refodpt/backend/domain/model/tran"

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
	conn := "root:zaqroot@tcp(" + config.Config.DB.Host + ":" + config.Config.DB.Port + ")/" + config.Config.DB.Name + "?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open("mysql", conn)
	if err != nil {
		panic(err.Error())
	}
	db.DB().SetMaxIdleConns(config.Config.DB.Pool)
	db.LogMode(false)
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
	)

	cr := repository.NewRepository(db)

	registry := registry.NewRegistry(cr)
	handler := handler.NewHandler(registry)

	i := importer.NewImporter(registry)
	i.Run()
	// jobrunner.Start()
	// jobrunner.Schedule("@every 1m", importer.NewImporter(registry))

	// ch := make(chan int, 2)
	// wg := sync.WaitGroup{}

	// for i := 0; i < 10; i++ {
	// 	ch <- 1
	// 	wg.Add(1)
	// 	go func(s string) {
	// 		log.Println("Hello", s)
	// 		time.Sleep(time.Second) //コピペして実行する時はスリープ入れると理解しやすい
	// 		<-ch
	// 		wg.Done()
	// 	}("World" + strconv.Itoa(i))
	// }
	// wg.Wait()

	router := gin.Default()
	router.Use(cors.Default())

	router.GET("/healthz", func(c *gin.Context) { c.String(http.StatusOK, "OK") })
	router.GET("/jobrunner", func(c *gin.Context) { c.JSON(http.StatusOK, jobrunner.StatusJson()) })
	train := router.Group("/train")
	{
		train.GET("/railways", func(c *gin.Context) { handler.GetRailwayMasters(c) })
		train.GET("/railways/:sameAs", func(c *gin.Context) { handler.GetRailwayMaster(c) })
		train.GET("/stations/:sameAs", func(c *gin.Context) { handler.GetStationMaster(c) })
		train.GET("/passengerSurveys/:sameAs", func(c *gin.Context) { handler.GetPassengerSurveyMaster(c) })
		train.GET("/stationTimetables/:sameAs", func(c *gin.Context) { handler.GetStationTimetableMaster(c) })
		train.GET("/stationTimetableObjects/:id", func(c *gin.Context) { handler.GetStationTimetableObjectMaster(c) })
		train.GET("/trains", func(c *gin.Context) { handler.GetTrains(c) })
		train.GET("/trains/:sameAs", func(c *gin.Context) { handler.GetTrain(c) })
	}

	router.Run(":" + config.Config.Server.Port)
}
