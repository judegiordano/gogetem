package mongo

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strings"
	"time"

	"github.com/judegiordano/gogetem/pkg/dotenv"
	"github.com/judegiordano/gogetem/pkg/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/mongo/driver/connstring"
)

var Client *mongo.Client
var Database *string

type Bson bson.M

func init() {
	if Client != nil {
		logger.Debug("[MONGO CONNECTION]", "reusing existing client")
		return
	}
	uri := dotenv.String("MONGO_URI")
	if uri == nil {
		logger.Fatal("[MONGO CONNECTION]", "no connection string set")
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	// init
	var err error
	url, err := connstring.ParseAndValidate(*uri)
	if err != nil {
		logger.Fatal("[MONGO CONNECTION]", "error parsing connection string:", err)
	}
	Database = &url.Database
	opts := options.Client().ApplyURI(*uri)
	logger.Debug("connecting to", *opts.AppName)
	Client, err = mongo.Connect(ctx, opts)
	if err != nil {
		logger.Fatal("[MONGO CONNECTION]", err)
	}
}

var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

func toSnakeCase(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

func collectionName[model interface{}]() string {
	name := reflect.TypeOf((*model)(nil)).Elem().Name()
	snake := toSnakeCase(name)
	if !strings.HasSuffix(snake, "s") {
		snake = fmt.Sprintf("%vs", snake)
	}
	return snake
}

func collection[model interface{}]() (*mongo.Collection, context.Context, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	name := collectionName[model]()
	coll := Client.Database(*Database).Collection(name)
	return coll, ctx, cancel
}

func List[model interface{}](filter interface{}, opts ...*options.FindOptions) ([]model, error) {
	coll, ctx, cancel := collection[model]()
	defer cancel()
	var results []model
	cursor, err := coll.Find(ctx, filter, opts...)
	if err != nil {
		logger.Error("[MONGO LIST]", err)
		return results, err
	}
	if err = cursor.All(ctx, &results); err != nil {
		logger.Error("[MONGO LIST CURSOR]", err)
		return results, err
	}
	return results, nil
}

func Insert[model interface{}](document model, opts ...*options.InsertOneOptions) (*model, error) {
	coll, ctx, cancel := collection[model]()
	defer cancel()
	_, err := coll.InsertOne(ctx, document, opts...)
	if err != nil {
		logger.Error("[MONGO INSERT]", err)
		return nil, err
	}
	return &document, nil
}

func Read[model interface{}](filter interface{}, opts ...*options.FindOneOptions) (*model, error) {
	coll, ctx, cancel := collection[model]()
	defer cancel()
	var out model
	result := coll.FindOne(ctx, filter, opts...)
	if result == nil {
		logger.Error("[MONGO READ]", "no documents returned")
		return nil, errors.New("no document found")
	}
	if err := result.Decode(&out); err != nil {
		logger.Error("[MONGO READ]", err)
		return nil, err
	}
	return &out, nil
}

func ReadById[model interface{}](_id string, opts ...*options.FindOneOptions) (*model, error) {
	coll, ctx, cancel := collection[model]()
	defer cancel()
	var out model
	result := coll.FindOne(ctx, Bson{"_id": _id}, opts...)
	if result == nil {
		logger.Error("[MONGO READ_BY_ID]", "no document returned")
		return nil, errors.New("no document found")
	}
	if err := result.Decode(&out); err != nil {
		logger.Error("[MONGO READ_BY_ID]", err)
		return nil, err
	}
	return &out, nil
}

func InsertMany[model interface{}](docs []model, opts ...*options.InsertManyOptions) ([]model, error) {
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

func UpdateOne[model interface{}](filter interface{}, updates interface{}, opts ...*options.FindOneAndUpdateOptions) (*model, error) {
	coll, ctx, cancel := collection[model]()
	defer cancel()
	var out model
	after := options.After
	options := append(opts, &options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
	})
	result := coll.FindOneAndUpdate(ctx, filter, updates, options...)
	if result == nil {
		logger.Error("[MONGO UPDATE_ONE]", "no documents found")
		return nil, errors.New("no document found")
	}
	if err := result.Decode(&out); err != nil {
		logger.Error("[MONGO UPDATE_ONE]", err)
		return nil, err
	}
	return &out, nil
}

func UpdateMany[model interface{}](filter interface{}, updates interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	coll, ctx, cancel := collection[model]()
	defer cancel()
	result, err := coll.UpdateMany(ctx, filter, updates, opts...)
	if err != nil {
		logger.Error("[MONGO UPDATE_MANY]", err)
		return nil, err
	}
	if result == nil {
		logger.Error("[MONGO UPDATE_MANY]", "error updating many")
		return nil, errors.New("error updating many")
	}
	return result, nil
}

func Delete[model interface{}](filter interface{}, opts ...*options.FindOneAndDeleteOptions) (*model, error) {
	coll, ctx, cancel := collection[model]()
	defer cancel()
	var out model
	result := coll.FindOneAndDelete(ctx, filter, opts...)
	if result == nil {
		logger.Error("[MONGO DELETE]", "no documents returned")
		return nil, errors.New("no document found")
	}
	if err := result.Decode(&out); err != nil {
		logger.Error("[MONGO DELETE]", err)
		return nil, err
	}
	return &out, nil
}

func DeleteMany[model interface{}](filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error) {
	coll, ctx, cancel := collection[model]()
	defer cancel()
	result, err := coll.DeleteMany(ctx, filter, opts...)
	if err != nil {
		logger.Error("[MONGO DELETE_MANY]", err)
		return nil, err
	}
	if result == nil {
		logger.Error("[MONGO DELETE_MANY]", "error deleting many")
		return nil, errors.New("error deleting many")
	}
	return result, nil
}

func EstimatedCount[model interface{}](opts ...*options.EstimatedDocumentCountOptions) (*int64, error) {
	coll, ctx, cancel := collection[model]()
	defer cancel()
	result, err := coll.EstimatedDocumentCount(ctx, opts...)
	if err != nil {
		logger.Error("[MONGO ESTIMATED_COUNT]", err)
		return nil, err
	}
	return &result, nil
}

func Count[model interface{}](filter interface{}, opts ...*options.CountOptions) (*int64, error) {
	coll, ctx, cancel := collection[model]()
	defer cancel()
	result, err := coll.CountDocuments(ctx, filter, opts...)
	if err != nil {
		logger.Error("[MONGO COUNT]", err)
		return nil, err
	}
	return &result, nil
}
