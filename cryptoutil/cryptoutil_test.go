package cryptoutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMD5Str(t *testing.T) {
	assert.Equal(t, "b0804ec967f48520697662a204f5fe72", MD5Str("These pretzels are making me thirsty."))
}
