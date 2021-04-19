package base

import "github.com/go-redis/redis/v8"

type redisClient struct {
	client *redis.Client
}

type redisHandler interface {
	Get(key string) (string, error)
}

func (rClient redisClient) Get(key string) (string, error) {
	return rClient.client.Get(ctx, key).Result()
}

func (rClient redisClient) Set(key, value string) error {
	return rClient.client.Set(ctx, key, value, 0).Err()
}

func (rClient redisClient) Delete(key ...string) error {
	return rClient.client.Del(ctx, key...).Err()
}
