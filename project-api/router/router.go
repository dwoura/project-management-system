package router

import (
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"log"
	"net"
	"test.com/project-user/config"
	loginServiceV1 "test.com/project-user/pkg/service/login_service.v1"
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

type gRPCConfig struct {
	Addr         string
	RegisterFunc func(*grpc.Server)
}

func RegisterGrpc() *grpc.Server {
	c := gRPCConfig{
		Addr: config.C.GC.Addr,
		RegisterFunc: func(g *grpc.Server) {
			loginServiceV1.RegisterLoginServiceServer(g, loginServiceV1.New())
		},
	}
	s := grpc.NewServer()
	c.RegisterFunc(s)
	lis, err := net.Listen("tcp", c.Addr)
	if err != nil {
		log.Println("cannot listen")
	}
	go func() {
		err = s.Serve(lis)
		if err != nil {
			log.Println("server started error", err)
			return
		}
	}()
	return s
}
