package redislib

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (c *Client) SetByteWithExp(key string, value []byte, durationInSeconds int64) error {
	_, err := c.Conn.Do("SETEX", key, durationInSeconds, string(value))
	if err != nil {
		return status.Errorf(codes.Internal, "ERROR: fail set key %s, val %s, error %s", key, value, err.Error())
	}

	return nil
}

func (c *Client) SetStringWithExp(key string, value string, durationInSeconds int64) error {
	_, err := c.Conn.Do("SETEX", key, durationInSeconds, value)
	if err != nil {
		return status.Errorf(codes.Internal, "ERROR: fail set key %s, val %s, error %s", key, value, err.Error())
	}

	return nil
}

func (c *Client) SetIntWithExp(key string, value int, durationInSeconds int64) error {
	_, err := c.Conn.Do("SETEX", key, durationInSeconds, value)
	if err != nil {
		return status.Errorf(codes.Internal, "ERROR: fail set key %s, val %s, error %s", key, value, err.Error())
	}

	return nil
}

func (c *Client) SetInt64WithExp(key string, value int64, durationInSeconds int64) error {
	_, err := c.Conn.Do("SETEX", key, durationInSeconds, value)
	if err != nil {
		return status.Errorf(codes.Internal, "ERROR: fail set key %s, val %s, error %s", key, value, err.Error())
	}

	return nil
}

func (c *Client) SetByte(key string, value []byte) error {
	_, err := c.Conn.Do("SET", key, string(value))
	if err != nil {
		return status.Errorf(codes.Internal, "ERROR: fail set key %s, val %s, error %s", key, value, err.Error())
	}

	return nil
}

func (c *Client) SetString(key string, value string) error {
	_, err := c.Conn.Do("SET", key, value)
	if err != nil {
		return status.Errorf(codes.Internal, "ERROR: fail set key %s, val %s, error %s", key, value, err.Error())
	}

	return nil
}

func (c *Client) SetInt(key string, value int) error {
	_, err := c.Conn.Do("SET", key, value)
	if err != nil {
		return status.Errorf(codes.Internal, "ERROR: fail set key %s, val %s, error %s", key, value, err.Error())
	}

	return nil
}

func (c *Client) SetInt64(key string, value int64) error {
	_, err := c.Conn.Do("SET", key, value)
	if err != nil {
		return status.Errorf(codes.Internal, "ERROR: fail set key %s, val %s, error %s", key, value, err.Error())
	}

	return nil
}
