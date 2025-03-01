package mongo

import (
	"github.com/judegiordano/gogetem/pkg/logger"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type InsertManyOptions = options.InsertManyOptions

func InsertMany[T Model](docs []T, opts ...*InsertManyOptions) ([]T, error) {
	var model T
	coll, ctx, cancel := collection(model.CollectionName())
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
