package main

import (
	"./handler"
	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default() // 生成一个引擎
	//engine.LoadHTMLGlob("template/*") // 加载前端
	for _, h := range handler.HandlerList { // 注册各类handler
		engine.Handle(h.GetMethod(), h.GetPath(), h.GetHandle())
	}
	engine.Run() // 启动引擎
}
