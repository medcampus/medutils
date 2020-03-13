package redislib

import "github.com/gomodule/redigo/redis"

func (c *Client) IsExists(key string) (bool, error) {
	return redis.Bool(c.Conn.Do("EXISTS", key))
}
