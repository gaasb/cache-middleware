package cache_middleware

import (
	"context"
	"github.com/redis/go-redis/v9"
)

type TestRedis redis.Client

func (r TestRedis) Get(key string) (*Item, error) {
	//TODO implement me
	panic("implement me")
}

func (r TestRedis) Set(key string, data any) error {
	//TODO implement me
	panic("implement me")
}

func (r TestRedis) Remove(key string) error {
	_, err := r.Get(key)
	return err
}

func (r TestRedis) Flush() {
	r.Flush()
	//TODO implement me
	panic("implement me")
}

func NewRedis() {
	rdb := redis.NewClient(
		&redis.Options{
			Network:               "",
			Addr:                  "",
			ClientName:            "",
			Dialer:                nil,
			OnConnect:             nil,
			Username:              "",
			Password:              "",
			CredentialsProvider:   nil,
			DB:                    0,
			MaxRetries:            0,
			MinRetryBackoff:       0,
			MaxRetryBackoff:       0,
			DialTimeout:           0,
			ReadTimeout:           0,
			WriteTimeout:          0,
			ContextTimeoutEnabled: false,
			PoolFIFO:              false,
			PoolSize:              0,
			PoolTimeout:           0,
			MinIdleConns:          0,
			MaxIdleConns:          0,
			ConnMaxIdleTime:       0,
			ConnMaxLifetime:       0,
			TLSConfig:             nil,
			Limiter:               nil,
		},
	)
	rdb.Ping(context.Background())
}
