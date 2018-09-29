package main

import (
	"fmt"
	"net/http"
	"runtime"
	"testing"

	"github.com/bamzi/jobrunner"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/yuki-toida/refodpt/backend/config"
	"github.com/yuki-toida/refodpt/backend/infrastructure/cache"
	"github.com/yuki-toida/refodpt/backend/infrastructure/db"
	"github.com/yuki-toida/refodpt/backend/infrastructure/job/importer"
	"github.com/yuki-toida/refodpt/backend/infrastructure/repository"
	admin "github.com/yuki-toida/refodpt/backend/interface/admin/handler"
	app "github.com/yuki-toida/refodpt/backend/interface/app/handler"
)

func init() {
	config.Init("./", "./.env")
}

func main() {
	db := db.Connect()
	db.LogMode(true)

	cc := cache.NewCache()
	tr := testing.Benchmark(func(b *testing.B) {
		cc.Init(db)
	})
	fmt.Printf("[Benchmark] %v\n", tr)

	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("[MemStats] Alloc:%v TotalAlloc:%v HeapAlloc:%v HeapSys:%v\n", m.Alloc, m.TotalAlloc, m.HeapAlloc, m.HeapSys)

	repo := repository.NewRepository(db, cc)
	jobrunner.Start()
	jobrunner.Schedule("@every 15m", importer.NewImporter(repo))

	r := gin.Default()
	r.Use(cors.Default())

	r.GET("/", func(c *gin.Context) { c.String(http.StatusOK, "OK") })
	r.GET("/healthz", func(c *gin.Context) { c.String(http.StatusOK, "OK") })

	ap := r.Group("/app")
	{
		aph := app.NewHandler(repo)
		ap.GET("/connectings", func(c *gin.Context) { aph.GetConnectingStations(c) })
		ap.POST("/routes", func(c *gin.Context) { aph.GetRoutes(c) })
	}

	ad := r.Group("/admin")
	{
		adh := admin.NewHandler(repo)
		ad.GET("/jobrunner", func(c *gin.Context) { c.JSON(http.StatusOK, jobrunner.StatusJson()) })
		ad.GET("/time", func(c *gin.Context) { adh.GetUpdateTime(c) })
		ad.GET("/railways", func(c *gin.Context) { adh.GetRailways(c) })
		ad.GET("/railways/:sameAs", func(c *gin.Context) { adh.GetRailway(c) })
		ad.GET("/stations/:sameAs", func(c *gin.Context) { adh.GetStation(c) })
		ad.GET("/passengerSurveys/:sameAs", func(c *gin.Context) { adh.GetPassengerSurvey(c) })
		ad.GET("/stationTimetables/:sameAs", func(c *gin.Context) { adh.GetStationTimetable(c) })
		ad.GET("/stationTimetableObjects/:id", func(c *gin.Context) { adh.GetStationTimetableObject(c) })
		ad.GET("/trains", func(c *gin.Context) { adh.GetTrains(c) })
		ad.GET("/trains/:sameAs", func(c *gin.Context) { adh.GetTrain(c) })
		ad.GET("/trainTimetables/:trainSameAs", func(c *gin.Context) { adh.GetTrainTimetable(c) })
		ad.GET("/trainInformations", func(c *gin.Context) { adh.GetTrainInformations(c) })
		ad.GET("/trainInformations/:sameAs", func(c *gin.Context) { adh.GetTrainInformation(c) })
	}

	r.Run(":" + config.Config.APIServer.Port)
}
