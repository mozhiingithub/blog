package cache

import (
	"fmt"
	"sync"

	"github.com/gomodule/redigo/redis"
)

var cache *redis.Conn
var lock *sync.Mutex = &sync.Mutex{}

const (
	host    = "localhost"
	network = "tcp"
	port    = 6379
)

// GetInstance 获取redis连接池实例
func GetInstance() *redis.Conn {
	if nil == cache {
		lock.Lock()
		defer lock.Unlock()
		if nil == db {
			info := fmt.Sprintf("%s:%d", host, port)
			cache, _ = redis.Dial(network, info)
		}
	}
	return cache
}

// CloseInstance 关闭redis连接池实例
func CloseInstance() {
	if nil != cache {
		cache.Close()
	}
}
