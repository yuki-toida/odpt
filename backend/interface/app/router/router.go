package router

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/yuki-toida/refodpt/backend/infrastructure/repository"
	"github.com/yuki-toida/refodpt/backend/interface/app/handler"
)

func Create(r *repository.Repository) http.Handler {
	h := handler.NewHandler(r)

	e := gin.Default()
	e.Use(cors.Default())
	e.GET("/", func(c *gin.Context) { c.String(http.StatusOK, "OK") })
	e.GET("/healthz", func(c *gin.Context) { c.String(http.StatusOK, "OK") })

	t := e.Group("/train")
	{
		t.GET("/railways", func(c *gin.Context) { h.GetRailways(c) })
	}

	return e
}
