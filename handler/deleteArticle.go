package handler

import (
	"database/sql"
	"net/http"
	"strconv"

	ch "../cache"
	d "../database"
	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
)

var deleteArticle = handler{
	path:   "/article",
	method: "DELETE",
	handle: func(c *gin.Context) {
		result := "删除失败"
		for {
			articleIDStr := c.Query("id")
			articleID, _ := strconv.Atoi(articleIDStr)
			var db *sql.DB = d.GetInstance() // 获取数据库连接池实例
			tx, err := db.Begin()            // 开启事务
			if nil != err {                  //无法开启事务
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
			// 删除完数据库相关内容，提交事务前，先行删除redis中的缓存
			var rs redis.Conn = ch.GetInstance() // 获取redis实例
			_, err = rs.Do("del", articleID)     // 删除缓存
			if nil != err {                      // 缓存删除失败
				result = "无法删除缓存"
				tx.Rollback()
				break
			}
			_, err = rs.Do("hdel", "count", articleID) // 删除博文阅读量数据
			if nil != err {                            // 阅读量删除失败
				result = "无法删除阅读量"
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
