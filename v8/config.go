package redis_helpers

import "fmt"

type RedisConfig struct {
	Host     string
	Port     int
	Database int
	Password string
	SSLEnabled         bool `mapstructure:"ssl_enabled"`
	ClusterEnabled bool `mapstructure:"cluster_enabled"`
}

func (c RedisConfig) URL() string {
	return fmt.Sprintf("redis://:%s@%s:%d/%d", c.Password, c.Host, c.Port, c.Database)
}
