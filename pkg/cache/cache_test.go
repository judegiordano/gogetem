package cache

import (
	"errors"
	"testing"
	"time"

	"github.com/judegiordano/gogetem/pkg/nanoid"
	"github.com/stretchr/testify/assert"
)

func TestSet(t *testing.T) {
	exp := time.Second * 30
	k, _ := nanoid.New()
	v, _ := nanoid.New()
	size := Set(k, v, exp)
	assert.True(t, size >= 1)
}

func TestGet(t *testing.T) {
	exp := time.Second * 30
	k, _ := nanoid.New()
	v, _ := nanoid.New()
	size := Set(k, v, exp)

	assert.True(t, size >= 1)
	found, err := Get[string](k)
	assert.Nil(t, err)
	assert.NotNil(t, found)
	assert.Equal(t, *found, v)
}

func TestCount(t *testing.T) {
	exp := time.Second * 30
	for i := 0; i < 10; i++ {
		k, _ := nanoid.New()
		v, _ := nanoid.New()
		Set(k, v, exp)
	}
	c := Size()
	assert.True(t, c >= 10)
}

func TestEmpty(t *testing.T) {
	exp := time.Second * 30
	for i := 0; i < 10; i++ {
		k, _ := nanoid.New()
		v, _ := nanoid.New()
		Set(k, v, exp)
	}
	c := Size()
	assert.True(t, c >= 10)

	empty := Empty()
	assert.True(t, empty == 0)
}

func TestGetInt(t *testing.T) {
	exp := time.Second * 30
	k, _ := nanoid.New()
	v := 22
	size := Set(k, v, exp)

	assert.True(t, size >= 1)
	found, err := Get[int](k)
	assert.Nil(t, err)
	assert.NotNil(t, found)
	assert.Equal(t, *found, v)
}

func TestGetStruct(t *testing.T) {
	exp := time.Second * 30
	type Data struct {
		Id   string
		Name string
	}
	k, _ := nanoid.New()
	id, _ := nanoid.New()
	name, _ := nanoid.New()
	v := Data{
		Id:   id,
		Name: name,
	}
	size := Set(k, v, exp)

	assert.True(t, size >= 1)
	found, err := Get[Data](k)
	assert.Nil(t, err)
	assert.NotNil(t, found)
	assert.Equal(t, *found, v)
	assert.Equal(t, found.Id, v.Id)
	assert.Equal(t, found.Name, v.Name)
}

func TestGetBool(t *testing.T) {
	exp := time.Second * 30
	k, _ := nanoid.New()
	v := false
	size := Set(k, v, exp)

	assert.True(t, size >= 1)
	found, err := Get[bool](k)
	assert.Nil(t, err)
	assert.NotNil(t, found)
	assert.Equal(t, *found, v)
}

func TestGetSlice(t *testing.T) {
	exp := time.Second * 30
	k, _ := nanoid.New()
	v := []string{"cat", "dog"}
	size := Set(k, v, exp)

	assert.True(t, size >= 1)
	found, err := Get[[]string](k)
	assert.Nil(t, err)
	assert.NotNil(t, found)
	assert.Equal(t, *found, v)
}

func TestRemove(t *testing.T) {
	exp := time.Second * 30
	k, _ := nanoid.New()
	v, _ := nanoid.New()
	size := Set(k, v, exp)

	assert.True(t, size >= 1)
	found, err := Get[string](k)
	assert.Nil(t, err)
	assert.NotNil(t, found)
	assert.Equal(t, *found, v)

	removed, err := Remove(k)
	assert.Nil(t, err)
	assert.NotNil(t, removed)
	assert.Equal(t, *removed, k)

	gone, err := Get[string](k)
	assert.Nil(t, gone)
	assert.NotNil(t, err)
	assert.Equal(t, err, errors.New("Key not found in cache"))
}

func TestErrorParsing(t *testing.T) {
	exp := time.Second * 30
	k, _ := nanoid.New()
	v, _ := nanoid.New()
	size := Set(k, v, exp)

	assert.True(t, size >= 1)
	found, err := Get[int](k)
	assert.Nil(t, found)
	assert.NotNil(t, err)
	assert.Equal(t, err, errors.New("error parsing value"))
}

func TestExpiration(t *testing.T) {
	Empty()
	exp := time.Microsecond
	k, _ := nanoid.New()
	v, _ := nanoid.New()
	size := Set(k, v, exp)
	assert.True(t, size >= 1)
	// pause
	time.Sleep(time.Millisecond)

	found, err := Get[string](k)
	assert.Nil(t, found)
	assert.NotNil(t, err)
	assert.Equal(t, err, errors.New("Key not found in cache"))
	assert.True(t, Size() < size)
}
