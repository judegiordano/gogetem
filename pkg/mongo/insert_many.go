package mongo

import (
	"github.com/judegiordano/gogetem/pkg/logger"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type InsertManyOptions = options.InsertManyOptions

func InsertMany[model interface{}](docs []model, opts ...*InsertManyOptions) ([]model, error) {
	coll, ctx, cancel := collection[model]()
	defer cancel()
	bson := make([]interface{}, len(docs))
	for i, v := range docs {
		bson[i] = v
	}
	_, err := coll.InsertMany(ctx, bson, opts...)
	if err != nil {
		logger.Error("[MONGO INSERT MANY]", err)
		return nil, err
	}
	return docs, nil
}
