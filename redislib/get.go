package redislib

import (
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

	return reply.([]byte), nil
}

func (c *Client) GetString(key string) (string, error) {
	reply, err := c.Conn.Do("GET", key)
	if err != nil {
		return "", status.Errorf(codes.Internal, "ERROR: fail get key %s, error %s", key, err.Error())
	}

	if reply == nil {
		return "", status.Errorf(codes.InvalidArgument, "Error key not found, key %v", key)
	}

	return reply.(string), nil
}

func (c *Client) GetInt32(key string) (int32, error) {
	reply, err := c.Conn.Do("GET", key)
	if err != nil {
		return 0, status.Errorf(codes.Internal, "ERROR: fail get key %s, error %s", key, err.Error())
	}

	if reply == nil {
		return 0, status.Errorf(codes.InvalidArgument, "Error key not found, key %v", key)
	}

	return reply.(int32), nil
}

func (c *Client) GetInt64(key string) (int64, error) {
	reply, err := c.Conn.Do("GET", key)
	if err != nil {
		return 0, status.Errorf(codes.Internal, "ERROR: fail get key %s, error %s", key, err.Error())
	}

	if reply == nil {
		return 0, status.Errorf(codes.InvalidArgument, "Error key not found, key %v", key)
	}

	return reply.(int64), nil
}
