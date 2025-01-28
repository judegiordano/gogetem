package cache

func Empty() int {
	cache.Flush()
	return cache.Count()
}
