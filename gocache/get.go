package gocache

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (c *Client) Get(key string) ([]byte, error) {
	reply, err := c.Conn.Do("GET", key)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "ERROR: fail get key %s, error %s", key, err.Error())
	}

	if reply == nil {
		return nil, status.Errorf(codes.InvalidArgument, "Error key not found, key %v", key)
	}

	return reply.([]byte), nil
}
