package mongo

import (
	"github.com/judegiordano/gogetem/pkg/logger"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type AggregateOptions = options.AggregateOptions

func Aggregate[T Model](pipeline interface{}, opts ...*AggregateOptions) ([]T, error) {
	var model T
	coll, ctx, cancel := collection(model.CollectionName())
	defer cancel()
	cursor, err := coll.Aggregate(ctx, pipeline, opts...)
	var results = []T{}
	if err != nil {
		logger.Error("[MONGO AGGREGATE]", err)
		return results, err
	}
	if err = cursor.All(ctx, &results); err != nil {
		logger.Error("[MONGO AGGREGATE CURSOR]", err)
		return results, err
	}
	return results, nil
}
