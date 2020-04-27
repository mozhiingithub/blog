package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type handler interface {
	GetAddress() string
	GetMethod() string
	Handle(c *gin.Context)
}

type simplehandler struct{}

func (s simplehandler) GetAddress() string {
	return "/"
}

func (s simplehandler) GetMethod() string {
	return "GET"
}

func (s simplehandler) Handle(c *gin.Context) {
	c.String(200, "123")
}

type imghandler struct{}

func (s imghandler) GetAddress() string {
	return "/imgswskjw"
}

func (s imghandler) GetMethod() string {
	return "POST"
}

func (s imghandler) Handle(c *gin.Context) {
	c.String(200, "456")
}

var handleList = []handler{simplehandler{}, imghandler{}}

func HandleList(engine *gin.Engine) {
	for _, h := range handleList {
		fmt.Println(h.GetMethod(), h.GetAddress())
	}
}
