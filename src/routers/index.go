package routers

import (
	"github.com/gin-gonic/gin"
)

// 将每个模块的路由路径进行拆分，在这儿统一初始化，gin框架的路由启动
func InitRouter(r *gin.Engine) *gin.Engine {
	group := r.Group("api")
	{
		// 用户表的路由
		UserRouter(group.Group("student"))
		// 新文档的路由
		PostRouter(group.Group("post"))
		// 统计信息的路由
		StatsRouter(group.Group("stats"))
	}
	return r
}
