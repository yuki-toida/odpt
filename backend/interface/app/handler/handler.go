package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yuki-toida/refodpt/backend/infrastructure/repository"
	"github.com/yuki-toida/refodpt/backend/usecase/cache"
)

type Handler struct {
	r *repository.Repository
}

func NewHandler(r *repository.Repository) *Handler {
	return &Handler{
		r: r,
	}
}

func handle(c *gin.Context, data interface{}, err error) {
	res := gin.H{"data": data}
	if err != nil {
		res["error"] = err.Error()
	}
	c.JSON(http.StatusOK, res)
}

func (h *Handler) GetRailways(c *gin.Context) {
	data := cache.NewUseCase(h.r.Cache).GetRailways()
	handle(c, data, nil)
}
