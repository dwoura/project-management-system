package main

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

func main() {
	r := gin.Default()
	//r.Run(":8080") //不优雅

	//优雅启停项目
	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}
	go func() {
		//使用协程不会阻塞程序
		log.Printf("web server running in %s \n", srv.Addr)
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
	log.Println("Shutting Down project web server...")
	//由于关闭进行了延迟，我们需要一个上下文ctx
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalln("web server Shutdown, cause by : ", err)
	}
	select {
	case <-ctx.Done():
		log.Println("wait timeout...")
	}
	log.Println("web server stop success...")
}
