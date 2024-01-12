package main

import (
	"github.com/gin-gonic/gin"
	_ "test.com/project-api/api"
	srv "test.com/project-common"
	"test.com/project-user/config"
	"test.com/project-user/router"
)

func main() {
	r := gin.Default()
	//从配置中读取日志配置，初始化日志
	router.InitRouter(r)

	srv.Run(r, config.C.SC.Name, config.C.SC.Addr, nil)
}
