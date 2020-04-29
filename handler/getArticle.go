package handler

import (
	"net/http"
	"strconv"

	d "../database"
	"github.com/gin-gonic/gin"
)

var getArticle = handler{
	path:   "/article",
	method: "GET",
	handle: func(c *gin.Context) {
		for {
			articleIDStr := c.Query("id") // 获取博文id参数
			if "" == articleIDStr {
				break
			}
			articleID, err := strconv.Atoi(articleIDStr)
			if nil != err {
				break
			}
			title, ok := d.GetTitle(articleID) // 获取此id对应的博文标题
			if !ok {
				break
			}
			t, ok := d.GetTime(articleID) // 获取此id对应的博文发布时间
			if !ok {
				break
			}
			content, ok := d.GetContent(articleID) // 获取此id对应的博文正文
			if !ok {
				break
			}
			c.HTML(http.StatusOK, "article.html", gin.H{
				"title":   title,
				"t":       t,
				"content": content,
			})
			return
		}
		c.HTML(http.StatusNotFound, "result.html", gin.H{
			"result": "博文不存在",
		})
		return
	},
}
