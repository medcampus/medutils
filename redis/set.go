package redis

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

func (c *Client) SetByte(key string, value []byte, duration time.Duration) error {
	_, err := c.Conn.Do("SETEX", key, duration, string(value))
	if err != nil {
		return status.Errorf(codes.Internal, "ERROR: fail set key %s, val %s, error %s", key, value, err.Error())
	}

	return nil
}

func (c *Client) SetString(key string, value string, duration time.Duration) error {
	_, err := c.Conn.Do("SETEX", key, duration, value)
	if err != nil {
		return status.Errorf(codes.Internal, "ERROR: fail set key %s, val %s, error %s", key, value, err.Error())
	}

	return nil
}

func (c *Client) SetInt32(key string, value int32, duration time.Duration) error {
	_, err := c.Conn.Do("SETEX", key, duration, value)
	if err != nil {
		return status.Errorf(codes.Internal, "ERROR: fail set key %s, val %s, error %s", key, value, err.Error())
	}

	return nil
}

func (c *Client) SetInt64(key string, value int64, duration time.Duration) error {
	_, err := c.Conn.Do("SETEX", key, duration, value)
	if err != nil {
		return status.Errorf(codes.Internal, "ERROR: fail set key %s, val %s, error %s", key, value, err.Error())
	}

	return nil
}
