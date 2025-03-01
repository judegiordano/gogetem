package mongo

import (
	"errors"

	"github.com/judegiordano/gogetem/pkg/logger"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type FindOneAndDeleteOptions = options.FindOneAndDeleteOptions

func Delete[T Model](filter interface{}, opts ...*FindOneAndDeleteOptions) (*T, error) {
	var model T
	coll, ctx, cancel := collection(model.CollectionName())
	defer cancel()
	result := coll.FindOneAndDelete(ctx, filter, opts...)
	if result == nil {
		logger.Error("[MONGO DELETE]", "no documents returned")
		return nil, errors.New("no document found")
	}
	if err := result.Decode(&model); err != nil {
		logger.Error("[MONGO DELETE]", err)
		return nil, err
	}
	return &model, nil
}
