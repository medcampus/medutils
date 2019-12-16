package redislib

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (c *Client) SetByte(key string, value []byte, durationInSeconds int64) error {
	_, err := c.Conn.Do("SETEX", key, durationInSeconds, string(value))
	if err != nil {
		return status.Errorf(codes.Internal, "ERROR: fail set key %s, val %s, error %s", key, value, err.Error())
	}

	return nil
}

func (c *Client) SetString(key string, value string, durationInSeconds int64) error {
	_, err := c.Conn.Do("SETEX", key, durationInSeconds, value)
	if err != nil {
		return status.Errorf(codes.Internal, "ERROR: fail set key %s, val %s, error %s", key, value, err.Error())
	}

	return nil
}

func (c *Client) SetInt32(key string, value int32, durationInSeconds int64) error {
	_, err := c.Conn.Do("SETEX", key, durationInSeconds, value)
	if err != nil {
		return status.Errorf(codes.Internal, "ERROR: fail set key %s, val %s, error %s", key, value, err.Error())
	}

	return nil
}

func (c *Client) SetInt64(key string, value int64, durationInSeconds int64) error {
	_, err := c.Conn.Do("SETEX", key, durationInSeconds, value)
	if err != nil {
		return status.Errorf(codes.Internal, "ERROR: fail set key %s, val %s, error %s", key, value, err.Error())
	}

	return nil
}
