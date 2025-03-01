package mongo

import (
	"errors"

	"github.com/judegiordano/gogetem/pkg/logger"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type FindOneOptions = options.FindOneOptions

func ReadById[T Model](_id string, opts ...*FindOneOptions) (*T, error) {
	var model T
	coll, ctx, cancel := collection(model.CollectionName())
	defer cancel()
	result := coll.FindOne(ctx, Bson{"_id": _id}, opts...)
	if result == nil {
		logger.Error("[MONGO READ_BY_ID]", "no document returned")
		return nil, errors.New("no document found")
	}
	if err := result.Decode(&model); err != nil {
		logger.Error("[MONGO READ_BY_ID]", err)
		return nil, err
	}
	return &model, nil
}
