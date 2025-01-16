package mongo

import (
	"context"
	"fmt"
	"reflect"
	"regexp"
	"strings"
	"time"

	"github.com/judegiordano/gogetem/pkg/dotenv"
	"github.com/judegiordano/gogetem/pkg/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/mongo/driver/connstring"
)

var Client *mongo.Client
var Database *string

func init() {
	if Client != nil {
		logger.Debug("reusing existing client")
		return
	}
	uri := dotenv.String("MONGO_URI")
	if uri == nil {
		logger.Fatal("no connection string set")
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	// init
	var err error
	url, err := connstring.ParseAndValidate(*uri)
	if err != nil {
		logger.Fatal("error parsing connection string %v", err)
	}
	Database = &url.Database
	opts := options.Client().ApplyURI(*uri)
	logger.Info("connecting to %v...", *opts.AppName)
	Client, err = mongo.Connect(ctx, opts)
	if err != nil {
		logger.Fatal("error connecting to mongo %v", err)
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

func List[model interface{}](filter interface{}, opts ...*options.FindOptions) ([]model, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	collName := collectionName[model]()
	coll := Client.Database(*Database).Collection(collName)
	var results []model
	cursor, err := coll.Find(ctx, filter, opts...)
	if err != nil {
		return results, err
	}
	if err = cursor.All(ctx, &results); err != nil {
		return results, err
	}
	return results, nil
}

func Insert[model interface{}](document model, opts ...*options.InsertOneOptions) (*model, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	collName := collectionName[model]()
	coll := Client.Database(*Database).Collection(collName)
	_, err := coll.InsertOne(ctx, document, opts...)
	if err != nil {
		return nil, err
	}
	return &document, nil
}
