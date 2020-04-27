package main

import (
	d "./database"
	"./handler"
	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default() // 生成一个引擎
	//engine.LoadHTMLGlob("template/*") // 加载前端
	for _, h := range handler.HandlerList { // 注册各类handler
		engine.Handle(h.GetMethod(), h.GetPath(), h.GetHandle())
	}
	defer d.CloseInstance() // 后台结束时，关闭数据库连接池
	engine.Run()            // 启动引擎
}
