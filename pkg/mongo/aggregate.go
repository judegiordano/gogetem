package mongo

import (
	"github.com/judegiordano/gogetem/pkg/logger"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type AggregateOptions = options.AggregateOptions

func Aggregate[model, out interface{}](pipeline interface{}, opts ...*AggregateOptions) ([]out, error) {
	coll, ctx, cancel := collection[model]()
	defer cancel()
	cursor, err := coll.Aggregate(ctx, pipeline, opts...)
	var results = []out{}
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
