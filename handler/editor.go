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
	method: "POST",
	handle: func(c *gin.Context) {
		articleIDStr := c.PostForm("id") // 获取需编辑博文的id。发布新博文时，则不发送id号
		articleID, _ := strconv.Atoi(articleIDStr)
		articleURL := fmt.Sprintf("/article?id=%d", articleID) // 博文URI
		method := "POST"                                       // 对博文的方法。POST对应发布新博文，PUT对应更新原有博文
		titleValue := ""
		contentValue := ""
		for 0 != articleID { // id非0,表明不是发布新博文
			var ok bool
			titleValue, ok = d.GetTitle(articleID)
			if !ok { // 库中没有该id
				break
			}
			contentValue, ok = d.GetContent(articleID)
			if !ok { // 库中没有该id对应的正文
				break
			}
			method = "PUT" // 成功获取id对应标题和正文，将方法给为PUT，对应更新操作
			break
		}
		c.HTML(http.StatusOK, "editor.html", gin.H{
			"articleURL":   articleURL,
			"method":       method,
			"titleValue":   titleValue,
			"contentValue": contentValue,
		})
	},
}
