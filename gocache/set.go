package gocache

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

func (c *Client) Set(key string, value []byte, duration time.Duration) error {
	_, err := c.Conn.Do("SET", key, string(value))
	if err != nil {
		return status.Errorf(codes.Internal, "ERROR: fail set key %s, val %s, error %s", key, value, err.Error())
	}

	return nil
}
