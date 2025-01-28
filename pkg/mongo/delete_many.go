package mongo

import (
	"errors"

	"github.com/judegiordano/gogetem/pkg/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DeleteOptions = options.DeleteOptions
type DeleteResult = mongo.DeleteResult

func DeleteMany[model interface{}](filter interface{}, opts ...*DeleteOptions) (*DeleteResult, error) {
	coll, ctx, cancel := collection[model]()
	defer cancel()
	result, err := coll.DeleteMany(ctx, filter, opts...)
	if err != nil {
		logger.Error("[MONGO DELETE_MANY]", err)
		return nil, err
	}
	if result == nil {
		logger.Error("[MONGO DELETE_MANY]", "error deleting many")
		return nil, errors.New("error deleting many")
	}
	return result, nil
}
