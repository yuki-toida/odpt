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

func handleError(c *gin.Context, err error) {
	c.JSON(http.StatusOK, gin.H{"error": err.Error()})
}

// GetCalendar func
func (h *Handler) GetCalendar(c *gin.Context) {
	data, err := shared.NewUseCase().GetCalendar()
	if err != nil {
		handleError(c, err)
	} else {
		c.JSON(http.StatusOK, gin.H{"data": data})
	}
}

// GetOperator func
func (h *Handler) GetOperator(c *gin.Context) {
	data, err := shared.NewUseCase().GetOperator()
	if err != nil {
		handleError(c, err)
	} else {
		c.JSON(http.StatusOK, gin.H{"data": data})
	}
}

// GetPassengerSurvey func
func (h *Handler) GetPassengerSurvey(c *gin.Context) {
	data, err := train.NewUseCase().GetPassengerSurvey()
	if err != nil {
		handleError(c, err)
	} else {
		c.JSON(http.StatusOK, gin.H{"data": data})
	}
}

// GetRailDirection func
func (h *Handler) GetRailDirection(c *gin.Context) {
	data, err := train.NewUseCase().GetRailDirection()
	if err != nil {
		handleError(c, err)
	} else {
		c.JSON(http.StatusOK, gin.H{"data": data})
	}
}

// GetRailway func
func (h *Handler) GetRailway(c *gin.Context) {
	data, err := train.NewUseCase().GetRailway()
	if err != nil {
		handleError(c, err)
	} else {
		c.JSON(http.StatusOK, gin.H{"data": data})
	}
}
