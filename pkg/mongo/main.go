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
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/mongo/driver/connstring"
)

var Client *mongo.Client
var Database *string

type Bson bson.M

type Document struct {
	Id        string    `bson:"_id,omitempty" json:"id"`
	CreatedAt time.Time `bson:"created_at,omitempty" json:"created_at"`
	UpdatedAt time.Time `bson:"updated_at,omitempty" json:"updated_at"`
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
