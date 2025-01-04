package cache

import (
	"dreon_ecommerce_server/shared/interfaces"
	"encoding/json"
	"fmt"
	"time"

	"github.com/patrickmn/go-cache"
)

type appCache struct {
	cache  *cache.Cache
	logger interfaces.ILogger
	interfaces.ICache
}

func NewAppCache(logger interfaces.ILogger, defaultExpire, cleanupInterval *time.Duration) *appCache {
	if defaultExpire == nil {
		defaultExpire = new(time.Duration)
		*defaultExpire = time.Hour * 24
	}
	if cleanupInterval == nil {
		cleanupInterval = new(time.Duration)
		*cleanupInterval = time.Hour * 24
	}

	logger.Info("Connecting to cache server")
	return &appCache{
		cache:  cache.New(*defaultExpire, *cleanupInterval),
		logger: logger,
	}
}

func (c *appCache) Set(key string, value interface{}, expireTime *time.Duration) error {
	action := "appCache.Set"
	c.logger.Infof("[%s] set key %s", action, key)
	if expireTime == nil {
		expireTime = new(time.Duration)
		*expireTime = cache.DefaultExpiration
	}
	b, err := json.Marshal(value)
	if err != nil {
		return err
	}
	c.cache.Set(fmt.Sprintf("dreon_ecommerce:%s", key), b, *expireTime)
	return nil
}

func (c *appCache) Get(key string) (interface{}, error) {
	action := "appCache.Get"
	c.logger.Infof("[%s] get key %s", action, key)
	value, found := c.cache.Get(fmt.Sprintf("dreon_ecommerce:%s", key))
	if !found {
		return nil, nil
	}
	var result interface{}
	err := json.Unmarshal(value.([]byte), &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *appCache) Delete(key string) error {
	action := "appCache.Delete"
	c.logger.Infof("[%s] delete key %s", action, key)
	c.cache.Delete(fmt.Sprintf("dreon_ecommerce:%s", key))
	return nil
}

func (c *appCache) Clear() error {
	action := "appCache.Clear"
	c.logger.Infof("[%s] clear cache", action)
	c.cache.Flush()
	return nil
}

func (c *appCache) ClearWithPrefix(prefix string) error {
	return nil
}
