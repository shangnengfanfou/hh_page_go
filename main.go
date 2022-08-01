package main

import (
	"hh_page_go/src/routers"
	"hh_page_go/src/utils"
	"log"

	"github.com/gin-gonic/gin"
	"gopkg.in/ini.v1"
)

func main() {
	r := gin.Default()

	// 读取配置文件
	conf, err := ini.Load("/data/apps/gosrv/config/my.ini")
	if err != nil {
		log.Fatal("配置文件读取失败, err = ", err)
	}

	// 初始化路由路径
	routers.InitRouter(r)

	// 初始化数据库
	db := utils.SqlClient(conf)
	defer db.Close()

	// 项目的启动端口在这儿设置，也可以从配置文件中读取到这儿
	port := ":" + conf.Section("").Key("httpport").String()
	if err := r.Run(port); err != nil {
		log.Fatal("程序启动失败:", err)
	}
}
