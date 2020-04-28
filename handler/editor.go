package handler

import (
	"fmt"
	"net/http"
	"strconv"

	d "../database"
	"github.com/gin-gonic/gin"
)

var editor = handler{
	path:   "/editor",
	method: "GET",
	handle: func(c *gin.Context) {
		articleIDStr := c.DefaultQuery("id", "0") // 获取需编辑博文的id。发布新博文时，则不发送id号
		articleID, _ := strconv.Atoi(articleIDStr)
		articleURL := fmt.Sprintf("/article?id=%d", articleID) // 博文URI
		method := "POST"                                       // 对博文的方法。POST对应发布新博文，PUT对应更新原有博文
		titleValue := "在此处输入标题。。。"
		contentValue := "在此处输入正文。。。"
		if 0 != articleID { // id非0,表明不是发布新博文
			db := d.GetInstance()                                                   // 获取数据库连接池实例
			rows, _ := db.Query("select title from titles where id = ?", articleID) // 查询此博文是否存在
			if rows.Next() {                                                        // 查询结果非空，表明此博文存在
				rows.Scan(&titleValue)                                                     // 获取博文标题
				rows, _ = db.Query("select content from contents where id = ?", articleID) // 查询博文正文
				if rows.Next() {
					rows.Scan(&contentValue) // 获取博文正文
					method = "PUT"           // 确认此博文存在，将http方法改为PUT
				}
			}
		}
		c.HTML(http.StatusOK, "editor.html", gin.H{
			"articleURL":   articleURL,
			"method":       method,
			"titleValue":   titleValue,
			"contentValue": contentValue,
		})
	},
}
