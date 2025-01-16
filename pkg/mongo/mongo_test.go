package mongo

import (
	"errors"
	"testing"
	"time"

	"github.com/judegiordano/gogetem/pkg/logger"
	"github.com/judegiordano/gogetem/pkg/nanoid"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
)

type Address struct {
	Street  string `bson:"street,omitempty" json:"street,omitempty"`
	Address int    `bson:"address,omitempty" json:"address,omitempty"`
}

type Profile struct {
	Username    string  `bson:"username,omitempty" json:"username,omitempty"`
	Avavtar     int     `bson:"avatar,omitempty" json:"avatar,omitempty"`
	OldUsername *string `bson:"old_username" json:"old_username"`
}

type User struct {
	Id        string    `bson:"_id,omitempty" json:"id,omitempty"`
	Name      string    `bson:"name,omitempty" json:"name,omitempty"`
	Age       int       `bson:"age,omitempty" json:"age,omitempty"`
	Enabled   bool      `bson:"enabled,omitempty" json:"enabled,omitempty"`
	Hobbies   []string  `bson:"hobbies,omitempty" json:"hobbies,omitempty"`
	Address   Address   `bson:"address,omitempty" json:"address,omitempty"`
	Profiles  []Profile `bson:"profiles,omitempty" json:"profiles,omitempty"`
	Balance   float64   `bson:"balance,omitempty" json:"balance,omitempty"`
	CreatedAt time.Time `bson:"created_at,omitempty" json:"created_at,omitempty"`
	UpdatedAt time.Time `bson:"updated_at,omitempty" json:"updated_at,omitempty"`
}

func mockUser() User {
	id, err := nanoid.New()
	if err != nil {
		logger.Fatal("error mocking user: ", err)
	}
	u := new(string)
	*u = "old_username"
	return User{
		Id:      id,
		Name:    "judeboy",
		Age:     27,
		Enabled: true,
		Hobbies: []string{"code", "video games", "travel"},
		Address: Address{
			Street:  "Fake St",
			Address: 123,
		},
		Profiles: []Profile{{
			Username:    "judeboy",
			OldUsername: nil,
			Avavtar:     162534,
		}, {
			Username:    "judeboy_2",
			OldUsername: u,
			Avavtar:     8945769,
		}},
		Balance:   200.64,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}
}

func TestCollectionName(t *testing.T) {
	name := collectionName[User]()
	assert.Equal(t, name, "users")
}

func TestClientConnection(t *testing.T) {
	assert.NotNil(t, Client)
}

func TestDb(t *testing.T) {
	assert.NotNil(t, Database)
}

func TestInsertOne(t *testing.T) {
	user := mockUser()
	inserted, err := Insert[User](user)
	assert.Nil(t, err)
	assert.NotNil(t, inserted)
	assert.Equal(t, inserted.Id, user.Id)
}

func TestList(t *testing.T) {
	docs, err := List[User](bson.D{})
	assert.Nil(t, err)
	for _, doc := range docs {
		assert.NotNil(t, doc.Id)
	}
}

func TestRead(t *testing.T) {
	user := mockUser()
	inserted, err := Insert[User](user)
	assert.Nil(t, err)
	// read
	filter := bson.M{"_id": inserted.Id}
	doc, err := Read[User](filter)
	assert.Nil(t, err)
	assert.Equal(t, doc.Id, inserted.Id, user.Id)
	logger.Info("document found:", *doc)
}

func TestReadNil(t *testing.T) {
	filter := bson.M{"_id": "NOT_FOUND"}
	doc, err := Read[User](filter)
	assert.NotNil(t, err)
	assert.Nil(t, doc)
	assert.Equal(t, err, errors.New("mongo: no documents in result"))
}

func TestListIn(t *testing.T) {
	for i := 0; i < 10; i++ {
		user := mockUser()
		if i%2 == 0 {
			user.Name = "even_name"
		} else {
			user.Name = "odd_name"
		}
		_, err := Insert(user)
		assert.Nil(t, err)
	}
	names := []string{"even_name", "odd_name"}
	filter := bson.M{"name": bson.M{"$in": names}}
	docs, err := List[User](filter)
	assert.Nil(t, err)
	assert.NotNil(t, docs)
	for _, doc := range docs {
		assert.NotNil(t, docs)
		if doc.Name != "even_name" && doc.Name != "odd_name" {
			logger.Error(doc)
			t.Error("doc.Name should match filter")
		}
	}
}

func TestInsertMany(t *testing.T) {
	var users []User
	for i := 0; i < 10; i++ {
		user := mockUser()
		user.Name = "bulk_written_username"
		users = append(users, user)
	}
	inserted, err := InsertMany[User](users)
	assert.Nil(t, err)
	assert.NotNil(t, inserted)
	for _, doc := range inserted {
		assert.Equal(t, doc.Name, "bulk_written_username")
	}
}
