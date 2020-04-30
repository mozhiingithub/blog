package handler

import (
	"net/http"
	"strconv"

	ch "../cache"
	d "../database"
	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
)

const expireSecond = 60

var getArticle = handler{
	path:   "/article",
	method: "GET",
	handle: func(c *gin.Context) {
		result := "博文不存在"
		for {
			articleIDStr := c.Query("id") // 获取博文id参数
			if "" == articleIDStr {
				break
			}
			articleID, err := strconv.Atoi(articleIDStr)
			if nil != err {
				break
			}

			for { // 先判断redis中是否有此博文的缓存
				var rs redis.Conn = ch.GetInstance() // 获取redis实例
				ok, _ := redis.Bool(rs.Do("exists", articleID))
				if !ok { // redis中没有该博文缓存
					break
				}
				m, err := redis.StringMap(rs.Do("hgetall", articleID)) // 获取博文的标题、时间、正文
				if 0 == len(m) {                                       // 获取失败
					break
				}
				// 获取博文阅读量并加一。
				// 假设count中没有该id的阅读量值，则新建一个并加一
				count, err := redis.Int(rs.Do("hincrby", "count", articleID, 1))
				if nil != err { // 获取失败
					break
				}
				m["count"] = strconv.Itoa(count)         // 将阅读量添加到m中
				c.HTML(http.StatusOK, "article.html", m) // 将缓存传到博文展示页前端
				return
			}

			// 因各种原因，无法获取博文缓存。现从数据库中获取
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

			// 从数据库中获取博文内容，将其写入缓存
			var rs redis.Conn = ch.GetInstance()                                          // 获取redis实例
			err = rs.Send("hmset", articleID, "title", title, "t", t, "content", content) // 以id为key，写入一个hash
			if nil != err {                                                               // 写入失败
				result = "博文写入缓存失败"
				break
			}
			count, err := redis.Int(rs.Do("hincrby", "count", articleID, 1)) // 获取阅读量，加一
			if nil != err {                                                  // 获取失败
				result = "获取阅读量失败"
				break
			}
			rs.Send("expire", articleID, expireSecond) // 为该博文的缓存设置过期时间

			// 将内容传到博文展示页前端
			c.HTML(http.StatusOK, "article.html", gin.H{
				"title":   title,
				"t":       t,
				"content": content,
				"count":   strconv.Itoa(count),
			})
			return
		}
		c.HTML(http.StatusNotFound, "result.html", gin.H{
			"result": result,
		})
		return
	},
}
