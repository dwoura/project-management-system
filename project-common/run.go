package common

import (
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Run(r *gin.Engine, srvName string, addr string, stop func()) {
	//优雅启停项目
	srv := &http.Server{
		Addr:    addr,
		Handler: r,
	}
	go func() {
		//使用协程不会阻塞程序
		log.Printf("web server: %s running in %s \n", srvName, srv.Addr)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalln(err)
		}
	}()

	//等待关闭信号，阻塞程序
	//SIGINT 用户发送INTR字符(Ctrl+C)触发
	//SIGTERM 结束程序(可以被捕获、阻塞或忽略)
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit //阻塞进程直到有信号
	log.Printf("Shutting Down project web server: %s ...\n", srvName)
	//由于关闭进行了延迟，我们需要一个上下文ctx
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	if stop != nil {
		stop()
	}
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("web server: %s Shutdown, cause by : \n", srvName, err)
	}
	select {
	case <-ctx.Done():
		log.Println("wait timeout...")
	}
	log.Println("web server: %s stop success... \n", srvName)
}
