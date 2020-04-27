package main

import (
	"./handler"
	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default() // 生成一个引擎
	//engine.LoadHTMLGlob("template/*") // 加载前端
	handler.HandleList(engine)
	//engine.Run() // 启动引擎
}
