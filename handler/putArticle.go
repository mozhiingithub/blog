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

var putArticle = handler{
	path:   "/article",
	method: "PUT",
	handle: func(c *gin.Context) {
		result := "修改失败"
		for {
			articleIDStr := c.Query("id")
			articleID, _ := strconv.Atoi(articleIDStr)
			title := c.PostForm("title")
			content := c.PostForm("content")
			if "" == title || "" == content { // 标题或正文部分为空
				result = "标题或正文部分为空"
				break
			}
			var db *sql.DB = d.GetInstance() // 获取数据库连接池实例
			tx, err := db.Begin()            // 开启事务
			if nil != err {                  //无法开启事务
				result = "无法开启事务"
				break
			}
			stmt, err := tx.Prepare("update titles set title = ? where id = ?") // 准备修改标题
			if nil != err {                                                     // 准备失败
				result = "标题修改准备失败"
				tx.Rollback()
				break
			}
			_, err = stmt.Exec(title, articleID) // 修改标题
			if nil != err {                      // 修改失败
				result = "标题修改失败"
				tx.Rollback()
				break
			}
			stmt, err = tx.Prepare("update contents set content = ? where id = ?") // 准备修改正文
			if nil != err {                                                        // 准备失败
				result = "正文修改准备失败"
				tx.Rollback()
				break
			}
			_, err = stmt.Exec(content, articleID) // 修改正文
			if nil != err {                        // 修改失败
				result = "正文修改失败"
				tx.Rollback()
				break
			}
			stmt, err = tx.Prepare("update ts set t = now() where id = ?") // 准备更新发布时间
			if nil != err {                                                // 准备失败
				result = "发布时间修改准备失败"
				tx.Rollback()
				break
			}
			_, err = stmt.Exec(articleID) // 修改发布时间
			if nil != err {               // 修改插入失败
				result = "发布时间修改失败"
				tx.Rollback()
				break
			}
			// 修改完数据库内容，提交事务前，先行删除redis中的缓存
			var rs redis.Conn = ch.GetInstance() // 获取redis实例
			err = rs.Send("del", articleID)      // 删除缓存
			if nil != err {                      // 缓存删除失败
				result = "无法删除缓存"
				tx.Rollback()
				break
			}
			err = tx.Commit() // 提交事务
			if nil != err {   // 提交失败
				result = "事务提交失败"
				tx.Rollback()
				break
			}
			result = "修改成功"
			break
		}
		c.HTML(http.StatusOK, "result.html", gin.H{
			"result": result,
		})
	},
}
