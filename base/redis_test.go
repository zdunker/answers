package base

import (
	"os"
	"testing"

	"github.com/go-redis/redis/v8"
	"github.com/stretchr/testify/assert"
)

var (
	rdb *redis.Client
)

func TestRedisHandler(t *testing.T) {
	client := redisClient{
		rdb,
	}
	value, _ := client.Get("test")
	assert.Equal(t, "", value)

	err := client.Set("test", "nonono")
	assert.Nil(t, err)

	value, err = client.Get("test")
	assert.Equal(t, "nonono", value)
	assert.Nil(t, err)

	err = client.Delete("test")
	assert.Nil(t, err)

	value, err = client.Get("test")
	assert.Equal(t, "", value)
	assert.Equal(t, redis.Nil, err)
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	os.Exit(code)
}

func setup() {
	rdb = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		DB:   0,
	})
}
