package redis_helpers_test

import (
	"os"
	"strconv"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	rh "github.com/nrfta/go-redis-helpers"
)

func getEnv(key string, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}

var _ = Describe("Connection Test", func() {
	var port, _ = strconv.Atoi(getEnv("REDIS_PORT", "6379"))
	var db, _ = strconv.Atoi(getEnv("REDIS_DATABASE", "0"))

	var (
		testConfig = rh.RedisConfig{
			Host:     getEnv("REDIS_HOST", "localhost"),
			Port:     port,
			Password: getEnv("REDIS_PASSWORD", "password"),
			Database: db,
		}
	)
	Context("Redis v8", func() {
		It("should connect to a database", func() {
			rdb, err := rh.ConnectRedisV8(testConfig)
			Expect(err).To(BeNil())
			Expect(rdb).To(Not(BeNil()))
		})
	})

	Context("Redis v7", func() {
		It("should connect to a database", func() {
			rdb, err := rh.ConnectRedisV7(testConfig)
			Expect(err).To(BeNil())
			Expect(rdb).To(Not(BeNil()))
		})
	})
})
