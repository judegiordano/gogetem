package cache

import "time"

func Set(key Key, val Value, expiresIn time.Duration) int {
	c := cache.Count()
	cache.Add(key, expiresIn, val)
	return c + 1
}
