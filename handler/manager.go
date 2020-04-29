package handler

import (
	"fmt"
	"net/http"

	d "../database"
	"github.com/gin-gonic/gin"
)

var manager = handler{
	path:   "/manager",
	method: "GET",
	handle: func(c *gin.Context) {
		blogList := make([]blog, 0)
		db := d.GetInstance() // 获取数据库连接池实例
		rows, _ := db.Query("select * from titles order by id desc")
		for rows.Next() {
			var rowID int
			var rowTitle string
			rows.Scan(&rowID, &rowTitle)
			blogList = append(blogList, blog{
				Id:    fmt.Sprintf("/article?id=%d", rowID),
				Edit:  rowID,
				Title: rowTitle,
			})
		}
		c.HTML(http.StatusOK, "manager.html", gin.H{
			"blogList": blogList,
		})
	},
}

type blog struct {
	Id    string
	Edit  int
	Title string
}
