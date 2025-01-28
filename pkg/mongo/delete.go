package mongo

import (
	"errors"

	"github.com/judegiordano/gogetem/pkg/logger"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type FindOneAndDeleteOptions = options.FindOneAndDeleteOptions

func Delete[model interface{}](filter interface{}, opts ...*FindOneAndDeleteOptions) (*model, error) {
	coll, ctx, cancel := collection[model]()
	defer cancel()
	var out model
	result := coll.FindOneAndDelete(ctx, filter, opts...)
	if result == nil {
		logger.Error("[MONGO DELETE]", "no documents returned")
		return nil, errors.New("no document found")
	}
	if err := result.Decode(&out); err != nil {
		logger.Error("[MONGO DELETE]", err)
		return nil, err
	}
	return &out, nil
}
