package handler

import (
	"net/http"

	d "../database"
	"github.com/gin-gonic/gin"
)

var postArticle = handler{
	path:   "/article",
	method: "POST",
	handle: func(c *gin.Context) {
		result := "发布失败"
		for {
			title := c.PostForm("title")
			content := c.PostForm("content")
			if "" == title || "" == content { // 标题或正文部分为空
				break
			}
			db := d.GetInstance() // 获取数据库连接池实例
			tx, err := db.Begin() // 开启事务
			if nil != err {       //无法开启事务
				break
			}
			stmt, err := tx.Prepare("insert into titles(title) values(?)") // 准备将标题插入数据库
			if nil != err {                                                // 准备失败
				tx.Rollback()
				break
			}
			res, err := stmt.Exec(title) // 插入标题
			if nil != err {              // 插入失败
				tx.Rollback()
				break
			}
			id, err := res.LastInsertId() // 获取自增id
			if nil != err {               // 获取失败
				tx.Rollback()
				break
			}
			stmt, err = tx.Prepare("insert into contents(id,content) values(?,?)") // 准备将正文插入数据库
			if nil != err {                                                        // 准备失败
				tx.Rollback()
				break
			}
			res, err = stmt.Exec(id, content) // 插入正文
			if nil != err {                   // 插入失败
				tx.Rollback()
				break
			}
			stmt, err = tx.Prepare("insert into ts(id) values(?)") // 准备将发布时间插入数据库
			if nil != err {                                        // 准备失败
				tx.Rollback()
				break
			}
			res, err = stmt.Exec(id) // 插入发布时间
			if nil != err {          // 插入失败
				tx.Rollback()
				break
			}
			err = tx.Commit() // 提交事务
			if nil != err {   // 提交失败
				tx.Rollback()
				break
			}
			result = "发布成功"
			break
		}
		c.HTML(http.StatusOK, "result.html", gin.H{ // 跳转到处理结果界面
			"result": result,
		})
	},
}
