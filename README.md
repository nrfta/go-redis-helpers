# go-redis-helpers

![CI](https://github.com/nrfta/go-redis-helpers/workflows/CI/badge.svg?branch=master)

Provides common config and connect for Redis database.

## Config

```go
type RedisConfig struct {
	Host     string
	Port     int
	Database int
	Password string
}
```

Default `Port` is the standard Redis port: 6379

`Database` in Redis is an integer.

`Password` is optional if using Redis in passwordless mode.

## Connect

`Connect` will attempt to ping the Redis database for up to 5 seconds. If there is
no response after 5 seconds, then the function returns an error.
