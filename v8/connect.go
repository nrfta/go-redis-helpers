package redis_helpers

import (
	"context"
	"crypto/tls"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

const DefaultPort = 6379

func ConnectRedis(c RedisConfig) (*redis.Client, error) {
	port := DefaultPort
	if c.Port > 0 {
		port = c.Port
	}
	o := &redis.Options{
		Addr: fmt.Sprintf("%s:%d", c.Host, port),
		DB:   c.Database,
	}
	if c.Password != "" {
		o.Password = c.Password
	}
	if c.SSLEnabled {
		o.TLSConfig = &tls.Config{
			ServerName: c.Host,
		}
	}

	rdb := redis.NewClient(o)

	// try pinging for 5 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	for {
		_, err := rdb.Ping(ctx).Result()
		if err == nil {
			return rdb, nil
		}

		select {
		case <-ctx.Done():
			return nil, fmt.Errorf("failed to connect to redis; no ping response")
		default:
			time.Sleep(100 * time.Millisecond)
			continue
		}
	}
}
