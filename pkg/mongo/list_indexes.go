package mongo

import (
	"github.com/judegiordano/gogetem/pkg/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ListIndexesOptions = options.ListIndexesOptions

type IndexBatch struct {
	v    int
	key  Bson
	name string
	ns   string
}

func ListIndexes[model interface{}](opts ...*ListIndexesOptions) ([]IndexBatch, error) {
	coll, ctx, cancel := collection[model]()
	defer cancel()
	var idx_raw []bson.M
	cursor, err := coll.Indexes().List(ctx, opts...)
	if err != nil {
		logger.Error("[MONGO LIST INDEXES]", err)
		return nil, err
	}
	if err = cursor.All(ctx, &idx_raw); err != nil {
		logger.Error("[MONGO LIST INDEXES CURSOR]", err)
		return nil, err
	}

	var idx_batches []IndexBatch
	for _, index := range idx_raw {
		v, _ := index["v"].(int32)
		key, _ := index["key"].(Bson)
		name, _ := index["name"].(string)
		ns, _ := index["ns"].(string)

		idx_batches = append(idx_batches, IndexBatch{
			v:    int(v),
			key:  key,
			name: name,
			ns:   ns,
		})
	}
	return idx_batches, nil
}
