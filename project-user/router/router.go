package router

import (
	"github.com/gin-gonic/gin"
)

// Router 接口
type Router interface {
	Route(r *gin.Engine)
}

// 路由注册器实体
type RegisterRouter struct {
}

// 路由注册器实例化
func New() *RegisterRouter {
	return &RegisterRouter{}
}

// 路由注册器的注册方法
func (*RegisterRouter) Route(ro Router, r *gin.Engine) {
	ro.Route(r)
}

var routers []Router

func Register(ro ...Router) {
	routers = append(routers, ro...)
}

func InitRouter(r *gin.Engine) {
	//rg := New()
	//为每个用户注册一个路由，但是每次要操作这一步，还是有点麻烦了
	//rg.Route(&user.RouterUser{}, r)

	//注册所有路由
	for _, ro := range routers {
		ro.Route(r)
	}
}
