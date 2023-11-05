package main

import (
	"github.com/gin-gonic/gin"
	srv "test.com/project-common"
	"test.com/project-user/router"
)

func main() {
	r := gin.Default()
	//r.Run(":8080") //不优雅
	router.InitRouter(r)
	//优雅启停项目
	srv.Run(r, "my-projects-user", ":8082")
}
