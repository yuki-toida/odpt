package main

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/yuki-toida/refodpt/backend/config"
	"github.com/yuki-toida/refodpt/backend/interface/handler"
	"github.com/yuki-toida/refodpt/backend/interface/registry"
)

func init() {
	config.Init()
}

func main() {
	// conn := "root:zaqroot@tcp(" + config.Config.Db.Host + ":" + config.Config.Db.Port + ")/" + config.Config.Db.Name + "?charset=utf8&parseTime=True&loc=Local"
	// db, err := gorm.Open("mysql", conn)
	// if err != nil {
	// 	panic(err.Error())
	// }
	// db.LogMode(true)
	// db.AutoMigrate(&model.User{}, &model.Event{})

	registry := registry.NewRegistry()
	handler := handler.NewHandler(registry)

	router := gin.Default()
	router.Use(cors.Default())

	router.GET("/healthz", func(c *gin.Context) { c.String(http.StatusOK, "OK") })

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
