package database

import (
	"github.com/gomodule/redigo/redis"
	"github.com/soveran/redisurl"
)

var redisPool *redis.Pool

func SetupRedisDB() {
	pool := &redis.Pool{
		MaxActive: 5,
		MaxIdle:   5,
		Wait:      true,
		Dial:      GetRedisConnection,
	}

	redisPool = pool
}

func GetRedisConnection() (redis.Conn, error) {
	return redisurl.Connect()
}

func GetRedisPool() *redis.Pool {
	return redisPool
}
