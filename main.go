package main

import (
	"net/http"
	"os"
	"os/signal"

	"./cache"
	d "./database"
	"./handler"
	"github.com/gin-gonic/gin"
)

const port = "8080"

func main() {
	engine := gin.Default()                 // 生成一个引擎
	engine.LoadHTMLGlob("template/*")       // 加载前端
	for _, h := range handler.HandlerList { // 注册各类handler
		engine.Handle(h.GetMethod(), h.GetPath(), h.GetHandle())
	}
	srv := &http.Server{ //定义一个原生server，将gin引擎设为该sever的handler
		Addr:    ":" + port,
		Handler: engine,
	}
	c := make(chan os.Signal, 1)            // 定义一个通道以获取终止信号
	signal.Notify(c, os.Interrupt, os.Kill) // 当系统对程序发出终止或杀死信号时，对通道c写入该信号
	go srv.ListenAndServe()                 // 在协程中监听服务
	<-c                                     // 主线程等待终止信号
	srv.Close()                             // 收到信号后，关闭server
	cache.CloseInstance()                   // 关闭redis连接池
	d.CloseInstance()                       // 关闭数据库连接池
}
