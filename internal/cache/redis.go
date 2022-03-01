package cache

import (
	"altair-backend/config"
	"github.com/garyburd/redigo/redis"
	"log"
	"time"
)

var (
	pool      *redis.Pool
	redisHost = config.ServerConfig.RedisHost
	password  = config.ServerConfig.RedisPassword
)

//newRedisPool:创建redis连接池
func newRedisPool() *redis.Pool {
	return &redis.Pool{
		MaxIdle:     30,
		MaxActive:   60,
		IdleTimeout: 120 * time.Second,
		Wait:        true,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", redisHost, redis.DialPassword(password))
			if err != nil {
				log.Println("Redis 连接失败", err.Error())
				return nil, err
			}
			return c, nil
		},
		//定时检查redis是否出状况
		TestOnBorrow: func(conn redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			res, err := conn.Do("PING")
			log.Println("Redis 健康检查", res)
			return err
		},
	}
}

//初始化redis连接池
func init() {
	pool = newRedisPool()
}

//对外暴露连接池
func RedisPool() *redis.Pool {
	return pool
}
