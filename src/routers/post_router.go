package routers

import (
	"hh_page_go/src/apis"

	"github.com/gin-gonic/gin"
)

func PostRouter(g *gin.RouterGroup) {
	c := apis.NewArticleApi()
	{
		g.POST("add", c.Add)
		g.POST("page", c.Page)
		g.GET("view/:uid", c.IncrViewsCount)
		g.GET("info", c.Info)
	}
}
