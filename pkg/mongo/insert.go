package mongo

import (
	"github.com/judegiordano/gogetem/pkg/logger"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type InsertOneOptions = options.InsertOneOptions

func Insert[T Model](document T, opts ...*InsertOneOptions) (*T, error) {
	coll, ctx, cancel := collection(document.CollectionName())
	defer cancel()
	_, err := coll.InsertOne(ctx, document, opts...)
	if err != nil {
		logger.Error("[MONGO INSERT]", err)
		return nil, err
	}
	return &document, nil
}
