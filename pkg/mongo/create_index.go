package mongo

import (
	"github.com/judegiordano/gogetem/pkg/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type CreateIndexesOptions = options.CreateIndexesOptions
type IndexModel = mongo.IndexModel
type IndexOptions = options.IndexOptions

func CreateIndex[model interface{}](index IndexModel, opts ...*CreateIndexesOptions) (*string, error) {
	coll, ctx, cancel := collection[model]()
	defer cancel()
	idx, err := coll.Indexes().CreateOne(ctx, index, opts...)
	if err != nil {
		logger.Error("[MONGO CREATE INDEX]", err)
		return nil, err
	}
	return &idx, nil
}
