package handler

import (
	"github.com/gin-gonic/gin"
)

type handler struct {
	path   string
	method string
	handle func(c *gin.Context)
}

func (h handler) GetPath() string {
	return h.path
}

func (h handler) GetMethod() string {
	return h.method
}

func (h handler) GetHandle() func(c *gin.Context) {
	return h.handle
}

// HandlerList 储存handler的名目
var HandlerList = []handler{manager, editor}
