package mongo

import (
	"errors"

	"github.com/judegiordano/gogetem/pkg/logger"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type FindOneAndUpdateOptions = options.FindOneAndUpdateOptions

func UpdateOne[model interface{}](filter interface{}, updates interface{}, opts ...*FindOneAndUpdateOptions) (*model, error) {
	coll, ctx, cancel := collection[model]()
	defer cancel()
	var out model
	after := options.After
	options := append(opts, &FindOneAndUpdateOptions{
		ReturnDocument: &after,
	})
	result := coll.FindOneAndUpdate(ctx, filter, updates, options...)
	if result == nil {
		logger.Error("[MONGO UPDATE_ONE]", "no documents found")
		return nil, errors.New("no document found")
	}
	if err := result.Decode(&out); err != nil {
		logger.Error("[MONGO UPDATE_ONE]", err)
		return nil, err
	}
	return &out, nil
}
