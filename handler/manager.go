package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var manager = handler{
	path:   "/manager",
	method: "GET",
	handle: func(c *gin.Context) {
		c.HTML(http.StatusOK, "manager.html", gin.H{})
	},
}
