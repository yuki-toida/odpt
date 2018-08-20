package main

import (
	"fmt"
	"net/http"
	"runtime"

	"github.com/bamzi/jobrunner"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/yuki-toida/refodpt/backend/config"
	"github.com/yuki-toida/refodpt/backend/infrastructure/cache"
	"github.com/yuki-toida/refodpt/backend/infrastructure/db"
	"github.com/yuki-toida/refodpt/backend/infrastructure/repository"
	"github.com/yuki-toida/refodpt/backend/interface/handler"
	"github.com/yuki-toida/refodpt/backend/interface/job/importer"
	"github.com/yuki-toida/refodpt/backend/interface/registry"
)

func init() {
	config.Init("./", "./.env")
}

func main() {
	db := db.Connect()
	db.LogMode(true)

	cc := cache.NewCache()
	cc.Init(db)
	r := repository.NewRepository(db, cc)
	jobrunner.Start()
	jobrunner.Schedule("@every 2m", importer.NewImporter(r))

	h := handler.NewHandler(registry.NewRegistry(r))
	router := gin.Default()
	router.Use(cors.Default())

	router.GET("/", func(c *gin.Context) { c.String(http.StatusOK, "OK") })
	router.GET("/healthz", func(c *gin.Context) { c.String(http.StatusOK, "OK") })
	router.GET("/jobrunner", func(c *gin.Context) { c.JSON(http.StatusOK, jobrunner.StatusJson()) })

	t := router.Group("/train")
	{
		t.GET("/railways", func(c *gin.Context) { h.GetRailwayMasters(c) })
		t.GET("/railways/:sameAs", func(c *gin.Context) { h.GetRailwayMaster(c) })
		t.GET("/stations/:sameAs", func(c *gin.Context) { h.GetStationMaster(c) })
		t.GET("/passengerSurveys/:sameAs", func(c *gin.Context) { h.GetPassengerSurveyMaster(c) })
		t.GET("/stationTimetables/:sameAs", func(c *gin.Context) { h.GetStationTimetableMaster(c) })
		t.GET("/stationTimetableObjects/:id", func(c *gin.Context) { h.GetStationTimetableObjectMaster(c) })
		t.GET("/trains", func(c *gin.Context) { h.GetTrains(c) })
		t.GET("/trains/:sameAs", func(c *gin.Context) { h.GetTrain(c) })
		t.GET("/trainTimetables/:trainSameAs", func(c *gin.Context) { h.GetTrainTimetable(c) })
		t.GET("/trainInformations", func(c *gin.Context) { h.GetTrainInformations(c) })
		t.GET("/trainInformations/:sameAs", func(c *gin.Context) { h.GetTrainInformation(c) })
	}

	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("Alloc:%v TotalAlloc:%v HeapAlloc:%v HeapSys:%v", m.Alloc, m.TotalAlloc, m.HeapAlloc, m.HeapSys)

	router.Run(":" + config.Config.Server.Port)
}
