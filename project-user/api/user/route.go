package user

import "github.com/gin-gonic/gin"

type RouterUser struct {
}

// 路径绑定
func (*RouterUser) Route(r *gin.Engine) {
	h := &HandlerUser{}
	//handler带有context参数,只需要传入方法名
	r.POST("/project/login/getCaptcha", h.getCaptcha)
}
