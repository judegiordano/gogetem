package mongo

import (
	"errors"

	"github.com/judegiordano/gogetem/pkg/logger"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type FindOneAndUpdateOptions = options.FindOneAndUpdateOptions

func UpdateOne[T Model](filter interface{}, updates interface{}, opts ...*FindOneAndUpdateOptions) (*T, error) {
	var model T
	coll, ctx, cancel := collection(model.CollectionName())
	defer cancel()
	after := options.After
	options := append(opts, &FindOneAndUpdateOptions{
		ReturnDocument: &after,
	})
	result := coll.FindOneAndUpdate(ctx, filter, updates, options...)
	if result == nil {
		logger.Error("[MONGO UPDATE_ONE]", "no documents found")
		return nil, errors.New("no document found")
	}
	if err := result.Decode(&model); err != nil {
		logger.Error("[MONGO UPDATE_ONE]", err)
		return nil, err
	}
	return &model, nil
}
