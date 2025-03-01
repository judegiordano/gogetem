package mongo

import (
	"context"
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

type Model interface {
	CollectionName() string
}

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
	Client, err = mongo.Connect(ctx, opts)
	if err != nil {
		logger.Fatal("[MONGO CONNECTION]", err)
	}
}

func collection(name string) (*mongo.Collection, context.Context, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	coll := Client.Database(*Database).Collection(name)
	return coll, ctx, cancel
}
