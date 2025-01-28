package mongo

import (
	"github.com/judegiordano/gogetem/pkg/logger"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type CountOptions = options.CountOptions

func Count[model interface{}](filter interface{}, opts ...*CountOptions) (*int64, error) {
	coll, ctx, cancel := collection[model]()
	defer cancel()
	result, err := coll.CountDocuments(ctx, filter, opts...)
	if err != nil {
		logger.Error("[MONGO COUNT]", err)
		return nil, err
	}
	return &result, nil
}
