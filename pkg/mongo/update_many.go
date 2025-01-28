package mongo

import (
	"errors"

	"github.com/judegiordano/gogetem/pkg/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UpdateOptions = options.UpdateOptions
type UpdateResult = mongo.UpdateResult

func UpdateMany[model interface{}](filter interface{}, updates interface{}, opts ...*UpdateOptions) (*UpdateResult, error) {
	coll, ctx, cancel := collection[model]()
	defer cancel()
	result, err := coll.UpdateMany(ctx, filter, updates, opts...)
	if err != nil {
		logger.Error("[MONGO UPDATE_MANY]", err)
		return nil, err
	}
	if result == nil {
		logger.Error("[MONGO UPDATE_MANY]", "error updating many")
		return nil, errors.New("error updating many")
	}
	return result, nil
}
