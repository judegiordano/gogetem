package mongo

import (
	"github.com/judegiordano/gogetem/pkg/logger"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type FindOptions = options.FindOptions

func List[T Model](filter interface{}, opts ...*FindOptions) ([]T, error) {
	var model T
	coll, ctx, cancel := collection(model.CollectionName())
	defer cancel()
	cursor, err := coll.Find(ctx, filter, opts...)
	var results = []T{}
	if err != nil {
		logger.Error("[MONGO LIST]", err)
		return results, err
	}
	if err = cursor.All(ctx, &results); err != nil {
		logger.Error("[MONGO LIST CURSOR]", err)
		return results, err
	}
	return results, nil
}
