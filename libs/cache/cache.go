package cache

import (
	"context"
	"dreon_ecommerce_server/shared/interfaces"
	"fmt"
	"time"

	"dreon_ecommerce_server/configs"

	"github.com/redis/go-redis/v9"
)

type appCache struct {
	logger      interfaces.ILogger
	redisClient *redis.Client
	interfaces.ICache
}

func NewAppCache(c *configs.AppConfig, logger interfaces.ILogger) *appCache {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     c.Cache.RedisHost + ":" + c.Cache.RedisPort,
		Password: c.Cache.RedisPassword,
		DB:       c.Cache.RedisDB,
	})

	logger.Info("Connected to cache server")
	return &appCache{
		logger:      logger,
		redisClient: redisClient,
	}
}

func (c *appCache) Set(key string, value interface{}, expireTime *time.Duration) error {
	action := "appCache.Set"
	c.logger.Infof("[%s] set key %s", action, key)
	rKey := fmt.Sprintf("dreon_ecommerce:%s", key)
	c.redisClient.Set(context.Background(), rKey, value, *expireTime)
	return nil
}

func (c *appCache) Get(key string) (interface{}, error) {
	action := "appCache.Get"
	c.logger.Infof("[%s] get key %s", action, key)
	rKey := fmt.Sprintf("dreon_ecommerce:%s", key)
	val, err := c.redisClient.Get(context.Background(), rKey).Result()
	if err != nil {
		return nil, err
	}

	return val, nil
}

func (c *appCache) Delete(key string) error {
	action := "appCache.Delete"
	c.logger.Infof("[%s] delete key %s", action, key)
	rKey := fmt.Sprintf("dreon_ecommerce:%s", key)
	c.redisClient.Del(context.Background(), rKey)
	return nil
}

func (c *appCache) Clear() error {
	action := "appCache.Clear"
	c.logger.Infof("[%s] clear cache", action)
	c.redisClient.FlushAll(context.Background())
	return nil
}

func (c *appCache) ClearWithPrefix(prefix string) error {
	action := "appCache.ClearWithPrefix"
	c.logger.Infof("[%s] clear cache with prefix %s", action, prefix)
	keys, err := c.redisClient.Keys(context.Background(), fmt.Sprintf("dreon_ecommerce:%s*", prefix)).Result()
	if err != nil {
		return err
	}

	for _, key := range keys {
		c.redisClient.Del(context.Background(), key)
	}

	return nil
}
