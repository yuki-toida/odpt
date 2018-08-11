package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yuki-toida/refodpt/backend/interface/registry"
	"github.com/yuki-toida/refodpt/backend/usecase/shared"
	"github.com/yuki-toida/refodpt/backend/usecase/train"
)

// Handler type
type Handler struct {
	registry *registry.Registry
}

// NewHandler func
func NewHandler(r *registry.Registry) *Handler {
	return &Handler{
		registry: r,
	}
}

func handle(c *gin.Context, data interface{}, err error) {
	res := gin.H{"data": data}
	if err != nil {
		res["error"] = err.Error()
	}
	c.JSON(http.StatusOK, res)
}

// GetCategoryMasters func
func (h *Handler) GetCategoryMasters(c *gin.Context) {
	data := shared.NewUseCase(h.registry.Repository).GetCategoryMasters()
	handle(c, data, nil)
}

// GetRailwayMasters func
func (h *Handler) GetRailwayMasters(c *gin.Context) {
	data := train.NewUseCase(h.registry.Repository).GetRailwayMasters()
	handle(c, data, nil)
}

// GetRailwayMaster func
func (h *Handler) GetRailwayMaster(c *gin.Context) {
	sameAs := c.Param("sameAs")
	data, err := train.NewUseCase(h.registry.Repository).GetRailwayMaster(sameAs)
	handle(c, data, err)
}

// GetStationMasters func
func (h *Handler) GetStationMasters(c *gin.Context) {
	data := train.NewUseCase(h.registry.Repository).GetStationMasters()
	handle(c, data, nil)
}

// GetStationMaster func
func (h *Handler) GetStationMaster(c *gin.Context) {
	sameAs := c.Param("sameAs")
	data, err := train.NewUseCase(h.registry.Repository).GetStationMaster(sameAs)
	handle(c, data, err)
}
