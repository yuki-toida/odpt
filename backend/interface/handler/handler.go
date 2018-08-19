package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/yuki-toida/refodpt/backend/interface/registry"
	"github.com/yuki-toida/refodpt/backend/usecase/cache"
	"github.com/yuki-toida/refodpt/backend/usecase/tran"
)

type Handler struct {
	registry *registry.Registry
}

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

func (h *Handler) GetPassengerSurveyMaster(c *gin.Context) {
	sameAs := c.Param("sameAs")
	data, err := cache.NewUseCase(h.registry.Repository.Cache).GetPassengerSurvey(sameAs)
	handle(c, data, err)
}

func (h *Handler) GetRailwayMasters(c *gin.Context) {
	data := cache.NewUseCase(h.registry.Repository.Cache).GetRailways()
	handle(c, data, nil)
}

func (h *Handler) GetRailwayMaster(c *gin.Context) {
	sameAs := c.Param("sameAs")
	data, err := cache.NewUseCase(h.registry.Repository.Cache).GetRailway(sameAs)
	handle(c, data, err)
}

func (h *Handler) GetStationMaster(c *gin.Context) {
	sameAs := c.Param("sameAs")
	data, err := cache.NewUseCase(h.registry.Repository.Cache).GetStation(sameAs)
	handle(c, data, err)
}

func (h *Handler) GetStationTimetableMaster(c *gin.Context) {
	sameAs := c.Param("sameAs")
	data, err := tran.NewUseCase(h.registry.Repository).GetStationTimetableMaster(sameAs)
	handle(c, data, err)
}

func (h *Handler) GetStationTimetableObjectMaster(c *gin.Context) {
	ID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		handle(c, nil, err)
		return
	}
	data, err := tran.NewUseCase(h.registry.Repository).GetStationTimetableMasterObject(ID)
	handle(c, data, err)
}

func (h *Handler) GetTrains(c *gin.Context) {
	data := tran.NewUseCase(h.registry.Repository).GetTrains()
	handle(c, data, nil)
}

func (h *Handler) GetTrain(c *gin.Context) {
	sameAs := c.Param("sameAs")
	data, err := tran.NewUseCase(h.registry.Repository).GetTrain(sameAs)
	handle(c, data, err)
}

func (h *Handler) GetTrainTimetable(c *gin.Context) {
	trainSameAs := c.Param("trainSameAs")
	data, err := tran.NewUseCase(h.registry.Repository).GetTrainTimetable(trainSameAs)
	handle(c, data, err)
}

func (h *Handler) GetTrainInformations(c *gin.Context) {
	data := tran.NewUseCase(h.registry.Repository).GetTrainInformations()
	handle(c, data, nil)
}

func (h *Handler) GetTrainInformation(c *gin.Context) {
	sameAs := c.Param("sameAs")
	data, err := tran.NewUseCase(h.registry.Repository).GetTrainInformation(sameAs)
	handle(c, data, err)
}
