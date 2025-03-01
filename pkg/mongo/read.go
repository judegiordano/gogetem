package mongo

import (
	"errors"

	"github.com/judegiordano/gogetem/pkg/logger"
)

func Read[T Model](filter interface{}, opts ...*FindOneOptions) (*T, error) {
	var model T
	coll, ctx, cancel := collection(model.CollectionName())
	defer cancel()
	result := coll.FindOne(ctx, filter, opts...)
	if result == nil {
		logger.Error("[MONGO READ]", "no documents returned")
		return nil, errors.New("no document found")
	}
	if err := result.Decode(&model); err != nil {
		logger.Error("[MONGO READ]", err)
		return nil, err
	}
	return &model, nil
}
