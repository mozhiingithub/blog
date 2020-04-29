package handler

import (
	"github.com/gin-gonic/gin"
)

var editArticle = handler{
	path:   "/article",
	method: "POST",
	handle: func(c *gin.Context) {
		postMethod := c.PostForm("_method")
		switch postMethod {
		case "POST":
			postArticle.handle(c)
		case "PUT":
			putArticle.handle(c)
		case "DELETE":
			deleteArticle.handle(c)
		}
	},
}
