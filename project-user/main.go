package main

import (
	"github.com/gin-gonic/gin"
	srv "test.com/project-common"
	_ "test.com/project-user/api/user" //用于触发user包中的init函数
	"test.com/project-user/router"
)

func main() {
	r := gin.Default()
	//r.Run(":8080") //不优雅
	router.InitRouter(r)
	//优雅启停项目
	srv.Run(r, "my-projects-user", ":8082")
}
