package cache

import "github.com/muesli/cache2go"

var cache *cache2go.CacheTable

type Key = interface{}
type Value = interface{}

func init() {
	cache = cache2go.Cache("gogetem:cache")
}
