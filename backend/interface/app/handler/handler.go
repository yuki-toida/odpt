package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yuki-toida/refodpt/backend/infrastructure/repository"
	"github.com/yuki-toida/refodpt/backend/usecase/route"
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

func (h *Handler) GetConnectingStations(c *gin.Context) {
	data := route.NewUseCase(h.r).GetConnectingStations()
	handle(c, data, nil)
}

func (h *Handler) GetRoutes(c *gin.Context) {
	params := &struct{ From, To string }{}
	if err := c.BindJSON(params); err != nil {
		handle(c, nil, err)
	} else {
		data, err := route.NewUseCase(h.r).GetRoutes(params.From, params.To)
		handle(c, data, err)
	}
}
