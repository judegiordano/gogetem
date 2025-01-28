package mongo

import (
	"github.com/judegiordano/gogetem/pkg/logger"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type EstimatedDocumentCountOptions = options.EstimatedDocumentCountOptions

func EstimatedCount[model interface{}](opts ...*EstimatedDocumentCountOptions) (*int64, error) {
	coll, ctx, cancel := collection[model]()
	defer cancel()
	result, err := coll.EstimatedDocumentCount(ctx, opts...)
	if err != nil {
		logger.Error("[MONGO ESTIMATED_COUNT]", err)
		return nil, err
	}
	return &result, nil
}
