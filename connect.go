package redis_helpers

import (
	"context"
	"crypto/tls"
	"fmt"
	"strings"
	"time"

	"github.com/go-redis/redis/v8"
)

const DefaultPort = 6379

func ConnectRedis(c RedisConfig) (redis.UniversalClient, error) {
	if c.ClusterEnabled {
		return connectRedisCluster(c)
	}

	port := DefaultPort
	if c.Port > 0 {
		port = c.Port
	}

	o := &redis.Options{
		Addr:     fmt.Sprintf("%s:%d", c.Host, port),
		DB:       c.Database,
		Password: c.Password,
	}

	if c.SSLEnabled {
		o.TLSConfig = &tls.Config{
			ServerName: c.Host,
		}
	}

	rdb := redis.NewClient(o)

	if err := testConnection(rdb); err != nil {
		return nil, err
	}
	return rdb, nil
}

func connectRedisCluster(c RedisConfig) (*redis.ClusterClient, error) {
	if !c.ClusterEnabled {
		return nil, fmt.Errorf("config does not specify cluster enabled")
	}

	port := DefaultPort
	if c.Port > 0 {
		port = c.Port
	}

	hosts := strings.Split(c.Host, ",")
	for i, h := range hosts {
		hosts[i] = fmt.Sprintf("%s:%d", h, port)
	}

	o := &redis.ClusterOptions{
		Addrs:      hosts,
		Password:   c.Password,
		MaxRetries: 10,
	}

	if c.SSLEnabled {
		o.TLSConfig = &tls.Config{
			ServerName: c.Host,
		}
	}

	rdb := redis.NewClusterClient(o)

	if err := testConnection(rdb); err != nil {
		return nil, err
	}
	return rdb, nil
}

func testConnection(rdb redis.UniversalClient) error {
	// try pinging for 5 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	for {
		_, err := rdb.Ping(ctx).Result()
		if err == nil {
			return nil
		}

		select {
		case <-ctx.Done():
			return fmt.Errorf("failed to connect to redis; no ping response")
		default:
			time.Sleep(100 * time.Millisecond)
			continue
		}
	}
}
