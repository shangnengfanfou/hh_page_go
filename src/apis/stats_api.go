package apis

import (
	"hh_page_go/src/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type StatsApi struct {
	statsService *services.StatsService
}

func NewStatsApi() *StatsApi {
	return &StatsApi{statsService: services.NewStatsService()}
}

func (h *StatsApi) CountAndSumByTime(c *gin.Context) {
	sTime, err := strconv.Atoi(c.Query("startTime"))
	eTime, err := strconv.Atoi(c.Query("endTime"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": "500", "msg": err.Error()})
		return
	}
	ret, err := h.statsService.CountAndSumByTime(sTime, eTime)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": "500", "msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": "200", "msg": "OK", "data": ret})
}
