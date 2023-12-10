package lib_test

import (
	"arithmos/lib"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMD5(t *testing.T) {
	hash := lib.MD5([]byte("The quick brown fox jumps over the lazy dog"))
	assert.Equal(t, "9e107d9d372bb6826bd81d3542a419d6", fmt.Sprintf("%x", hash))
	hash = lib.MD5([]byte("The quick brown fox jumps over the lazy dog."))
	assert.Equal(t, "e4d909c290d0fb1ca068ffaddf22cbd0", fmt.Sprintf("%x", hash))
	hash = lib.MD5([]byte(""))
	assert.Equal(t, "d41d8cd98f00b204e9800998ecf8427e", fmt.Sprintf("%x", hash))
}
