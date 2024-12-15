package redis

import "time"

type Tool interface {
	Get(key string) (string, error)
	SetEX(key string, value interface{}, expiration time.Duration) error
	LPush(key string, value ...interface{}) error
	RPush(key string, value ...interface{}) error
	LRange(key string, start int, stop int) []string
	HSet(key string, value ...interface{}) error
}
