package router

import (
	"net/http"

	"github.com/bamzi/jobrunner"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/yuki-toida/refodpt/backend/infrastructure/repository"
	"github.com/yuki-toida/refodpt/backend/interface/admin/handler"
)

func Create(r *repository.Repository) http.Handler {
	h := handler.NewHandler(r)

	e := gin.Default()
	e.Use(cors.Default())
	e.GET("/", func(c *gin.Context) { c.String(http.StatusOK, "OK") })
	e.GET("/healthz", func(c *gin.Context) { c.String(http.StatusOK, "OK") })
	e.GET("/jobrunner", func(c *gin.Context) { c.JSON(http.StatusOK, jobrunner.StatusJson()) })
	e.GET("/time", func(c *gin.Context) { h.GetUpdateTime(c) })

	t := e.Group("/train")
	{
		t.GET("/railways", func(c *gin.Context) { h.GetRailways(c) })
		t.GET("/railways/:sameAs", func(c *gin.Context) { h.GetRailway(c) })
		t.GET("/stations/:sameAs", func(c *gin.Context) { h.GetStation(c) })
		t.GET("/passengerSurveys/:sameAs", func(c *gin.Context) { h.GetPassengerSurvey(c) })
		t.GET("/stationTimetables/:sameAs", func(c *gin.Context) { h.GetStationTimetable(c) })
		t.GET("/stationTimetableObjects/:id", func(c *gin.Context) { h.GetStationTimetableObject(c) })
		t.GET("/trains", func(c *gin.Context) { h.GetTrains(c) })
		t.GET("/trains/:sameAs", func(c *gin.Context) { h.GetTrain(c) })
		t.GET("/trainTimetables/:trainSameAs", func(c *gin.Context) { h.GetTrainTimetable(c) })
		t.GET("/trainInformations", func(c *gin.Context) { h.GetTrainInformations(c) })
		t.GET("/trainInformations/:sameAs", func(c *gin.Context) { h.GetTrainInformation(c) })
	}

	return e
}
