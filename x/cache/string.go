package cache

import (
	"fmt"
	"time"
)

type String struct{}

func NewString() *String {
	return &String{}
}

func (s *String) Set(key, value string, cost int64, ttl time.Duration) bool {
	cache.Wait()
	return cache.SetWithTTL(key, value, cost, ttl)
}

func (s *String) Get(key string) (interface{}, bool) {
	cache.Wait()
	return cache.Get(key)
}

func (s *String) Key(key string) string {
	return fmt.Sprintf("string:data:bind:%s", key)
}
