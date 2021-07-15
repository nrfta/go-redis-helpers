package redis_helpers_test

import (
	"os"
	"strconv"

	"github.com/nrfta/go-redis-helpers/v7"
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

	Context("Redis v7", func() {
		It("should connect to a database", func() {
			testConfig := redis_helpers.RedisConfig{
				Host:     getEnv("REDIS_HOST", "localhost"),
				Port:     port,
				Password: getEnv("REDIS_PASSWORD", "password"),
				Database: db,
			}
			rdb, err := redis_helpers.ConnectRedis(testConfig)
			Expect(err).To(BeNil())
			Expect(rdb).To(Not(BeNil()))
		})
	})
})
