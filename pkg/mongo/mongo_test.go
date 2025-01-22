package mongo

import (
	"errors"
	"testing"
	"time"

	"github.com/golang-module/carbon"
	"github.com/judegiordano/gogetem/pkg/logger"
	"github.com/judegiordano/gogetem/pkg/nanoid"
	"github.com/stretchr/testify/assert"
)

type Address struct {
	Street  string `bson:"street,omitempty" json:"street,omitempty"`
	Address int    `bson:"address,omitempty" json:"address,omitempty"`
}

type Profile struct {
	Username    string  `bson:"username,omitempty" json:"username,omitempty"`
	Avatar      int     `bson:"avatar,omitempty" json:"avatar,omitempty"`
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
		Name:    "john",
		Age:     27,
		Enabled: true,
		Hobbies: []string{"code", "video games", "travel"},
		Address: Address{
			Street:  "Fake St",
			Address: 123,
		},
		Profiles: []Profile{{
			Username:    "john",
			OldUsername: nil,
			Avatar:      162534,
		}, {
			Username:    "john_2",
			OldUsername: u,
			Avatar:      8945769,
		}},
		Balance:   200.64,
		CreatedAt: carbon.Now().Carbon2Time().UTC(),
		UpdatedAt: carbon.Now().Carbon2Time().UTC(),
	}
}

func TestObjectId(t *testing.T) {
	assert.NotEqual(t, ObjectId(), ObjectId())
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
	docs, err := List[User](Bson{})
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
	filter := Bson{"_id": inserted.Id}
	doc, err := Read[User](filter)
	assert.Nil(t, err)
	assert.Equal(t, doc.Id, inserted.Id, user.Id)
	logger.Info("document found:", *doc)
}

func TestReadById(t *testing.T) {
	user := mockUser()
	inserted, err := Insert[User](user)
	assert.Nil(t, err)
	// read
	doc, err := ReadById[User](inserted.Id)
	assert.Nil(t, err)
	assert.Equal(t, doc.Id, inserted.Id, user.Id)
	logger.Info("document found:", *doc)
}

func TestReadNil(t *testing.T) {
	filter := Bson{"_id": "NOT_FOUND"}
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
	filter := Bson{"name": Bson{"$in": names}}
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

func TestUpdateOne(t *testing.T) {
	user := mockUser()
	user.CreatedAt = carbon.Yesterday().Carbon2Time().UTC()
	user.UpdatedAt = carbon.Yesterday().Carbon2Time().UTC()
	inserted, err := Insert[User](user)
	assert.Nil(t, err)
	assert.NotNil(t, inserted)
	// update
	filter := Bson{"_id": inserted.Id}
	updates := Bson{
		"$set": Bson{
			"name":       "new_name",
			"updated_at": carbon.Now().Carbon2Time().UTC(),
		},
	}
	updated, err := UpdateOne[User](filter, updates)
	assert.Nil(t, err)
	assert.NotNil(t, updated)
	assert.Equal(t, updated.Name, "new_name")
	assert.Equal(t, updated.Id, inserted.Id)
	assert.True(t, updated.UpdatedAt.After(inserted.UpdatedAt))
}

func TestPushToArray(t *testing.T) {
	user := mockUser()
	user.CreatedAt = carbon.Yesterday().Carbon2Time().UTC()
	user.UpdatedAt = carbon.Yesterday().Carbon2Time().UTC()
	inserted, err := Insert[User](user)
	assert.Nil(t, err)
	assert.NotNil(t, inserted)
	// update
	filter := Bson{"_id": inserted.Id}
	updates := Bson{
		"$set": Bson{
			"name":       "new_array_name",
			"updated_at": carbon.Now().Carbon2Time().UTC(),
		},
		"$push": Bson{
			"hobbies": Bson{"$each": []string{"ayo", "pause"}},
		},
	}
	updated, err := UpdateOne[User](filter, updates)
	assert.Nil(t, err)
	assert.NotNil(t, updated)
	assert.Equal(t, updated.Name, "new_array_name")
	assert.Equal(t, updated.Id, inserted.Id)
	assert.True(t, updated.UpdatedAt.After(inserted.UpdatedAt))
	assert.True(t, len(updated.Hobbies) == 5)
	assert.Equal(t, updated.Hobbies, []string{"code", "video games", "travel", "ayo", "pause"})
}

func TestIncrement(t *testing.T) {
	user := mockUser()
	user.CreatedAt = carbon.Yesterday().Carbon2Time().UTC()
	user.UpdatedAt = carbon.Yesterday().Carbon2Time().UTC()
	inserted, err := Insert[User](user)
	assert.Nil(t, err)
	assert.NotNil(t, inserted)
	// update
	filter := Bson{"_id": inserted.Id}
	updates := Bson{
		"$set": Bson{
			"name":       "new_int_array",
			"updated_at": carbon.Now().Carbon2Time().UTC(),
		},
		"$inc": Bson{"age": 1},
	}
	updated, err := UpdateOne[User](filter, updates)
	assert.Nil(t, err)
	assert.NotNil(t, updated)
	assert.Equal(t, updated.Name, "new_int_array")
	assert.Equal(t, updated.Id, inserted.Id)
	assert.True(t, updated.UpdatedAt.After(inserted.UpdatedAt))
	assert.True(t, updated.Age == inserted.Age+1)
}

func TestUpdateMany(t *testing.T) {
	var users []User
	n, _ := nanoid.New()
	for i := 0; i < 10; i++ {
		user := mockUser()
		user.Name = n
		users = append(users, user)
	}
	InsertMany[User](users)

	filter := Bson{"name": n}
	newName, _ := nanoid.New()
	updates := Bson{
		"$set": Bson{
			"name":       newName,
			"updated_at": carbon.Now().Carbon2Time().UTC(),
		},
	}
	updated, err := UpdateMany[User](filter, updates)
	assert.Nil(t, err)
	assert.Equal(t, updated.MatchedCount, int64(10))
	assert.Equal(t, updated.ModifiedCount, int64(10))
}

func TestDeleteOne(t *testing.T) {
	user := mockUser()
	n, _ := nanoid.New()
	user.Name = n
	inserted, err := Insert[User](user)
	assert.Nil(t, err)
	// delete
	filter := Bson{"name": inserted.Name}
	removed, err := Delete[User](filter)
	assert.Nil(t, err)
	assert.Equal(t, removed.Id, inserted.Id)
	assert.Equal(t, removed.Name, inserted.Name)

	// user should not exist
	doc, err := Read[User](filter)
	assert.Nil(t, doc)
	assert.NotNil(t, err)
	assert.Equal(t, err, errors.New("mongo: no documents in result"))
}

func TestDeleteMany(t *testing.T) {
	var users []User
	n, _ := nanoid.New()
	for i := 0; i < 10; i++ {
		user := mockUser()
		user.Name = n
		users = append(users, user)
	}
	InsertMany[User](users)

	// delete
	filter := Bson{"name": n}
	removed, err := DeleteMany[User](filter)
	assert.Nil(t, err)
	assert.Equal(t, removed.DeletedCount, int64(10))

	// user should not exist
	doc, err := Read[User](filter)
	assert.Nil(t, doc)
	assert.NotNil(t, err)
	assert.Equal(t, err, errors.New("mongo: no documents in result"))
}

func TestEstimatedCount(t *testing.T) {
	var users []User
	for i := 0; i < 10; i++ {
		user := mockUser()
		users = append(users, user)
	}
	InsertMany[User](users)

	count, err := EstimatedCount[User]()
	assert.Nil(t, err)
	assert.NotNil(t, count)
	assert.True(t, *count >= 10)
}

func TestCount(t *testing.T) {
	var users []User
	n, _ := nanoid.New()

	for i := 0; i < 10; i++ {
		user := mockUser()
		user.Name = n
		users = append(users, user)
	}
	InsertMany[User](users)

	filter := Bson{"name": n}
	count, err := Count[User](filter)
	assert.Nil(t, err)
	assert.NotNil(t, count)
	assert.True(t, *count == 10)
}
