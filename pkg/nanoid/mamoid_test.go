package nanoid

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNanoid(t *testing.T) {
	n1, err := New()
	assert.Nil(t, err)
	assert.NotNil(t, n1)
	assert.True(t, len(n1) > 1)
	// create second
	n2, err := New()
	assert.Nil(t, err)
	assert.NotNil(t, n2)
	assert.True(t, len(n2) > 1)
	// uniqueness
	assert.NotEqual(t, n1, n2)
}
