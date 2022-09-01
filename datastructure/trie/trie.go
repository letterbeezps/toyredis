package trie

type Trie[T any] struct {
	isEnd    bool
	value    T
	children map[rune]*Trie[T]
}

func New[T any]() *Trie[T] {
	return &Trie[T]{
		isEnd:    false,
		children: map[rune]*Trie[T]{},
	}
}

func (root *Trie[T]) Insert(word string, value T) {
	node := root
	for _, c := range word {
		if child, ok := node.children[c]; ok {
			node = child
		} else {
			newNode := &Trie[T]{
				children: map[rune]*Trie[T]{},
			}
			node.children[c] = newNode
			node = newNode
		}
	}

	node.isEnd = true
	node.value = value
}

func (root *Trie[T]) Search(word string) (isEnd bool, value T) {
	node := root
	for _, c := range word {
		if child, ok := node.children[c]; ok {
			node = child
			continue
		} else {
			isEnd = false
			return
		}
	}
	isEnd = node.isEnd
	value = node.value
	return
}

func (root *Trie[T]) StartWith(prefix string) bool {
	node := root
	for _, c := range prefix {
		if child, ok := node.children[c]; ok {
			node = child
			continue
		} else {
			return false
		}
	}
	return true
}
