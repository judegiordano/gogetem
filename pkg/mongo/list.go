package mongo

import (
	"github.com/judegiordano/gogetem/pkg/logger"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type FindOptions = options.FindOptions

func List[model interface{}](filter interface{}, opts ...*FindOptions) ([]model, error) {
	coll, ctx, cancel := collection[model]()
	defer cancel()
	cursor, err := coll.Find(ctx, filter, opts...)
	var results = []model{}
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
