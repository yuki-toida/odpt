package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/yuki-toida/refodpt/backend/interface/registry"
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

// GetStationMaster func
func (h *Handler) GetStationMaster(c *gin.Context) {
	sameAs := c.Param("sameAs")
	data, err := train.NewUseCase(h.registry.Repository).GetStationMaster(sameAs)
	handle(c, data, err)
}

// GetPassengerSurveyMaster func
func (h *Handler) GetPassengerSurveyMaster(c *gin.Context) {
	sameAs := c.Param("sameAs")
	data, err := train.NewUseCase(h.registry.Repository).GetPassengerSurveyMaster(sameAs)
	handle(c, data, err)
}

// GetStationTimetableMaster func
func (h *Handler) GetStationTimetableMaster(c *gin.Context) {
	sameAs := c.Param("sameAs")
	data, err := train.NewUseCase(h.registry.Repository).GetStationTimetableMaster(sameAs)
	handle(c, data, err)
}

// GetStationTimetableObjectMaster func
func (h *Handler) GetStationTimetableObjectMaster(c *gin.Context) {
	ID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		handle(c, nil, err)
		return
	}
	data, err := train.NewUseCase(h.registry.Repository).GetStationTimetableObjectMaster(ID)
	handle(c, data, err)
}

// GetTrains func
func (h *Handler) GetTrains(c *gin.Context) {
	data := train.NewUseCase(h.registry.Repository).GetTrains()
	handle(c, data, nil)
}

// GetTrain func
func (h *Handler) GetTrain(c *gin.Context) {
	sameAs := c.Param("sameAs")
	data, err := train.NewUseCase(h.registry.Repository).GetTrain(sameAs)
	handle(c, data, err)
}
