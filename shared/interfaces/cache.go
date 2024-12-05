package interfaces

import "time"

type ICache interface {
	Set(key string, value interface{}, expireTime *time.Duration) error
	Get(key string) (interface{}, error)
	Delete(key string) error
	Clear() error
	ClearWithPrefix(prefix string) error
}
