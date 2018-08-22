package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/yuki-toida/refodpt/backend/infrastructure/repository"
	"github.com/yuki-toida/refodpt/backend/usecase/cache"
	"github.com/yuki-toida/refodpt/backend/usecase/tran"
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

func (h *Handler) GetUpdateTime(c *gin.Context) {
	data := tran.NewUseCase(h.r).GetUpdateTime()
	handle(c, data, nil)
}

func (h *Handler) GetPassengerSurvey(c *gin.Context) {
	sameAs := c.Param("sameAs")
	data, err := cache.NewUseCase(h.r.Cache).GetPassengerSurvey(sameAs)
	handle(c, data, err)
}

func (h *Handler) GetRailways(c *gin.Context) {
	data := cache.NewUseCase(h.r.Cache).GetRailways()
	handle(c, data, nil)
}

func (h *Handler) GetRailway(c *gin.Context) {
	sameAs := c.Param("sameAs")
	data, err := cache.NewUseCase(h.r.Cache).GetRailway(sameAs)
	handle(c, data, err)
}

func (h *Handler) GetStation(c *gin.Context) {
	sameAs := c.Param("sameAs")
	data, err := cache.NewUseCase(h.r.Cache).GetStation(sameAs)
	handle(c, data, err)
}

func (h *Handler) GetStationTimetable(c *gin.Context) {
	sameAs := c.Param("sameAs")
	data, err := tran.NewUseCase(h.r).GetStationTimetable(sameAs)
	handle(c, data, err)
}

func (h *Handler) GetStationTimetableObject(c *gin.Context) {
	ID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		handle(c, nil, err)
		return
	}
	data, err := tran.NewUseCase(h.r).GetStationTimetableObject(ID)
	handle(c, data, err)
}

func (h *Handler) GetTrains(c *gin.Context) {
	data := tran.NewUseCase(h.r).GetTrains()
	handle(c, data, nil)
}

func (h *Handler) GetTrain(c *gin.Context) {
	sameAs := c.Param("sameAs")
	data, err := tran.NewUseCase(h.r).GetTrain(sameAs)
	handle(c, data, err)
}

func (h *Handler) GetTrainTimetable(c *gin.Context) {
	trainSameAs := c.Param("trainSameAs")
	data, err := tran.NewUseCase(h.r).GetTrainTimetable(trainSameAs)
	handle(c, data, err)
}

func (h *Handler) GetTrainInformations(c *gin.Context) {
	data := tran.NewUseCase(h.r).GetTrainInformations()
	handle(c, data, nil)
}

func (h *Handler) GetTrainInformation(c *gin.Context) {
	sameAs := c.Param("sameAs")
	data, err := tran.NewUseCase(h.r).GetTrainInformation(sameAs)
	handle(c, data, err)
}
