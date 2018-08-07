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
	db.AutoMigrate(model.Calendar{}, model.CalendarDay{}, model.Operator{})

	cr := repository.NewRepository(db)

	registry := registry.NewRegistry(cr)
	handler := handler.NewHandler(registry)

	jobrunner.Start()
	// jobrunner.Schedule("@every 1m", importer.NewImporter(registry))
	importer.NewImporter(registry).Run()

	router := gin.Default()
	router.Use(cors.Default())

	router.GET("/healthz", func(c *gin.Context) { c.String(http.StatusOK, "OK") })
	router.GET("/jobrunner", func(c *gin.Context) { c.JSON(http.StatusOK, jobrunner.StatusJson()) })

	shared := router.Group("/shared")
	{
		shared.GET("/calendar", func(c *gin.Context) { handler.GetCalendar(c) })
		shared.GET("/operator", func(c *gin.Context) { handler.GetOperator(c) })
	}

	train := router.Group("/train")
	{
		train.GET("/passengerSurvey", func(c *gin.Context) { handler.GetPassengerSurvey(c) })
		train.GET("/railDirection", func(c *gin.Context) { handler.GetRailDirection(c) })
		train.GET("/railway", func(c *gin.Context) { handler.GetRailway(c) })
	}

	router.Run(":" + config.Config.Server.Port)
}
