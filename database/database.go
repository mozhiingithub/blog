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
