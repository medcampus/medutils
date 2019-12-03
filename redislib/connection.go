package redislib

import (
	"github.com/gomodule/redigo/redis"
	"github.com/spf13/viper"
	"log"
	"os"
	"sync"
)

var (
	pool *redis.Pool
	once sync.Once
)

type Client struct {
	Conn redis.Conn
}

func initPool() *redis.Pool {
	return &redis.Pool{
		MaxIdle:   80,
		MaxActive: 12000,
		Dial: func() (redis.Conn, error) {
			conn, err := redis.Dial("tcp", viper.GetString("redis.host"))
			if err != nil {
				log.Printf("ERROR: fail init redis pool: %s", err.Error())
				os.Exit(1)
			}
			return conn, err
		},
	}
}

func GetRedisClient() *Client {
	once.Do(func() {
		pool = initPool()
	})

	return &Client{Conn: pool.Get()}
}

func (c *Client) Close() {
	defer c.Conn.Close()
}

func Ping(conn redis.Conn) {
	_, err := redis.String(conn.Do("PING"))
	if err != nil {
		log.Printf("ERROR: fail ping redis conn: %s", err.Error())
		os.Exit(1)
	}
}
