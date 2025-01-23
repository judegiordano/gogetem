package cache

import (
	"errors"
	"time"

	"github.com/muesli/cache2go"
)

var cache *cache2go.CacheTable

type Key = interface{}
type Value = interface{}

func Set(key Key, val Value, expiresIn time.Duration) int {
	c := cache.Count()
	cache.Add(key, expiresIn, val)
	return c + 1
}

func Size() int {
	return cache.Count()
}

func Empty() int {
	cache.Flush()
	return cache.Count()
}

func Get[T any](key Key) (*T, error) {
	v, err := cache.Value(key)
	if err != nil {
		return nil, err
	}
	data := v.Data()
	result, ok := data.(T)
	if !ok {
		return nil, errors.New("error parsing value")
	}
	return &result, nil
}

func Remove(key Key) (*Key, error) {
	_, err := cache.Delete(key)
	if err != nil {
		return nil, err
	}
	return &key, nil
}

func init() {
	cache = cache2go.Cache("gogetem:cache")
}
