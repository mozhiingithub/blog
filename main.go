package main

import "github.com/gin-gonic/gin"

func main() {
	engine := gin.Default()           // 生成一个引擎
	engine.LoadHTMLGlob("template/*") // 加载前端
	// myhandler.handle(engine) 注册各类handler
	engine.Run() // 启动引擎
}
