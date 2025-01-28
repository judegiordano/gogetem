package mongo

import (
	"github.com/judegiordano/gogetem/pkg/logger"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type InsertOneOptions = options.InsertOneOptions

func Insert[model interface{}](document model, opts ...*InsertOneOptions) (*model, error) {
	coll, ctx, cancel := collection[model]()
	defer cancel()
	_, err := coll.InsertOne(ctx, document, opts...)
	if err != nil {
		logger.Error("[MONGO INSERT]", err)
		return nil, err
	}
	return &document, nil
}
