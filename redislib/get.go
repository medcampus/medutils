package redislib

import (
	"github.com/gomodule/redigo/redis"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (c *Client) GetByte(key string) ([]byte, error) {
	reply, err := c.Conn.Do("GET", key)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "ERROR: fail get key %s, error %s", key, err.Error())
	}

	if reply == nil {
		return nil, status.Errorf(codes.InvalidArgument, "Error key not found, key %v", key)
	}

	return redis.Bytes(reply, nil)
}

func (c *Client) GetString(key string) (string, error) {
	reply, err := c.Conn.Do("GET", key)
	if err != nil {
		return "", status.Errorf(codes.Internal, "ERROR: fail get key %s, error %s", key, err.Error())
	}

	if reply == nil {
		return "", status.Errorf(codes.InvalidArgument, "Error key not found, key %v", key)
	}

	return redis.String(reply, nil)
}

func (c *Client) GetInt(key string) (int, error) {
	reply, err := c.Conn.Do("GET", key)
	if err != nil {
		return 0, status.Errorf(codes.Internal, "ERROR: fail get key %s, error %s", key, err.Error())
	}

	if reply == nil {
		return 0, status.Errorf(codes.InvalidArgument, "Error key not found, key %v", key)
	}

	return redis.Int(reply, nil)
}

func (c *Client) GetInt64(key string) (int64, error) {
	reply, err := c.Conn.Do("GET", key)
	if err != nil {
		return 0, status.Errorf(codes.Internal, "ERROR: fail get key %s, error %s", key, err.Error())
	}

	if reply == nil {
		return 0, status.Errorf(codes.InvalidArgument, "Error key not found, key %v", key)
	}

	return redis.Int64(reply, nil)
}
