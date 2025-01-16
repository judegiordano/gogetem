package mongo

import (
	"testing"
	"time"

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
	id, err := nanoid.New()
	assert.Nil(t, err)
	u := new(string)
	*u = "old_username"
	user := User{
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
