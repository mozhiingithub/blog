package handler

import (
	"net/http"
	"strconv"

	d "../database"
	"github.com/gin-gonic/gin"
)

var deleteArticle = handler{
	path:   "/article",
	method: "DELETE",
	handle: func(c *gin.Context) {
		result := "删除失败"
		for {
			articleIDStr := c.Query("id")
			articleID, _ := strconv.Atoi(articleIDStr)
			db := d.GetInstance() // 获取数据库连接池实例
			tx, err := db.Begin() // 开启事务
			if nil != err {       //无法开启事务
				result = "无法开启事务"
				break
			}
			stmt, err := tx.Prepare("delete from ts where id = ?") // 准备删除时间
			if nil != err {                                        // 准备失败
				result = "时间删除准备失败"
				tx.Rollback()
				break
			}
			_, err = stmt.Exec(articleID) // 删除时间
			if nil != err {               // 删除失败
				result = "时间删除失败"
				tx.Rollback()
				break
			}
			stmt, err = tx.Prepare("delete from contents where id = ?") // 准备删除正文
			if nil != err {                                             // 准备失败
				result = "正文删除准备失败"
				tx.Rollback()
				break
			}
			_, err = stmt.Exec(articleID) // 删除正文
			if nil != err {               // 删除失败
				result = "正文删除失败"
				tx.Rollback()
				break
			}
			stmt, err = tx.Prepare("delete from titles where id = ?") // 准备删除标题
			if nil != err {                                           // 准备失败
				result = "标题删除准备失败"
				tx.Rollback()
				break
			}
			_, err = stmt.Exec(articleID) // 删除标题
			if nil != err {               // 删除失败
				result = "标题删除失败"
				tx.Rollback()
				break
			}
			err = tx.Commit() // 提交事务
			if nil != err {   // 提交失败
				result = "事务提交失败"
				tx.Rollback()
				break
			}
			result = "删除成功"
			break
		}
		c.HTML(http.StatusOK, "result.html", gin.H{
			"result": result,
		})
	},
}
