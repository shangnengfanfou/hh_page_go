package routers

import (
	"hh_page_go/src/apis"

	"github.com/gin-gonic/gin"
)

func StatsRouter(g *gin.RouterGroup) {
	c := apis.NewStatsApi()
	{
		g.GET("/overview", c.CountAndSumByTime)
	}
}
