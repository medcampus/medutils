package redis

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (c *Client) Delete(pattern string, wildcard bool) error {
	_, err := c.Conn.Do("DEL", pattern)
	if err != nil {
		return status.Errorf(codes.Internal, "ERROR: fail del key %s, error %s", pattern, err.Error())
	}

	return nil
}

// TODO implement wild card functionality
