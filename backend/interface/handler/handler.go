package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yuki-toida/refodpt/backend/interface/registry"
	"github.com/yuki-toida/refodpt/backend/usecase/raw"
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
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"data": data})
	}
}

// GetCalendar func
func (h *Handler) GetCalendar(c *gin.Context) {
	data, err := raw.NewUseCase().GetCalendar()
	handle(c, data, err)
}

// GetOperator func
func (h *Handler) GetOperator(c *gin.Context) {
	data, err := raw.NewUseCase().GetOperator()
	handle(c, data, err)
}

// GetPassengerSurvey func
func (h *Handler) GetPassengerSurvey(c *gin.Context) {
	data, err := raw.NewUseCase().GetPassengerSurvey()
	handle(c, data, err)
}

// GetRailDirection func
func (h *Handler) GetRailDirection(c *gin.Context) {
	data, err := raw.NewUseCase().GetRailDirection()
	handle(c, data, err)
}

// GetRailway func
func (h *Handler) GetRailway(c *gin.Context) {
	data, err := raw.NewUseCase().GetRailway()
	handle(c, data, err)
}

// GetRailwayFare func
func (h *Handler) GetRailwayFare(c *gin.Context) {
	data, err := raw.NewUseCase().GetRailwayFare()
	handle(c, data, err)
}

// GetStation func
func (h *Handler) GetStation(c *gin.Context) {
	data, err := raw.NewUseCase().GetStation()
	handle(c, data, err)
}

// GetStationTimetable func
func (h *Handler) GetStationTimetable(c *gin.Context) {
	data, err := raw.NewUseCase().GetStationTimetable()
	handle(c, data, err)
}
