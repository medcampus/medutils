package medutils

import (
	"errors"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"sync"
)

var pool sync.Map

type (
	GrpcServiceConfig struct {
		Key    string
		Target string //e.g: infomed:8081
		Opts   []grpc.DialOption
	}

	// grpcPool is that gRPC connection pool by buffered channel
	grpcPool struct {
		conns  chan *grpc.ClientConn
		cg     connGenerator
		target string
		opts   []grpc.DialOption
	}

	// ConnGenerator is function type to generate a grpc connection function
	connGenerator func(target string, opts ...grpc.DialOption) (conn *grpc.ClientConn, err error)
)

// Create a bundle of connection pool instance
func Create(cg connGenerator, initialConn, maxConn int, GrpcServiceConfigs ...GrpcServiceConfig) error {
	if cg == nil || initialConn <= 0 || maxConn <= 0 || initialConn > maxConn || len(GrpcServiceConfigs) < 0 {
		return errors.New("invalid arguments")
	}

	for _, GrpcServiceConfig := range GrpcServiceConfigs {
		if _, ok := pool.Load(GrpcServiceConfig.Key); !ok {
			cp := &grpcPool{
				conns:  make(chan *grpc.ClientConn, maxConn),
				cg:     cg,
				target: GrpcServiceConfig.Target,
				opts:   GrpcServiceConfig.Opts,
			}
			for i := 0; i < initialConn; i++ {
				c, err := cg(GrpcServiceConfig.Target, GrpcServiceConfig.Opts...)
				if err != nil {
					return err
				}
				cp.conns <- c
			}
			pool.Store(GrpcServiceConfig.Key, cp)
		}
	}
	return nil
}

// Get is that try to get a grpc connection from grpc pool by specified key
func Get(key string) (*grpc.ClientConn, error) {
	if val, ok := pool.Load(key); ok {
		return val.(*grpcPool).get()
	}
	return nil, errors.New("invalid key: " + key)
}

// PutBack is that give back a specific gRPC connection to gRPC pool
func PutBack(key string, conn *grpc.ClientConn) error {
	if val, ok := pool.Load(key); ok {
		return val.(*grpcPool).putBack(conn)
	}
	return errors.New("invalid key: " + key)
}

// Close the gRPC pool totally
func Close() {
	pool.Range(func(key, value interface{}) bool {
		grpcPool, ok := value.(*grpcPool)
		if !ok {
			return ok
		}
		grpcPool.close()
		for conn := range grpcPool.conns {
			if err := conn.Close(); err != nil {
				return false
			}
		}
		return true
	})
}

func (c *grpcPool) get() (*grpc.ClientConn, error) {
	select {
	case conn := <-c.conns:
		if conn == nil {
			return nil, errors.New("connection is closed")
		}
		return conn, nil
	default:
		// channel is empty
		conn, err := c.cg(c.target, c.opts...)
		if err != nil {
			return nil, err
		}
		return conn, nil
	}
}

func (c *grpcPool) putBack(conn *grpc.ClientConn) (err error) {
	if conn == nil {
		return errors.New("conn is nil")
	}

	switch conn.GetState().String() {
	case "READY":
		c.conns <- conn
		return nil
	default:
		err = conn.Close()
		if err != nil {
			return err
		}

		conn, err = c.cg(c.target, c.opts...)
		if err != nil {
			return err
		}
		c.conns <- conn
		return nil
	}
}

func (c *grpcPool) close() {
	close(c.conns)
}

func DeferGrpcConnection(key string, conn *grpc.ClientConn) {
	if err := PutBack(key, conn); err != nil {
		logrus.Errorf("error PutBack key %s, clientCon %v, clientCon state %s, error %v", key, conn.Target(), conn.GetState().String(), err)
	}
}
