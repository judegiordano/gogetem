package mongo

import (
	"github.com/judegiordano/gogetem/pkg/logger"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DropIndexesOptions = options.DropIndexesOptions

func DropIndex[T Model](index string, opts ...*DropIndexesOptions) (*string, error) {
	var model T
	coll, ctx, cancel := collection(model.CollectionName())
	defer cancel()
	_, err := coll.Indexes().DropOne(ctx, index, opts...)
	if err != nil {
		logger.Error("[MONGO DROP INDEX]", err)
		return nil, err
	}
	return &index, nil
}
