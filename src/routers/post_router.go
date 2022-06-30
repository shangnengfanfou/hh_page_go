package routers

import (
	"github.com/gin-gonic/gin"
)

func PostRouter(g *gin.RouterGroup) {
	// 因为我在controller将每个controller类实现了初始化，完成多态的实现，这样即使多个不同类中的方法名一样也没有问题，详情见controller层代码
}
