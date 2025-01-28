package cache

func Remove(key Key) (*Key, error) {
	_, err := cache.Delete(key)
	if err != nil {
		return nil, err
	}
	return &key, nil
}
