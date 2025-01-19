package slug

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNormalizeString(t *testing.T) {
	i := "  hErE iS a meSSy sTrinG   "
	assert.Equal(t, "here-is-a-messy-string", Slugify(i))
}
