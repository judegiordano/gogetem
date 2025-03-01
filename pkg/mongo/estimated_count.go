package mongo

import (
	"github.com/judegiordano/gogetem/pkg/logger"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type EstimatedDocumentCountOptions = options.EstimatedDocumentCountOptions

func EstimatedCount[T Model](opts ...*EstimatedDocumentCountOptions) (*int64, error) {
	var model T
	coll, ctx, cancel := collection(model.CollectionName())
	defer cancel()
	result, err := coll.EstimatedDocumentCount(ctx, opts...)
	if err != nil {
		logger.Error("[MONGO ESTIMATED_COUNT]", err)
		return nil, err
	}
	return &result, nil
}
