package redis_helpers_test

import (
	rh7 "github.com/nrfta/go-redis-helpers/v7"
	rh8 "github.com/nrfta/go-redis-helpers/v8"
	"os"
	"strconv"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
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

	Context("Redis v8", func() {
		It("should connect to a database", func() {
			testConfig := rh8.RedisConfig{
				Host:     getEnv("REDIS_HOST", "localhost"),
				Port:     port,
				Password: getEnv("REDIS_PASSWORD", "password"),
				Database: db,
			}

			rdb, err := rh8.ConnectRedis(testConfig)
			Expect(err).To(BeNil())
			Expect(rdb).To(Not(BeNil()))
		})
	})

	Context("Redis v7", func() {
		It("should connect to a database", func() {
			testConfig := rh7.RedisConfig{
				Host:     getEnv("REDIS_HOST", "localhost"),
				Port:     port,
				Password: getEnv("REDIS_PASSWORD", "password"),
				Database: db,
			}
			rdb, err := rh7.ConnectRedis(testConfig)
			Expect(err).To(BeNil())
			Expect(rdb).To(Not(BeNil()))
		})
	})
})
