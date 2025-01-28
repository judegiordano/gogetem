package mongo

import (
	"errors"

	"github.com/judegiordano/gogetem/pkg/logger"
)

func Read[model interface{}](filter interface{}, opts ...*FindOneOptions) (*model, error) {
	coll, ctx, cancel := collection[model]()
	defer cancel()
	var out model
	result := coll.FindOne(ctx, filter, opts...)
	if result == nil {
		logger.Error("[MONGO READ]", "no documents returned")
		return nil, errors.New("no document found")
	}
	if err := result.Decode(&out); err != nil {
		logger.Error("[MONGO READ]", err)
		return nil, err
	}
	return &out, nil
}
