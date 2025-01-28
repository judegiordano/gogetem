package cache

import "errors"

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
