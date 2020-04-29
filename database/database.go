package database

import (
	"database/sql"
	"fmt"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var lock *sync.Mutex = &sync.Mutex{}

const (
	host         = "localhost"
	user         = "root"
	psw          = ""
	network      = "tcp"
	name         = "blog"
	port         = 3306
	maxLifeTime  = 10
	maxOpenConns = 20
	maxIdleConns = 10
)

// GetInstance 获取数据库连接池实例
func GetInstance() *sql.DB {
	if nil == db {
		lock.Lock()
		defer lock.Unlock()
		if nil == db {
			info := fmt.Sprintf("%s:%s@%s(%s:%d)/%s", user, psw, network, host, port, name)
			db, _ = sql.Open("mysql", info)
			db.SetConnMaxLifetime(maxLifeTime * time.Second)
			db.SetMaxOpenConns(maxOpenConns)
			db.SetMaxIdleConns(maxIdleConns)
		}
	}
	return db
}

// CloseInstance 关闭数据库连接池实例
func CloseInstance() {
	if nil != db {
		db.Close()
	}
}

// GetTitle 检查该博文ID是否真实存在。若存在，返回博文对应标题
func GetTitle(id int) (title string, ok bool) {
	rows, _ := GetInstance().Query("select title from titles where id = ?", id) // 查询此博文是否存在
	ok = rows.Next()
	if ok { // 查询结果非空，表明此博文存在
		rows.Scan(&title) // 获取博文标题
	}
	return
}

// GetContent 根据ID，获取博文正文
func GetContent(id int) (content string, ok bool) {
	rows, _ := GetInstance().Query("select content from contents where id = ?", id) // 查询博文正文
	ok = rows.Next()
	if ok {
		rows.Scan(&content) // 获取博文正文
	}
	return
}

// GetTime 根据ID，获取博文发布时间
func GetTime(id int) (t string, ok bool) {
	rows, _ := GetInstance().Query("select t from ts where id = ?", id) // 查询博文发布时间
	ok = rows.Next()
	if ok {
		rows.Scan(&t) // 获取博文发布时间
	}
	return
}
