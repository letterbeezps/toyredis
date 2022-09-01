package hash

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -run TestGETSET
func TestGETSET(t *testing.T) {
	key := "firstHash"
	hash := New[[]byte]()

	// HSET
	value1 := []byte("value1")
	ret1 := hash.Hset(key, "field1", value1)

	fmt.Println(ret1)
	assert.Equal(t, ret1, 1)

	// KEYS
	keysret := hash.Keys()
	assert.Equal(t, keysret, []string{"firstHash"})

	// HGET
	ret2, ok := hash.Hget(key, "field1")
	assert.Equal(t, ok, true)
	fmt.Println(ret2)
	assert.Equal(t, ret2, []byte("value1"))

	// HSETNX
	value1copy := []byte("value1copy")
	ret3 := hash.Hsetnx(key, "field1", value1copy)
	assert.Equal(t, ret3, 0)
}
