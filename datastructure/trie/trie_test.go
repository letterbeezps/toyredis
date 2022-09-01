package trie

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -run TestTrie
func TestTrie(t *testing.T) {
	trie := New[int]()

	word1, value1 := "letter", 25

	trie.Insert(word1, value1)

	ok1, ret1 := trie.Search("lette")

	assert.Equal(t, ok1, false)
	assert.Equal(t, ret1, 0)

	ok2 := trie.StartWith("lette")
	assert.Equal(t, ok2, true)

	ok3, ret2 := trie.Search("letter")

	assert.Equal(t, ok3, true)
	assert.Equal(t, ret2, 25)

	word2, value2 := "lette", 26

	trie.Insert(word2, value2)

	ok4, ret3 := trie.Search("lette")
	assert.Equal(t, ok4, true)
	assert.Equal(t, ret3, 26)

}
